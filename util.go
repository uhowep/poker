package poker

// sort desc
func quickSort(slices []int) {
	if len(slices) <= 1 {
		return
	}
	mark := slices[0]
	start, end := 0, len(slices)-1
	for i := 1; i <= end; {
		// put large value into left area and less value into right area
		if slices[i] >= mark {
			slices[start], slices[i] = slices[i], slices[start]
			start++
			i++
		} else {
			slices[end], slices[i] = slices[i], slices[end]
			end--
		}
	}
	// recursive left area
	quickSort(slices[:start])
	// recursive right area
	quickSort(slices[start+1:])
	// return
	return
}

/*
 * 对无序的数组按poker规则进行提权排序
 * 如35732->33752
 */
func sortByPoker(slices []int) int {
	// 逆序排序
	quickSort(slices)
	// key为相同的次数,value存储face切片,相比较下使用切片/数组比map快(数组相比切片会更快一点)
	sameFaceCount := [][]int{
		0: []int{},
		1: []int{},
		2: []int{},
		3: []int{},
	}
	// 找相同
	sliceLen := len(slices)
	for i := 0; i < sliceLen; {
		j := i + 1
		for ; j < sliceLen; j++ {
			if slices[j] != slices[i] {
				break
			}
		}
		sameCount := j - i
		// 相同数量范围1~4,数组下标0~3: sameFaceCount[1]=>[5,2]表示有2张5和2张3
		sameFaceCount[sameCount-1] = append(sameFaceCount[sameCount-1], slices[i])
		i = j
	}
	// deal with result
	index := 0
	allSameTimes := 0		// n(n-1)/2
	//resultSlice := make([]int, sliceLen)
	// 相同的从高到低遍历
	for i := 4; i >= 1; i-- {
		sfcLen := len(sameFaceCount[i-1])
		// 如果该组中有牌面值,则添加到结果数组中,比如2=>[5,3]表示有2张5和2张3
		if sfcLen > 0 {
			// 该组下的相同次数,如组号为4相同,则 4*(4-1)/2 = 6
			tmpST := (i) * (i - 1) / 2
			for _, face := range sameFaceCount[i-1] {
				// 这里表示插入i张牌(使用该循环避免append进行扩容)
				for n := 1; n <= i; n++ {
					slices[index] = face
					index++
				}
				// 维护all same times
				allSameTimes += tmpST
			}
		}
	}
	// return
	return allSameTimes
}

/*
 * 对逆序排好序的牌面值进行按poker大小进行排序
 * 如: 63332 -> 33362 / 98844 -> 88449
 * 此段可读性更高,但是效率会低点
 */
/*
func sortByPoker(slices []int) []int {
	sliceLen := len(slices)
	// {same_count:[same_face1,same_face2...]} : 63332 -> {1:[6,2], 2:[], 3:[3,3,3], 4:[]}
	// 目前固定5张牌,要通用的话则对参数取长度sliceLen,进而创建map的长度为1~sliceLen
	//midMap := map[int][]int{
	//	1: []int{},
	//	2: []int{},
	//	3: []int{},
	//	4: []int{},
	//}
	// 相比较下使用切片/数组比map快(数组相比切片会更快一点)
	midMap := [][]int{
		0: []int{},
		1: []int{},
		2: []int{},
		3: []int{},
	}
	for i := 0; i < sliceLen; {
		j := i + 1
		for ; j < sliceLen; j++ {
			if slices[j] != slices[i] {
				break
			}
		}
		sameCount := j - i
		// 相同数量范围1~4,数组下标0~3
		midMap[sameCount-1] = append(midMap[sameCount-1], slices[i:j]...)
		i = j
	}
	// deal with result
	resultSlice := make([]int, 0)
	for i := 3; i >= 0; i-- {
		if len(midMap[i]) != 0 {
			resultSlice = append(resultSlice, midMap[i]...)
		}
	}
	// return
	return resultSlice
}
*/