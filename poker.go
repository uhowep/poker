package poker

import (
	"errors"
	"strconv"
)

func NewPokerCard(origin string) *PokerCard {
	if len(origin) % 2 != 0 {
		panic("invalid card string length")
	}
	// define variable
	ghostCount := 0
	cards := make([]Card, 0)
	faces := make([]int, 0)
	colorMap := make(map[string][]int, 4)
	// parse string
	for i := 0; i < len(origin); {
		if origin[i] == 'X' {
			ghostCount++
		} else {
			// parse string into face
			face := 0
			switch origin[i] {
			case 'A':
				face = 14
			case 'T':
				face = 10
			case 'J':
				face = 11
			case 'Q':
				face = 12
			case 'K':
				face = 13
			default:
				var err error
				face, err = strconv.Atoi(string(origin[i]))
				if err != nil {
					panic("convert string to int failed with " + string(origin[i]))
				}
			}
			color := string(origin[i+1])
			// append into cards
			cards = append(cards, Card{Face:face, Color:color})
			// append into faces
			faces = append(faces, face)
			// deal with color map
			if _, ok := colorMap[color]; ok {
				colorMap[color] = append(colorMap[color], face)
			} else {
				colorMap[color] = make([]int, 0)
				colorMap[color] = append(colorMap[color], face)
			}
		}
		i += 2
	}
	// return
	return &PokerCard{
		Cards:cards,
		GhostCount:ghostCount,
		AuxiliaryData:AuxiliaryData{Faces:faces, ColorMap:colorMap},
	}
}

/*
 * 先判断皇同和同花顺,若未找到则进行查重,进而判断,最后进行顺子及同花的判断来"提权"
 * 未做过多分析来写条件以提升效率,util模块同理,尽可能的做到了可复用性
 * 能够判断包含N个癞子的N张手牌,不局限于7张
 */
func (pc *PokerCard) Parse() error {
	// sort desc
	quickSort(pc.AuxiliaryData.Faces)
	// is straight flush --副作用:对AuxiliaryData.ColorMap排序
	sfFaces, isStraightFlush := buildStraightFlush(pc.AuxiliaryData.Faces, pc.AuxiliaryData.ColorMap, pc.GhostCount)
	if isStraightFlush {
		// is royal flush
		if sfFaces[0] == 14 {
			pc.Type = ROYAL_FLUSH
		} else {
			pc.Type = STRAIGHT_FLUSH
		}
		pc.ResultFaces = sfFaces
		// return
		return nil
	}
	// could build same
	sameTimes := 0
	sameFaces, isSame := buildSame(pc.AuxiliaryData.Faces, pc.GhostCount)
	if isSame {
		// calculate same times
		faceLen := len(sameFaces)
		for i := 0; i < faceLen; i++ {
			for j := i+1; j < faceLen; j++ {
				if sameFaces[i] == sameFaces[j] {
					sameTimes++
				}
			}
		}
	}
	// judge
	switch sameTimes {
	case 0:
		if len(pc.AuxiliaryData.Faces) >= 5 {
			pc.Type = SIMPLE
			pc.ResultFaces = pc.AuxiliaryData.Faces[:5]
		} else {	// 基本不会跑到这里来,因为在new的时候就做了限制
			return errors.New("invalid simple parse")
		}
	case 1:
		// 一对
		pc.Type = ONE_PAIRS
		pc.ResultFaces = sameFaces
	case 2:
		// 两对
		pc.Type = TWO_PAIRS
		pc.ResultFaces = sameFaces
	case 3:
		// 三条
		pc.Type = THREE_SAME
		pc.ResultFaces = sameFaces
	case 4:
		// 葫芦
		pc.Type = GOURD
		pc.ResultFaces = sameFaces
	case 6:
		// 四条
		pc.Type = FOUR_SAME
		pc.ResultFaces = sameFaces
	default:
		return errors.New("invalid card type case")
	}
	// 额外判断同花和顺子,同花和顺子比4条和葫芦小,但是比三条两对一对单张都大(先判断小的)
	if sameTimes < 4 {
		// is straight
		straightFaces, isStraight := buildStraight(pc.AuxiliaryData.Faces, pc.GhostCount)
		if isStraight {
			pc.Type = STRAIGHT
			pc.ResultFaces = straightFaces
		}
		// is flush
		flushFaces, isFlush := buildFlush(pc.AuxiliaryData.ColorMap, pc.GhostCount)
		if isFlush {
			pc.Type = FLUSH
			pc.ResultFaces = flushFaces
		}
	}
	// return
	return nil
}
