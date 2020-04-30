package slice

// Return index of x in slice.
func Index(slice []interface{}, x interface{}) int {
	index := -1
	// 返回最终索引
	for i:=0; i<len(slice); i++ {
		if slice[i] == x {
			index = i
			break
		}
	}
	return index
}
