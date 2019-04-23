package triangle

import "math"

type Kind int

const (
    NaT Kind = iota
    Equ
    Iso
    Sca
)

func KindFromSides(a,b,c float64) Kind {
	switch {
	case a<=0,b<=0,c<=0,a+b<c,a+c<b,b+c<a,math.IsNaN(a+b+c),math.IsInf(a+b+c,1):return NaT
	case a==b&&a==c: return Equ
	case a==b,a==c,b==c:return Iso
	case a!=b&&a!=c:return Sca
	default:return NaT
	}
}
