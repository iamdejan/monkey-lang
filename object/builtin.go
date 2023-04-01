package object

const BuiltInObj = "BUILT_IN"

type BuiltInFunction func(args ...Object) Object

type BuiltIn struct {
	Fn BuiltInFunction
}

func (b *BuiltIn) Type() ObjectType {
	return BuiltInObj
}
func (b *BuiltIn) Inspect() string {
	return "built-in function"
}
