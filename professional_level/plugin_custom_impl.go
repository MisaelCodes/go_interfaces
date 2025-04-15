package professional_level

import "strings"

type Plugin interface {
	Execute(data string) string
}

type UpperCasePlugin struct{}

func (ucp *UpperCasePlugin) Execute(data string) string {
	return strings.ToUpper(data)
}

type ReversePluging struct{}

func (rp *ReversePluging) Execute(data string) string {
	newData := []byte(data)
	i := 0
	for i < len(data)-(i+1) {
		newData[i], newData[len(newData)-(i+1)] = newData[len(newData)-(i+1)], newData[i]
		i++
	}
	return string(newData)
}
