package object

import (
	"bytes"
	"strings"
)

const ArrayObj = "ARRAY"

type Array struct {
	Elements []Object
}

func (ao *Array) Type() ObjectType {
	return ArrayObj
}
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
