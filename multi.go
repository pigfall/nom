package nom

func SeparatedList0[O any,O2 any](separator ParseFn[O2],f ParseFn[O])ParseFn[[]O]{
	return func(input string)(*IResult[[]O],error){
		list := make([]O,0)
		for {
			res,err := f(input)
			if err != nil {
				return &IResult[[]O]{notParsed:input,produced:list},nil
			}
			input = res.notParsed
			list = append(list,res.produced)
			resB,err := separator(input)
			if err != nil{
				return &IResult[[]O]{notParsed:input,produced:list},nil
			}
			input = resB.notParsed
		}
	}

}
