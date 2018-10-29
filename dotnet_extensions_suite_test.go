package dotnet_extensions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDotnetExtensions(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DotnetExtensions Suite")
}
