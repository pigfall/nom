package nom

import(
	"testing"
"github.com/stretchr/testify/require"
)

func TestAlt(t *testing.T){
	res,err := Alt(
		[]ParseFn[string]{
			Tag("12"),
			Tag("13"),
			Tag("14"),
		},
	)("14")
	if err != nil{
		t.Fatal(err)
	}
	require.Equal(t,"14",res.produced)
}
