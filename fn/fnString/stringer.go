package fnString

import "strings"

type (
	Chain struct {
		v string
	}
)

func NewChain(v string) *Chain {
	return &Chain{
		v: v,
	}
}

func (x *Chain) SnakeCase() *Chain {
	x.v = ToSnakeCase(x.v)
	return x
}

// CamelCase
// 카멜케이스로 변경
// startWithUpper = true 인 경우 대문자로 시작하는 카멜케이스
func (x *Chain) CamelCase(startWithUpper ...bool) *Chain {
	upper := false
	if len(startWithUpper) != 0 {
		upper = startWithUpper[0]
	}
	x.v = ToCamelCase(x.v, upper)
	return x
}

// Plural
// 복수형 변형
func (x *Chain) Plural() *Chain {
	x.v = ToPlural(x.v)
	return x
}

func (x *Chain) String() string {
	return x.v
}

func (x *Chain) RemoveSpace() *Chain {
	x.v = strings.Replace(x.v, " ", "", -1)
	return x
}
