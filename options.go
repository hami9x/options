//Package options implements a simple, flexible and convenient method for specifying options,
//with the help of reflections.
package options

import (
	"fmt"
	"reflect"
)

//Spec specifies the list of options.
//A spec should be implemented as a struct which have Option fields,
//each representing an option.
//Each Option must be a different struct, multiple fields with the same struct
//type will cause unwanted behavior (only the last field receives the value).
type Spec interface{}

//Option represents an option.
//Each option must be a different struct, and should be defined
//as a struct with a Value field.
//
//For example:
//	type RouteName struct { Value: string }
type Option interface{}

type OptionsProvider struct {
	spec Spec
	m    map[string]string
	set  map[string]bool //the values has been set or not
}

//NewOptions creates a new OptionsProvider.
func NewOptions(spec Spec) *OptionsProvider {
	o := &OptionsProvider{spec, make(map[string]string), make(map[string]bool)}
	v := reflect.ValueOf(spec)
	if v.CanSet() {
		panic(fmt.Sprintf("The spec passed in must be a pointer, got type %v",
			v.Type().Name()))
	}
	v = v.Elem()
	specType := v.Type()
	for i := 0; i < specType.NumField(); i++ {
		field := specType.Field(i)
		o.m[field.Type.Name()] = field.Name
	}
	return o
}

//Options set the options in spec according to the opts passed in.
func (o *OptionsProvider) Options(opts ...Option) *OptionsProvider {
	for _, opt := range opts {
		optType := reflect.TypeOf(opt)
		if _, ok := optType.FieldByName("Value"); !ok {
			panic(fmt.Sprintf("Option %v doesn't have a Value field.", optType.Name()))
		}
		fieldName := o.m[optType.Name()]
		field := reflect.ValueOf(o.spec).Elem().FieldByName(fieldName)
		if !field.CanSet() || !field.IsValid() {
			panic(fmt.Sprintf("There is no option %v.", optType.Name()))
		}
		field.Set(reflect.ValueOf(opt))
		o.set[fieldName] = true
	}
	return o
}

//Get returns the options data, which could be casted
//to the original spec type and the value could be get from there.
func (o *OptionsProvider) Get() interface{} {
	return o.spec
}

//Check if the field is set
func (o *OptionsProvider) IsSet(field string) bool {
	return o.set[field]
}

//ExportToMap exports the options to a map
func (o *OptionsProvider) ExportToMap() map[string]interface{} {
	return o.ExportToMapWithTag("")
}

//ExportToMapWithTag exports the options to a map,
//the key names in the map are determined by a specific struct tag.
//The tags are set in the options spec.
//Configs that have not been set don't appear in the map.
func (o *OptionsProvider) ExportToMapWithTag(tag string) map[string]interface{} {
	spec := reflect.ValueOf(o.spec).Elem()
	specType := spec.Type()
	m := make(map[string]interface{})
	for i := 0; i < specType.NumField(); i++ {
		exportedName := specType.Field(i).Name
		if !o.IsSet(exportedName) {
			continue
		}
		if tag != "" {
			taggedName := specType.Field(i).Tag.Get(tag)
			if taggedName != "" {
				exportedName = taggedName
			}
		}
		m[exportedName] = spec.Field(i).FieldByName("Value").Interface()
	}
	return m
}
