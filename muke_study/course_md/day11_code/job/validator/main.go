package main

import (
	"fmt"
	"reflect"
)

type MyNested struct {
	Email string `validate:"empty"`
}
type TT struct {
	Age      int `validate:"empty"`
	MyNested MyNested
}

func validateEmpty(v interface{}) (bool, string) {
	validateResult := true
	errMsg := "success"
	vt := reflect.TypeOf(v)
	vv := reflect.ValueOf(v)
	for i := 0; i < vt.NumField(); i++ {
		fieldVal := vv.Field(i)
		tagContent := vt.Field(i).Tag.Get("validate")
		field := vt.Field(i).Name
		k := fieldVal.Kind()
/*		fmt.Println(tagContent)
		fmt.Println(fieldVal)
		fmt.Println(k)
*/
		switch k {
		case reflect.Int:
			val := fieldVal.Int()
			tagValStr := tagContent
			switch tagValStr {
			case "empty":
				if val == 0 {
					errMsg = fmt.Sprintf("validate field " + field + " empty failed, field val is: %v", val)
					validateResult = false
				}
			}
		case reflect.String:
			val := fieldVal.String()
			tagValStr := tagContent
			switch tagValStr {
			case "empty":
				if val == "" {
					errMsg = "validate field " + field + " empty failed, field val is: " + val
					validateResult = false
				}
			}
		case reflect.Struct:
			// 如果有内嵌的 struct，那么深度优先遍历
			// 就是一个递归过程
			valInter := fieldVal.Interface()
			MyNestedResult, msg := validateEmpty(valInter)
			if MyNestedResult == false {
				validateResult = false
				errMsg = msg
			}
		}
	}
	return validateResult, errMsg
}

func main() {
	var a = TT{Age: 0, MyNested: MyNested{Email: "abc@abc.com"}}

	validateResult, errMsg := validateEmpty(a)
	fmt.Println(validateResult, errMsg)
}
