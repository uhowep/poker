package poker


func Compare(cLeft *PokerCard, cRight *PokerCard) int {
	var result int
	// compare type
	if cLeft.CardType > cRight.CardType {
		result = 1
	} else if cLeft.CardType < cRight.CardType {
		result = 2
	} else {
		// compare face when type is same
		for i := 0; i < len(cLeft.Faces); i++ {
			if cLeft.Faces[i] > cRight.Faces[i] {
				return 1
			} else if cLeft.Faces[i] < cRight.Faces[i] {
				return 2
			}
		}
	}
	// return
	return result
}
