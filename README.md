[![Build Status](https://travis-ci.org/miclip/dotnet-extensions.svg?branch=master)](https://travis-ci.org/miclip/dotnet-extensions) [![codecov](https://codecov.io/gh/miclip/dotnet-extensions/branch/master/graph/badge.svg)](https://codecov.io/gh/miclip/dotnet-extensions)

### Dotnet CLI Extensions

This project is a plugin for the dotnet cli to fill gaps when working for CI/CD platforms that are stateless like [Concourse](https://concourse-ci.org/).

It's main target is being able to package an published application as a nuget package and push to a nuget feed. This enables nuget as an immutable application artifact repository whilst leveraging the existing nuget tooling.

Currently there are two supported commands:

`dotnet nuget-ext version`

This command queuries the provided nuget feed for a particlar package and returns either the latest or lists all the current versions. This enables CI/CD pipelines to increment the version without introducing complex version numbers with dates and/or times.

#### Latest
``` sh
dotnet nuget-ext versions --name Steeltoe.Common --latest --source https://api.nuget.org/v3/index.json
```

returns:
``` sh
2.1.1
```
#### List All
```sh
dotnet nuget-ext versions --name Steeltoe.Common --source https://api.nuget.org/v3/index.json
```

```sh
2.0.0
2.1.0
2.1.1
```

### Usage
```
Usage:
  dotnet nuget-ext versions [flags]
```
#### Flags
| Short | Long  |  Type |  Description|   
|---|---|---|---|
| -h | --help |   | help for versions |
| -l | --latest | bool | Returns only latest version | 
| -n | --name |  string |  Package Name |
| -p | --prerelease | bool |  Includes prerelease packages in the list. |
| -s | --source | string |  Specifies the nuget feed URL |

### Examples

#### Package Library increment version

The following command demonstrates packaging a library and incremeting the version based on the `version-spec` and the latest version found in the nuget feed. 

The `version-spec` controls which value is incremented by 1. For example a spec of `1.0.*` and a latest nuget version of `2.0.1` will result in the version `2.0.2`. Version suffixes like `-alpha` are supported provided the spec matches.

~~~ sh
dotnet pack ./DotnetResource.TestLibraryOne.csproj -p:PackageVersion=$(dotnet nuget-ext versions --name DotnetResource.TestLibraryOne --latest --source https://www.myget.org/F/dotnet-resource-test/api/v3/index.json --autoincrement --version-spec "1.0.*")
~~~

The second command:

`dotnet pack-ext simple`


