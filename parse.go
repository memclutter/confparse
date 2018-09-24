package confparse

import (
	"flag"
	"reflect"
)

// Parse config container. A container is a pointer to a structure with tags.
func Parse(container interface{}) error {

	val, typ := determine(container)

	// Find all fields and read tag
	for i := 0; i < typ.NumField(); i++ {

		name, value, usage := extractTags(typ.Field(i).Tag)
		addr := val.Field(i).Addr().Interface()

		if err := declareFlag(name, value, usage, addr); err != nil {
			return err
		}
	}

	// Parse all defined arguments
	flag.Parse()

	return nil
}

// Determine container value and type
func determine(container interface{}) (val reflect.Value, typ reflect.Type) {
	val = reflect.ValueOf(container).Elem()
	typ = val.Type()
	return
}

// Extract container field tags values
func extractTags(tags reflect.StructTag) (name, value, help string) {
	name = tags.Get("name")
	value = tags.Get("value")
	help = tags.Get("usage")
	return
}

// Declare CLI argument
func declareFlag(name, value, usage string, addr interface{}) error {

	switch ptr := addr.(type) {
	case *string:
		flag.StringVar(ptr, name, value, usage)
	}

	return nil
}
