/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"

	"github.com/nlopes/slack"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	crdv1 "kube.op/controller/api/v1"
)

// PodTrackerReconciler reconciles a PodTracker object
type PodTrackerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=crd.kube.op,resources=podtrackers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=crd.kube.op,resources=podtrackers/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=crd.kube.op,resources=podtrackers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the PodTracker object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.20.2/pkg/reconcile
func (r *PodTrackerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	logger.Info("Controller Triggered!!!")

	// var podTracker crdv1.PodTracker
	// if err := r.Get(ctx, req.NamespacedName, &podTracker); err != nil {
	// 	logger.Error(err, "unable to fetch pod tracker")
	// 	return ctrl.Result{}, client.IgnoreNotFound(err)
	// }

	// logger.V(1).Info("Found tracker", "name", podTracker.Spec.Name)

	var podTrackerList crdv1.PodTrackerList

	if err := r.List(ctx, &podTrackerList); err != nil {
		logger.Error(err, "Can't find pod tracker list")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if len(podTrackerList.Items) == 0 {
		logger.V(1).Info("no pod trackers configured")
		return ctrl.Result{}, nil
	} else {
		var podObject v1.Pod
		err := r.Get(context.Background(), req.NamespacedName, &podObject)
		if err != nil {
			return ctrl.Result{}, client.IgnoreNotFound(err)
		}
		logger.V(1).Info("Found reporter configured. Sending report")
		report(podTrackerList.Items[0], podObject)
	}

	return ctrl.Result{}, nil
}

func report(reporter crdv1.PodTracker, pod v1.Pod) {
	// Report to Slack
	log.Log.V(1).Info("Reporting to reporter", "name", reporter.Spec.Name, "endpoint", reporter.Spec.Report.Key)
	slackChannel := reporter.Spec.Report.Channel
	app := slack.New(reporter.Spec.Report.Key, slack.OptionDebug(true))

	message := fmt.Sprintf("New pod created: %s", pod.Name)
	msgText := slack.NewTextBlockObject("mrkdwn", message, false, false)
	msgSection := slack.NewSectionBlock(msgText, nil, nil)

	msg := slack.MsgOptionBlocks(
		msgSection,
	)

	fmt.Print(msg)

	log.Log.V(1).Info("Reporting", "message", "", "channel", slackChannel)
	_, _, _, err := app.SendMessage(slackChannel, msg)

	if err != nil {
		log.Log.V(1).Info(err.Error())
	}
}

func (r *PodTrackerReconciler) HandlePodEvents(ctx context.Context, pod client.Object) []reconcile.Request {

	if pod.GetNamespace() != "default" {
		return []reconcile.Request{}
	}

	namespacedName := types.NamespacedName{
		Namespace: pod.GetNamespace(),
		Name:      pod.GetName(),
	}

	var podObject v1.Pod
	err := r.Get(context.Background(), namespacedName, &podObject)

	if err != nil {
		return []reconcile.Request{}
	}

	if len(podObject.GetAnnotations()) == 0 {
		log.Log.V(1).Info("No annotations set, so this pod is becoming a tracked one now", "pod", podObject.Name)
	} else if podObject.GetAnnotations()["exampleAnnotation"] == "crd.kube.op" {
		log.Log.V(1).Info("Found a managed pod, lets report it", "pod", podObject.Name)
	} else {
		return []reconcile.Request{}
	}

	podObject.SetAnnotations(map[string]string{
		"SourceCRD": "crd.kube.op",
	})

	if err := r.Update(context.TODO(), &podObject); err != nil {
		log.Log.V(1).Info("Error trying to update pod", "err", err)
	}

	requests := []reconcile.Request{
		{
			NamespacedName: namespacedName,
		},
	}

	return requests
}

// SetupWithManager sets up the controller with the Manager.
func (r *PodTrackerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&crdv1.PodTracker{}).
		Watches(
			&v1.Pod{},
			handler.EnqueueRequestsFromMapFunc(r.HandlePodEvents),
			builder.WithPredicates(predicate.ResourceVersionChangedPredicate{}),
		).
		Complete(r)
}
