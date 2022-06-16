package common

func SliceDifference(first []interface{}, second []interface{}, equals func(interface{}, interface{}) bool) []interface{} {
	res := make([]interface{}, 0)

	for i := range first {
		has := false

		for j := range second {
			if equals(first[i], second[j]) {
				has = true
				break
			}
		}

		if !has {
			res = append(res, first[i])
		}
	}

	return res
}

func SliceDifferenceInt64(first []int64, second []int64) []int64 {
	res := make([]int64, 0)

	for i := range first {
		has := false

		for j := range second {
			if first[i] == second[j] {
				has = true
				break
			}
		}

		if !has {
			res = append(res, first[i])
		}
	}

	return res
}

func SliceDifferenceString(first []string, second []string) []string {
	res := make([]string, 0)

	for i := range first {
		has := false

		for j := range second {
			if first[i] == second[j] {
				has = true
				break
			}
		}

		if !has {
			res = append(res, first[i])
		}
	}

	return res
}

func StringInterfaceEquality(p1 interface{}, p2 interface{}) bool {
	b1, ok1 := p1.(string)
	b2, ok2 := p2.(string)
	if !ok1 || !ok2 {
		return false
	}
	return b1 == b2
}

func SliceStringToInterfaces(ss []string) []interface{} {
	ids := make([]interface{}, len(ss))
	for i := range ss {
		ids[i] = ss[i]
	}

	return ids
}

func SliceIntToInterfaces(si []int) []interface{} {
	ids := make([]interface{}, len(si))
	for i := range si {
		ids[i] = si[i]
	}

	return ids
}
