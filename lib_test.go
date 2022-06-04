package nom

import(
	"testing"
	"github.com/stretchr/testify/require"
)

func TestTag(t *testing.T){
	tests := []struct{
		tag string
		input string
		want func(result *IResult,err error)
	}{
		{
			tag:"tzz",
			input:"tzz33",
			want:func(result *IResult,err error){
				require.Equal(t,nil,err)
				require.Equal(t,"tzz",result.produced)
				require.Equal(t,"33",result.notParsed)
			},
		},
		{
			tag:" tzz",
			input:"tzz33",
			want:func(result *IResult,err error){
				require.Equal(t,MustParseErr(err).IsNotMatch(),true)
			},
		},
	}

	for _,test := range tests{
		var  r,e = Tag(test.tag)(test.input)
		test.want(&r,e)
	}
}
