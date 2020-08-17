package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/vladimirdotk/news-bot/internal/domain"
)

// MessageHandlerMock implements telegram.MessageHandler
type MessageHandlerMock struct {
	t minimock.Tester

	funcHandle          func(message *domain.IncomingMessage) (err error)
	inspectFuncHandle   func(message *domain.IncomingMessage)
	afterHandleCounter  uint64
	beforeHandleCounter uint64
	HandleMock          mMessageHandlerMockHandle
}

// NewMessageHandlerMock returns a mock for telegram.MessageHandler
func NewMessageHandlerMock(t minimock.Tester) *MessageHandlerMock {
	m := &MessageHandlerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.HandleMock = mMessageHandlerMockHandle{mock: m}
	m.HandleMock.callArgs = []*MessageHandlerMockHandleParams{}

	return m
}

type mMessageHandlerMockHandle struct {
	mock               *MessageHandlerMock
	defaultExpectation *MessageHandlerMockHandleExpectation
	expectations       []*MessageHandlerMockHandleExpectation

	callArgs []*MessageHandlerMockHandleParams
	mutex    sync.RWMutex
}

// MessageHandlerMockHandleExpectation specifies expectation struct of the MessageHandler.Handle
type MessageHandlerMockHandleExpectation struct {
	mock    *MessageHandlerMock
	params  *MessageHandlerMockHandleParams
	results *MessageHandlerMockHandleResults
	Counter uint64
}

// MessageHandlerMockHandleParams contains parameters of the MessageHandler.Handle
type MessageHandlerMockHandleParams struct {
	message *domain.IncomingMessage
}

// MessageHandlerMockHandleResults contains results of the MessageHandler.Handle
type MessageHandlerMockHandleResults struct {
	err error
}

// Expect sets up expected params for MessageHandler.Handle
func (mmHandle *mMessageHandlerMockHandle) Expect(message *domain.IncomingMessage) *mMessageHandlerMockHandle {
	if mmHandle.mock.funcHandle != nil {
		mmHandle.mock.t.Fatalf("MessageHandlerMock.Handle mock is already set by Set")
	}

	if mmHandle.defaultExpectation == nil {
		mmHandle.defaultExpectation = &MessageHandlerMockHandleExpectation{}
	}

	mmHandle.defaultExpectation.params = &MessageHandlerMockHandleParams{message}
	for _, e := range mmHandle.expectations {
		if minimock.Equal(e.params, mmHandle.defaultExpectation.params) {
			mmHandle.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmHandle.defaultExpectation.params)
		}
	}

	return mmHandle
}

// Inspect accepts an inspector function that has same arguments as the MessageHandler.Handle
func (mmHandle *mMessageHandlerMockHandle) Inspect(f func(message *domain.IncomingMessage)) *mMessageHandlerMockHandle {
	if mmHandle.mock.inspectFuncHandle != nil {
		mmHandle.mock.t.Fatalf("Inspect function is already set for MessageHandlerMock.Handle")
	}

	mmHandle.mock.inspectFuncHandle = f

	return mmHandle
}

// Return sets up results that will be returned by MessageHandler.Handle
func (mmHandle *mMessageHandlerMockHandle) Return(err error) *MessageHandlerMock {
	if mmHandle.mock.funcHandle != nil {
		mmHandle.mock.t.Fatalf("MessageHandlerMock.Handle mock is already set by Set")
	}

	if mmHandle.defaultExpectation == nil {
		mmHandle.defaultExpectation = &MessageHandlerMockHandleExpectation{mock: mmHandle.mock}
	}
	mmHandle.defaultExpectation.results = &MessageHandlerMockHandleResults{err}
	return mmHandle.mock
}

//Set uses given function f to mock the MessageHandler.Handle method
func (mmHandle *mMessageHandlerMockHandle) Set(f func(message *domain.IncomingMessage) (err error)) *MessageHandlerMock {
	if mmHandle.defaultExpectation != nil {
		mmHandle.mock.t.Fatalf("Default expectation is already set for the MessageHandler.Handle method")
	}

	if len(mmHandle.expectations) > 0 {
		mmHandle.mock.t.Fatalf("Some expectations are already set for the MessageHandler.Handle method")
	}

	mmHandle.mock.funcHandle = f
	return mmHandle.mock
}

