package nom

import(
	"fmt"
	"strings"
	"unicode"
	"log"
)

// The input data will be compared to tag combinators's argument and will return the part of the input that matches the argument
func Tag(tag string)ParseFn[string]{
	return func(input string)(*IResult[string],error){
		log.Println("tag input ",input)
		if strings.HasPrefix(input,tag){
			return &IResult[string]{
				notParsed:input[len(tag):],
				produced:tag,
			},nil
		}

		return &IResult[string]{},parseErr(fmt.Sprintf("tag [%s] not match,input is %s",tag,input)).notMatch().build()
	}
}

// Recognize zero or more lowcase and uppercase ASCII alphabetic characters
func Alpha0()ParseFn[string]{
	return func(input string)(*IResult[string],error){
		if len(input) == 0{
			return &IResult[string]{
				notParsed:input,
			},nil
		}

		for i,r := range []rune(input){
			if unicode.IsLetter(r){
				continue
			}else{
				return &IResult[string]{
					notParsed:string([]rune(input)[i:]),
					produced:string([]rune(input)[:i]),
				},nil
			}
		}

		return &IResult[string]{
			notParsed:"",
			produced:input,
		},nil
	}
}

// Recognize one or more lowercase and uppercase ASCII alphabetic characters
func Alpha1()ParseFn[string]{
	return func(input string)(*IResult[string],error){
		i,err :=Alpha0()(input)
		if err != nil{
			return &IResult[string]{},err
		}
		if len(i.produced) == 0{
			return &IResult[string]{notParsed:input},parseErr("need at least one alpha").notMatch().build()
		}
		return i,nil
	}
}

func AlphaNumeric0()ParseFn[string]{
	return func(input string)(*IResult[string],error){
		if len(input) == 0{
			return &IResult[string]{
				notParsed:input,
			},nil
		}

		for i,r := range []rune(input) {
			if unicode.IsLetter(r) || unicode.IsNumber(r){
				continue
			}else{
				return &IResult[string]{
					notParsed:string([]rune(input)[i:]),
					produced:string([]rune(input)[:i]),
				},nil
			}
		}

		return &IResult[string]{
			notParsed:"",
			produced:input,
		},nil
	}
}


func AlphaNumeric1()ParseFn[string]{
	return func(input string)(*IResult[string],error){
		i,err :=AlphaNumeric0()(input)
		if err != nil{
			return &IResult[string]{},err
		}
		if len(i.produced) == 0{
			return &IResult[string]{notParsed:input},parseErr("need at least one alpha or number").notMatch().build()
		}
		return i,nil
	}
}

func TakeWhile(predicate func(r rune)bool)ParseFn[string]{
	return func(input string)(*IResult[string],error){
		for i,r := range []rune(input){
			if !predicate(r){
				return &IResult[string]{
					notParsed:input[i:],
					produced: input[:i],
				},nil
			}
		}

		return &IResult[string]{
			notParsed:"",
			produced: input,
		},nil
	}
}

func TakeWhile1(predicate func(r rune)bool)ParseFn[string]{
	return func(input string)(*IResult[string],error){
		if len(input) == 0{
			return nil,parseErr("take whilte not match").notMatch().build()
		}
		for i,r := range []rune(input){
			if !predicate(r){
				if i == 0{
					return nil,parseErr("take whilte not match").notMatch().build()
				}
				return &IResult[string]{
					notParsed:input[i:],
					produced: input[:i],
				},nil
			}
		}

		return &IResult[string]{
			notParsed:"",
			produced: input,
		},nil
	}
}



func Space0()ParseFn[string]{
	return func(input string)(*IResult[string],error){
		return  TakeWhile(func(r rune)bool{
			return strings.ContainsAny(string(r),"\t\n \r")
		})(input)
	}
}
