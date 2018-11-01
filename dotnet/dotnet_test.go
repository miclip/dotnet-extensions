package dotnet_test


import (
	"github.com/miclip/dotnet-extensions/fakes"
	
	//"os/exec"
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
	"github.com/miclip/dotnet-extensions/dotnet"
)



var _ = Describe("dotnetclient", func() {
	BeforeEach(func(){
		dotnet.ExecCommand = fakes.FakeExecCommand
	})

	It("should execute dotnet build command", func() {
		// defer func() { dotnet.ExecCommand = exec.Command }()
		// fakes.MockedExitStatus = 0
		// fakes.MockedStdout = ""
		// expectedCommand := "dotnet build tmp/source/path/project.csproj -f netcoreapp2.1 -r ubuntu.14.04-x64"		
		
		// client := dotnet.NewDotnetClient("/path/project.csproj","netcoreapp2.1","ubuntu.14.04-x64", "tmp/source")
		// _, err := client.Build()
		// Ω(fakes.CommandString).Should(Equal(expectedCommand))
		// Ω(err).ShouldNot(HaveOccurred())		
	})

	
})