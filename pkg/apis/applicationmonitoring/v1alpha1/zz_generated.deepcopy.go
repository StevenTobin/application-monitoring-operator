// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationMonitoring) DeepCopyInto(out *ApplicationMonitoring) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationMonitoring.
func (in *ApplicationMonitoring) DeepCopy() *ApplicationMonitoring {
	if in == nil {
		return nil
	}
	out := new(ApplicationMonitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationMonitoring) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationMonitoringList) DeepCopyInto(out *ApplicationMonitoringList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApplicationMonitoring, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationMonitoringList.
func (in *ApplicationMonitoringList) DeepCopy() *ApplicationMonitoringList {
	if in == nil {
		return nil
	}
	out := new(ApplicationMonitoringList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationMonitoringList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationMonitoringSpec) DeepCopyInto(out *ApplicationMonitoringSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationMonitoringSpec.
func (in *ApplicationMonitoringSpec) DeepCopy() *ApplicationMonitoringSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationMonitoringSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationMonitoringStatus) DeepCopyInto(out *ApplicationMonitoringStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationMonitoringStatus.
func (in *ApplicationMonitoringStatus) DeepCopy() *ApplicationMonitoringStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationMonitoringStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlackboxTarget) DeepCopyInto(out *BlackboxTarget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlackboxTarget.
func (in *BlackboxTarget) DeepCopy() *BlackboxTarget {
	if in == nil {
		return nil
	}
	out := new(BlackboxTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BlackboxTarget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlackboxTargetList) DeepCopyInto(out *BlackboxTargetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BlackboxTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlackboxTargetList.
func (in *BlackboxTargetList) DeepCopy() *BlackboxTargetList {
	if in == nil {
		return nil
	}
	out := new(BlackboxTargetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BlackboxTargetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlackboxTargetSpec) DeepCopyInto(out *BlackboxTargetSpec) {
	*out = *in
	if in.BlackboxTargets != nil {
		in, out := &in.BlackboxTargets, &out.BlackboxTargets
		*out = make([]BlackboxtargetData, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlackboxTargetSpec.
func (in *BlackboxTargetSpec) DeepCopy() *BlackboxTargetSpec {
	if in == nil {
		return nil
	}
	out := new(BlackboxTargetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlackboxTargetStatus) DeepCopyInto(out *BlackboxTargetStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlackboxTargetStatus.
func (in *BlackboxTargetStatus) DeepCopy() *BlackboxTargetStatus {
	if in == nil {
		return nil
	}
	out := new(BlackboxTargetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlackboxtargetData) DeepCopyInto(out *BlackboxtargetData) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlackboxtargetData.
func (in *BlackboxtargetData) DeepCopy() *BlackboxtargetData {
	if in == nil {
		return nil
	}
	out := new(BlackboxtargetData)
	in.DeepCopyInto(out)
	return out
}
