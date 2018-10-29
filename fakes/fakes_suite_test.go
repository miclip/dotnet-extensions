package fakes_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFakes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fakes Suite")
}
