package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// QueueServiceMock implements command.QueueService
type QueueServiceMock struct {
	t minimock.Tester

	funcPublish          func(topic string, data interface{}) (err error)
	inspectFuncPublish   func(topic string, data interface{})
	afterPublishCounter  uint64
	beforePublishCounter uint64
	PublishMock          mQueueServiceMockPublish
}

// NewQueueServiceMock returns a mock for command.QueueService
func NewQueueServiceMock(t minimock.Tester) *QueueServiceMock {
	m := &QueueServiceMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.PublishMock = mQueueServiceMockPublish{mock: m}
	m.PublishMock.callArgs = []*QueueServiceMockPublishParams{}

	return m
}

type mQueueServiceMockPublish struct {
	mock               *QueueServiceMock
	defaultExpectation *QueueServiceMockPublishExpectation
	expectations       []*QueueServiceMockPublishExpectation

	callArgs []*QueueServiceMockPublishParams
	mutex    sync.RWMutex
}

// QueueServiceMockPublishExpectation specifies expectation struct of the QueueService.Publish
type QueueServiceMockPublishExpectation struct {
	mock    *QueueServiceMock
	params  *QueueServiceMockPublishParams
	results *QueueServiceMockPublishResults
	Counter uint64
}

// QueueServiceMockPublishParams contains parameters of the QueueService.Publish
type QueueServiceMockPublishParams struct {
	topic string
	data  interface{}
}

// QueueServiceMockPublishResults contains results of the QueueService.Publish
type QueueServiceMockPublishResults struct {
	err error
}

// Expect sets up expected params for QueueService.Publish
func (mmPublish *mQueueServiceMockPublish) Expect(topic string, data interface{}) *mQueueServiceMockPublish {
	if mmPublish.mock.funcPublish != nil {
		mmPublish.mock.t.Fatalf("QueueServiceMock.Publish mock is already set by Set")
	}

	if mmPublish.defaultExpectation == nil {
		mmPublish.defaultExpectation = &QueueServiceMockPublishExpectation{}
	}

	mmPublish.defaultExpectation.params = &QueueServiceMockPublishParams{topic, data}
	for _, e := range mmPublish.expectations {
		if minimock.Equal(e.params, mmPublish.defaultExpectation.params) {
			mmPublish.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmPublish.defaultExpectation.params)
		}
	}

	return mmPublish
}

// Inspect accepts an inspector function that has same arguments as the QueueService.Publish
func (mmPublish *mQueueServiceMockPublish) Inspect(f func(topic string, data interface{})) *mQueueServiceMockPublish {
	if mmPublish.mock.inspectFuncPublish != nil {
		mmPublish.mock.t.Fatalf("Inspect function is already set for QueueServiceMock.Publish")
	}

	mmPublish.mock.inspectFuncPublish = f

	return mmPublish
}

// Return sets up results that will be returned by QueueService.Publish
func (mmPublish *mQueueServiceMockPublish) Return(err error) *QueueServiceMock {
	if mmPublish.mock.funcPublish != nil {
		mmPublish.mock.t.Fatalf("QueueServiceMock.Publish mock is already set by Set")
	}

	if mmPublish.defaultExpectation == nil {
		mmPublish.defaultExpectation = &QueueServiceMockPublishExpectation{mock: mmPublish.mock}
	}
	mmPublish.defaultExpectation.results = &QueueServiceMockPublishResults{err}
	return mmPublish.mock
}

//Set uses given function f to mock the QueueService.Publish method
func (mmPublish *mQueueServiceMockPublish) Set(f func(topic string, data interface{}) (err error)) *QueueServiceMock {
	if mmPublish.defaultExpectation != nil {
		mmPublish.mock.t.Fatalf("Default expectation is already set for the QueueService.Publish method")
	}

	if len(mmPublish.expectations) > 0 {
		mmPublish.mock.t.Fatalf("Some expectations are already set for the QueueService.Publish method")
	}

	mmPublish.mock.funcPublish = f
	return mmPublish.mock
}

