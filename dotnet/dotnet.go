package dotnet

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"github.com/miclip/dotnet-extensions"
)

// DotnetClient ...
type DotnetClient interface {
	SimplePack(packageID string, version string, sourceDir string, outputDir string) (string, error)
}

type dotnetclient struct {
	path      string
	framework string
	runtime   string
}

// NewDotnetClient ...
func NewDotnetClient(
	path string,
	framework string,
	runtime string,
) DotnetClient {
	projectPath := path
	targetFramework := framework
	targetRuntime := runtime
	return &dotnetclient{
		path:      projectPath,
		framework: targetFramework,
		runtime:   targetRuntime,
	}
}

func (client *dotnetclient) SimplePack(packageID string, version string, sourceDir string, outputDir string) (string, error) {
	packageName := fmt.Sprintf("%s.%s.nupkg", packageID, version)
	packagePath := path.Join(outputDir, packageName)
	files := []string{}
	err := filepath.Walk(sourceDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Mode().IsRegular() {
				files = append(files, path)
			}			
			return nil
		})	
	err = utils.ZipFolder(packagePath, files, sourceDir)
	if err != nil {
		return packagePath, err
	}

	return packagePath, nil
}