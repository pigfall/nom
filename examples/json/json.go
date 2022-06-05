package main

import (
	"fmt"
"github.com/pigfall/nom"
)

func main() {
	s := `{"t":"tzz"}`
	fmt.Println(root(s))
}

func root(input string)(json,error){
	res,err := nom.Delimited(
		nom.Space0(),
		nom.Alt(
			[]nom.ParseFn[json]{
			 hash,
			},
		),
		nom.Space0(),
	)(input)
	if err != nil{
		return nil,err
	}
	return res.Produced(),nil
}

func hash(input string)(*nom.IResult[json],error){
	return nom.Preceded(
			nom.Tag("{"),
			nom.Terminated(
				nom.Map(
					nom.SeparatedList0(
						nom.Preceded(nom.Space0(),nom.Tag(",")),
						keyValueParse,
					),
					func(kvs []*keyValue)json{
						m := make(map[string]interface{})
						for _,kv := range kvs{
							m[kv.key] = kv.value
						}
						return m
					}),
				nom.Preceded(nom.Space0(),nom.Tag("}")),
			),
	)(input)
}

func keyValueParse(input string)(*nom.IResult[*keyValue],error){
	 return nom.Map(
		nom.SeparatedPair(
		nom.Preceded(nom.Space0(),str),
		nom.Preceded(nom.Space0(),nom.Tag(":")),
		nom.Preceded(nom.Space0(),str),
	),	
	func(p nom.Pair[string,string])*keyValue{
		return &keyValue{
			key:p.Left,
			value:p.Right,
		}
	},
	)(input) 
}

func str(input string)(*nom.IResult[string],error){
	return nom.Preceded(
		nom.Tag("\""),
		nom.Terminated(
			nom.Alpha1(),
			nom.Tag("\""),
		),
	)(input)
}

type keyValue struct{
	key string
	value interface{}
}

type json interface{}
