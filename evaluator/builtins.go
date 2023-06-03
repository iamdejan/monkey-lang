package evaluator

import "monkey/object"

var builtIns = map[string]*object.BuiltIn{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong argument count for `len` function. expected=`1`, actual=`%d`", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` method is not supported. actual=`%s`", args[0].Type())
			}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong argument count for `first` function. expected=`1`, actual=`%d`", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				str := arg.Value
				if str == "" {
					return Null
				}
				return &object.String{Value: string(arg.Value[0])}
			case *object.Array:
				elements := arg.Elements
				if len(elements) == 0 {
					return Null
				}
				return elements[0]
			default:
				return newError("argument to `first` method is not supported. actual=`%T`", args[0].Type())
			}
		},
	},
}
