package object

const ReturnValueObj = "RETURN_VALUE"

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType {
	return ReturnValueObj
}
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}
