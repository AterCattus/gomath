package main

func Gcd_TYPENAME_(a, b _TYPE_) (res _TYPE_) {
	res = 1
	for a != 0 && b != 0 {
		if am, bm := (a%2) == 0, (b%2) == 0; am && bm {
			res, a, b = 2*res, a/2, b/2
		} else if am && !bm {
			a /= 2
		} else if !am && bm {
			b /= 2
		} else if a >= b {
			a -= b
		} else {
			b -= a
		}
	}
	if a == 0 {
		res *= b
	} else {
		res *= a
	}
	return res
}
