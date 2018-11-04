package dotnet

import (
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
