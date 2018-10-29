package fakes_test

import (
	"time"

	"github.com/miclip/dotnet-extensions/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("FakeUi", func() {
	Context("Fake UI", func() {
		var (
			bOut *gbytes.Buffer
			bErr *gbytes.Buffer
			bIn  *gbytes.Buffer
		)
		BeforeEach(func() {
			bOut, bErr, bIn = gbytes.NewBuffer(), gbytes.NewBuffer(), gbytes.NewBuffer()
			_, _ = gbytes.TimeoutWriter(bOut, time.Second), gbytes.TimeoutWriter(bOut, time.Second)
			_ = gbytes.TimeoutReader(bOut, time.Second)
		})
		It("creates a fake ui", func() {
			ui := fakes.NewFakeUI(bOut, bErr, bIn)
			Ω(ui).ShouldNot(BeNil())
		})
	})

})
