package nom

import(
	"strconv"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMap(t *testing.T){
	res,err := Map(
		Tag("12"),
		func(s string)int{
			v,err := strconv.ParseInt(s,10,64)
			if err != nil{
				panic(err)
			}
			return int(v)
		},
	)("12")
	if err != nil{
		t.Fatal(err)
	}
	require.Equal(t,12,res.produced)

}
