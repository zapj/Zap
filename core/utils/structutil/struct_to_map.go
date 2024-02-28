package structutil

import (
	"reflect"
)

func Convert(struc interface{}) map[string]interface{} {

	returnMap := make(map[string]interface{})

	sType := getStructType(struc)

	if sType.Kind() != reflect.Struct {
		return returnMap
	}

	for i := 0; i < sType.NumField(); i++ {
		structFieldName := sType.Field(i).Name
		structVal := reflect.ValueOf(struc)
		returnMap[structFieldName] = structVal.FieldByName(structFieldName).Interface()
	}

	return returnMap
}

func getStructType(struc interface{}) reflect.Type {
	sType := reflect.TypeOf(struc)
	if sType.Kind() == reflect.Ptr {
		sType = sType.Elem()
	}

	return sType
}
