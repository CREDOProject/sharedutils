package filter

func Filter[T any](list []T, test ...func(T) bool) (ret []T) {
	for _, item := range list {
		var check bool = true
		for _, t := range test {
			if !check {
				break
			}
			check = check && t(item)
		}
		if check {
			ret = append(ret, item)
		}
	}
	return
}
