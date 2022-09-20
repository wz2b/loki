//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureGates) DeepCopyInto(out *FeatureGates) {
	*out = *in
	out.OpenShift = in.OpenShift
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureGates.
func (in *FeatureGates) DeepCopy() *FeatureGates {
	if in == nil {
		return nil
	}
	out := new(FeatureGates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenShiftFeatureGates) DeepCopyInto(out *OpenShiftFeatureGates) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenShiftFeatureGates.
func (in *OpenShiftFeatureGates) DeepCopy() *OpenShiftFeatureGates {
	if in == nil {
		return nil
	}
	out := new(OpenShiftFeatureGates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectConfig) DeepCopyInto(out *ProjectConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ControllerManagerConfigurationSpec.DeepCopyInto(&out.ControllerManagerConfigurationSpec)
	out.Gates = in.Gates
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectConfig.
func (in *ProjectConfig) DeepCopy() *ProjectConfig {
	if in == nil {
		return nil
	}
	out := new(ProjectConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSProfileSpec) DeepCopyInto(out *TLSProfileSpec) {
	*out = *in
	if in.Ciphers != nil {
		in, out := &in.Ciphers, &out.Ciphers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSProfileSpec.
func (in *TLSProfileSpec) DeepCopy() *TLSProfileSpec {
	if in == nil {
		return nil
	}
	out := new(TLSProfileSpec)
	in.DeepCopyInto(out)
	return out
}