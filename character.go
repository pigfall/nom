package nom

import(
	"strings"
	"fmt"
)

type CharErr struct {

}

func (this *CharErr) Error() string{
	return fmt.Sprintf("char error")
}

func Char(char rune)( func(input string)(remainingInput string,matched rune,err error) ){
	return  func(input string)(remainingInput string,matched rune,err error) {
		if len(input) == 0{
			return "",0,&CharErr{}
		}
		if []rune(input)[0] == char{
			return input[1:],char,nil
		}

		return "",0, &CharErr{}
	}
}

type TagErr struct{}

func (this *TagErr) Error() string{
	return "tag not matched"
}

func Tag(t string)(func(input string)(remainingInput string,matched string,err error)){
	return func(input string)(remainingInput string,matched string,err error){
		if strings.HasPrefix(input,t){
			return input[len(t):],t,nil
		}
		return "","", &TagErr{}
	}
}

type AltErr struct{

}

func (this *AltErr) Error() string{
	return "alt err"
}

func Alt(l []func(input string)(remainingInput string,matched string,err error))func(string)(remainingInput string,matched string,err error){
	return func(i string)(remainingInput string,matched string,err error){
		for _,v := range l {
			r,m,err := v(i)
			if err == nil{
				return r,m,nil
			}
		}
		return "","",&AltErr{}
	}

}
