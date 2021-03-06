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
			versions := []nuget.Version{
				nuget.Version{
					ID:        "http://some.com/somefeed/api/v3/registration1/dotnetresource.testlibraryone/1.0.1.json",
					Version:   "1.0.23",
					Downloads: 10,
				},
			}
			nugetClientv3.GetPackageVersionsReturns(versions, nil)
			cmd.PackageName = "A.Package.Name"
			cmd.Prerelease = true
			cmd.Latest = true
			cmd.FeedUrl = "https://somefeed.com/feed.json"
			cmd.HandleVersions(ui, nugetClientv3, nil)
			Ω(nugetClientv3.GetPackageVersionsCallCount()).Should(Equal(1))
			Ω(bOut).To(gbytes.Say("1.0.23"))
		})
		It("returns an error", func() {

			cmd.PackageName = "Wrong.Package.Name"
			cmd.Prerelease = true
			cmd.Latest = true
			cmd.FeedUrl = "https://somefeed.com/feed.json"
			nugetClientv3.GetPackageVersionsReturns([]nuget.Version{}, fmt.Errorf("Package not found name: %s prerelease: %t", cmd.PackageName, cmd.Prerelease))
			cmd.HandleVersions(ui, nugetClientv3, nil)
			Ω(bOut).To(gbytes.Say(""))
			Ω(bErr).To(gbytes.Say("error retrieving latest version from feed. Package not found name: Wrong.Package.Name prerelease: true"))
		})
		It("it can't find a matching package", func() {
			nugetClientv3.GetPackageVersionsReturns([]nuget.Version{}, nil)
			cmd.HandleVersions(ui, nugetClientv3, nil)
			Ω(bOut).To(gbytes.Say(""))
			Ω(nugetClientv3.GetPackageVersionsCallCount()).Should(Equal(1))
			Ω(bErr).To(gbytes.Say("no version details found."))
		})
	})
	Context("versions lists all", func() {
		It("returns all the versions including prerelease", func() {
			versions := []nuget.Version{
				nuget.Version{
					ID:        "http://some.com/somefeed/api/v3/registration1/dotnetresource.testlibraryone/1.0.1.json",
					Version:   "1.0.23",
					Downloads: 10,
				},
				nuget.Version{
					ID:        "http://some.com/somefeed/api/v3/registration1/dotnetresource.testlibraryone/1.0.1.json",
					Version:   "1.0.22",
					Downloads: 10,
				},
			}
			nugetClientv3.GetPackageVersionsReturns(versions, nil)
			cmd.PackageName = "A.Package.Name"
			cmd.Prerelease = true
			cmd.Latest = false
			cmd.FeedUrl = "https://somefeed.com/feed.json"
			cmd.HandleVersions(ui, nugetClientv3, nil)			
			Ω(nugetClientv3.GetPackageVersionsCallCount()).Should(Equal(1))
			Ω(bOut).To(gbytes.Say("1.0.23\n1.0.22\n"))
		})
	})

})
