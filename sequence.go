package nom


func Delimited[T any](first ParseFn[string],second ParseFn[T],third ParseFn[string])ParseFn[T]{
	return func(input string)(*IResult[T],error){
		r,err := first(input)
		if err != nil{
			return nil,err
		}
		result,err :=second(r.notParsed)
		if err != nil{
			return result,err
		}
		resultB,err :=third(result.notParsed)
		if err != nil{
			return nil,err
		}
		return &IResult[T]{produced:result.produced,notParsed:resultB.notParsed},nil
	}
}
