/*
Copyright 2019 The Crossplane Authors.

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

package resource

import (
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/crossplaneio/crossplane-runtime/apis/core/v1alpha1"
	"github.com/crossplaneio/crossplane-runtime/pkg/meta"
)

const RemoveDefaultAnnotationsKey = "templatestacks.crossplane.io/remove-defaulting-annotations"
const RemoveDefaultAnnotationsTrueValue = "true"

// NewOwnerReferenceAdder returns a new *OwnerReferenceAdder
func NewOwnerReferenceAdder() OwnerReferenceAdder {
	return OwnerReferenceAdder{}
}

// OwnerReferenceAdder adds owner reference of ParentResource to all ChildResources
// except the Providers since their deletion should be delayed until all resources
// refer to them are deleted.
type OwnerReferenceAdder struct{}

func (lo OwnerReferenceAdder) Patch(cr ParentResource, list []ChildResource) ([]ChildResource, error) {
	ref := metav1.OwnerReference{
		APIVersion: cr.GetObjectKind().GroupVersionKind().GroupVersion().String(),
		Kind:       cr.GetObjectKind().GroupVersionKind().Kind,
		Name:       cr.GetName(),
		UID:        cr.GetUID(),
	}
	for _, o := range list {
		// TODO(muvaf): Provider kind resources are special in the sense that
		// their deletion should be blocked until all resources provisioned with
		// them are deleted. Since we let Kubernetes garbage collector clean the
		// resources, we skip deletion of Provider kind resources for deletions
		// to success.
		// Find a way to realize that dependency without bringing in too much
		// complexity.
		if isProvider(o) {
			continue
		}
		meta.AddOwnerReference(o, ref)
	}
	return list, nil
}

// NewDefaultingAnnotationRemover returns a new DefaultingAnnotationRemover
func NewDefaultingAnnotationRemover() DefaultingAnnotationRemover {
	return DefaultingAnnotationRemover{}
}

// DefaultingAnnotationRemover removes the defaulting annotation on the resources
// if not explicitly specified otherwise.
type DefaultingAnnotationRemover struct{}

func (lo DefaultingAnnotationRemover) Patch(cr ParentResource, list []ChildResource) ([]ChildResource, error) {
	if cr.GetAnnotations()[RemoveDefaultAnnotationsKey] != RemoveDefaultAnnotationsTrueValue {
		return list, nil
	}
	for _, o := range list {
		meta.RemoveAnnotations(o, v1alpha1.AnnotationDefaultClassKey)
	}
	return list, nil
}

// NewNamespacePatcher returns a new NamespacePatcher
func NewNamespacePatcher() NamespacePatcher {
	return NamespacePatcher{}
}

// NamespacePatcher patches the child resources whose metadata.namespace is empty
// with namespace of the parent resource.
type NamespacePatcher struct{}

func (lo NamespacePatcher) Patch(cr ParentResource, list []ChildResource) ([]ChildResource, error) {
	if cr.GetNamespace() == "" {
		return list, nil
	}
	for _, o := range list {
		if o.GetNamespace() == "" {
			o.SetNamespace(cr.GetNamespace())
		}
	}
	return list, nil
}

// todo: temp solution to detect provider kind.
func isProvider(o runtime.Object) bool {
	gvk := o.GetObjectKind().GroupVersionKind()
	return strings.HasSuffix(gvk.Group, "crossplane.io") && strings.ToLower(gvk.Kind) == "provider"
}
