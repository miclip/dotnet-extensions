package dotnet

import (
	"os/exec"
	"strings"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/miclip/dotnet-extensions"
)

// DotnetClient ...
type DotnetClient interface {
	SimplePack(packageID string, version string, sourceDir string, outputDir string) (string, error)
	Publish(projectPath string, outputdir string) ([]byte, error)
}

type dotnetclient struct {
	framework string
	runtime   string
}

// NewDotnetClient ...
func NewDotnetClient(
	framework string,
	runtime string,
) DotnetClient {
	targetFramework := framework
	targetRuntime := runtime
	return &dotnetclient{
		framework: targetFramework,
		runtime:   targetRuntime,
	}
}

var ExecCommand = exec.Command

func (client *dotnetclient) SimplePack(packageID string, version string, sourceDir string, outputDir string) (string, error) {
	packageName := fmt.Sprintf("%s.%s.nupkg", packageID, version)
	packagePath := path.Join(outputDir, packageName)
	files := []string{}
	isSubDir := false
	subDirRootPath := ""
	err := filepath.Walk(sourceDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// list all files in root of sourceDir and only first level subdirectories, not files contained within
			// the zip library duplicates the subdirectories
			if !isSubDir  {
				files = append(files, path)
			}

			if !isSubDir && info.IsDir() && path != sourceDir {
				isSubDir = true
				subDirRootPath = path
			}

			if isSubDir && !strings.Contains(path, subDirRootPath){
				isSubDir = false
				subDirRootPath = ""
			}			
			return nil
		})
	err = utils.MakeZip(packagePath, files)
	if err != nil {
		return packagePath, err
	}

	return packagePath, nil
}

func (client *dotnetclient) Publish(projectPath string, outputdir string) ([]byte, error) {
	cmd := ExecCommand("dotnet", "publish", projectPath, "--no-build", "--no-restore", "-f", client.framework, "-r", client.runtime, "-o", outputdir)
	out, err := cmd.CombinedOutput()
	return out, err
}