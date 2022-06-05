package nom

type IResult[P any] struct {
	notParsed string
	produced P
}

func (this *IResult[P]) NotParsed()string{
	return this.notParsed
}

func (this *IResult[P]) Produced()P{
	return this.produced
}

type ParseFn[P any] func(input string)(*IResult[P],error)


