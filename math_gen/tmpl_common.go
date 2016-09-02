package main

func Min_TYPENAME_(a, b _TYPE_) _TYPE_ {
	if a < b {
		return a
	}
	return b
}

func Max_TYPENAME_(a, b _TYPE_) _TYPE_ {
	if a >= b {
		return a
	}
	return b
}

func Min_TYPENAME_s(a []_TYPE_) (min _TYPE_) {
	if l := len(a); l > 0 {
		min = a[0]
		for i := 1; i < l; i++ {
			if min > a[i] {
				min = a[i]
			}
		}
	}
	return
}

func Max_TYPENAME_s(a []_TYPE_) (max _TYPE_) {
	if l := len(a); l > 0 {
		max = a[0]
		for i := 1; i < l; i++ {
			if max < a[i] {
				max = a[i]
			}
		}
	}
	return
}
