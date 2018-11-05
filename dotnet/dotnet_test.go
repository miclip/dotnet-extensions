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

	It("should execute dotnet publish command", func() {
		defer func() { dotnet.ExecCommand = exec.Command }()
		fakes.MockedExitStatus = 0
		fakes.MockedStdout = ""
		expectedCommand := "dotnet publish /path/project.csproj --no-build --no-restore -f netcoreapp2.1 -r ubuntu.14.04-x64 -o /tmp/publish"		
		client := dotnet.NewDotnetClient("netcoreapp2.1","ubuntu.14.04-x64")
		_, err := client.Publish("/path/project.csproj","/tmp/publish")
		Ω(fakes.CommandString).Should(Equal(expectedCommand))
		Ω(err).ShouldNot(HaveOccurred())		
	})

	
})