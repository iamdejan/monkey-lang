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
				if len(elements) <= 0 {
					return Null
				}
				return elements[0]
			default:
				return newError("argument to `first` method is not supported. actual=`%T`", args[0].Type())
			}
		},
	},
	"last": {
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
				l := len(str)
				return &object.String{Value: string(arg.Value[l-1])}
			case *object.Array:
				elements := arg.Elements
				l := len(elements)
				if l <= 0 {
					return Null
				}
				return elements[l-1]
			default:
				return newError("argument to `last` method is not supported. actual=`%T`", args[0].Type())
			}
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. expected=`2`, actual=`%d`", len(args))
			}

			if args[0].Type() != object.ArrayObj {
				return newError("first argument to `push` method is not supported. expected=`%s`, actual=`%s`", object.ArrayObj, args[0].Type())
			}

			elem := args[1]
			arr, _ := args[0].(*object.Array)
			length := len(arr.Elements)
			if length == 0 {
				newArr := make([]object.Object, 1)
				newArr[0] = elem
				return &object.Array{Elements: newArr}
			}

			newArr := make([]object.Object, length+1)
			copy(newArr, arr.Elements)
			newArr[length] = elem
			return &object.Array{Elements: newArr}
		},
	},
}
