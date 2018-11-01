// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"os"

	"github.com/miclip/dotnet-extensions/dotnet"
	"github.com/miclip/dotnet-extensions/nuget"
	"github.com/miclip/dotnet-extensions/ui"
	"github.com/spf13/cobra"
)

var (
	ProjectFile   string
	Version       string
	AutoIncrement bool
	FeedUrl       string
	OutputDir     string
	NoBuild       bool
	NoRestore     bool
	Runtime       string
	Framework     string
)

// simpleCmd represents the recipe command
var simpleCmd = &cobra.Command{
	Use:   "simple <PROJECT>",
	Short: "Create a simple package",
	Long: `Create a simple package

	Create a simple package, intended for Applications that have been published`,
	Run: func(cmd *cobra.Command, args []string) {
		nugetClientv3 := nuget.NewNugetClientv3(FeedUrl)
		dotnetClient := dotnet.NewDotnetClient(args[0], Framework, Runtime, "")
		ui := ui.NewConsoleUI()
		HandleSimplePack(ui, nugetClientv3, dotnetClient, args)
	},
}

// HandleSimplePack ...
func HandleSimplePack(ui ui.UI, nugetClientv3 nuget.NugetClientv3, dotnetClient dotnet.DotnetClient, args []string) {
	ProjectFile = args[0]
	if _, err := os.Stat(ProjectFile); os.IsNotExist(err) {
		ui.ErrorLinef("project %s not found \n %v", ProjectFile, err)
		return
	}
	nuspec, _ := nugetClientv3.CreateNuspecFromProject(ProjectFile, "")
	
	if AutoIncrement {
		versions, err := nugetClientv3.GetPackageVersions(context.Background(), nuspec.ID, true)
		if err != nil {
			ui.ErrorLinef("error retrieving latest version from feed. %v", err)
			return
		}
		newVersion, err := nugetClientv3.AutoIncrementVersion(Version,versions[len(versions)-1].Version)
		if err != nil {
			ui.ErrorLinef("error incrementing version %v", err)
			return
		}
		nuspec.Version = newVersion
	}
	ui.PrintLinef("Package %s Version: %s", nuspec.ID, nuspec.Version)
}

func init() {
	rootCmd.AddCommand(simpleCmd)
	simpleCmd.Flags().StringVarP(&FeedUrl, "source", "s", "", "Specifies the nuget server URL.")
	simpleCmd.Flags().StringVarP(&Version, "version", "v", "", "Version of the package")
	simpleCmd.Flags().BoolVarP(&AutoIncrement, "autoincrement", "a", false, "Automatically increments the version based on latest from nuget feed, requires --source")
	simpleCmd.Flags().StringVarP(&OutputDir, "output", "o", "", "The output directory to place built packages in.")
	simpleCmd.Flags().BoolVarP(&NoBuild, "no-build", "", false, "Do not build the project before packing. Implies --no-restore.")
	simpleCmd.Flags().BoolVarP(&NoRestore, "no-restore", "", false, "Do not restore the project before building.")
	simpleCmd.Flags().StringVarP(&Runtime, "runtime", "r", "", "The target runtime to restore packages for.")
	simpleCmd.Flags().StringVarP(&Framework, "framework", "f", "", "The target framework")
}
