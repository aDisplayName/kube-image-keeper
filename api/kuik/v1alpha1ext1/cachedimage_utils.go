package v1alpha1ext1

import (
	"context"

	"github.com/adisplayname/kube-image-keeper/internal/registry"
	"github.com/distribution/reference"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func CachedImageNameFromSourceImage(sourceImage string) (string, error) {
	named, err := reference.ParseNormalizedNamed(sourceImage)
	if err != nil {
		return "", err
	}

	tag := "latest"
	if tagged, ok := named.(reference.Tagged); ok {
		tag = tagged.Tag()
	} else if digested, ok := named.(reference.Digested); ok {
		tag = digested.Digest().String()
	}

	return registry.SanitizeName(named.Name() + "-" + tag), nil
}

func (r *CachedImage) Repository() (reference.Named, error) {
	named, err := reference.ParseNormalizedNamed(r.Spec.SourceImage)
	if err != nil {
		return nil, err
	}

	return named, nil
}

func (r *CachedImage) GetPullSecrets(apiReader client.Reader) ([]corev1.Secret, error) {
	named, err := r.Repository()
	if err != nil {
		return nil, err
	}

	repository := Repository{}
	err = apiReader.Get(context.TODO(), types.NamespacedName{Name: registry.SanitizeName(named.Name())}, &repository)
	if err != nil && !apierrors.IsNotFound(err) {
		return nil, err
	}

	pullSecrets, err := repository.GetPullSecrets(apiReader)
	if err != nil {
		return nil, err
	}

	return pullSecrets, nil
}

func (r *CachedImage) Upstream() (string, error) {
	named, err := r.Repository()
	if err != nil {
		return "", err
	}

	return reference.Domain(named), nil
}
