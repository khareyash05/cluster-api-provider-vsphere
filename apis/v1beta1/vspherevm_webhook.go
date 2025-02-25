/*
Copyright 2021 The Kubernetes Authors.

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

package v1beta1

import (
	"fmt"
	"net"
	"reflect"

	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *VSphereVM) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// +kubebuilder:webhook:verbs=create;update,path=/validate-infrastructure-cluster-x-k8s-io-v1beta1-vspherevm,mutating=false,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=vspherevms,versions=v1beta1,name=validation.vspherevm.infrastructure.x-k8s.io,sideEffects=None,admissionReviewVersions=v1beta1
// +kubebuilder:webhook:verbs=create;update,path=/mutate-infrastructure-cluster-x-k8s-io-v1beta1-vspherevm,mutating=true,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=vspherevms,versions=v1beta1,name=default.vspherevm.infrastructure.x-k8s.io,sideEffects=None,admissionReviewVersions=v1beta1

// Default implements webhook.Defaulter so a webhook will be registered for the type.
func (r *VSphereVM) Default() {
	// Set Linux as default OS value
	if r.Spec.OS == "" {
		r.Spec.OS = Linux
	}
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type.
func (r *VSphereVM) ValidateCreate() error {
	var allErrs field.ErrorList
	spec := r.Spec

	if spec.Network.PreferredAPIServerCIDR != "" {
		allErrs = append(allErrs, field.Invalid(field.NewPath("spec", "PreferredAPIServerCIDR"), spec.Network.PreferredAPIServerCIDR, "cannot be set, as it will be removed and is no longer used"))
	}

	for i, device := range spec.Network.Devices {
		for j, ip := range device.IPAddrs {
			if _, _, err := net.ParseCIDR(ip); err != nil {
				allErrs = append(allErrs, field.Invalid(field.NewPath("spec", "network", fmt.Sprintf("devices[%d]", i), fmt.Sprintf("ipAddrs[%d]", j)), ip, "ip addresses should be in the CIDR format"))
			}
		}
	}

	if r.Spec.OS == Windows && len(r.Name) > 15 {
		allErrs = append(allErrs, field.Invalid(field.NewPath("name"), r.Name, "name has to be less than 16 characters for Windows VM"))
	}
	return aggregateObjErrors(r.GroupVersionKind().GroupKind(), r.Name, allErrs)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type.
//
//nolint:forcetypeassert
func (r *VSphereVM) ValidateUpdate(old runtime.Object) error {
	newVSphereVM, err := runtime.DefaultUnstructuredConverter.ToUnstructured(r)
	if err != nil {
		return apierrors.NewInternalError(errors.Wrap(err, "failed to convert new VSphereVM to unstructured object"))
	}
	oldVSphereVM, err := runtime.DefaultUnstructuredConverter.ToUnstructured(old)
	if err != nil {
		return apierrors.NewInternalError(errors.Wrap(err, "failed to convert old VSphereVM to unstructured object"))
	}

	var allErrs field.ErrorList

	newVSphereVMSpec := newVSphereVM["spec"].(map[string]interface{})
	oldVSphereVMSpec := oldVSphereVM["spec"].(map[string]interface{})

	// allow changes to biosUUID, bootstrapRef, thumbprint
	keys := []string{"biosUUID", "bootstrapRef", "thumbprint"}
	// allow changes to os only if the old spec has empty OS field
	if _, ok := oldVSphereVMSpec["os"]; !ok {
		keys = append(keys, "os")
	}
	r.deleteSpecKeys(oldVSphereVMSpec, keys)
	r.deleteSpecKeys(newVSphereVMSpec, keys)

	newVSphereVMNetwork := newVSphereVMSpec["network"].(map[string]interface{})
	oldVSphereVMNetwork := oldVSphereVMSpec["network"].(map[string]interface{})

	// allow changes to the network devices
	networkKeys := []string{"devices"}
	r.deleteSpecKeys(oldVSphereVMNetwork, networkKeys)
	r.deleteSpecKeys(newVSphereVMNetwork, networkKeys)

	if !reflect.DeepEqual(oldVSphereVMSpec, newVSphereVMSpec) {
		allErrs = append(allErrs, field.Forbidden(field.NewPath("spec"), "cannot be modified"))
	}

	return aggregateObjErrors(r.GroupVersionKind().GroupKind(), r.Name, allErrs)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type.
func (r *VSphereVM) ValidateDelete() error {
	return nil
}

func (r *VSphereVM) deleteSpecKeys(spec map[string]interface{}, keys []string) {
	if len(spec) == 0 || len(keys) == 0 {
		return
	}

	for _, key := range keys {
		delete(spec, key)
	}
}
