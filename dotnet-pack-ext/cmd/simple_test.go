package cmd_test

import (
	"time"

	"github.com/miclip/dotnet-extensions/dotnet-pack-ext/cmd"
	"github.com/miclip/dotnet-extensions/fakes"
	"github.com/miclip/dotnet-extensions/nuget"
	"github.com/miclip/dotnet-extensions/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Simple cmd", func() {
	var (
		nugetClientv3 *fakes.FakeNugetClientv3
		dotnetClient *fakes.FakeDotnetClient
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
		dotnetClient = &fakes.FakeDotnetClient{}
	})
	Context("simple pack prepublished with autoincrement", func() {
		It("increments the version from nuget and creates a package", func() {
			versions := []nuget.Version{
				nuget.Version{
					ID:        "http://some.com/somefeed/api/v3/registration1/dotnetresource.testlibraryone/1.0.1.json",
					Version:   "1.0.23",
					Downloads: 10,
				},
			}
			nugetClientv3.GetPackageVersionsReturns(versions, nil)
			nugetClientv3.AutoIncrementVersionReturns("1.0.24", nil)
			nugetClientv3.CreateNuspecFromProjectReturns(nuget.Nuspec{ID: "A.Package", Version:"1.0.0"}, nil)
			dotnetClient.SimplePackReturns("./A.Package.1.0.24.nupkg", nil)
			args := []string{"../../test_files/TestApplication.csproj"}
			cmd.AutoIncrement = true
			cmd.Version = "1.0.*"
			cmd.BasePath = "/tmp/"
			cmd.NoPublish = true
			cmd.FeedUrl = "https://somefeed.com/feed.json"
			cmd.HandleSimplePack(ui, nugetClientv3, dotnetClient, args)
			Ω(bErr).NotTo(gbytes.Say("."))
			Ω(nugetClientv3.GetPackageVersionsCallCount()).Should(Equal(1))
			Ω(nugetClientv3.AutoIncrementVersionCallCount()).Should(Equal(1))
			Ω(dotnetClient.SimplePackCallCount()).Should(Equal(1))			
			Ω(bOut).To(gbytes.Say("Package A.Package Version: 1.0.24\nPackage Path ./A.Package.1.0.24.nupkg"))			
		})		
	})
	Context("simple pack prepublished without autoincrement", func() {
		It("increments the version from nuget and creates a package", func() {
			nugetClientv3.CreateNuspecFromProjectReturns(nuget.Nuspec{ID: "A.Package", Version:"1.0.0"}, nil)
			dotnetClient.SimplePackReturns("./A.Package.1.0.0.nupkg", nil)
			args := []string{"../../test_files/TestApplication.csproj"}
			cmd.AutoIncrement = false
			cmd.Version = "1.0.*"
			cmd.BasePath = "/tmp/"
			cmd.NoPublish = true
			cmd.FeedUrl = "https://somefeed.com/feed.json"
			cmd.HandleSimplePack(ui, nugetClientv3, dotnetClient, args)
			Ω(bErr).NotTo(gbytes.Say("."))
			Ω(nugetClientv3.GetPackageVersionsCallCount()).Should(Equal(0))
			Ω(nugetClientv3.AutoIncrementVersionCallCount()).Should(Equal(0))
			Ω(dotnetClient.SimplePackCallCount()).Should(Equal(1))			
			Ω(bOut).To(gbytes.Say("Package A.Package Version: 1.0.0\nPackage Path ./A.Package.1.0.0.nupkg"))			
		})		
	})
	

})
