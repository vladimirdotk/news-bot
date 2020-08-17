package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/vladimirdotk/news-bot/internal/domain"
)

// CommandExecutorMock implements redis.CommandExecutor
type CommandExecutorMock struct {
	t minimock.Tester

	funcExec          func(message domain.IncomingMessage) (err error)
	inspectFuncExec   func(message domain.IncomingMessage)
	afterExecCounter  uint64
	beforeExecCounter uint64
	ExecMock          mCommandExecutorMockExec
}

// NewCommandExecutorMock returns a mock for redis.CommandExecutor
func NewCommandExecutorMock(t minimock.Tester) *CommandExecutorMock {
	m := &CommandExecutorMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ExecMock = mCommandExecutorMockExec{mock: m}
	m.ExecMock.callArgs = []*CommandExecutorMockExecParams{}

	return m
}

type mCommandExecutorMockExec struct {
	mock               *CommandExecutorMock
	defaultExpectation *CommandExecutorMockExecExpectation
	expectations       []*CommandExecutorMockExecExpectation

	callArgs []*CommandExecutorMockExecParams
	mutex    sync.RWMutex
}

// CommandExecutorMockExecExpectation specifies expectation struct of the CommandExecutor.Exec
type CommandExecutorMockExecExpectation struct {
	mock    *CommandExecutorMock
	params  *CommandExecutorMockExecParams
	results *CommandExecutorMockExecResults
	Counter uint64
}

// CommandExecutorMockExecParams contains parameters of the CommandExecutor.Exec
type CommandExecutorMockExecParams struct {
	message domain.IncomingMessage
}

// CommandExecutorMockExecResults contains results of the CommandExecutor.Exec
type CommandExecutorMockExecResults struct {
	err error
}

// Expect sets up expected params for CommandExecutor.Exec
func (mmExec *mCommandExecutorMockExec) Expect(message domain.IncomingMessage) *mCommandExecutorMockExec {
	if mmExec.mock.funcExec != nil {
		mmExec.mock.t.Fatalf("CommandExecutorMock.Exec mock is already set by Set")
	}

	if mmExec.defaultExpectation == nil {
		mmExec.defaultExpectation = &CommandExecutorMockExecExpectation{}
	}

	mmExec.defaultExpectation.params = &CommandExecutorMockExecParams{message}
	for _, e := range mmExec.expectations {
		if minimock.Equal(e.params, mmExec.defaultExpectation.params) {
			mmExec.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmExec.defaultExpectation.params)
		}
	}

	return mmExec
}

// Inspect accepts an inspector function that has same arguments as the CommandExecutor.Exec
func (mmExec *mCommandExecutorMockExec) Inspect(f func(message domain.IncomingMessage)) *mCommandExecutorMockExec {
	if mmExec.mock.inspectFuncExec != nil {
		mmExec.mock.t.Fatalf("Inspect function is already set for CommandExecutorMock.Exec")
	}

	mmExec.mock.inspectFuncExec = f

	return mmExec
}

// Return sets up results that will be returned by CommandExecutor.Exec
func (mmExec *mCommandExecutorMockExec) Return(err error) *CommandExecutorMock {
	if mmExec.mock.funcExec != nil {
		mmExec.mock.t.Fatalf("CommandExecutorMock.Exec mock is already set by Set")
	}

	if mmExec.defaultExpectation == nil {
		mmExec.defaultExpectation = &CommandExecutorMockExecExpectation{mock: mmExec.mock}
	}
	mmExec.defaultExpectation.results = &CommandExecutorMockExecResults{err}
	return mmExec.mock
}

//Set uses given function f to mock the CommandExecutor.Exec method
func (mmExec *mCommandExecutorMockExec) Set(f func(message domain.IncomingMessage) (err error)) *CommandExecutorMock {
	if mmExec.defaultExpectation != nil {
		mmExec.mock.t.Fatalf("Default expectation is already set for the CommandExecutor.Exec method")
	}

	if len(mmExec.expectations) > 0 {
		mmExec.mock.t.Fatalf("Some expectations are already set for the CommandExecutor.Exec method")
	}

	mmExec.mock.funcExec = f
	return mmExec.mock
}

