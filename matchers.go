package gkube

import (
	gtypes "github.com/onsi/gomega/types"
	"github.com/testernetes/gkube/matchers"
	"k8s.io/apimachinery/pkg/runtime"
)

type IgnorePaths = matchers.IgnorePaths
type AllowPaths = matchers.AllowPaths

var IgnoreAutogeneratedMetadata = matchers.IgnoreAutogeneratedMetadata

func EqualObject(original runtime.Object, opts ...matchers.EqualObjectMatchOption) gtypes.GomegaMatcher {
	return matchers.NewEqualObjectMatcher(original, opts...)
}

func HaveJSONPath(path string, matcher gtypes.GomegaMatcher) gtypes.GomegaMatcher {
	return matchers.NewHaveJSONPathMatcher(path, matcher)
}
