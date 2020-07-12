package functions

import (
	"github.com/programmer-richa/utility/validators"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
)

//ConvertToMap converts a struct to bson.M
// This is used for updating the data
// In this function removes empty fields and id from the struct as per user choice
func ConvertToMap(model interface{}, removeID bool, removeEmptyField bool) bson.M {
	ret := bson.M{}

	modelReflect := reflect.ValueOf(model)

	if modelReflect.Kind() == reflect.Ptr {
		modelReflect = modelReflect.Elem()
	}

	modelRefType := modelReflect.Type()
	fieldsCount := modelReflect.NumField()

	var fieldData interface{}

	for i := 0; i < fieldsCount; i++ {
		field := modelReflect.Field(i)

		switch field.Kind() {
		case reflect.Struct:
			fallthrough
		case reflect.Ptr:
			fieldData = ConvertToMap(field.Interface(), removeID, removeEmptyField)
		default:
			fieldData = field.Interface()
		}
		if (removeEmptyField && !validators.NotBlank(fieldData)) || (removeID && modelRefType.Field(i).Name == "Id") {
			continue
		}
		ret[modelRefType.Field(i).Name] = fieldData
	}

	return ret
}
