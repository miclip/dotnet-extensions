package dotnet_test


import (
	"github.com/miclip/dotnet-extensions/fakes"
	
	"os/exec"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/miclip/dotnet-extensions/dotnet"
)

var _ = Describe("dotnetclient", func() {
	BeforeEach(func(){
		dotnet.ExecCommand = fakes.FakeExecCommand
	})

	It("should execute dotnet publish command netcoreapp and ubuntu", func() {
		defer func() { dotnet.ExecCommand = exec.Command }()
		fakes.MockedExitStatus = 0
		fakes.MockedStdout = ""
		expectedCommand := "dotnet publish /path/project.csproj -o /tmp/publish --no-build --no-restore -f netcoreapp2.1 -r ubuntu.14.04-x64"		
		client := dotnet.NewDotnetClient("netcoreapp2.1","ubuntu.14.04-x64")
		_, err := client.Publish("/path/project.csproj","/tmp/publish", true, true)
		Ω(fakes.CommandString).Should(Equal(expectedCommand))
		Ω(err).ShouldNot(HaveOccurred())		
	})
	It("should execute dotnet publish command netcoreapp and windows", func() {
		defer func() { dotnet.ExecCommand = exec.Command }()
		fakes.MockedExitStatus = 0
		fakes.MockedStdout = ""
		expectedCommand := "dotnet publish /path/project.csproj -o /tmp/publish --no-build --no-restore -f netcoreapp2.1 -r win10-x64"		
		client := dotnet.NewDotnetClient("netcoreapp2.1","win10-x64")
		_, err := client.Publish("/path/project.csproj","/tmp/publish", true, true)
		Ω(fakes.CommandString).Should(Equal(expectedCommand))
		Ω(err).ShouldNot(HaveOccurred())		
	})
	It("should execute dotnet publish command without a framework and runtime", func() {
		defer func() { dotnet.ExecCommand = exec.Command }()
		fakes.MockedExitStatus = 0
		fakes.MockedStdout = ""
		expectedCommand := "dotnet publish /path/project.csproj -o /tmp/publish --no-build --no-restore"		
		client := dotnet.NewDotnetClient("","")
		_, err := client.Publish("/path/project.csproj","/tmp/publish", true, true)
		Ω(fakes.CommandString).Should(Equal(expectedCommand))
		Ω(err).ShouldNot(HaveOccurred())		
	})

	
})