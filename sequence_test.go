package nom

import(
	"testing"
	"github.com/stretchr/testify/require"
)

func TestDelimited(t *testing.T){
	tests :=[]struct{
		first ParseFn[string]
		second ParseFn[string]
		third ParseFn[string]
		input string
		wanted func(res *IResult[string],err error)
	}{
		{
			first: Tag("12"),
			second:Tag("34"),
			third:Tag("56"),
			input:"123456",
			wanted:func(res *IResult[string],err error){
				require.Equal(t,nil,err)
				require.Equal(t,"34",res.produced)
			},
		},
	}

	for _,test := range tests{
		res,err := Delimited(test.first,test.second,test.third)(test.input)
		test.wanted(res,err)
	}

}
