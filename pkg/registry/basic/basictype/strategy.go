package basictype

import (
	"context"

	"github.com/grepory/gentest/pkg/api/scheme"
	basicv1 "github.com/grepory/gentest/pkg/apis/basic/v1"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage/names"
)

type strategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

// Strategy is the RESTCreateStrategy for this type.
var Strategy = strategy{scheme.Scheme, names.SimpleNameGenerator}

var _ rest.RESTCreateStrategy = Strategy

func (strategy) NamespaceScoped() bool {
	return true
}

func (strategy) AllowCreateOrUpdate() bool {
	return true
}

func (strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_ = obj.(*basicv1.BasicType)
}

func (strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	bt := obj.(*basicv1.BasicType)
	return validation.ValidateBasicType(bt)
}

func (strategy) Canonicalize(obj runtime.Object) {
	_ = obj.(*basicv1.BasicType)
}

func (strategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	newObj := obj.(*basicv1.BasicType)
	errorList := validation.ValidateBasicType(newObj)
	return append(errorList, validation.ValidateBasicTypeUpdate(newObj, old.(*basicv1.BasicType))...)
}
