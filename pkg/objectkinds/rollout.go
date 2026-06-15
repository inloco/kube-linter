package objectkinds

import (
	argoRolloutsV1alpha1 "github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	Rollout = "Rollout"
)

var (
	rolloutGVK = argoRolloutsV1alpha1.SchemeGroupVersion.WithKind("Rollout")
)

func init() {
	RegisterObjectKind(Rollout, MatcherFunc(func(gvk schema.GroupVersionKind) bool {
		return gvk == rolloutGVK
	}))
}
