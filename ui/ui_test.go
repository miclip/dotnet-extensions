package ui_test

import (
	"time"

	"github.com/miclip/dotnet-extensions/fakes"
	"github.com/miclip/dotnet-extensions/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("UI", func() {
	Context("Console UI", func() {
		var (
			bOut *gbytes.Buffer
		)
		BeforeEach(func() {
			bOut, _ = gbytes.NewBuffer(), gbytes.NewBuffer()
			_, _ = gbytes.TimeoutWriter(bOut, time.Second), gbytes.TimeoutWriter(bOut, time.Second)
		})
		It("Creates a console ui", func() {
			ui := ui.NewConsoleUI()
			Ω(ui).ShouldNot(BeNil())
		})
	})
	Context("Console UI", func() {
		var (
			bOut *gbytes.Buffer
			bErr *gbytes.Buffer
			bIn  *gbytes.Buffer
			ui   ui.UI
		)
		BeforeEach(func() {
			bOut, bErr = gbytes.NewBuffer(), gbytes.NewBuffer()
			_, _ = gbytes.TimeoutWriter(bOut, time.Second), gbytes.TimeoutWriter(bOut, time.Second)
			_ = gbytes.TimeoutReader(bIn, time.Second)
			ui = fakes.NewFakeUI(bOut, bErr, bIn)
		})
		It("it writes a print line format", func() {
			Ω(ui).ShouldNot(BeNil())
			ui.PrintLinef("test line %s", "arg")
			Ω(bOut).To(gbytes.Say("test line arg"))
			Ω(bOut).To(gbytes.Say("\n"))
		})
		It("it writes a print format", func() {
			Ω(ui).ShouldNot(BeNil())
			ui.Printf("test line %s", "arg")
			Ω(bOut).To(gbytes.Say("test line arg"))
			Ω(bOut).ToNot(gbytes.Say("\n"))
		})
		It("it writes an error line format", func() {
			Ω(ui).ShouldNot(BeNil())
			ui.ErrorLinef("test line %s", "arg")
			Ω(bErr).To(gbytes.Say("test line arg"))
			Ω(bErr).To(gbytes.Say("\n"))
		})
		It("it writes an error format", func() {
			Ω(ui).ShouldNot(BeNil())
			ui.Errorf("test line %s", "arg")
			Ω(bErr).To(gbytes.Say("test line arg"))
			Ω(bErr).ToNot(gbytes.Say("\n"))
		})
	})

})
