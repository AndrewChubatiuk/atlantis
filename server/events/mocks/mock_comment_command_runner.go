// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: CommentCommandRunner)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	command "github.com/runatlantis/atlantis/server/events/command"
	"reflect"
	"time"
)

type MockCommentCommandRunner struct {
	fail func(message string, callerSkip ...int)
}

func NewMockCommentCommandRunner(options ...pegomock.Option) *MockCommentCommandRunner {
	mock := &MockCommentCommandRunner{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockCommentCommandRunner) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockCommentCommandRunner) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockCommentCommandRunner) Run(_param0 *command.Context, _param1 *command.Comment) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommentCommandRunner().")
	}
	params := []pegomock.Param{_param0, _param1}
	pegomock.GetGenericMockFrom(mock).Invoke("Run", params, []reflect.Type{})
}

func (mock *MockCommentCommandRunner) VerifyWasCalledOnce() *VerifierMockCommentCommandRunner {
	return &VerifierMockCommentCommandRunner{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockCommentCommandRunner) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockCommentCommandRunner {
	return &VerifierMockCommentCommandRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockCommentCommandRunner) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockCommentCommandRunner {
	return &VerifierMockCommentCommandRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockCommentCommandRunner) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockCommentCommandRunner {
	return &VerifierMockCommentCommandRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockCommentCommandRunner struct {
	mock                   *MockCommentCommandRunner
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockCommentCommandRunner) Run(_param0 *command.Context, _param1 *command.Comment) *MockCommentCommandRunner_Run_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Run", params, verifier.timeout)
	return &MockCommentCommandRunner_Run_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommentCommandRunner_Run_OngoingVerification struct {
	mock              *MockCommentCommandRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommentCommandRunner_Run_OngoingVerification) GetCapturedArguments() (*command.Context, *command.Comment) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockCommentCommandRunner_Run_OngoingVerification) GetAllCapturedArguments() (_param0 []*command.Context, _param1 []*command.Comment) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*command.Context, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(*command.Context)
		}
		_param1 = make([]*command.Comment, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(*command.Comment)
		}
	}
	return
}
