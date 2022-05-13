package nom

import (
	"errors"
	"testing"
)


func TestChar (t *testing.T ){
	char := 't'
	r,m,err := Char(char)("tzz")
	if err != nil {
		t.Fatal(err)
	}
	if r != "zz"{
		t.Fatalf("unexpect remaining %s",r)
	}
	if m != 't'{
		t.Fatalf("unexpect matched %v",m)
	}

	_,_,err = Char(char)("")
	var targetErr *CharErr
	if !errors.As(err,&targetErr){
		t.Fatal("expect CharErr")
	}

	_,_,err = Char(char)("bzz")
	 targetErr = nil
	if !errors.As(err,&targetErr){
		t.Fatal("expect CharErr")
	}
}

func TestTag(t *testing.T){
	input:= "abc"
	r,m,err := Tag("ab")(input)
	if err != nil{
		t.Fatal(err)
	}
	if r != "c"{
		t.Fatal("unexpected remaining")
	}
	if m != "ab"{
		t.Fatal("unexptecd matched")
	}

	_,_,err = Tag("c")(input)
	if err == nil{
		t.Fatal("unexptecd")
	}
}
