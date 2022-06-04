package nom

import(
	"fmt"
	"strings"
	"unicode"
)

func Tag(tag string)ParseFn{
	return func(input string)(IResult,error){
		if strings.HasPrefix(input,tag){
			return IResult{
				notParsed:input[len(tag):],
				produced:tag,
			},nil
		}

		return IResult{},parseErr(fmt.Sprintf("tag [%s]not match",tag)).notMatch().build()
	}
}

func Alpha0()ParseFn{
	return func(input string)(IResult,error){
		if len(input) == 0{
			return IResult{
				notParsed:input,
			},nil
		}

		for i,r := range []rune(input){
			if unicode.IsLetter(r){
				continue
			}else{
				return IResult{
					notParsed:string([]rune(input)[i:]),
					produced:string([]rune(input)[:i]),
				},nil
			}
		}

		return IResult{
			notParsed:"",
			produced:input,
		},nil
	}
}

func Alpha1()ParseFn{
	return func(input string)(IResult,error){
		i,err :=Alpha0()(input)
		if err != nil{
			return IResult{},err
		}
		if len(i.produced) == 0{
			return IResult{notParsed:input},parseErr("need at least one alpha").notMatch().build()
		}
		return i,nil
	}
}

func AlphaNumeric0()ParseFn{
	return func(input string)(IResult,error){
		if len(input) == 0{
			return IResult{
				notParsed:input,
			},nil
		}

		for i,r := range []rune(input) {
			if unicode.IsLetter(r) || unicode.IsNumber(r){
				continue
			}else{
				return IResult{
					notParsed:string([]rune(input)[i:]),
					produced:string([]rune(input)[:i]),
				},nil
			}
		}

		return IResult{
			notParsed:"",
			produced:input,
		},nil
	}
}


func AlphaNumeric1()ParseFn{
	return func(input string)(IResult,error){
		i,err :=AlphaNumeric0()(input)
		if err != nil{
			return IResult{},err
		}
		if len(i.produced) == 0{
			return IResult{notParsed:input},parseErr("need at least one alpha or number").notMatch().build()
		}
		return i,nil
	}
}

func TakeWhile(predicate func(r rune)bool)ParseFn{
	return func(input string)(IResult,error){
		for i,r := range []rune(input){
			if !predicate(r){
				return IResult{
					notParsed:input[i:],
					produced: input[:i],
				},nil
			}
		}

		return IResult{
			notParsed:"",
			produced: input,
		},nil
	}
}

func TakeWhile1(predicate func(r rune)bool)ParseFn{
	return func(input string)(IResult,error){
		if len(input) == 0{
			return IResult{},parseErr("take whilte not match").notMatch().build()
		}
		for i,r := range []rune(input){
			if !predicate(r){
				if i == 0{
					return IResult{},parseErr("take whilte not match").notMatch().build()
				}
				return IResult{
					notParsed:input[i:],
					produced: input[:i],
				},nil
			}
		}

		return IResult{
			notParsed:"",
			produced: input,
		},nil
	}
}