// When sets expectation for the MessageHandler.Handle which will trigger the result defined by the following
// Then helper
func (mmHandle *mMessageHandlerMockHandle) When(message *domain.IncomingMessage) *MessageHandlerMockHandleExpectation {
	if mmHandle.mock.funcHandle != nil {
		mmHandle.mock.t.Fatalf("MessageHandlerMock.Handle mock is already set by Set")
	}

	expectation := &MessageHandlerMockHandleExpectation{
		mock:   mmHandle.mock,
		params: &MessageHandlerMockHandleParams{message},
	}
	mmHandle.expectations = append(mmHandle.expectations, expectation)
	return expectation
}

// Then sets up MessageHandler.Handle return parameters for the expectation previously defined by the When method
func (e *MessageHandlerMockHandleExpectation) Then(err error) *MessageHandlerMock {
	e.results = &MessageHandlerMockHandleResults{err}
	return e.mock
}

// Handle implements telegram.MessageHandler
func (mmHandle *MessageHandlerMock) Handle(message *domain.IncomingMessage) (err error) {
	mm_atomic.AddUint64(&mmHandle.beforeHandleCounter, 1)
	defer mm_atomic.AddUint64(&mmHandle.afterHandleCounter, 1)

	if mmHandle.inspectFuncHandle != nil {
		mmHandle.inspectFuncHandle(message)
	}

	mm_params := &MessageHandlerMockHandleParams{message}

	// Record call args
	mmHandle.HandleMock.mutex.Lock()
	mmHandle.HandleMock.callArgs = append(mmHandle.HandleMock.callArgs, mm_params)
	mmHandle.HandleMock.mutex.Unlock()

	for _, e := range mmHandle.HandleMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmHandle.HandleMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmHandle.HandleMock.defaultExpectation.Counter, 1)
		mm_want := mmHandle.HandleMock.defaultExpectation.params
		mm_got := MessageHandlerMockHandleParams{message}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmHandle.t.Errorf("MessageHandlerMock.Handle got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmHandle.HandleMock.defaultExpectation.results
		if mm_results == nil {
			mmHandle.t.Fatal("No results are set for the MessageHandlerMock.Handle")
		}
		return (*mm_results).err
	}
	if mmHandle.funcHandle != nil {
		return mmHandle.funcHandle(message)
	}
	mmHandle.t.Fatalf("Unexpected call to MessageHandlerMock.Handle. %v", message)
	return
}

// HandleAfterCounter returns a count of finished MessageHandlerMock.Handle invocations
func (mmHandle *MessageHandlerMock) HandleAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmHandle.afterHandleCounter)
}

// HandleBeforeCounter returns a count of MessageHandlerMock.Handle invocations
func (mmHandle *MessageHandlerMock) HandleBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmHandle.beforeHandleCounter)
}

// Calls returns a list of arguments used in each call to MessageHandlerMock.Handle.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmHandle *mMessageHandlerMockHandle) Calls() []*MessageHandlerMockHandleParams {
	mmHandle.mutex.RLock()

	argCopy := make([]*MessageHandlerMockHandleParams, len(mmHandle.callArgs))
	copy(argCopy, mmHandle.callArgs)

	mmHandle.mutex.RUnlock()

	return argCopy
}

// MinimockHandleDone returns true if the count of the Handle invocations corresponds
// the number of defined expectations
func (m *MessageHandlerMock) MinimockHandleDone() bool {
	for _, e := range m.HandleMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.HandleMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterHandleCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcHandle != nil && mm_atomic.LoadUint64(&m.afterHandleCounter) < 1 {
		return false
	}
	return true
}

// MinimockHandleInspect logs each unmet expectation
func (m *MessageHandlerMock) MinimockHandleInspect() {
	for _, e := range m.HandleMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MessageHandlerMock.Handle with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.HandleMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterHandleCounter) < 1 {
		if m.HandleMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MessageHandlerMock.Handle")
		} else {
			m.t.Errorf("Expected call to MessageHandlerMock.Handle with params: %#v", *m.HandleMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcHandle != nil && mm_atomic.LoadUint64(&m.afterHandleCounter) < 1 {
		m.t.Error("Expected call to MessageHandlerMock.Handle")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MessageHandlerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockHandleInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MessageHandlerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *MessageHandlerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockHandleDone()
}
