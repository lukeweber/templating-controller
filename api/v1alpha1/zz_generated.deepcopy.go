// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Behavior) DeepCopyInto(out *Behavior) {
	*out = *in
	out.CRD = in.CRD
	in.EngineConfiguration.DeepCopyInto(&out.EngineConfiguration)
	out.Source = in.Source
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Behavior.
func (in *Behavior) DeepCopy() *Behavior {
	if in == nil {
		return nil
	}
	out := new(Behavior)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EngineConfiguration) DeepCopyInto(out *EngineConfiguration) {
	*out = *in
	if in.KustomizeConfiguration != nil {
		in, out := &in.KustomizeConfiguration, &out.KustomizeConfiguration
		*out = new(KustomizeConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EngineConfiguration.
func (in *EngineConfiguration) DeepCopy() *EngineConfiguration {
	if in == nil {
		return nil
	}
	out := new(EngineConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FieldBinding) DeepCopyInto(out *FieldBinding) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FieldBinding.
func (in *FieldBinding) DeepCopy() *FieldBinding {
	if in == nil {
		return nil
	}
	out := new(FieldBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KustomizeConfiguration) DeepCopyInto(out *KustomizeConfiguration) {
	*out = *in
	if in.Overlays != nil {
		in, out := &in.Overlays, &out.Overlays
		*out = make([]Overlay, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Kustomization != nil {
		in, out := &in.Kustomization, &out.Kustomization
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KustomizeConfiguration.
func (in *KustomizeConfiguration) DeepCopy() *KustomizeConfiguration {
	if in == nil {
		return nil
	}
	out := new(KustomizeConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Overlay) DeepCopyInto(out *Overlay) {
	*out = *in
	out.ObjectReference = in.ObjectReference
	if in.Bindings != nil {
		in, out := &in.Bindings, &out.Bindings
		*out = make([]FieldBinding, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Overlay.
func (in *Overlay) DeepCopy() *Overlay {
	if in == nil {
		return nil
	}
	out := new(Overlay)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Source) DeepCopyInto(out *Source) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Source.
func (in *Source) DeepCopy() *Source {
	if in == nil {
		return nil
	}
	out := new(Source)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateStack) DeepCopyInto(out *TemplateStack) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateStack.
func (in *TemplateStack) DeepCopy() *TemplateStack {
	if in == nil {
		return nil
	}
	out := new(TemplateStack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TemplateStack) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateStackList) DeepCopyInto(out *TemplateStackList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TemplateStack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateStackList.
func (in *TemplateStackList) DeepCopy() *TemplateStackList {
	if in == nil {
		return nil
	}
	out := new(TemplateStackList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TemplateStackList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateStackSpec) DeepCopyInto(out *TemplateStackSpec) {
	*out = *in
	in.AppMetadataSpec.DeepCopyInto(&out.AppMetadataSpec)
	in.Permissions.DeepCopyInto(&out.Permissions)
	in.Behavior.DeepCopyInto(&out.Behavior)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateStackSpec.
func (in *TemplateStackSpec) DeepCopy() *TemplateStackSpec {
	if in == nil {
		return nil
	}
	out := new(TemplateStackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateStackStatus) DeepCopyInto(out *TemplateStackStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateStackStatus.
func (in *TemplateStackStatus) DeepCopy() *TemplateStackStatus {
	if in == nil {
		return nil
	}
	out := new(TemplateStackStatus)
	in.DeepCopyInto(out)
	return out
}
