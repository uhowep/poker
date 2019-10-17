package poker

// quick sort desc with int
func quickSort(slices []int) {
	if len(slices) <= 1 {
		return
	}
	start, end := 0, len(slices) - 1
	mark := slices[0]
	// compare and move
	for i := 1; i <= end; {
		if slices[i] > mark {
			// put max value into left area
			slices[i], slices[start] = slices[start], slices[i]
			// move left area
			start++
			// next compare value index
			i++
		} else {
			// put min value into right area
			slices[i], slices[end] = slices[end], slices[i]
			// move right area
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
 * judge if it can be built maximum straight with ghost which contain 5 cards finally
 * slices must already be sorted desc
 * condition for success: found continuous face count which equal (5 - ghostCount)
 * condition for failed: traversed the end of slice/array or the length of slice less than (5 - ghostCount)
 * state: 1.use ghost first: e.g.[9,8,7,6],2 => [11,10,9,8,7], e.g.[9,7,6,5],2 => [10,9,8,7,6]
 *      : 2.reach top | surplus ghost: e.g.[12,11,10,9],3 => [14,13,12,11,10]
 *      : 3.transfer 14 to 1: e.g.[14,5,3,2],1 => [5,4,3,2,1]
 *      : 4.see test unit for others
 */
func buildStraight(slices []int, ghostCount int) (result []int, isSuccess bool) {
	result = make([]int, 0)
	// define variable
	shouldFound := 5 - ghostCount		// find continuous face count
	if shouldFound <= 0 {
		return []int{14,13,12,11,10}, true
	}
	// not enough length
	if len(slices) < shouldFound {
		return
	}
	// if has 14/A, then append 1/A at the end of slices
	if slices[0] == 14 {
		slices = append(slices, 1)
	}
	// iterate slices
	for i := 0; len(slices) - i >= shouldFound; i++ {
		remainGhost := ghostCount
		thisVal := slices[i]	// current value for comparing it is continuous
		alreadyFound := 1		// already found continuous faces
		// search for continuous
		for j := i + 1; j < len(slices) && alreadyFound < shouldFound; {
			// found continuous
			if thisVal - slices[j] == 1 {
				alreadyFound++
				thisVal = slices[j]
				j++		// next
			} else if slices[j] == thisVal {
				// found same and then continue
				j++
			} else if remainGhost < 1 {
				// not found and not enough ghost to use
				break
			} else {
				// use ghost --don't increase alreadyFound caused of it must be in slices
				//alreadyFound++
				remainGhost--
				thisVal = thisVal - 1
			}
		}
		// check meet the condition for success and build straight
		if alreadyFound >= shouldFound {
			isSuccess = true
			// append remain ghost to left
			for maxVal := slices[i]+1; remainGhost > 0 && maxVal <= 14; {
				result = append([]int{maxVal}, result...)
				remainGhost--
				maxVal++
			}
			// append found
			for tVal := slices[i]; tVal >= thisVal; tVal-- {
				result = append(result, tVal)
			}
			// append remain ghost to right
			for minVal := thisVal-1; remainGhost > 0 && minVal >= 1; {
				result = append(result, minVal)
				remainGhost--
				minVal--
			}
			// break
			break
		}
	}
	// return
	return
}

/*
 * judge if it can be built maximum flush with ghost which contain 5 cards finally
 * the slices could be not sorted --now is already sorted
 */
func buildFlush(colorMap map[string][]int, ghostCount int) (result []int, isSuccess bool) {
	result = make([]int, 0)
	// define variable
	shouldFound := 5 - ghostCount		// find same color count
	maxFaces := make([]int, 0)
	// search for flush, ignore color
	for _, slices := range colorMap {
		faceLen := len(slices)
		// found it
		if faceLen >= shouldFound {
			isSuccess = true
			// if the slices are already be sorted desc outside
			//quickSort(slices)		--偷懒: 外层已排好
			if len(maxFaces) > 0 {
				// found max face
				for i := 0; i < shouldFound; i++ {
					if slices[i] > maxFaces[i] {
						maxFaces = slices
						break
					} else if slices[i] < maxFaces[i] {
						break
					} else {
						// if same and then compare next
						continue
					}
				}
			} else {
				maxFaces = slices
			}
		} else {
			continue
		}
	}
	if isSuccess {
		// append ghost which change to 14
		i := 0
		for ; i < ghostCount; i++ {
			result = append(result, 14)
		}
		// append max face and cut
		result = append(result, maxFaces...)
		result = result[:5]
		// deal with 14,14,14,14,14 case
		if result[4] == 14 {
			result[4] = 13
		}
	}
	// return
	return
}

/*
 * judge if it can be build maximum straight flush with ghost which contain 5 cards finally
 * the slices must already be sorted desc
 */
func buildStraightFlush(slices []int, face2ColorMap map[string][]int, ghostCount int) (result []int, isSuccess bool) {
	result = make([]int, 0)
	// define variable
	shouldFound := 5 - ghostCount
	straightFaces := make([][]int, 0)
	straightMaxFaces := 0
	// straight flush
	if len(slices) < shouldFound {
		// not enough length
		return
	} else if shouldFound <= 0 {
		// royal flush
		result = []int{14,13,12,11,10}
		isSuccess = true
		return
	} else if shouldFound == 1 {
		// straight flush
		result, isSuccess = buildStraight(slices, ghostCount)
		return
	}
	// get same color face slices and judge that whether build straight
	for _, faces := range face2ColorMap {
		if len(faces) < shouldFound {
			continue
		}
		// sort by desc
		quickSort(faces)
		// judge it could be built straight
		faceResult, isStraight := buildStraight(faces, ghostCount)
		if isStraight {
			straightFaces = append(straightFaces, faceResult)
		}
	}
	// get max straight flush faces
	for _, faces := range straightFaces {
		isSuccess = true		// success
		if faces[0] > straightMaxFaces {
			straightMaxFaces = faces[0]
			// put the max straight flush faces to result
			result = faces
		}
	}
	// return
	return
}

/*
 * 构建相同:总张数5,以构建最大且最多相同为目标,最大权重大于最多,如优先构建4相同,然后是3相同2相同,然后是3相同1相同1相同,以此类推
 * 大致流程:
   已组装的长度,记alreadyBuild; 需要寻找拥有相同数超过（4-癞子数）,记needFound
   1.若(5-alreadyBuild)<=0,表示不需要,则直接14,14,14,14,13
   2.若(5-alreadyBuild)=1,表示找最大的那张即可,然后14,14,14,14,X (X == 14 ? 13 : X )
   3.其他情况,寻找构建最优,即根据needSame来构建最优:从sameMap中取key不小于needSame的face组,取其中最大的,如9,9,3,3,3不取3而取9(因为癞子够)
   3-a-1.若取到,则face记为tmFace,sameTimes记为tmFaceTimes
   3-a-2.若未取到,则从sameMap中取key小于needSame的下一个不空的face组,取第一个作为tmFace和tmFaceTimes
   3-b.将tmFaceTimes个tmFace封装进结果中
   3-c.查看是否还有剩余癞子,若有则选择合适的癞子封装进结果里(考虑是变为14还是tmFace)
   4.重复3直到组装满5个,然后最后处理下数组
 * 注:癞子不会去考虑变为次要相同的牌,仅考虑主要,如[9,5]优先考虑9不考虑5,[9,9,5,5]也一样
*/
func buildSame(slices []int, ghostCount int) (result []int, isSuccess bool) {
	result = make([]int, 0)
	// judge failure case and special case
	if ghostCount >= 5 {
		// judge special case
		result = []int{14,14,14,14,13}
		isSuccess = true
		return
	} else if len(slices) < 5 - ghostCount {
		// failure case
		isSuccess = false
		return
	} else if ghostCount == 4 {
		if slices[0] == 14 {
			result = []int{14,14,14,14,13}
		} else {
			result = []int{14,14,14,14,slices[0]}
		}
		isSuccess = true
		return
	}
	// build same map
	sameMap := map[int][]int{1:[]int{}, 2:[]int{}, 3:[]int{}, 4:[]int{}}
	for i := 0; i < len(slices); {
		j := i + 1
		for ; j < len(slices); j++ {
			if slices[j] != slices[i] {
				// not find continuous same face
				break
			}
		}
		// calculate same times and append
		sameTimes := j - i
		sameMap[sameTimes] = append(sameMap[sameTimes], slices[i])
		// move to next for comparing
		i = j
	}
	// no repeat face
	if ghostCount == 0 && len(slices) == len(sameMap[1]) {
		isSuccess = false
		return
	}
	// start to build
	remainGhost := ghostCount
	alreadyBuild := 0
	needSame := 4 - ghostCount
	for alreadyBuild < 5 {
		tmFace := 0		// this max face
		tmFaceTimes := 0	// the same times of this max face
		// find the max face which its same times more than need same times
		for n := needSame; n <= 4; n++ {
			thisFaces, ok := sameMap[n]
			if ok && len(thisFaces) > 0 && thisFaces[0] > tmFace{
				tmFace = thisFaces[0]
				tmFaceTimes = n
			}
		}
		// if not find and then find max face on next level
		if tmFace == 0 {
			for n := needSame - 1; n >= 1; n-- {
				thisFaces, ok := sameMap[n]
				if ok && len(thisFaces) > 0 {
					tmFace = thisFaces[0]
					tmFaceTimes = n
					break	// break when found it
				}
			}
		}
		// use tmFace/tmFaceTimes to append
		for n := 0; n < tmFaceTimes && n < 4; n++ {
			result = append(result, tmFace)
			alreadyBuild++
		}
		// use ghost to append if it has more
		for remainGhost > 0 {
			ghostFace := 0		// change ghost to someone face
			if alreadyBuild >= 4 {		// --can be commented and append tmFace directly
				ghostFace = 14
			} else {
				ghostFace = tmFace
			}
			// append
			result = append(result, ghostFace)
			alreadyBuild++
			remainGhost--
		}
		// delete the found face in sameMap in order to keep the first face in slices is maximum
		if len(sameMap[tmFaceTimes]) >= 1 {
			sameMap[tmFaceTimes] = sameMap[tmFaceTimes][1:]
		}
		// next round need same times to find
		needSame = 5 - alreadyBuild - remainGhost
	}
	// deal with result
	result = result[:5]
	if result[4] == result[0] {
		// caused of judge and set 14 at the end in line:280
		result[4] = 13
	}
	// return
	isSuccess = true
	return
}