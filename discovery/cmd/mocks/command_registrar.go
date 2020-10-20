// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import common "github.com/tw-bc-group/fabric-gm/cmd/common"

import kingpin "gopkg.in/alecthomas/kingpin.v2"
import mock "github.com/stretchr/testify/mock"

// CommandRegistrar is an autogenerated mock type for the CommandRegistrar type
type CommandRegistrar struct {
	mock.Mock
}

// Command provides a mock function with given fields: name, help, onCommand
func (_m *CommandRegistrar) Command(name string, help string, onCommand common.CLICommand) *kingpin.CmdClause {
	ret := _m.Called(name, help, onCommand)

	var r0 *kingpin.CmdClause
	if rf, ok := ret.Get(0).(func(string, string, common.CLICommand) *kingpin.CmdClause); ok {
		r0 = rf(name, help, onCommand)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kingpin.CmdClause)
		}
	}

	return r0
}
