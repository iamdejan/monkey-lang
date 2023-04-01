package evaluator

import "monkey/object"

var builtIns = map[string]*object.BuiltIn{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong argument count for `len` function. expected=`1`, actual=`%d`", len(args))
			}

			arg := args[0]
			if arg.Type() != object.StringObj {
				return newError("wrong argument type for `len` function. expected=`%s`, actual=`%s`", object.StringObj, arg.Type())
			}

			str, ok := arg.(*object.String)
			if !ok {
				return newError("parse error for argument")
			}
			return &object.Integer{Value: int64(len(str.Value))}
		},
	},
}
