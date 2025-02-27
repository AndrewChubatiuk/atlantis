// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events/webhooks (interfaces: Sender)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	webhooks "github.com/runatlantis/atlantis/server/events/webhooks"
	logging "github.com/runatlantis/atlantis/server/logging"
	"reflect"
	"time"
)

type MockSender struct {
	fail func(message string, callerSkip ...int)
}

func NewMockSender(options ...pegomock.Option) *MockSender {
	mock := &MockSender{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockSender) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockSender) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockSender) Send(log logging.Logger, applyResult webhooks.ApplyResult) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSender().")
	}
	params := []pegomock.Param{log, applyResult}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Send", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockSender) VerifyWasCalledOnce() *VerifierMockSender {
	return &VerifierMockSender{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockSender) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockSender {
	return &VerifierMockSender{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockSender) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockSender {
	return &VerifierMockSender{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockSender) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockSender {
	return &VerifierMockSender{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockSender struct {
	mock                   *MockSender
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockSender) Send(log logging.Logger, applyResult webhooks.ApplyResult) *MockSender_Send_OngoingVerification {
	params := []pegomock.Param{log, applyResult}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Send", params, verifier.timeout)
	return &MockSender_Send_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSender_Send_OngoingVerification struct {
	mock              *MockSender
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSender_Send_OngoingVerification) GetCapturedArguments() (logging.Logger, webhooks.ApplyResult) {
	log, applyResult := c.GetAllCapturedArguments()
	return log[len(log)-1], applyResult[len(applyResult)-1]
}

func (c *MockSender_Send_OngoingVerification) GetAllCapturedArguments() (_param0 []logging.Logger, _param1 []webhooks.ApplyResult) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]logging.Logger, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(logging.Logger)
		}
		_param1 = make([]webhooks.ApplyResult, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(webhooks.ApplyResult)
		}
	}
	return
}
