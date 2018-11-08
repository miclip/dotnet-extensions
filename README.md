[![Build Status](https://travis-ci.org/miclip/dotnet-extensions.svg?branch=master)](https://travis-ci.org/miclip/dotnet-extensions) [![codecov](https://codecov.io/gh/miclip/dotnet-extensions/branch/master/graph/badge.svg)](https://codecov.io/gh/miclip/dotnet-extensions)

# Dotnet CLI Extensions

This project is a plugin for the dotnet cli to package the pubished artifacts of applications into nuget packages. It aims make it easier when working with CI/CD platforms that don't have internal storage for artifacts like [Concourse](https://concourse-ci.org/).

It's main target is being able to package an published application as a nuget package and push to a nuget feed. This enables nuget as an immutable application artifact repository whilst leveraging the existing nuget tooling.

Currently there are two supported commands:

## Dotnet nuget-ext

`dotnet nuget-ext version`

This command queries the provided nuget feed for a particular package and returns either the latest or lists all the current versions. This enables CI/CD pipelines to increment the version without introducing complex version numbers with dates and/or times.

### Key Features

#### List All

```sh
dotnet nuget-ext versions --name Steeltoe.Common --source https://api.nuget.org/v3/index.json
```

```sh
2.0.0
2.1.0
2.1.1
```

#### Latest (--latest)

``` sh
dotnet nuget-ext versions --name Steeltoe.Common --latest --source https://api.nuget.org/v3/index.json
```

returns:

``` sh
2.1.1
```

### Usage

``` sh
Usage:
  dotnet nuget-ext versions [flags]
```

#### Flags

| Short | Long  |  Type |  Description|
|---|---|---|---|
| -h | --help |   | help for versions |
| -l | --latest | bool | Returns only latest version vs listing all versions.|
| -n | --name |  string |  Package Name e.g. `Newtonsoft.Json` |
| -p | --prerelease | bool |  Includes prerelease packages in the list. |
| -s | --source | string |  Specifies the nuget feed URL |
| -i | --increment | bool |  increments the version number based on latest in feed and `version-spec`, version-spec is required |
| -m | --version-spec | string | string that controls which section of the version to increment by one or more asterisks .e.g. `1.0.*` increments the patch, `1.*.*` increments the minor and patch. |

### Examples

#### Package Library increment version

The following command demonstrates packaging a library and incrementing the version based on the `version-spec` and the latest version found in the nuget feed.

The `version-spec` controls which value is incremented by 1. For example a spec of `1.0.*` and a latest nuget version of `2.0.1` will result in the version `2.0.2`. Version suffixes like `-alpha` are supported provided the spec matches `1.0.1-alpha.*` => `2.0.2-alpha.234`.

~~~ sh
dotnet pack ./DotnetResource.TestLibraryOne.csproj -p:PackageVersion=$(dotnet nuget-ext versions 
--name DotnetResource.TestLibraryOne --latest 
--source https://www.myget.org/F/dotnet-resource-test/api/v3/index.json 
--increment --version-spec "1.0.*")
~~~

## Dotnet pack-ext

`dotnet pack-ext simple`

The simple pack command targets packaging the published output of an application into a nuget package. The standard `dotnet pack` command only supports libraries and requires files to be organized into content and libs etc. This is complicated to manage and is overkill for the published output as all that's needed as an immutable package of the runtime artifacts.

Combined with the `dotnet nuget-ext versions` command this plugin provides an easy way to manage application runtime artifacts via any nuget feed.

### Key Features

#### Version (--version)

Provide a static version to use as with `dotnet pack`. The command will run the `dotnet publish` command, package the `--output` directory using the version provided.

~~~ sh
dotnet pack-ext simple ./TestApplication/TestApplication.csproj -o ./_publish --version 1.0.1 -r ubuntu.14.04-x64 -f netcoreapp2.1
~~~

After publishing, packaging and versioning automatically push the package to the feed.

~~~ sh
dotnet pack-ext simple ./TestApplication/TestApplication.csproj -o ./_publish --version 1.0.1 -r ubuntu.14.04-x64 -f netcoreapp2.1 --push-package --source https://www.nuget.org/myfeed/api/v3/index.json -k [APIKEY]
~~~

#### Increment (--increment)

Rather then provide a version the plugin can generate the version based on the latest package in the nuget feed and rules provided in the `--version` [spec]. Note the Asterisks in the version.

~~~ sh
dotnet pack-ext simple ./TestApplication/TestApplication.csproj -o ./_publish --version 1.0.* -r ubuntu.14.04-x64 -f netcoreapp2.1 --push-package --source https://www.nuget.org/myfeed/api/v3/index.json -k [APIKEY] --increment
~~~

#### Prepublish

The `simple` command can package a pre published application by setting the `--basepath` flag to the location of the published output. Note that the `csproj` is still required as it contains the information for the nuspec file.

~~~ sh
dotnet pack-ext simple ./TestApplication/TestApplication.csproj -o ./packages --version 1.0.* -r ubuntu.14.04-x64 -f netcoreapp2.1 --push-package --source https://www.nuget.org/myfeed/api/v3/index.json -k [APIKEY] --increment --basepath ./_publish
~~~

### Usage

``` sh
Usage:
  dotnet pack-ext simple [flags]
```

#### Flags

| Short | Long  |  Type |  Description|
|---|---|---|---|
| -h | --help |  bool | help for versions |
| -k | --api-key | string | The API key for the nuget server.|
| -i | --increment | bool |  increments the version number based on latest in feed and `version-spec`, `version-spec` is required when `increment` is true|
| -v | --version | string |  Version or version-spec of the package |
| -b | --basepath | string | Location of published artifacts, required when `--no-pubish` is true |
| -f | --framework | string | The target framework |
| | --no-build | bool | Do not build the project before packing. Implies `--no-restore`. |
| | --no-publish | bool | Do not publish the project before packing. Implies --no-build & --no-restore. Requires a `--basepath` to the published output |
| | --no-restore | bool | Do not publish the project before packing. Implies --no-build & --no-restore. Requires a `--basepath` to the published output |
| -o | --output | string | The output directory to place built packages in. |
| -p | --push-package | bool | Push the package to nuget feed, `--source` and --api-key are required. |
| -r | --runtime | bool | The target runtime to restore packages for. |
| -s | --source | string | Specifies the nuget server URL. |

## Installation

//TODO Self install

### Homebrew

### Chocolatey

### Apt-get
