// Code generated by counterfeiter. DO NOT EDIT.
package execfakes

import (
	"io"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/exec"
	"github.com/concourse/concourse/vars"
)

type FakeCheckDelegate struct {
	ErroredStub        func(lager.Logger, string)
	erroredMutex       sync.RWMutex
	erroredArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	FinishedStub        func(lager.Logger, bool)
	finishedMutex       sync.RWMutex
	finishedArgsForCall []struct {
		arg1 lager.Logger
		arg2 bool
	}
	ImageSourceRedactionStub        func(atc.Source) (atc.Source, error)
	imageSourceRedactionMutex       sync.RWMutex
	imageSourceRedactionArgsForCall []struct {
		arg1 atc.Source
	}
	imageSourceRedactionReturns struct {
		result1 atc.Source
		result2 error
	}
	imageSourceRedactionReturnsOnCall map[int]struct {
		result1 atc.Source
		result2 error
	}
	ImageVersionDeterminedStub        func(db.UsedResourceCache) error
	imageVersionDeterminedMutex       sync.RWMutex
	imageVersionDeterminedArgsForCall []struct {
		arg1 db.UsedResourceCache
	}
	imageVersionDeterminedReturns struct {
		result1 error
	}
	imageVersionDeterminedReturnsOnCall map[int]struct {
		result1 error
	}
	InitializingStub        func(lager.Logger)
	initializingMutex       sync.RWMutex
	initializingArgsForCall []struct {
		arg1 lager.Logger
	}
	SaveVersionsStub        func(db.SpanContext, []atc.Version) error
	saveVersionsMutex       sync.RWMutex
	saveVersionsArgsForCall []struct {
		arg1 db.SpanContext
		arg2 []atc.Version
	}
	saveVersionsReturns struct {
		result1 error
	}
	saveVersionsReturnsOnCall map[int]struct {
		result1 error
	}
	StartingStub        func(lager.Logger)
	startingMutex       sync.RWMutex
	startingArgsForCall []struct {
		arg1 lager.Logger
	}
	StderrStub        func() io.Writer
	stderrMutex       sync.RWMutex
	stderrArgsForCall []struct {
	}
	stderrReturns struct {
		result1 io.Writer
	}
	stderrReturnsOnCall map[int]struct {
		result1 io.Writer
	}
	StdoutStub        func() io.Writer
	stdoutMutex       sync.RWMutex
	stdoutArgsForCall []struct {
	}
	stdoutReturns struct {
		result1 io.Writer
	}
	stdoutReturnsOnCall map[int]struct {
		result1 io.Writer
	}
	VariablesStub        func() vars.CredVarsTracker
	variablesMutex       sync.RWMutex
	variablesArgsForCall []struct {
	}
	variablesReturns struct {
		result1 vars.CredVarsTracker
	}
	variablesReturnsOnCall map[int]struct {
		result1 vars.CredVarsTracker
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCheckDelegate) Errored(arg1 lager.Logger, arg2 string) {
	fake.erroredMutex.Lock()
	fake.erroredArgsForCall = append(fake.erroredArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Errored", []interface{}{arg1, arg2})
	fake.erroredMutex.Unlock()
	if fake.ErroredStub != nil {
		fake.ErroredStub(arg1, arg2)
	}
}

func (fake *FakeCheckDelegate) ErroredCallCount() int {
	fake.erroredMutex.RLock()
	defer fake.erroredMutex.RUnlock()
	return len(fake.erroredArgsForCall)
}

func (fake *FakeCheckDelegate) ErroredCalls(stub func(lager.Logger, string)) {
	fake.erroredMutex.Lock()
	defer fake.erroredMutex.Unlock()
	fake.ErroredStub = stub
}

func (fake *FakeCheckDelegate) ErroredArgsForCall(i int) (lager.Logger, string) {
	fake.erroredMutex.RLock()
	defer fake.erroredMutex.RUnlock()
	argsForCall := fake.erroredArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCheckDelegate) Finished(arg1 lager.Logger, arg2 bool) {
	fake.finishedMutex.Lock()
	fake.finishedArgsForCall = append(fake.finishedArgsForCall, struct {
		arg1 lager.Logger
		arg2 bool
	}{arg1, arg2})
	fake.recordInvocation("Finished", []interface{}{arg1, arg2})
	fake.finishedMutex.Unlock()
	if fake.FinishedStub != nil {
		fake.FinishedStub(arg1, arg2)
	}
}

func (fake *FakeCheckDelegate) FinishedCallCount() int {
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	return len(fake.finishedArgsForCall)
}

func (fake *FakeCheckDelegate) FinishedCalls(stub func(lager.Logger, bool)) {
	fake.finishedMutex.Lock()
	defer fake.finishedMutex.Unlock()
	fake.FinishedStub = stub
}

func (fake *FakeCheckDelegate) FinishedArgsForCall(i int) (lager.Logger, bool) {
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	argsForCall := fake.finishedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCheckDelegate) ImageSourceRedaction(arg1 atc.Source) (atc.Source, error) {
	fake.imageSourceRedactionMutex.Lock()
	ret, specificReturn := fake.imageSourceRedactionReturnsOnCall[len(fake.imageSourceRedactionArgsForCall)]
	fake.imageSourceRedactionArgsForCall = append(fake.imageSourceRedactionArgsForCall, struct {
		arg1 atc.Source
	}{arg1})
	fake.recordInvocation("ImageSourceRedaction", []interface{}{arg1})
	fake.imageSourceRedactionMutex.Unlock()
	if fake.ImageSourceRedactionStub != nil {
		return fake.ImageSourceRedactionStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.imageSourceRedactionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCheckDelegate) ImageSourceRedactionCallCount() int {
	fake.imageSourceRedactionMutex.RLock()
	defer fake.imageSourceRedactionMutex.RUnlock()
	return len(fake.imageSourceRedactionArgsForCall)
}

func (fake *FakeCheckDelegate) ImageSourceRedactionCalls(stub func(atc.Source) (atc.Source, error)) {
	fake.imageSourceRedactionMutex.Lock()
	defer fake.imageSourceRedactionMutex.Unlock()
	fake.ImageSourceRedactionStub = stub
}

func (fake *FakeCheckDelegate) ImageSourceRedactionArgsForCall(i int) atc.Source {
	fake.imageSourceRedactionMutex.RLock()
	defer fake.imageSourceRedactionMutex.RUnlock()
	argsForCall := fake.imageSourceRedactionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCheckDelegate) ImageSourceRedactionReturns(result1 atc.Source, result2 error) {
	fake.imageSourceRedactionMutex.Lock()
	defer fake.imageSourceRedactionMutex.Unlock()
	fake.ImageSourceRedactionStub = nil
	fake.imageSourceRedactionReturns = struct {
		result1 atc.Source
		result2 error
	}{result1, result2}
}

func (fake *FakeCheckDelegate) ImageSourceRedactionReturnsOnCall(i int, result1 atc.Source, result2 error) {
	fake.imageSourceRedactionMutex.Lock()
	defer fake.imageSourceRedactionMutex.Unlock()
	fake.ImageSourceRedactionStub = nil
	if fake.imageSourceRedactionReturnsOnCall == nil {
		fake.imageSourceRedactionReturnsOnCall = make(map[int]struct {
			result1 atc.Source
			result2 error
		})
	}
	fake.imageSourceRedactionReturnsOnCall[i] = struct {
		result1 atc.Source
		result2 error
	}{result1, result2}
}

func (fake *FakeCheckDelegate) ImageVersionDetermined(arg1 db.UsedResourceCache) error {
	fake.imageVersionDeterminedMutex.Lock()
	ret, specificReturn := fake.imageVersionDeterminedReturnsOnCall[len(fake.imageVersionDeterminedArgsForCall)]
	fake.imageVersionDeterminedArgsForCall = append(fake.imageVersionDeterminedArgsForCall, struct {
		arg1 db.UsedResourceCache
	}{arg1})
	fake.recordInvocation("ImageVersionDetermined", []interface{}{arg1})
	fake.imageVersionDeterminedMutex.Unlock()
	if fake.ImageVersionDeterminedStub != nil {
		return fake.ImageVersionDeterminedStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.imageVersionDeterminedReturns
	return fakeReturns.result1
}

func (fake *FakeCheckDelegate) ImageVersionDeterminedCallCount() int {
	fake.imageVersionDeterminedMutex.RLock()
	defer fake.imageVersionDeterminedMutex.RUnlock()
	return len(fake.imageVersionDeterminedArgsForCall)
}

func (fake *FakeCheckDelegate) ImageVersionDeterminedCalls(stub func(db.UsedResourceCache) error) {
	fake.imageVersionDeterminedMutex.Lock()
	defer fake.imageVersionDeterminedMutex.Unlock()
	fake.ImageVersionDeterminedStub = stub
}

func (fake *FakeCheckDelegate) ImageVersionDeterminedArgsForCall(i int) db.UsedResourceCache {
	fake.imageVersionDeterminedMutex.RLock()
	defer fake.imageVersionDeterminedMutex.RUnlock()
	argsForCall := fake.imageVersionDeterminedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCheckDelegate) ImageVersionDeterminedReturns(result1 error) {
	fake.imageVersionDeterminedMutex.Lock()
	defer fake.imageVersionDeterminedMutex.Unlock()
	fake.ImageVersionDeterminedStub = nil
	fake.imageVersionDeterminedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCheckDelegate) ImageVersionDeterminedReturnsOnCall(i int, result1 error) {
	fake.imageVersionDeterminedMutex.Lock()
	defer fake.imageVersionDeterminedMutex.Unlock()
	fake.ImageVersionDeterminedStub = nil
	if fake.imageVersionDeterminedReturnsOnCall == nil {
		fake.imageVersionDeterminedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.imageVersionDeterminedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCheckDelegate) Initializing(arg1 lager.Logger) {
	fake.initializingMutex.Lock()
	fake.initializingArgsForCall = append(fake.initializingArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("Initializing", []interface{}{arg1})
	fake.initializingMutex.Unlock()
	if fake.InitializingStub != nil {
		fake.InitializingStub(arg1)
	}
}

func (fake *FakeCheckDelegate) InitializingCallCount() int {
	fake.initializingMutex.RLock()
	defer fake.initializingMutex.RUnlock()
	return len(fake.initializingArgsForCall)
}

func (fake *FakeCheckDelegate) InitializingCalls(stub func(lager.Logger)) {
	fake.initializingMutex.Lock()
	defer fake.initializingMutex.Unlock()
	fake.InitializingStub = stub
}

func (fake *FakeCheckDelegate) InitializingArgsForCall(i int) lager.Logger {
	fake.initializingMutex.RLock()
	defer fake.initializingMutex.RUnlock()
	argsForCall := fake.initializingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCheckDelegate) SaveVersions(arg1 db.SpanContext, arg2 []atc.Version) error {
	var arg2Copy []atc.Version
	if arg2 != nil {
		arg2Copy = make([]atc.Version, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.saveVersionsMutex.Lock()
	ret, specificReturn := fake.saveVersionsReturnsOnCall[len(fake.saveVersionsArgsForCall)]
	fake.saveVersionsArgsForCall = append(fake.saveVersionsArgsForCall, struct {
		arg1 db.SpanContext
		arg2 []atc.Version
	}{arg1, arg2Copy})
	fake.recordInvocation("SaveVersions", []interface{}{arg1, arg2Copy})
	fake.saveVersionsMutex.Unlock()
	if fake.SaveVersionsStub != nil {
		return fake.SaveVersionsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.saveVersionsReturns
	return fakeReturns.result1
}

func (fake *FakeCheckDelegate) SaveVersionsCallCount() int {
	fake.saveVersionsMutex.RLock()
	defer fake.saveVersionsMutex.RUnlock()
	return len(fake.saveVersionsArgsForCall)
}

func (fake *FakeCheckDelegate) SaveVersionsCalls(stub func(db.SpanContext, []atc.Version) error) {
	fake.saveVersionsMutex.Lock()
	defer fake.saveVersionsMutex.Unlock()
	fake.SaveVersionsStub = stub
}

func (fake *FakeCheckDelegate) SaveVersionsArgsForCall(i int) (db.SpanContext, []atc.Version) {
	fake.saveVersionsMutex.RLock()
	defer fake.saveVersionsMutex.RUnlock()
	argsForCall := fake.saveVersionsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCheckDelegate) SaveVersionsReturns(result1 error) {
	fake.saveVersionsMutex.Lock()
	defer fake.saveVersionsMutex.Unlock()
	fake.SaveVersionsStub = nil
	fake.saveVersionsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCheckDelegate) SaveVersionsReturnsOnCall(i int, result1 error) {
	fake.saveVersionsMutex.Lock()
	defer fake.saveVersionsMutex.Unlock()
	fake.SaveVersionsStub = nil
	if fake.saveVersionsReturnsOnCall == nil {
		fake.saveVersionsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.saveVersionsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCheckDelegate) Starting(arg1 lager.Logger) {
	fake.startingMutex.Lock()
	fake.startingArgsForCall = append(fake.startingArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("Starting", []interface{}{arg1})
	fake.startingMutex.Unlock()
	if fake.StartingStub != nil {
		fake.StartingStub(arg1)
	}
}

func (fake *FakeCheckDelegate) StartingCallCount() int {
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	return len(fake.startingArgsForCall)
}

func (fake *FakeCheckDelegate) StartingCalls(stub func(lager.Logger)) {
	fake.startingMutex.Lock()
	defer fake.startingMutex.Unlock()
	fake.StartingStub = stub
}

func (fake *FakeCheckDelegate) StartingArgsForCall(i int) lager.Logger {
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	argsForCall := fake.startingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCheckDelegate) Stderr() io.Writer {
	fake.stderrMutex.Lock()
	ret, specificReturn := fake.stderrReturnsOnCall[len(fake.stderrArgsForCall)]
	fake.stderrArgsForCall = append(fake.stderrArgsForCall, struct {
	}{})
	fake.recordInvocation("Stderr", []interface{}{})
	fake.stderrMutex.Unlock()
	if fake.StderrStub != nil {
		return fake.StderrStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.stderrReturns
	return fakeReturns.result1
}

func (fake *FakeCheckDelegate) StderrCallCount() int {
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	return len(fake.stderrArgsForCall)
}

func (fake *FakeCheckDelegate) StderrCalls(stub func() io.Writer) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = stub
}

func (fake *FakeCheckDelegate) StderrReturns(result1 io.Writer) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = nil
	fake.stderrReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeCheckDelegate) StderrReturnsOnCall(i int, result1 io.Writer) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = nil
	if fake.stderrReturnsOnCall == nil {
		fake.stderrReturnsOnCall = make(map[int]struct {
			result1 io.Writer
		})
	}
	fake.stderrReturnsOnCall[i] = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeCheckDelegate) Stdout() io.Writer {
	fake.stdoutMutex.Lock()
	ret, specificReturn := fake.stdoutReturnsOnCall[len(fake.stdoutArgsForCall)]
	fake.stdoutArgsForCall = append(fake.stdoutArgsForCall, struct {
	}{})
	fake.recordInvocation("Stdout", []interface{}{})
	fake.stdoutMutex.Unlock()
	if fake.StdoutStub != nil {
		return fake.StdoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.stdoutReturns
	return fakeReturns.result1
}

func (fake *FakeCheckDelegate) StdoutCallCount() int {
	fake.stdoutMutex.RLock()
	defer fake.stdoutMutex.RUnlock()
	return len(fake.stdoutArgsForCall)
}

func (fake *FakeCheckDelegate) StdoutCalls(stub func() io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = stub
}

func (fake *FakeCheckDelegate) StdoutReturns(result1 io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = nil
	fake.stdoutReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeCheckDelegate) StdoutReturnsOnCall(i int, result1 io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = nil
	if fake.stdoutReturnsOnCall == nil {
		fake.stdoutReturnsOnCall = make(map[int]struct {
			result1 io.Writer
		})
	}
	fake.stdoutReturnsOnCall[i] = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeCheckDelegate) Variables() vars.CredVarsTracker {
	fake.variablesMutex.Lock()
	ret, specificReturn := fake.variablesReturnsOnCall[len(fake.variablesArgsForCall)]
	fake.variablesArgsForCall = append(fake.variablesArgsForCall, struct {
	}{})
	fake.recordInvocation("Variables", []interface{}{})
	fake.variablesMutex.Unlock()
	if fake.VariablesStub != nil {
		return fake.VariablesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.variablesReturns
	return fakeReturns.result1
}

func (fake *FakeCheckDelegate) VariablesCallCount() int {
	fake.variablesMutex.RLock()
	defer fake.variablesMutex.RUnlock()
	return len(fake.variablesArgsForCall)
}

func (fake *FakeCheckDelegate) VariablesCalls(stub func() vars.CredVarsTracker) {
	fake.variablesMutex.Lock()
	defer fake.variablesMutex.Unlock()
	fake.VariablesStub = stub
}

func (fake *FakeCheckDelegate) VariablesReturns(result1 vars.CredVarsTracker) {
	fake.variablesMutex.Lock()
	defer fake.variablesMutex.Unlock()
	fake.VariablesStub = nil
	fake.variablesReturns = struct {
		result1 vars.CredVarsTracker
	}{result1}
}

func (fake *FakeCheckDelegate) VariablesReturnsOnCall(i int, result1 vars.CredVarsTracker) {
	fake.variablesMutex.Lock()
	defer fake.variablesMutex.Unlock()
	fake.VariablesStub = nil
	if fake.variablesReturnsOnCall == nil {
		fake.variablesReturnsOnCall = make(map[int]struct {
			result1 vars.CredVarsTracker
		})
	}
	fake.variablesReturnsOnCall[i] = struct {
		result1 vars.CredVarsTracker
	}{result1}
}

func (fake *FakeCheckDelegate) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.erroredMutex.RLock()
	defer fake.erroredMutex.RUnlock()
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	fake.imageSourceRedactionMutex.RLock()
	defer fake.imageSourceRedactionMutex.RUnlock()
	fake.imageVersionDeterminedMutex.RLock()
	defer fake.imageVersionDeterminedMutex.RUnlock()
	fake.initializingMutex.RLock()
	defer fake.initializingMutex.RUnlock()
	fake.saveVersionsMutex.RLock()
	defer fake.saveVersionsMutex.RUnlock()
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	fake.stdoutMutex.RLock()
	defer fake.stdoutMutex.RUnlock()
	fake.variablesMutex.RLock()
	defer fake.variablesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCheckDelegate) recordInvocation(key string, args []interface{}) {
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

var _ exec.CheckDelegate = new(FakeCheckDelegate)
