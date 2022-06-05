package nom

import(
	"strconv"
"github.com/stretchr/testify/require"
	"testing"
)

func TestSepateList0(t *testing.T){
	res,err := SeparatedList0(
		Tag(","),
		Tag("1"),
	)("1,1,1")
	if err != nil{
		t.Fatal(err)
	}
	require.Equal(t,[]string{"1","1","1"},res.produced)

	resB,err := SeparatedList0(
		Tag(","),
		Map(Tag("1"),func(s string)int{
			v,err := strconv.ParseInt(s,10,64)
			if err != nil{
				t.Fatal(err)
			}
			return int(v)
		}),
	)("1,1,1")
	if err != nil{
		t.Fatal(err)
	}
	require.Equal(t,[]int{1,1,1},resB.produced)
}
