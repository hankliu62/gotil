package gotil

// SliceSplice delete or insert slice item
func SliceSplice(values []interface{}, index, howmany int, insert ...interface{}) []interface{} {
	var newValues []interface{}

	if index >= len(values) {
		newValues = values
	} else {
		newValues = append(values[:index], values[index+howmany:]...)
	}

	if insert != nil && len(insert) > 0 {
		rear := append([]interface{}{}, newValues[index:]...)
		newValues = append(newValues[0:index], insert...)
		newValues = append(newValues, rear...)
	}

	return newValues
}
