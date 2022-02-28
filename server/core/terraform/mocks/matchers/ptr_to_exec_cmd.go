// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"

	"github.com/petergtz/pegomock"

	exec "os/exec"
)

func AnyPtrToExecCmd() *exec.Cmd {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*exec.Cmd))(nil)).Elem()))
	var nullValue *exec.Cmd
	return nullValue
}

func EqPtrToExecCmd(value *exec.Cmd) *exec.Cmd {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *exec.Cmd
	return nullValue
}

func NotEqPtrToExecCmd(value *exec.Cmd) *exec.Cmd {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue *exec.Cmd
	return nullValue
}

func PtrToExecCmdThat(matcher pegomock.ArgumentMatcher) *exec.Cmd {
	pegomock.RegisterMatcher(matcher)
	var nullValue *exec.Cmd
	return nullValue
}
