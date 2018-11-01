package dotnet_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDotnet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dotnet Suite")
}
