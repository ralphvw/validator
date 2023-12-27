package validator

import "reflect"

func Validate(obj interface{}, fields []string) (bool, []string) {
	objValue := reflect.ValueOf(obj)
	var missingFields []string

	for _, field := range fields {
		fieldValue := objValue.FieldByName(field)

		if !fieldValue.IsValid() {
			missingFields = append(missingFields, field)
		} else {
			zeroValue := reflect.Zero(fieldValue.Type())
			if reflect.DeepEqual(fieldValue.Interface(), zeroValue.Interface()) {
				missingFields = append(missingFields, field)
			}
		}
	}

	fieldsExist := len(missingFields) == 0

	return fieldsExist, missingFields
}
