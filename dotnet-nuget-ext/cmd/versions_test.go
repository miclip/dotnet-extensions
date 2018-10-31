package cmd_test

import (
	"fmt"
	"time"

	"github.com/miclip/dotnet-extensions/dotnet-nuget-ext/cmd"
	"github.com/miclip/dotnet-extensions/fakes"
	"github.com/miclip/dotnet-extensions/nuget"
	"github.com/miclip/dotnet-extensions/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Versions cmd", func() {
	var (
		nugetClientv3 *fakes.FakeNugetClientv3
		bOut          *gbytes.Buffer
		bErr          *gbytes.Buffer
		bIn           *gbytes.Buffer
		ui            ui.UI
	)
	BeforeEach(func() {
		bOut, bErr, bIn = gbytes.NewBuffer(), gbytes.NewBuffer(), gbytes.NewBuffer()
		_, _ = gbytes.TimeoutWriter(bOut, time.Second), gbytes.TimeoutWriter(bOut, time.Second)
		_ = gbytes.TimeoutReader(bIn, time.Second)
		ui = fakes.NewFakeUI(bOut, bErr, bIn)
		nugetClientv3 = &fakes.FakeNugetClientv3{}
	})
	Context("versions latest", func() {
		It("prints the latest prelease version", func() {
			packageVersion := nuget.PackageVersion{
				ID:          "A.Package.Name",
				Version:     "1.0.23",
				Description: "A descriptive description",
			}
			nugetClientv3.GetPackageVersionReturns(&packageVersion, nil)
			cmd.PackageName = "A.Package.Name"
			cmd.Prerelease = true
			cmd.Latest = true
			cmd.FeedUrl = "https://somefeed.com/feed.json"
			cmd.HandleVersions(ui, nugetClientv3, nil)
			Ω(bOut).To(gbytes.Say("1.0.23"))
		})
		It("returns an error", func() {
			
			cmd.PackageName = "Wrong.Package.Name"
			cmd.Prerelease = true
			cmd.Latest = true
			cmd.FeedUrl = "https://somefeed.com/feed.json"
			nugetClientv3.GetPackageVersionReturns(nil, fmt.Errorf("Package not found name: %s prerelease: %t", cmd.PackageName, cmd.Prerelease))
			cmd.HandleVersions(ui, nugetClientv3, nil)
			Ω(bOut).To(gbytes.Say(""))
			Ω(bErr).To(gbytes.Say("error retrieving latest version from feed. Package not found name: Wrong.Package.Name prerelease: true"))
		})
		It("it can't find a matching package", func() {
			nugetClientv3.GetPackageVersionReturns(nil, nil)
			cmd.HandleVersions(ui, nugetClientv3, nil)
			Ω(bOut).To(gbytes.Say(""))
			Ω(bErr).To(gbytes.Say("no version details found."))
		})
	})

})
