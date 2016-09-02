package main

func Avg_TYPENAME_(a []_TYPE_) (res _TYPE_) {
	cnt := uint64(len(a))
	if cnt == 0 {
		return 0
	}

	acc := uint64(0)
	for _, val := range a {
		valDiv, valMod := uint64(val)/cnt, uint64(val)%cnt
		res += _TYPE_(valDiv)

		if acc+valMod >= cnt {
			res++
			acc -= cnt - valMod
		} else {
			acc += valMod
		}
	}

	return
}
