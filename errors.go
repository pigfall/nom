package nom

import(
	"errors"
)


const (
	errUndefiend = iota
	errNotMatch 
)

type ParseErr struct{
	desc string
	errCode int
}

func (this *ParseErr) Error()string{
	return this.desc
}

func (this *ParseErr) IsNotMatch()bool{
	return this.errCode == errNotMatch
}

type parseErrBuilder struct{
	desc string
	errCode int
}

func parseErr(desc string)*parseErrBuilder{
	return &parseErrBuilder{desc:desc}
}

func (this *parseErrBuilder) notMatch()*parseErrBuilder{
	this.errCode = errNotMatch
	return this
}

func (this *parseErrBuilder) build()*ParseErr{
	if this.errCode == errUndefiend{
		panic("unreachable")
	}
	return &ParseErr{
		desc:this.desc,
		errCode: this.errCode,
	}
}

func MustParseErr(err error)*ParseErr{
	var parseErr *ParseErr
	if errors.As(err,&parseErr){
		return parseErr
	}
	panic("not parseErr")
}
