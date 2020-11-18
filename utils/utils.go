package utils

import (
	"log"
	"reflect"
	"strings"
)

func ExtractOplogUpdatedFields(oplogUpdatedFields map[string]interface{}) []string {
	/*
		Allows to extract and format the mongodb fields and subfields of the map containing the fields which have received an update

		oplogUpdatedFields := map[string]interface{}{"field1": "value1", "field2.6": map[string]interface{}{"subfield21": "value21", "subfield22": "value22"}}

		// map[field1:value1 field2.6:map[subfield22:value22 subfield21:value21]]

		ExtractOplogUpdatedFields(oplogUpdatedFields)
		// [field1 field2.subfield21 field2.subfield22]

	*/
	var updatedFields []string

	ExtractSubFields(oplogUpdatedFields, &updatedFields, "")

	return RemoveDuplicates(updatedFields)
}

func ExtractSubFields(inputMap map[string]interface{}, updatedFields *[]string, parentKey string) {
	for k, v := range inputMap {
		var subKey string
		if parentKey == "" {
			subKey = strings.Split(k, ".")[0]
		} else {
			subKey = parentKey + "." + strings.Split(k, ".")[0]
		}
		if reflect.TypeOf(v).Kind() == reflect.Map {
			ExtractSubFields(v.(map[string]interface{}), updatedFields, subKey)
		} else {
			*updatedFields = append(*updatedFields, subKey)
		}
	}

}

func RemoveDuplicates(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
