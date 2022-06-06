package nom

import(
	"log"
)

func init(){
	log.SetFlags(log.Llongfile)
}


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

func Preceded[O1 any,O2 any](first ParseFn[O1],second ParseFn[O2])ParseFn[O2]{
	return func(input string)(*IResult[O2],error){
		res,err :=first(input)
		if err != nil{
			return nil,err
		}
		return second(res.notParsed)
	}
}

func Terminated[O1,O2 any](first ParseFn[O1],second ParseFn[O2])ParseFn[O1]{
	return func(input string)(*IResult[O1],error){
		res,err := first(input)
		if err != nil{
			return nil,err
		}
		r,err := second(res.notParsed)
		if err != nil{
			return nil,err
		}
		return &IResult[O1]{
			notParsed:r.notParsed,
			produced:res.produced,
		},nil
	}

}

type Pair[O1,O2 any] struct{
	Left O1
	Right O2
}

func SeparatedPair[O1,O2,O3 any](
	first ParseFn[O1],
	sep ParseFn[O2],
	second ParseFn[O3],
)ParseFn[Pair[O1,O3]]{
	return func(input string)(*IResult[Pair[O1,O3]],error){
		log.Println("pair ",input)
		r1,err := first(input)
		if err != nil{
			log.Println(err)
			return nil,err
		}
		log.Println("r1 notParsed ",r1.notParsed)
		r2,err := sep(r1.notParsed)
		if err != nil{
			log.Println(err)
			return nil,err
		}
		log.Println("pair2 success",input)
		r3,err :=second(r2.notParsed)
		if err != nil{
			return nil,err
		}
		log.Println("pair success",input)
		return &IResult[Pair[O1,O3]]{
			produced:Pair[O1,O3]{Left:r1.produced,Right:r3.produced},
			notParsed: r3.notParsed,
		},nil
	}
}
