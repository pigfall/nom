package nom

import(
	"testing"
	"unicode"
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

func TestAlpha0(t *testing.T){
	tests := []struct{
		tag string
		input string
		want func(result *IResult,err error)
	}{
		{
			input:"tzz33",
			want:func(result *IResult,err error){
				require.Equal(t,nil,err)
				require.Equal(t,"tzz",result.produced)
				require.Equal(t,"33",result.notParsed)
			},
		},
		{
			input:"tzztzz",
			want:func(result *IResult,err error){
				require.Equal(t,nil,err)
				require.Equal(t,"tzztzz",result.produced)
				require.Equal(t,"",result.notParsed)
			},
		},
		{
			input:"t2zz33",
			want:func(result *IResult,err error){
				require.Equal(t,nil,err)
				require.Equal(t,"2zz33",result.notParsed)
				require.Equal(t,"t",result.produced)
			},
		},
		{
			input:" t2zz33",
			want:func(result *IResult,err error){
				require.Equal(t,nil,err)
				require.Equal(t," t2zz33",result.notParsed)
				require.Equal(t,"",result.produced)
			},
		},
	}

	for _,test := range tests{
		var  r,e = Alpha0()(test.input)
		test.want(&r,e)
	}
}

func TestAlpha1(t *testing.T){
	tests := []struct{
		tag string
		input string
		want func(result *IResult,err error)
	}{
		{
			input:"tzz33",
			want:func(result *IResult,err error){
				require.Equal(t,nil,err)
				require.Equal(t,"tzz",result.produced)
				require.Equal(t,"33",result.notParsed)
			},
		},
		{
			input:"t2zz33",
			want:func(result *IResult,err error){
				require.Equal(t,nil,err)
				require.Equal(t,"2zz33",result.notParsed)
				require.Equal(t,"t",result.produced)
			},
		},
		{
			input:" t2zz33",
			want:func(result *IResult,err error){
				require.Equal(t,true,MustParseErr(err).IsNotMatch())
				require.Equal(t," t2zz33",result.notParsed)
			},
		},
	}

	for _,test := range tests{
		var  r,e = Alpha1()(test.input)
		test.want(&r,e)
	}
}

func TestAlphaNumeric0(t *testing.T){
	tests := []struct{
		tag string
		input string
		want func(result *IResult,err error)
	}{
		{
			input:"tzz33",
			want:func(result *IResult,err error){
				require.Equal(t,nil,err)
				require.Equal(t,"tzz33",result.produced)
				require.Equal(t,"",result.notParsed)
			},
		},
		{
			input:"tzz_33",
			want:func(result *IResult,err error){
				require.Equal(t,nil,err)
				require.Equal(t,"tzz",result.produced)
				require.Equal(t,"_33",result.notParsed)
			},
		},
	}

	for _,test := range tests{
		var  r,e = AlphaNumeric0()(test.input)
		test.want(&r,e)
	}
}

func TestAlphaNumeric1(t *testing.T){
	tests := []struct{
		tag string
		input string
		want func(result *IResult,err error)
	}{
		{
			input:"tzz33",
			want:func(result *IResult,err error){
				require.Equal(t,nil,err)
				require.Equal(t,"tzz33",result.produced)
				require.Equal(t,"",result.notParsed)
			},
		},
		{
			input:"tzz_33",
			want:func(result *IResult,err error){
				require.Equal(t,nil,err)
				require.Equal(t,"tzz",result.produced)
				require.Equal(t,"_33",result.notParsed)
			},
		},
		{
			input:"_33",
			want:func(result *IResult,err error){
				require.Equal(t,true,MustParseErr(err).IsNotMatch())
			},
		},
	}

	for _,test := range tests{
		var  r,e = AlphaNumeric1()(test.input)
		test.want(&r,e)
	}
}

func TestTakeWhile(t *testing.T){
	tests := []struct{
		input string
		want func(result *IResult,err error)
		prediacte func(r rune)bool
	}{
		{
			input:"tzz33",
			prediacte:func(r rune)bool{
				return unicode.IsLetter(r)
			},
			want:func(result *IResult,err error){
				require.Equal(t,nil,err)
				require.Equal(t,"tzz",result.produced)
				require.Equal(t,"33",result.notParsed)
			},
		},
	}

	for _,test := range tests{
		var  r,e = TakeWhile(test.prediacte)(test.input)
		test.want(&r,e)
	}
}


