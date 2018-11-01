// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	context "context"
	sync "sync"

	nuget "github.com/miclip/dotnet-extensions/nuget"
)

type FakeNugetClientv3 struct {
	AutoIncrementVersionStub        func(string, string) (string, error)
	autoIncrementVersionMutex       sync.RWMutex
	autoIncrementVersionArgsForCall []struct {
		arg1 string
		arg2 string
	}
	autoIncrementVersionReturns struct {
		result1 string
		result2 error
	}
	autoIncrementVersionReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	CreateNuspecStub        func(string, string, string, string, string) nuget.Nuspec
	createNuspecMutex       sync.RWMutex
	createNuspecArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}
	createNuspecReturns struct {
		result1 nuget.Nuspec
	}
	createNuspecReturnsOnCall map[int]struct {
		result1 nuget.Nuspec
	}
	CreateNuspecFromProjectStub        func(string, string) (nuget.Nuspec, error)
	createNuspecFromProjectMutex       sync.RWMutex
	createNuspecFromProjectArgsForCall []struct {
		arg1 string
		arg2 string
	}
	createNuspecFromProjectReturns struct {
		result1 nuget.Nuspec
		result2 error
	}
	createNuspecFromProjectReturnsOnCall map[int]struct {
		result1 nuget.Nuspec
		result2 error
	}
	DownloadPackageStub        func(context.Context, string, string, string) error
	downloadPackageMutex       sync.RWMutex
	downloadPackageArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}
	downloadPackageReturns struct {
		result1 error
	}
	downloadPackageReturnsOnCall map[int]struct {
		result1 error
	}
	GetNugetApiEndPointStub        func(context.Context, string) (string, error)
	getNugetApiEndPointMutex       sync.RWMutex
	getNugetApiEndPointArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getNugetApiEndPointReturns struct {
		result1 string
		result2 error
	}
	getNugetApiEndPointReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetPackageVersionsStub        func(context.Context, string, bool) ([]nuget.Version, error)
	getPackageVersionsMutex       sync.RWMutex
	getPackageVersionsArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 bool
	}
	getPackageVersionsReturns struct {
		result1 []nuget.Version
		result2 error
	}
	getPackageVersionsReturnsOnCall map[int]struct {
		result1 []nuget.Version
		result2 error
	}
	GetServiceIndexStub        func(context.Context) (*nuget.ServiceIndex, error)
	getServiceIndexMutex       sync.RWMutex
	getServiceIndexArgsForCall []struct {
		arg1 context.Context
	}
	getServiceIndexReturns struct {
		result1 *nuget.ServiceIndex
		result2 error
	}
	getServiceIndexReturnsOnCall map[int]struct {
		result1 *nuget.ServiceIndex
		result2 error
	}
	SearchQueryServiceStub        func(context.Context, string, string, bool) (*nuget.SearchResults, error)
	searchQueryServiceMutex       sync.RWMutex
	searchQueryServiceArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 bool
	}
	searchQueryServiceReturns struct {
		result1 *nuget.SearchResults
		result2 error
	}
	searchQueryServiceReturnsOnCall map[int]struct {
		result1 *nuget.SearchResults
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNugetClientv3) AutoIncrementVersion(arg1 string, arg2 string) (string, error) {
	fake.autoIncrementVersionMutex.Lock()
	ret, specificReturn := fake.autoIncrementVersionReturnsOnCall[len(fake.autoIncrementVersionArgsForCall)]
	fake.autoIncrementVersionArgsForCall = append(fake.autoIncrementVersionArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("AutoIncrementVersion", []interface{}{arg1, arg2})
	fake.autoIncrementVersionMutex.Unlock()
	if fake.AutoIncrementVersionStub != nil {
		return fake.AutoIncrementVersionStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.autoIncrementVersionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNugetClientv3) AutoIncrementVersionCallCount() int {
	fake.autoIncrementVersionMutex.RLock()
	defer fake.autoIncrementVersionMutex.RUnlock()
	return len(fake.autoIncrementVersionArgsForCall)
}

func (fake *FakeNugetClientv3) AutoIncrementVersionCalls(stub func(string, string) (string, error)) {
	fake.autoIncrementVersionMutex.Lock()
	defer fake.autoIncrementVersionMutex.Unlock()
	fake.AutoIncrementVersionStub = stub
}

func (fake *FakeNugetClientv3) AutoIncrementVersionArgsForCall(i int) (string, string) {
	fake.autoIncrementVersionMutex.RLock()
	defer fake.autoIncrementVersionMutex.RUnlock()
	argsForCall := fake.autoIncrementVersionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNugetClientv3) AutoIncrementVersionReturns(result1 string, result2 error) {
	fake.autoIncrementVersionMutex.Lock()
	defer fake.autoIncrementVersionMutex.Unlock()
	fake.AutoIncrementVersionStub = nil
	fake.autoIncrementVersionReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeNugetClientv3) AutoIncrementVersionReturnsOnCall(i int, result1 string, result2 error) {
	fake.autoIncrementVersionMutex.Lock()
	defer fake.autoIncrementVersionMutex.Unlock()
	fake.AutoIncrementVersionStub = nil
	if fake.autoIncrementVersionReturnsOnCall == nil {
		fake.autoIncrementVersionReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.autoIncrementVersionReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeNugetClientv3) CreateNuspec(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) nuget.Nuspec {
	fake.createNuspecMutex.Lock()
	ret, specificReturn := fake.createNuspecReturnsOnCall[len(fake.createNuspecArgsForCall)]
	fake.createNuspecArgsForCall = append(fake.createNuspecArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("CreateNuspec", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.createNuspecMutex.Unlock()
	if fake.CreateNuspecStub != nil {
		return fake.CreateNuspecStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createNuspecReturns
	return fakeReturns.result1
}

func (fake *FakeNugetClientv3) CreateNuspecCallCount() int {
	fake.createNuspecMutex.RLock()
	defer fake.createNuspecMutex.RUnlock()
	return len(fake.createNuspecArgsForCall)
}

func (fake *FakeNugetClientv3) CreateNuspecCalls(stub func(string, string, string, string, string) nuget.Nuspec) {
	fake.createNuspecMutex.Lock()
	defer fake.createNuspecMutex.Unlock()
	fake.CreateNuspecStub = stub
}

func (fake *FakeNugetClientv3) CreateNuspecArgsForCall(i int) (string, string, string, string, string) {
	fake.createNuspecMutex.RLock()
	defer fake.createNuspecMutex.RUnlock()
	argsForCall := fake.createNuspecArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeNugetClientv3) CreateNuspecReturns(result1 nuget.Nuspec) {
	fake.createNuspecMutex.Lock()
	defer fake.createNuspecMutex.Unlock()
	fake.CreateNuspecStub = nil
	fake.createNuspecReturns = struct {
		result1 nuget.Nuspec
	}{result1}
}

func (fake *FakeNugetClientv3) CreateNuspecReturnsOnCall(i int, result1 nuget.Nuspec) {
	fake.createNuspecMutex.Lock()
	defer fake.createNuspecMutex.Unlock()
	fake.CreateNuspecStub = nil
	if fake.createNuspecReturnsOnCall == nil {
		fake.createNuspecReturnsOnCall = make(map[int]struct {
			result1 nuget.Nuspec
		})
	}
	fake.createNuspecReturnsOnCall[i] = struct {
		result1 nuget.Nuspec
	}{result1}
}

func (fake *FakeNugetClientv3) CreateNuspecFromProject(arg1 string, arg2 string) (nuget.Nuspec, error) {
	fake.createNuspecFromProjectMutex.Lock()
	ret, specificReturn := fake.createNuspecFromProjectReturnsOnCall[len(fake.createNuspecFromProjectArgsForCall)]
	fake.createNuspecFromProjectArgsForCall = append(fake.createNuspecFromProjectArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreateNuspecFromProject", []interface{}{arg1, arg2})
	fake.createNuspecFromProjectMutex.Unlock()
	if fake.CreateNuspecFromProjectStub != nil {
		return fake.CreateNuspecFromProjectStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createNuspecFromProjectReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNugetClientv3) CreateNuspecFromProjectCallCount() int {
	fake.createNuspecFromProjectMutex.RLock()
	defer fake.createNuspecFromProjectMutex.RUnlock()
	return len(fake.createNuspecFromProjectArgsForCall)
}

func (fake *FakeNugetClientv3) CreateNuspecFromProjectCalls(stub func(string, string) (nuget.Nuspec, error)) {
	fake.createNuspecFromProjectMutex.Lock()
	defer fake.createNuspecFromProjectMutex.Unlock()
	fake.CreateNuspecFromProjectStub = stub
}

func (fake *FakeNugetClientv3) CreateNuspecFromProjectArgsForCall(i int) (string, string) {
	fake.createNuspecFromProjectMutex.RLock()
	defer fake.createNuspecFromProjectMutex.RUnlock()
	argsForCall := fake.createNuspecFromProjectArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNugetClientv3) CreateNuspecFromProjectReturns(result1 nuget.Nuspec, result2 error) {
	fake.createNuspecFromProjectMutex.Lock()
	defer fake.createNuspecFromProjectMutex.Unlock()
	fake.CreateNuspecFromProjectStub = nil
	fake.createNuspecFromProjectReturns = struct {
		result1 nuget.Nuspec
		result2 error
	}{result1, result2}
}

func (fake *FakeNugetClientv3) CreateNuspecFromProjectReturnsOnCall(i int, result1 nuget.Nuspec, result2 error) {
	fake.createNuspecFromProjectMutex.Lock()
	defer fake.createNuspecFromProjectMutex.Unlock()
	fake.CreateNuspecFromProjectStub = nil
	if fake.createNuspecFromProjectReturnsOnCall == nil {
		fake.createNuspecFromProjectReturnsOnCall = make(map[int]struct {
			result1 nuget.Nuspec
			result2 error
		})
	}
	fake.createNuspecFromProjectReturnsOnCall[i] = struct {
		result1 nuget.Nuspec
		result2 error
	}{result1, result2}
}

func (fake *FakeNugetClientv3) DownloadPackage(arg1 context.Context, arg2 string, arg3 string, arg4 string) error {
	fake.downloadPackageMutex.Lock()
	ret, specificReturn := fake.downloadPackageReturnsOnCall[len(fake.downloadPackageArgsForCall)]
	fake.downloadPackageArgsForCall = append(fake.downloadPackageArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("DownloadPackage", []interface{}{arg1, arg2, arg3, arg4})
	fake.downloadPackageMutex.Unlock()
	if fake.DownloadPackageStub != nil {
		return fake.DownloadPackageStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.downloadPackageReturns
	return fakeReturns.result1
}

func (fake *FakeNugetClientv3) DownloadPackageCallCount() int {
	fake.downloadPackageMutex.RLock()
	defer fake.downloadPackageMutex.RUnlock()
	return len(fake.downloadPackageArgsForCall)
}

func (fake *FakeNugetClientv3) DownloadPackageCalls(stub func(context.Context, string, string, string) error) {
	fake.downloadPackageMutex.Lock()
	defer fake.downloadPackageMutex.Unlock()
	fake.DownloadPackageStub = stub
}

func (fake *FakeNugetClientv3) DownloadPackageArgsForCall(i int) (context.Context, string, string, string) {
	fake.downloadPackageMutex.RLock()
	defer fake.downloadPackageMutex.RUnlock()
	argsForCall := fake.downloadPackageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeNugetClientv3) DownloadPackageReturns(result1 error) {
	fake.downloadPackageMutex.Lock()
	defer fake.downloadPackageMutex.Unlock()
	fake.DownloadPackageStub = nil
	fake.downloadPackageReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNugetClientv3) DownloadPackageReturnsOnCall(i int, result1 error) {
	fake.downloadPackageMutex.Lock()
	defer fake.downloadPackageMutex.Unlock()
	fake.DownloadPackageStub = nil
	if fake.downloadPackageReturnsOnCall == nil {
		fake.downloadPackageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.downloadPackageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNugetClientv3) GetNugetApiEndPoint(arg1 context.Context, arg2 string) (string, error) {
	fake.getNugetApiEndPointMutex.Lock()
	ret, specificReturn := fake.getNugetApiEndPointReturnsOnCall[len(fake.getNugetApiEndPointArgsForCall)]
	fake.getNugetApiEndPointArgsForCall = append(fake.getNugetApiEndPointArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetNugetApiEndPoint", []interface{}{arg1, arg2})
	fake.getNugetApiEndPointMutex.Unlock()
	if fake.GetNugetApiEndPointStub != nil {
		return fake.GetNugetApiEndPointStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getNugetApiEndPointReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNugetClientv3) GetNugetApiEndPointCallCount() int {
	fake.getNugetApiEndPointMutex.RLock()
	defer fake.getNugetApiEndPointMutex.RUnlock()
	return len(fake.getNugetApiEndPointArgsForCall)
}

func (fake *FakeNugetClientv3) GetNugetApiEndPointCalls(stub func(context.Context, string) (string, error)) {
	fake.getNugetApiEndPointMutex.Lock()
	defer fake.getNugetApiEndPointMutex.Unlock()
	fake.GetNugetApiEndPointStub = stub
}

func (fake *FakeNugetClientv3) GetNugetApiEndPointArgsForCall(i int) (context.Context, string) {
	fake.getNugetApiEndPointMutex.RLock()
	defer fake.getNugetApiEndPointMutex.RUnlock()
	argsForCall := fake.getNugetApiEndPointArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNugetClientv3) GetNugetApiEndPointReturns(result1 string, result2 error) {
	fake.getNugetApiEndPointMutex.Lock()
	defer fake.getNugetApiEndPointMutex.Unlock()
	fake.GetNugetApiEndPointStub = nil
	fake.getNugetApiEndPointReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeNugetClientv3) GetNugetApiEndPointReturnsOnCall(i int, result1 string, result2 error) {
	fake.getNugetApiEndPointMutex.Lock()
	defer fake.getNugetApiEndPointMutex.Unlock()
	fake.GetNugetApiEndPointStub = nil
	if fake.getNugetApiEndPointReturnsOnCall == nil {
		fake.getNugetApiEndPointReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getNugetApiEndPointReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeNugetClientv3) GetPackageVersions(arg1 context.Context, arg2 string, arg3 bool) ([]nuget.Version, error) {
	fake.getPackageVersionsMutex.Lock()
	ret, specificReturn := fake.getPackageVersionsReturnsOnCall[len(fake.getPackageVersionsArgsForCall)]
	fake.getPackageVersionsArgsForCall = append(fake.getPackageVersionsArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 bool
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetPackageVersions", []interface{}{arg1, arg2, arg3})
	fake.getPackageVersionsMutex.Unlock()
	if fake.GetPackageVersionsStub != nil {
		return fake.GetPackageVersionsStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getPackageVersionsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNugetClientv3) GetPackageVersionsCallCount() int {
	fake.getPackageVersionsMutex.RLock()
	defer fake.getPackageVersionsMutex.RUnlock()
	return len(fake.getPackageVersionsArgsForCall)
}

func (fake *FakeNugetClientv3) GetPackageVersionsCalls(stub func(context.Context, string, bool) ([]nuget.Version, error)) {
	fake.getPackageVersionsMutex.Lock()
	defer fake.getPackageVersionsMutex.Unlock()
	fake.GetPackageVersionsStub = stub
}

func (fake *FakeNugetClientv3) GetPackageVersionsArgsForCall(i int) (context.Context, string, bool) {
	fake.getPackageVersionsMutex.RLock()
	defer fake.getPackageVersionsMutex.RUnlock()
	argsForCall := fake.getPackageVersionsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeNugetClientv3) GetPackageVersionsReturns(result1 []nuget.Version, result2 error) {
	fake.getPackageVersionsMutex.Lock()
	defer fake.getPackageVersionsMutex.Unlock()
	fake.GetPackageVersionsStub = nil
	fake.getPackageVersionsReturns = struct {
		result1 []nuget.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeNugetClientv3) GetPackageVersionsReturnsOnCall(i int, result1 []nuget.Version, result2 error) {
	fake.getPackageVersionsMutex.Lock()
	defer fake.getPackageVersionsMutex.Unlock()
	fake.GetPackageVersionsStub = nil
	if fake.getPackageVersionsReturnsOnCall == nil {
		fake.getPackageVersionsReturnsOnCall = make(map[int]struct {
			result1 []nuget.Version
			result2 error
		})
	}
	fake.getPackageVersionsReturnsOnCall[i] = struct {
		result1 []nuget.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeNugetClientv3) GetServiceIndex(arg1 context.Context) (*nuget.ServiceIndex, error) {
	fake.getServiceIndexMutex.Lock()
	ret, specificReturn := fake.getServiceIndexReturnsOnCall[len(fake.getServiceIndexArgsForCall)]
	fake.getServiceIndexArgsForCall = append(fake.getServiceIndexArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("GetServiceIndex", []interface{}{arg1})
	fake.getServiceIndexMutex.Unlock()
	if fake.GetServiceIndexStub != nil {
		return fake.GetServiceIndexStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getServiceIndexReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNugetClientv3) GetServiceIndexCallCount() int {
	fake.getServiceIndexMutex.RLock()
	defer fake.getServiceIndexMutex.RUnlock()
	return len(fake.getServiceIndexArgsForCall)
}

func (fake *FakeNugetClientv3) GetServiceIndexCalls(stub func(context.Context) (*nuget.ServiceIndex, error)) {
	fake.getServiceIndexMutex.Lock()
	defer fake.getServiceIndexMutex.Unlock()
	fake.GetServiceIndexStub = stub
}

func (fake *FakeNugetClientv3) GetServiceIndexArgsForCall(i int) context.Context {
	fake.getServiceIndexMutex.RLock()
	defer fake.getServiceIndexMutex.RUnlock()
	argsForCall := fake.getServiceIndexArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNugetClientv3) GetServiceIndexReturns(result1 *nuget.ServiceIndex, result2 error) {
	fake.getServiceIndexMutex.Lock()
	defer fake.getServiceIndexMutex.Unlock()
	fake.GetServiceIndexStub = nil
	fake.getServiceIndexReturns = struct {
		result1 *nuget.ServiceIndex
		result2 error
	}{result1, result2}
}

func (fake *FakeNugetClientv3) GetServiceIndexReturnsOnCall(i int, result1 *nuget.ServiceIndex, result2 error) {
	fake.getServiceIndexMutex.Lock()
	defer fake.getServiceIndexMutex.Unlock()
	fake.GetServiceIndexStub = nil
	if fake.getServiceIndexReturnsOnCall == nil {
		fake.getServiceIndexReturnsOnCall = make(map[int]struct {
			result1 *nuget.ServiceIndex
			result2 error
		})
	}
	fake.getServiceIndexReturnsOnCall[i] = struct {
		result1 *nuget.ServiceIndex
		result2 error
	}{result1, result2}
}

func (fake *FakeNugetClientv3) SearchQueryService(arg1 context.Context, arg2 string, arg3 string, arg4 bool) (*nuget.SearchResults, error) {
	fake.searchQueryServiceMutex.Lock()
	ret, specificReturn := fake.searchQueryServiceReturnsOnCall[len(fake.searchQueryServiceArgsForCall)]
	fake.searchQueryServiceArgsForCall = append(fake.searchQueryServiceArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 bool
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("SearchQueryService", []interface{}{arg1, arg2, arg3, arg4})
	fake.searchQueryServiceMutex.Unlock()
	if fake.SearchQueryServiceStub != nil {
		return fake.SearchQueryServiceStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.searchQueryServiceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNugetClientv3) SearchQueryServiceCallCount() int {
	fake.searchQueryServiceMutex.RLock()
	defer fake.searchQueryServiceMutex.RUnlock()
	return len(fake.searchQueryServiceArgsForCall)
}

func (fake *FakeNugetClientv3) SearchQueryServiceCalls(stub func(context.Context, string, string, bool) (*nuget.SearchResults, error)) {
	fake.searchQueryServiceMutex.Lock()
	defer fake.searchQueryServiceMutex.Unlock()
	fake.SearchQueryServiceStub = stub
}

func (fake *FakeNugetClientv3) SearchQueryServiceArgsForCall(i int) (context.Context, string, string, bool) {
	fake.searchQueryServiceMutex.RLock()
	defer fake.searchQueryServiceMutex.RUnlock()
	argsForCall := fake.searchQueryServiceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeNugetClientv3) SearchQueryServiceReturns(result1 *nuget.SearchResults, result2 error) {
	fake.searchQueryServiceMutex.Lock()
	defer fake.searchQueryServiceMutex.Unlock()
	fake.SearchQueryServiceStub = nil
	fake.searchQueryServiceReturns = struct {
		result1 *nuget.SearchResults
		result2 error
	}{result1, result2}
}

func (fake *FakeNugetClientv3) SearchQueryServiceReturnsOnCall(i int, result1 *nuget.SearchResults, result2 error) {
	fake.searchQueryServiceMutex.Lock()
	defer fake.searchQueryServiceMutex.Unlock()
	fake.SearchQueryServiceStub = nil
	if fake.searchQueryServiceReturnsOnCall == nil {
		fake.searchQueryServiceReturnsOnCall = make(map[int]struct {
			result1 *nuget.SearchResults
			result2 error
		})
	}
	fake.searchQueryServiceReturnsOnCall[i] = struct {
		result1 *nuget.SearchResults
		result2 error
	}{result1, result2}
}

func (fake *FakeNugetClientv3) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.autoIncrementVersionMutex.RLock()
	defer fake.autoIncrementVersionMutex.RUnlock()
	fake.createNuspecMutex.RLock()
	defer fake.createNuspecMutex.RUnlock()
	fake.createNuspecFromProjectMutex.RLock()
	defer fake.createNuspecFromProjectMutex.RUnlock()
	fake.downloadPackageMutex.RLock()
	defer fake.downloadPackageMutex.RUnlock()
	fake.getNugetApiEndPointMutex.RLock()
	defer fake.getNugetApiEndPointMutex.RUnlock()
	fake.getPackageVersionsMutex.RLock()
	defer fake.getPackageVersionsMutex.RUnlock()
	fake.getServiceIndexMutex.RLock()
	defer fake.getServiceIndexMutex.RUnlock()
	fake.searchQueryServiceMutex.RLock()
	defer fake.searchQueryServiceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNugetClientv3) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ nuget.NugetClientv3 = new(FakeNugetClientv3)
