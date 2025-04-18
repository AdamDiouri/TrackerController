//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodTracker) DeepCopyInto(out *PodTracker) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodTracker.
func (in *PodTracker) DeepCopy() *PodTracker {
	if in == nil {
		return nil
	}
	out := new(PodTracker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodTracker) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodTrackerList) DeepCopyInto(out *PodTrackerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodTracker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodTrackerList.
func (in *PodTrackerList) DeepCopy() *PodTrackerList {
	if in == nil {
		return nil
	}
	out := new(PodTrackerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodTrackerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodTrackerSpec) DeepCopyInto(out *PodTrackerSpec) {
	*out = *in
	out.Report = in.Report
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodTrackerSpec.
func (in *PodTrackerSpec) DeepCopy() *PodTrackerSpec {
	if in == nil {
		return nil
	}
	out := new(PodTrackerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodTrackerStatus) DeepCopyInto(out *PodTrackerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodTrackerStatus.
func (in *PodTrackerStatus) DeepCopy() *PodTrackerStatus {
	if in == nil {
		return nil
	}
	out := new(PodTrackerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Reporter) DeepCopyInto(out *Reporter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Reporter.
func (in *Reporter) DeepCopy() *Reporter {
	if in == nil {
		return nil
	}
	out := new(Reporter)
	in.DeepCopyInto(out)
	return out
}