// When sets expectation for the QueueService.Publish which will trigger the result defined by the following
// Then helper
func (mmPublish *mQueueServiceMockPublish) When(topic string, data interface{}) *QueueServiceMockPublishExpectation {
	if mmPublish.mock.funcPublish != nil {
		mmPublish.mock.t.Fatalf("QueueServiceMock.Publish mock is already set by Set")
	}

	expectation := &QueueServiceMockPublishExpectation{
		mock:   mmPublish.mock,
		params: &QueueServiceMockPublishParams{topic, data},
	}
	mmPublish.expectations = append(mmPublish.expectations, expectation)
	return expectation
}

// Then sets up QueueService.Publish return parameters for the expectation previously defined by the When method
func (e *QueueServiceMockPublishExpectation) Then(err error) *QueueServiceMock {
	e.results = &QueueServiceMockPublishResults{err}
	return e.mock
}

// Publish implements command.QueueService
func (mmPublish *QueueServiceMock) Publish(topic string, data interface{}) (err error) {
	mm_atomic.AddUint64(&mmPublish.beforePublishCounter, 1)
	defer mm_atomic.AddUint64(&mmPublish.afterPublishCounter, 1)

	if mmPublish.inspectFuncPublish != nil {
		mmPublish.inspectFuncPublish(topic, data)
	}

	mm_params := &QueueServiceMockPublishParams{topic, data}

	// Record call args
	mmPublish.PublishMock.mutex.Lock()
	mmPublish.PublishMock.callArgs = append(mmPublish.PublishMock.callArgs, mm_params)
	mmPublish.PublishMock.mutex.Unlock()

	for _, e := range mmPublish.PublishMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmPublish.PublishMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmPublish.PublishMock.defaultExpectation.Counter, 1)
		mm_want := mmPublish.PublishMock.defaultExpectation.params
		mm_got := QueueServiceMockPublishParams{topic, data}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmPublish.t.Errorf("QueueServiceMock.Publish got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmPublish.PublishMock.defaultExpectation.results
		if mm_results == nil {
			mmPublish.t.Fatal("No results are set for the QueueServiceMock.Publish")
		}
		return (*mm_results).err
	}
	if mmPublish.funcPublish != nil {
		return mmPublish.funcPublish(topic, data)
	}
	mmPublish.t.Fatalf("Unexpected call to QueueServiceMock.Publish. %v %v", topic, data)
	return
}

// PublishAfterCounter returns a count of finished QueueServiceMock.Publish invocations
func (mmPublish *QueueServiceMock) PublishAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPublish.afterPublishCounter)
}

// PublishBeforeCounter returns a count of QueueServiceMock.Publish invocations
func (mmPublish *QueueServiceMock) PublishBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPublish.beforePublishCounter)
}

// Calls returns a list of arguments used in each call to QueueServiceMock.Publish.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmPublish *mQueueServiceMockPublish) Calls() []*QueueServiceMockPublishParams {
	mmPublish.mutex.RLock()

	argCopy := make([]*QueueServiceMockPublishParams, len(mmPublish.callArgs))
	copy(argCopy, mmPublish.callArgs)

	mmPublish.mutex.RUnlock()

	return argCopy
}

// MinimockPublishDone returns true if the count of the Publish invocations corresponds
// the number of defined expectations
func (m *QueueServiceMock) MinimockPublishDone() bool {
	for _, e := range m.PublishMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PublishMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPublishCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPublish != nil && mm_atomic.LoadUint64(&m.afterPublishCounter) < 1 {
		return false
	}
	return true
}

// MinimockPublishInspect logs each unmet expectation
func (m *QueueServiceMock) MinimockPublishInspect() {
	for _, e := range m.PublishMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to QueueServiceMock.Publish with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PublishMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPublishCounter) < 1 {
		if m.PublishMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to QueueServiceMock.Publish")
		} else {
			m.t.Errorf("Expected call to QueueServiceMock.Publish with params: %#v", *m.PublishMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPublish != nil && mm_atomic.LoadUint64(&m.afterPublishCounter) < 1 {
		m.t.Error("Expected call to QueueServiceMock.Publish")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *QueueServiceMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockPublishInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *QueueServiceMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *QueueServiceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockPublishDone()
}
