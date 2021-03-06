package javaum

import (
	"reflect"

	"github.com/m-pavel/gojos/lib/javaos"
)

func Unmarshal(model *javaos.JavaModel, dest interface{}) error {
	rv := reflect.ValueOf(dest)
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}

	um := unmarshallerFor(rv.Type())
	um.Unmarshal(rv, model)

	return nil
}
