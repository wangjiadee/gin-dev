package service

import (
	"fmt"
	"reflect"

	"gopkg.in/go-playground/validator.v8"
)

// "gopkg.in/go-playground/validator.v9"

// func TopicUrl(v validator.FieldLevel) bool {
// 	if _, ok := v.Top().Interface().(*Topic); ok {
// 		// 正则匹配
// 		if matched, _ := regexp.MatchString(`\w[4,10]$`, v.Field().String()); matched {
// 			return true
// 		}
// 		fmt.Println(v.Field().String())
// 		//return str != "invalid"
// 	}
// 	return false
// }

func TopicUrl(v *validator.Validate, topStruct reflect.Value, current reflect.Value, filed reflect.Value, fType reflect.Type, filedKind reflect.Kind, param string) bool {
	fmt.Println(topStruct)
	fmt.Println(topStruct.Interface())
	return true
}
