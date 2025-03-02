package utils

import "fmt"

func PrintOutput(value Value) {
	_type := value.Typ
	switch _type {
	case "array":
		fmt.Println("array", value.Array)
	case "bulk":
		fmt.Println(value.Bulk)
	case "string":
		fmt.Println(value.Str)
	default:
		fmt.Println("value", value)
	}
}
