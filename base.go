package nom

type IResult struct {
	notParsed string
	produced string
}

type ParseFn func(input string)(IResult,error)