// When sets expectation for the CommandExecutor.Exec which will trigger the result defined by the following
// Then helper
func (mmExec *mCommandExecutorMockExec) When(message domain.IncomingMessage) *CommandExecutorMockExecExpectation {
	if mmExec.mock.funcExec != nil {
		mmExec.mock.t.Fatalf("CommandExecutorMock.Exec mock is already set by Set")
	}

	expectation := &CommandExecutorMockExecExpectation{
		mock:   mmExec.mock,
		params: &CommandExecutorMockExecParams{message},
	}
	mmExec.expectations = append(mmExec.expectations, expectation)
	return expectation
}

// Then sets up CommandExecutor.Exec return parameters for the expectation previously defined by the When method
func (e *CommandExecutorMockExecExpectation) Then(err error) *CommandExecutorMock {
	e.results = &CommandExecutorMockExecResults{err}
	return e.mock
}

// Exec implements redis.CommandExecutor
func (mmExec *CommandExecutorMock) Exec(message domain.IncomingMessage) (err error) {
	mm_atomic.AddUint64(&mmExec.beforeExecCounter, 1)
	defer mm_atomic.AddUint64(&mmExec.afterExecCounter, 1)

	if mmExec.inspectFuncExec != nil {
		mmExec.inspectFuncExec(message)
	}

	mm_params := &CommandExecutorMockExecParams{message}

	// Record call args
	mmExec.ExecMock.mutex.Lock()
	mmExec.ExecMock.callArgs = append(mmExec.ExecMock.callArgs, mm_params)
	mmExec.ExecMock.mutex.Unlock()

	for _, e := range mmExec.ExecMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmExec.ExecMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmExec.ExecMock.defaultExpectation.Counter, 1)
		mm_want := mmExec.ExecMock.defaultExpectation.params
		mm_got := CommandExecutorMockExecParams{message}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmExec.t.Errorf("CommandExecutorMock.Exec got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmExec.ExecMock.defaultExpectation.results
		if mm_results == nil {
			mmExec.t.Fatal("No results are set for the CommandExecutorMock.Exec")
		}
		return (*mm_results).err
	}
	if mmExec.funcExec != nil {
		return mmExec.funcExec(message)
	}
	mmExec.t.Fatalf("Unexpected call to CommandExecutorMock.Exec. %v", message)
	return
}

// ExecAfterCounter returns a count of finished CommandExecutorMock.Exec invocations
func (mmExec *CommandExecutorMock) ExecAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmExec.afterExecCounter)
}

// ExecBeforeCounter returns a count of CommandExecutorMock.Exec invocations
func (mmExec *CommandExecutorMock) ExecBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmExec.beforeExecCounter)
}

// Calls returns a list of arguments used in each call to CommandExecutorMock.Exec.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmExec *mCommandExecutorMockExec) Calls() []*CommandExecutorMockExecParams {
	mmExec.mutex.RLock()

	argCopy := make([]*CommandExecutorMockExecParams, len(mmExec.callArgs))
	copy(argCopy, mmExec.callArgs)

	mmExec.mutex.RUnlock()

	return argCopy
}

// MinimockExecDone returns true if the count of the Exec invocations corresponds
// the number of defined expectations
func (m *CommandExecutorMock) MinimockExecDone() bool {
	for _, e := range m.ExecMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ExecMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterExecCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcExec != nil && mm_atomic.LoadUint64(&m.afterExecCounter) < 1 {
		return false
	}
	return true
}

// MinimockExecInspect logs each unmet expectation
func (m *CommandExecutorMock) MinimockExecInspect() {
	for _, e := range m.ExecMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CommandExecutorMock.Exec with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ExecMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterExecCounter) < 1 {
		if m.ExecMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CommandExecutorMock.Exec")
		} else {
			m.t.Errorf("Expected call to CommandExecutorMock.Exec with params: %#v", *m.ExecMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcExec != nil && mm_atomic.LoadUint64(&m.afterExecCounter) < 1 {
		m.t.Error("Expected call to CommandExecutorMock.Exec")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *CommandExecutorMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockExecInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *CommandExecutorMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *CommandExecutorMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockExecDone()
}