package object

const StringObj = "STRING"

type String struct {
	Value string
}

func (s *String) Type() ObjectType {
	return StringObj
}
func (s *String) Inspect() string {
	return s.Value
}
