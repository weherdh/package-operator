//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObjectDeployment) DeepCopyInto(out *ClusterObjectDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObjectDeployment.
func (in *ClusterObjectDeployment) DeepCopy() *ClusterObjectDeployment {
	if in == nil {
		return nil
	}
	out := new(ClusterObjectDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterObjectDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObjectDeploymentList) DeepCopyInto(out *ClusterObjectDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterObjectDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObjectDeploymentList.
func (in *ClusterObjectDeploymentList) DeepCopy() *ClusterObjectDeploymentList {
	if in == nil {
		return nil
	}
	out := new(ClusterObjectDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterObjectDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObjectDeploymentSpec) DeepCopyInto(out *ClusterObjectDeploymentSpec) {
	*out = *in
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		*out = new(int32)
		**out = **in
	}
	in.Selector.DeepCopyInto(&out.Selector)
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObjectDeploymentSpec.
func (in *ClusterObjectDeploymentSpec) DeepCopy() *ClusterObjectDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterObjectDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObjectDeploymentStatus) DeepCopyInto(out *ClusterObjectDeploymentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CollisionCount != nil {
		in, out := &in.CollisionCount, &out.CollisionCount
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObjectDeploymentStatus.
func (in *ClusterObjectDeploymentStatus) DeepCopy() *ClusterObjectDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterObjectDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObjectSet) DeepCopyInto(out *ClusterObjectSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObjectSet.
func (in *ClusterObjectSet) DeepCopy() *ClusterObjectSet {
	if in == nil {
		return nil
	}
	out := new(ClusterObjectSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterObjectSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObjectSetList) DeepCopyInto(out *ClusterObjectSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterObjectSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObjectSetList.
func (in *ClusterObjectSetList) DeepCopy() *ClusterObjectSetList {
	if in == nil {
		return nil
	}
	out := new(ClusterObjectSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterObjectSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObjectSetPhase) DeepCopyInto(out *ClusterObjectSetPhase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObjectSetPhase.
func (in *ClusterObjectSetPhase) DeepCopy() *ClusterObjectSetPhase {
	if in == nil {
		return nil
	}
	out := new(ClusterObjectSetPhase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterObjectSetPhase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObjectSetPhaseList) DeepCopyInto(out *ClusterObjectSetPhaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterObjectSetPhase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObjectSetPhaseList.
func (in *ClusterObjectSetPhaseList) DeepCopy() *ClusterObjectSetPhaseList {
	if in == nil {
		return nil
	}
	out := new(ClusterObjectSetPhaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterObjectSetPhaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObjectSetPhaseSpec) DeepCopyInto(out *ClusterObjectSetPhaseSpec) {
	*out = *in
	if in.Previous != nil {
		in, out := &in.Previous, &out.Previous
		*out = make([]PreviousRevisionReference, len(*in))
		copy(*out, *in)
	}
	if in.AvailabilityProbes != nil {
		in, out := &in.AvailabilityProbes, &out.AvailabilityProbes
		*out = make([]ObjectSetProbe, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]ObjectSetObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObjectSetPhaseSpec.
func (in *ClusterObjectSetPhaseSpec) DeepCopy() *ClusterObjectSetPhaseSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterObjectSetPhaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObjectSetPhaseStatus) DeepCopyInto(out *ClusterObjectSetPhaseStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ControllerOf != nil {
		in, out := &in.ControllerOf, &out.ControllerOf
		*out = make([]ControlledObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObjectSetPhaseStatus.
func (in *ClusterObjectSetPhaseStatus) DeepCopy() *ClusterObjectSetPhaseStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterObjectSetPhaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObjectSetSpec) DeepCopyInto(out *ClusterObjectSetSpec) {
	*out = *in
	if in.Previous != nil {
		in, out := &in.Previous, &out.Previous
		*out = make([]PreviousRevisionReference, len(*in))
		copy(*out, *in)
	}
	in.ObjectSetTemplateSpec.DeepCopyInto(&out.ObjectSetTemplateSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObjectSetSpec.
func (in *ClusterObjectSetSpec) DeepCopy() *ClusterObjectSetSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterObjectSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObjectSetStatus) DeepCopyInto(out *ClusterObjectSetStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RemotePhases != nil {
		in, out := &in.RemotePhases, &out.RemotePhases
		*out = make([]RemotePhaseReference, len(*in))
		copy(*out, *in)
	}
	if in.ControllerOf != nil {
		in, out := &in.ControllerOf, &out.ControllerOf
		*out = make([]ControlledObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObjectSetStatus.
func (in *ClusterObjectSetStatus) DeepCopy() *ClusterObjectSetStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterObjectSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObjectSlice) DeepCopyInto(out *ClusterObjectSlice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]ObjectSetObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObjectSlice.
func (in *ClusterObjectSlice) DeepCopy() *ClusterObjectSlice {
	if in == nil {
		return nil
	}
	out := new(ClusterObjectSlice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterObjectSlice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObjectSliceList) DeepCopyInto(out *ClusterObjectSliceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterObjectSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObjectSliceList.
func (in *ClusterObjectSliceList) DeepCopy() *ClusterObjectSliceList {
	if in == nil {
		return nil
	}
	out := new(ClusterObjectSliceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterObjectSliceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPackage) DeepCopyInto(out *ClusterPackage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPackage.
func (in *ClusterPackage) DeepCopy() *ClusterPackage {
	if in == nil {
		return nil
	}
	out := new(ClusterPackage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterPackage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPackageList) DeepCopyInto(out *ClusterPackageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterPackage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPackageList.
func (in *ClusterPackageList) DeepCopy() *ClusterPackageList {
	if in == nil {
		return nil
	}
	out := new(ClusterPackageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterPackageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlledObjectReference) DeepCopyInto(out *ControlledObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlledObjectReference.
func (in *ControlledObjectReference) DeepCopy() *ControlledObjectReference {
	if in == nil {
		return nil
	}
	out := new(ControlledObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectDeployment) DeepCopyInto(out *ObjectDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectDeployment.
func (in *ObjectDeployment) DeepCopy() *ObjectDeployment {
	if in == nil {
		return nil
	}
	out := new(ObjectDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectDeploymentList) DeepCopyInto(out *ObjectDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ObjectDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectDeploymentList.
func (in *ObjectDeploymentList) DeepCopy() *ObjectDeploymentList {
	if in == nil {
		return nil
	}
	out := new(ObjectDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectDeploymentSpec) DeepCopyInto(out *ObjectDeploymentSpec) {
	*out = *in
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		*out = new(int32)
		**out = **in
	}
	in.Selector.DeepCopyInto(&out.Selector)
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectDeploymentSpec.
func (in *ObjectDeploymentSpec) DeepCopy() *ObjectDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectDeploymentStatus) DeepCopyInto(out *ObjectDeploymentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CollisionCount != nil {
		in, out := &in.CollisionCount, &out.CollisionCount
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectDeploymentStatus.
func (in *ObjectDeploymentStatus) DeepCopy() *ObjectDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(ObjectDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSet) DeepCopyInto(out *ObjectSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSet.
func (in *ObjectSet) DeepCopy() *ObjectSet {
	if in == nil {
		return nil
	}
	out := new(ObjectSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSetList) DeepCopyInto(out *ObjectSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ObjectSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSetList.
func (in *ObjectSetList) DeepCopy() *ObjectSetList {
	if in == nil {
		return nil
	}
	out := new(ObjectSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSetObject) DeepCopyInto(out *ObjectSetObject) {
	*out = *in
	in.Object.DeepCopyInto(&out.Object)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSetObject.
func (in *ObjectSetObject) DeepCopy() *ObjectSetObject {
	if in == nil {
		return nil
	}
	out := new(ObjectSetObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSetPhase) DeepCopyInto(out *ObjectSetPhase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSetPhase.
func (in *ObjectSetPhase) DeepCopy() *ObjectSetPhase {
	if in == nil {
		return nil
	}
	out := new(ObjectSetPhase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectSetPhase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSetPhaseList) DeepCopyInto(out *ObjectSetPhaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ObjectSetPhase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSetPhaseList.
func (in *ObjectSetPhaseList) DeepCopy() *ObjectSetPhaseList {
	if in == nil {
		return nil
	}
	out := new(ObjectSetPhaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectSetPhaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSetPhaseSpec) DeepCopyInto(out *ObjectSetPhaseSpec) {
	*out = *in
	if in.Previous != nil {
		in, out := &in.Previous, &out.Previous
		*out = make([]PreviousRevisionReference, len(*in))
		copy(*out, *in)
	}
	if in.AvailabilityProbes != nil {
		in, out := &in.AvailabilityProbes, &out.AvailabilityProbes
		*out = make([]ObjectSetProbe, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]ObjectSetObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSetPhaseSpec.
func (in *ObjectSetPhaseSpec) DeepCopy() *ObjectSetPhaseSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectSetPhaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSetPhaseStatus) DeepCopyInto(out *ObjectSetPhaseStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ControllerOf != nil {
		in, out := &in.ControllerOf, &out.ControllerOf
		*out = make([]ControlledObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSetPhaseStatus.
func (in *ObjectSetPhaseStatus) DeepCopy() *ObjectSetPhaseStatus {
	if in == nil {
		return nil
	}
	out := new(ObjectSetPhaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSetProbe) DeepCopyInto(out *ObjectSetProbe) {
	*out = *in
	if in.Probes != nil {
		in, out := &in.Probes, &out.Probes
		*out = make([]Probe, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Selector.DeepCopyInto(&out.Selector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSetProbe.
func (in *ObjectSetProbe) DeepCopy() *ObjectSetProbe {
	if in == nil {
		return nil
	}
	out := new(ObjectSetProbe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSetSpec) DeepCopyInto(out *ObjectSetSpec) {
	*out = *in
	if in.Previous != nil {
		in, out := &in.Previous, &out.Previous
		*out = make([]PreviousRevisionReference, len(*in))
		copy(*out, *in)
	}
	in.ObjectSetTemplateSpec.DeepCopyInto(&out.ObjectSetTemplateSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSetSpec.
func (in *ObjectSetSpec) DeepCopy() *ObjectSetSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSetStatus) DeepCopyInto(out *ObjectSetStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RemotePhases != nil {
		in, out := &in.RemotePhases, &out.RemotePhases
		*out = make([]RemotePhaseReference, len(*in))
		copy(*out, *in)
	}
	if in.ControllerOf != nil {
		in, out := &in.ControllerOf, &out.ControllerOf
		*out = make([]ControlledObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSetStatus.
func (in *ObjectSetStatus) DeepCopy() *ObjectSetStatus {
	if in == nil {
		return nil
	}
	out := new(ObjectSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSetTemplate) DeepCopyInto(out *ObjectSetTemplate) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSetTemplate.
func (in *ObjectSetTemplate) DeepCopy() *ObjectSetTemplate {
	if in == nil {
		return nil
	}
	out := new(ObjectSetTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSetTemplatePhase) DeepCopyInto(out *ObjectSetTemplatePhase) {
	*out = *in
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]ObjectSetObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Slices != nil {
		in, out := &in.Slices, &out.Slices
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSetTemplatePhase.
func (in *ObjectSetTemplatePhase) DeepCopy() *ObjectSetTemplatePhase {
	if in == nil {
		return nil
	}
	out := new(ObjectSetTemplatePhase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSetTemplateSpec) DeepCopyInto(out *ObjectSetTemplateSpec) {
	*out = *in
	if in.Phases != nil {
		in, out := &in.Phases, &out.Phases
		*out = make([]ObjectSetTemplatePhase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AvailabilityProbes != nil {
		in, out := &in.AvailabilityProbes, &out.AvailabilityProbes
		*out = make([]ObjectSetProbe, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSetTemplateSpec.
func (in *ObjectSetTemplateSpec) DeepCopy() *ObjectSetTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectSetTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSlice) DeepCopyInto(out *ObjectSlice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]ObjectSetObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSlice.
func (in *ObjectSlice) DeepCopy() *ObjectSlice {
	if in == nil {
		return nil
	}
	out := new(ObjectSlice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectSlice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSliceList) DeepCopyInto(out *ObjectSliceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ObjectSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSliceList.
func (in *ObjectSliceList) DeepCopy() *ObjectSliceList {
	if in == nil {
		return nil
	}
	out := new(ObjectSliceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectSliceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Package) DeepCopyInto(out *Package) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Package.
func (in *Package) DeepCopy() *Package {
	if in == nil {
		return nil
	}
	out := new(Package)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Package) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageList) DeepCopyInto(out *PackageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Package, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageList.
func (in *PackageList) DeepCopy() *PackageList {
	if in == nil {
		return nil
	}
	out := new(PackageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageProbeKindSpec) DeepCopyInto(out *PackageProbeKindSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageProbeKindSpec.
func (in *PackageProbeKindSpec) DeepCopy() *PackageProbeKindSpec {
	if in == nil {
		return nil
	}
	out := new(PackageProbeKindSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageSpec) DeepCopyInto(out *PackageSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageSpec.
func (in *PackageSpec) DeepCopy() *PackageSpec {
	if in == nil {
		return nil
	}
	out := new(PackageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageStatus) DeepCopyInto(out *PackageStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageStatus.
func (in *PackageStatus) DeepCopy() *PackageStatus {
	if in == nil {
		return nil
	}
	out := new(PackageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreviousRevisionReference) DeepCopyInto(out *PreviousRevisionReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreviousRevisionReference.
func (in *PreviousRevisionReference) DeepCopy() *PreviousRevisionReference {
	if in == nil {
		return nil
	}
	out := new(PreviousRevisionReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Probe) DeepCopyInto(out *Probe) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(ProbeConditionSpec)
		**out = **in
	}
	if in.FieldsEqual != nil {
		in, out := &in.FieldsEqual, &out.FieldsEqual
		*out = new(ProbeFieldsEqualSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Probe.
func (in *Probe) DeepCopy() *Probe {
	if in == nil {
		return nil
	}
	out := new(Probe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProbeConditionSpec) DeepCopyInto(out *ProbeConditionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProbeConditionSpec.
func (in *ProbeConditionSpec) DeepCopy() *ProbeConditionSpec {
	if in == nil {
		return nil
	}
	out := new(ProbeConditionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProbeFieldsEqualSpec) DeepCopyInto(out *ProbeFieldsEqualSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProbeFieldsEqualSpec.
func (in *ProbeFieldsEqualSpec) DeepCopy() *ProbeFieldsEqualSpec {
	if in == nil {
		return nil
	}
	out := new(ProbeFieldsEqualSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProbeSelector) DeepCopyInto(out *ProbeSelector) {
	*out = *in
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(PackageProbeKindSpec)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProbeSelector.
func (in *ProbeSelector) DeepCopy() *ProbeSelector {
	if in == nil {
		return nil
	}
	out := new(ProbeSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemotePhaseReference) DeepCopyInto(out *RemotePhaseReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemotePhaseReference.
func (in *RemotePhaseReference) DeepCopy() *RemotePhaseReference {
	if in == nil {
		return nil
	}
	out := new(RemotePhaseReference)
	in.DeepCopyInto(out)
	return out
}
