// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"github.com/petergtz/pegomock"
	valid "github.com/runatlantis/atlantis/server/core/config/valid"
	"reflect"
)

func AnySliceOfValidStep() []valid.Step {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*([]valid.Step))(nil)).Elem()))
	var nullValue []valid.Step
	return nullValue
}

func EqSliceOfValidStep(value []valid.Step) []valid.Step {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue []valid.Step
	return nullValue
}
