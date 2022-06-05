package nom

import(
	"log"
)

func Alt[T any](parsers []ParseFn[T])ParseFn[T]{
	return func(i string)(*IResult[T],error){
		for _,p := range parsers{
			res,err := p(i)
			if err == nil{
				return res,nil
			}
			log.Println(err)
		}
		return nil,parseErr("all alt failed").notMatch().build()
	}
}
