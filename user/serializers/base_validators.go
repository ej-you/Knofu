package serializers

import (
	"fmt"
	"reflect"
	"strings"

	validate "github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
)


func baseValidator(givenStruct validate.Validator, errors *validate.Errors) {
	// Получаем значение структуры (с разыменовыванием поинтера через Elem())
	structValue := reflect.ValueOf(givenStruct).Elem()
	// получаем кол-во полей в структуре
	structNumFields := structValue.NumField()

	var field reflect.StructField
	var tag string
	// перебираем поля и проверяем теги структуры
	for i:=0; i < structNumFields; i++ {
		field = structValue.Type().Field(i)
		tag = field.Tag.Get("validate")
		if tag != "" {
			for _, tagValue := range strings.Split(tag, "|") {
				switch tagValue {
					case "required":
						if structValue.Field(i).String() == "" {
							errors.Add(field.Tag.Get("json"), fmt.Sprintf("%s must not be blank!", field.Name))
						}
					case "email":
						errors.Append(validate.Validate(
							&validators.EmailIsPresent{
								Name: "Email",
								Field: structValue.Field(i).String(),
								Message: "Mail is not in the right format.",
							},
						))
				}
			}
		}		
	}
}
