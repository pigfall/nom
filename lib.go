package nom

import(
	"fmt"
	"strings"
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
