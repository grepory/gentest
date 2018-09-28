package validation

import (
	basicv1 "github.com/grepory/gentest/pkg/apis/basic/v1"
	metav1validation "k8s.io/apimachinery/pkg/apis/meta/v1/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	
)

func ValidateBasicType(bt *basicv1.BasicType) field.ErrorList {
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, metav1validation.ValidateObjectMeta(&bt.ObjectMeta, true, nil, field.NewPath("metadata"))...)

	if bt.Field == "" {
		allErrs = append(allErrs, &field.Error{
			Type: validation.ErrorTypeRequired,
			Field: "field",
			BadValue: "",
			Detail: "must supply a value for field",
		}
	}

	if len(allErrs) != 0 {
		return allErrs
	}

	return nil
}

func ValidateBasicTypeUpdate(old *basicv1.BasicType, new *basicv1.BasicType) field.ErrorList {
	allErrs := ValidateBasicType(new)
	allErrs = append(allErrs, metav1validation.ValidateObjectMetaUpdate(new, old, field.NewPath("metadata")...))

	return allErrs
}