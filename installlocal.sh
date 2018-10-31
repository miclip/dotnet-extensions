#!/bin/sh
set -x
go install ./dotnet-nuget-ext 
go install ./dotnet-pack-ext

dotnet pack-ext --version
dotnet nuget-ext --version