package basictype

import (
	"context"

	basicv1 "github.com/grepory/gentest/pkg/apis/basic/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apiserver/pkg/registry/rest"
)

// Registry is an interface for things that know how to store BasicTypes.
type Registry interface {
	GetBasicType(ctx context.Context, name string, options *metav1.GetOptions) *basicv1.BasicType
}

type storage struct {
	rest.Getter
}

// NewRegistry returns a new Registry for the given Storage.
func NewRegistry(s rest.StandardStorage) Registry {
	return &storage{s}
}

func (s *storage) GetBasicType(ctx context.Context, name string, options *metav1.GetOptions) *basicv1.BasicType {
	obj, err := s.Get(ctx, name, options)
	if err != nil {
		return nil, err
	}

	return obj.(*basicv1.BasicType), nil
}
