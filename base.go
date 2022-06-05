package nom

type IResult[P any] struct {
	notParsed string
	produced P
}

type ParseFn[P any] func(input string)(*IResult[P],error)


