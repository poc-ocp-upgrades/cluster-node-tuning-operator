package e2e

import (
	"testing"
	framework "github.com/operator-framework/operator-sdk/pkg/test"
)

func TestMain(m *testing.M) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	framework.MainEntry(m)
}
