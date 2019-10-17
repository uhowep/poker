package poker


func Compare(cLeft *PokerCard, cRight *PokerCard) int {
	var result int
	// compare type
	if cLeft.Type > cRight.Type {
		result = 1
	} else if cLeft.Type < cRight.Type {
		result = 2
	} else {
		// compare face when type is same
		for i := 0; i < len(cLeft.ResultFaces); i++ {
			if cLeft.ResultFaces[i] > cRight.ResultFaces[i] {
				return 1
			} else if cLeft.ResultFaces[i] < cRight.ResultFaces[i] {
				return 2
			}
		}
	}
	// return
	return result
}
