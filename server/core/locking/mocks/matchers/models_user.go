// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"

	"github.com/petergtz/pegomock"

	models "github.com/runatlantis/atlantis/server/events/models"
)

func AnyModelsUser() models.User {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(models.User))(nil)).Elem()))
	var nullValue models.User
	return nullValue
}

func EqModelsUser(value models.User) models.User {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue models.User
	return nullValue
}

func NotEqModelsUser(value models.User) models.User {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue models.User
	return nullValue
}

func ModelsUserThat(matcher pegomock.ArgumentMatcher) models.User {
	pegomock.RegisterMatcher(matcher)
	var nullValue models.User
	return nullValue
}
