package main

func Avg_TYPENAME_(a []_TYPE_) _TYPE_ {
	var resf float64 // max precision

	l := float64(len(a))
	for _, val := range a {
		resf += float64(val) / l
	}

	return _TYPE_(resf)
}
