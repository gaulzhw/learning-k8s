package webhook

import (
	"context"
	"net/http"

	admissionv1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// ValidatingAdmission validates Migrate object when creating/updating/deleting.
type ValidatingAdmission struct {
	decoder *admission.Decoder
}

// Check if our ValidatingAdmission implements necessary interface
var _ admission.Handler = &ValidatingAdmission{}
var _ admission.DecoderInjector = &ValidatingAdmission{}

// Handle implements admission.Handler interface.
// It yields a response to an AdmissionRequest.
func (v *ValidatingAdmission) Handle(ctx context.Context, req admission.Request) admission.Response {
	klog.V(2).InfoS("validating admission", "operation", req.Operation, "namespace", req.Namespace, "name", req.Name)

	var allErrors field.ErrorList
	var err error

	switch req.Operation {
	case admissionv1.Create:
		allErrors, err = v.validCreate(ctx, req)
	case admissionv1.Update:
	case admissionv1.Delete:
		allErrors, err = v.validDelete(ctx, req)
	}

	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	if len(allErrors) != 0 {
		klog.Error(allErrors.ToAggregate())
		return admission.Denied(allErrors.ToAggregate().Error())
	}

	return admission.Allowed("")
}

// InjectDecoder implements admission.DecoderInjector interface.
// A decoder will be automatically injected.
func (v *ValidatingAdmission) InjectDecoder(d *admission.Decoder) error {
	v.decoder = d
	return nil
}

func (v *ValidatingAdmission) validCreate(ctx context.Context, req admission.Request) (field.ErrorList, error) {
	var allErrors field.ErrorList
	return allErrors, nil
}

func (v *ValidatingAdmission) validDelete(ctx context.Context, req admission.Request) (field.ErrorList, error) {
	var allErrors field.ErrorList
	newPath := field.NewPath("migrate")

	if len(req.Options.Raw) == 0 {
		allErrors = append(allErrors, field.Required(newPath.Child("metadata").Child("deletionGracePeriodSeconds"), "delete options should not empty"))
		return allErrors, nil
	}

	options := &metav1.DeleteOptions{}
	err := v.decoder.DecodeRaw(req.Options, options)
	if err != nil {
		return allErrors, err
	}
	if options.GracePeriodSeconds == nil {
		allErrors = append(allErrors, field.Required(newPath.Child("metadata").Child("deletionGracePeriodSeconds"), "only force delete supported"))
	}
	return allErrors, nil
}
