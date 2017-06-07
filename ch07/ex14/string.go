// Copyright 2017 Ken Miura
package ex13

func (v Var) String() string {
	return string(v)
	//return Format(v) // fmt.Fprintf内で空インターフェースにVarを渡した際にString呼び出しの無限再帰になり、stackoverflowを引き起こすので、Formatを呼ばずにそのままstringにして返す。
}

func (l literal) String() string {
	return Format(l)
}
func (u unary) String() string {
	return Format(u)
}

func (b binary) String() string {
	return Format(b)
}

func (c call) String() string {
	return Format(c)
}

func (m min) String() string {
	return Format(m)
}
