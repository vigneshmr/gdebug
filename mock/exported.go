package mock

/*

- This file contains assertion parallels that are found in testify/mock
  to enable packages to be swapped.
- Methods that are overridden here will just passthough call into testify/require

*/

import (
	"github.com/stretchr/testify/mock"
)

type AnythingOfTypeArgument struct {
	mock.AnythingOfTypeArgument
}

type Arguments struct {
	*mock.Arguments
}

type Call struct {
	*mock.Call
}

type Mock struct {
	*mock.Mock
}

func (m *Mock) On(methodName string, arguments ...interface{}) *mock.Call {
	return m.Mock.On(methodName, arguments)
}

type TestingT struct {
	mock.TestingT
}

var Anything = mock.Anything

func AnythingOfType(t string) AnythingOfTypeArgument {
	res := mock.AnythingOfType(t)
	return AnythingOfTypeArgument{res}
}

func AssertExpectationsForObjects(t TestingT, testObjects ...interface{}) bool {
	return mock.AssertExpectationsForObjects(t, testObjects)
}
