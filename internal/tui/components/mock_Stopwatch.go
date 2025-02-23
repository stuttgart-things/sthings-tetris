// Code generated by mockery v2.50.1. DO NOT EDIT.

package components

import (
	tea "github.com/charmbracelet/bubbletea"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// MockStopwatch is an autogenerated mock type for the Stopwatch type
type MockStopwatch struct {
	mock.Mock
}

type MockStopwatch_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStopwatch) EXPECT() *MockStopwatch_Expecter {
	return &MockStopwatch_Expecter{mock: &_m.Mock}
}

// Elapsed provides a mock function with no fields
func (_m *MockStopwatch) Elapsed() time.Duration {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Elapsed")
	}

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// MockStopwatch_Elapsed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Elapsed'
type MockStopwatch_Elapsed_Call struct {
	*mock.Call
}

// Elapsed is a helper method to define mock.On call
func (_e *MockStopwatch_Expecter) Elapsed() *MockStopwatch_Elapsed_Call {
	return &MockStopwatch_Elapsed_Call{Call: _e.mock.On("Elapsed")}
}

func (_c *MockStopwatch_Elapsed_Call) Run(run func()) *MockStopwatch_Elapsed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStopwatch_Elapsed_Call) Return(_a0 time.Duration) *MockStopwatch_Elapsed_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStopwatch_Elapsed_Call) RunAndReturn(run func() time.Duration) *MockStopwatch_Elapsed_Call {
	_c.Call.Return(run)
	return _c
}

// ID provides a mock function with no fields
func (_m *MockStopwatch) ID() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ID")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockStopwatch_ID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ID'
type MockStopwatch_ID_Call struct {
	*mock.Call
}

// ID is a helper method to define mock.On call
func (_e *MockStopwatch_Expecter) ID() *MockStopwatch_ID_Call {
	return &MockStopwatch_ID_Call{Call: _e.mock.On("ID")}
}

func (_c *MockStopwatch_ID_Call) Run(run func()) *MockStopwatch_ID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStopwatch_ID_Call) Return(_a0 int) *MockStopwatch_ID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStopwatch_ID_Call) RunAndReturn(run func() int) *MockStopwatch_ID_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with no fields
func (_m *MockStopwatch) Init() tea.Cmd {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Init")
	}

	var r0 tea.Cmd
	if rf, ok := ret.Get(0).(func() tea.Cmd); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(tea.Cmd)
		}
	}

	return r0
}

// MockStopwatch_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockStopwatch_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
func (_e *MockStopwatch_Expecter) Init() *MockStopwatch_Init_Call {
	return &MockStopwatch_Init_Call{Call: _e.mock.On("Init")}
}

func (_c *MockStopwatch_Init_Call) Run(run func()) *MockStopwatch_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStopwatch_Init_Call) Return(_a0 tea.Cmd) *MockStopwatch_Init_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStopwatch_Init_Call) RunAndReturn(run func() tea.Cmd) *MockStopwatch_Init_Call {
	_c.Call.Return(run)
	return _c
}

// Reset provides a mock function with no fields
func (_m *MockStopwatch) Reset() tea.Cmd {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Reset")
	}

	var r0 tea.Cmd
	if rf, ok := ret.Get(0).(func() tea.Cmd); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(tea.Cmd)
		}
	}

	return r0
}

// MockStopwatch_Reset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reset'
type MockStopwatch_Reset_Call struct {
	*mock.Call
}

// Reset is a helper method to define mock.On call
func (_e *MockStopwatch_Expecter) Reset() *MockStopwatch_Reset_Call {
	return &MockStopwatch_Reset_Call{Call: _e.mock.On("Reset")}
}

func (_c *MockStopwatch_Reset_Call) Run(run func()) *MockStopwatch_Reset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStopwatch_Reset_Call) Return(_a0 tea.Cmd) *MockStopwatch_Reset_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStopwatch_Reset_Call) RunAndReturn(run func() tea.Cmd) *MockStopwatch_Reset_Call {
	_c.Call.Return(run)
	return _c
}

// SetInterval provides a mock function with given fields: _a0
func (_m *MockStopwatch) SetInterval(_a0 time.Duration) {
	_m.Called(_a0)
}

// MockStopwatch_SetInterval_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetInterval'
type MockStopwatch_SetInterval_Call struct {
	*mock.Call
}

