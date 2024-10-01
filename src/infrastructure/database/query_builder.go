package database

import (
	"reflect"
)

type QueryBuild struct {
	Tabela string
}

func (query QueryBuild) ColumnsAndValuesCount(v any) (string, string) {
	fields := reflect.ValueOf(v)
	selecto := ""
	count := ""
	declaredFieldsSize := fields.NumField()

	for i := 0; i < declaredFieldsSize; i++ {
		fieldValue := fields.Type().Field(i).Tag
		valor := fieldValue.Get("column")
		if valor != "" {
			selecto = selecto + valor
			count = count + "?"

			if i != declaredFieldsSize-1 {
				selecto = selecto + ", "
				count = count + ", "
			}

		}

	}
	return selecto, count
}
