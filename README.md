[![Build Status](https://travis-ci.org/miclip/dotnet-extensions.svg?branch=master)](https://travis-ci.org/miclip/dotnet-extensions) [![codecov](https://codecov.io/gh/miclip/dotnet-extensions/branch/master/graph/badge.svg)](https://codecov.io/gh/miclip/dotnet-extensions)

### Dotnet CLI Extensions

This project is a plugin for the dotnet cli to fill gaps when working for CI/CD platforms that are stateless like [Concourse](https://concourse-ci.org/). 

It's main target is being able to package an published application as a nuget package and push to a nuget feed. This enables nuget as an immutable application artifact repository whilst leveraging the existing nuget tooling.

Currently there are two  supported commands

`dotnet nuget-ext version` 

This command queuries the provided nuget feed for a particlar package and returns either the latest or lists all the current versions. 

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
| Short | long  |  Type |  Description|   
|---|---|---|---|
| -h | --help |   | help for versions |
| -l | --latest | bool | Returns only latest version | 
| -n | --name |  string |  Package Name |
| -p | --prerelease | bool |  Includes prerelease packages in the list. |
| -s | --source | string |  Specifies the nuget feed URL |

### Examples