// SetInterval is a helper method to define mock.On call
//   - _a0 time.Duration
func (_e *MockStopwatch_Expecter) SetInterval(_a0 interface{}) *MockStopwatch_SetInterval_Call {
	return &MockStopwatch_SetInterval_Call{Call: _e.mock.On("SetInterval", _a0)}
}

func (_c *MockStopwatch_SetInterval_Call) Run(run func(_a0 time.Duration)) *MockStopwatch_SetInterval_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Duration))
	})
	return _c
}

func (_c *MockStopwatch_SetInterval_Call) Return() *MockStopwatch_SetInterval_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStopwatch_SetInterval_Call) RunAndReturn(run func(time.Duration)) *MockStopwatch_SetInterval_Call {
	_c.Run(run)
	return _c
}

// Stop provides a mock function with no fields
func (_m *MockStopwatch) Stop() tea.Cmd {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Stop")
	}

	var r0 tea.Cmd
	if rf, ok := ret.Get(0).(func() tea.Cmd); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(tea.Cmd)
		}
	}

	return r0
}

// MockStopwatch_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockStopwatch_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockStopwatch_Expecter) Stop() *MockStopwatch_Stop_Call {
	return &MockStopwatch_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockStopwatch_Stop_Call) Run(run func()) *MockStopwatch_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStopwatch_Stop_Call) Return(_a0 tea.Cmd) *MockStopwatch_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStopwatch_Stop_Call) RunAndReturn(run func() tea.Cmd) *MockStopwatch_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// Toggle provides a mock function with no fields
func (_m *MockStopwatch) Toggle() tea.Cmd {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Toggle")
	}

	var r0 tea.Cmd
	if rf, ok := ret.Get(0).(func() tea.Cmd); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(tea.Cmd)
		}
	}

	return r0
}

// MockStopwatch_Toggle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Toggle'
type MockStopwatch_Toggle_Call struct {
	*mock.Call
}

// Toggle is a helper method to define mock.On call
func (_e *MockStopwatch_Expecter) Toggle() *MockStopwatch_Toggle_Call {
	return &MockStopwatch_Toggle_Call{Call: _e.mock.On("Toggle")}
}

func (_c *MockStopwatch_Toggle_Call) Run(run func()) *MockStopwatch_Toggle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStopwatch_Toggle_Call) Return(_a0 tea.Cmd) *MockStopwatch_Toggle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStopwatch_Toggle_Call) RunAndReturn(run func() tea.Cmd) *MockStopwatch_Toggle_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0
func (_m *MockStopwatch) Update(_a0 tea.Msg) (tea.Model, tea.Cmd) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 tea.Model
	var r1 tea.Cmd
	if rf, ok := ret.Get(0).(func(tea.Msg) (tea.Model, tea.Cmd)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(tea.Msg) tea.Model); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(tea.Model)
		}
	}

	if rf, ok := ret.Get(1).(func(tea.Msg) tea.Cmd); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(tea.Cmd)
		}
	}

	return r0, r1
}

// MockStopwatch_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockStopwatch_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 tea.Msg
func (_e *MockStopwatch_Expecter) Update(_a0 interface{}) *MockStopwatch_Update_Call {
	return &MockStopwatch_Update_Call{Call: _e.mock.On("Update", _a0)}
}

func (_c *MockStopwatch_Update_Call) Run(run func(_a0 tea.Msg)) *MockStopwatch_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(tea.Msg))
	})
	return _c
}

func (_c *MockStopwatch_Update_Call) Return(_a0 tea.Model, _a1 tea.Cmd) *MockStopwatch_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStopwatch_Update_Call) RunAndReturn(run func(tea.Msg) (tea.Model, tea.Cmd)) *MockStopwatch_Update_Call {
	_c.Call.Return(run)
	return _c
}

// View provides a mock function with no fields
func (_m *MockStopwatch) View() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for View")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockStopwatch_View_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'View'
type MockStopwatch_View_Call struct {
	*mock.Call
}

// View is a helper method to define mock.On call
func (_e *MockStopwatch_Expecter) View() *MockStopwatch_View_Call {
	return &MockStopwatch_View_Call{Call: _e.mock.On("View")}
}

func (_c *MockStopwatch_View_Call) Run(run func()) *MockStopwatch_View_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStopwatch_View_Call) Return(_a0 string) *MockStopwatch_View_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStopwatch_View_Call) RunAndReturn(run func() string) *MockStopwatch_View_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStopwatch creates a new instance of MockStopwatch. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStopwatch(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStopwatch {
	mock := &MockStopwatch{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
