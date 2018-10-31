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

	"github.com/miclip/dotnet-extensions/nuget"
	"github.com/miclip/dotnet-extensions/ui"
	"github.com/spf13/cobra"
)

var (
	FeedUrl     string
	PackageName string
	Prerelease  bool
	Latest      bool
)

// recipeCmd represents the recipe command
var versionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "List versions of a package",
	Long: `List versions of a package

	List versions of a package from a remote nuget feed`,
	Run: func(cmd *cobra.Command, args []string) {
		nugetClientv3 := nuget.NewNugetClientv3(FeedUrl)
		ui := ui.NewConsoleUI()
		HandleVersions(ui, nugetClientv3, args)
	},
}

// HandleVersions ...
func HandleVersions(ui ui.UI, nugetClientv3 nuget.NugetClientv3, args []string) {
	versions, err := nugetClientv3.GetPackageVersions(context.Background(), PackageName, Prerelease)
	if err != nil {
		ui.ErrorLinef("error retrieving latest version from feed. %v", err)
		return
	}
	if len(versions) == 0 {
		ui.ErrorLinef("no version details found.")
		return
	}
	if Latest {
		ui.PrintLinef(versions[len(versions)-1].Version)
		return
	}
	for _, v := range versions {
		ui.PrintLinef(v.Version)
	}

}

func init() {
	rootCmd.AddCommand(versionsCmd)
	versionsCmd.Flags().StringVarP(&FeedUrl, "source", "s", "", "Specifies the server URL.")
	versionsCmd.Flags().StringVarP(&PackageName, "name", "n", "", "Package name")
	versionsCmd.Flags().BoolVarP(&Prerelease, "prerelease", "p", false, "Includes prerelease packages in the list.")
	versionsCmd.Flags().BoolVarP(&Latest, "latest", "l", false, "Returns only latest version")
}
