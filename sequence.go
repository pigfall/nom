package nom


func Delimited(first,second,third ParseFn)ParseFn{
	return func(input string)(IResult,error){
		result,err := first(input)
		if err != nil{
			return result,err
		}
		result,err =second(result.notParsed)
		if err != nil{
			return result,err
		}
		resultB,err :=third(result.notParsed)
		if err != nil{
			return resultB,err
		}
		return IResult{produced:result.produced,notParsed:resultB.notParsed},nil
	}
}
