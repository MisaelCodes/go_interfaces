package intermediate_level

import "fmt"

func DescribeType(value interface{}) {
	valueType := fmt.Sprintf("%T", value)
	switch valueType {
	case "int":
		fmt.Println("This is an int")
	case "string":
		fmt.Println("This is a string")
	default:
		fmt.Println("Unknown type")
	}
}
