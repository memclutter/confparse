package confparse

import (
	"flag"
	"reflect"
	"strconv"
	"time"
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

	case *bool:
		if v, err := toBool(value); err != nil {
			return err
		} else {
			flag.BoolVar(ptr, name, v, usage)
		}

	// String argument type
	case *string:
		flag.StringVar(ptr, name, value, usage)

	// Integer argument type
	case *int:
		if v, err := toInt(value); err != nil {
			return err
		} else {
			flag.IntVar(ptr, name, v, usage)
		}

	// Time duration argument type
	case *time.Duration:
		if v, err := toTimeDuration(value); err != nil {
			return err
		} else {
			flag.DurationVar(ptr, name, v, usage)
		}

	}

	return nil
}

// From string value to boolean
func toBool(value string) (result bool, err error) {
	if value != "" {
		result, err = strconv.ParseBool(value)
	}

	return
}

// From string value to integer value
func toInt(value string) (result int, err error) {
	if value != "" {
		result, err = strconv.Atoi(value)
	}

	return
}

// From string value to time duration
func toTimeDuration(value string) (result time.Duration, err error) {
	if value != "" {
		result, err = time.ParseDuration(value)
	}

	return
}
