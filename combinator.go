package nom

func Map[O any,O2 any](parser ParseFn[O],f func(O)(O2))ParseFn[O2]{
	return func(input string)(*IResult[O2],error){
		res,err := parser(input)
		if err != nil{
			return nil,err
		}
		out := f(res.produced)
		return &IResult[O2]{
			notParsed:res.notParsed,
			produced:out,
		},nil
	}
}
