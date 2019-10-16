package poker

import (
	"errors"
	"strconv"
)

type PokerCard struct {
	/*
	 * 同牌型比较两种方案：
	 * 1. 计算得分,要求对Faces做特殊排序,大的比较顺序在前,如34944时Faces=44493而非94443,然后乘积因子为14,可以理解成变成14进制
	 * 2. 额外维护一个CompareFaces数组/切片,按比较顺序存储要比较的数值,如34944时Faces=94443,CompareFaces=[4,9,3]
	 * 方案1更具通用性,能够同时比较N手牌
	 */
	// face of card slice: sort by poker weight rules, 36439 -> 33964
	Faces []int
	// card type: enums of constant
	CardType int
	// score: calculate the score when on the same card type
	Score int
	// color count: different color type count (1~4)
	ColorCount int
	// same times: search for same times, n(n-1)/2
	sameTimes int
}

func NewPokerCard(cardStr string) *PokerCard {
	faces, colorCount, err := parseStringToCard(cardStr)
	if err != nil {
		panic(err)
	}
	// new
	pokerCard := &PokerCard{
		Faces: faces,
		ColorCount: colorCount,
	}
	// return
	return pokerCard
}

func (pc *PokerCard) Parse() error {
	// parse type and get card type
	err := pc.ParseType()
	if err != nil {
		return err
	}
	// return
	return nil
}

/*
 * 判断牌型: 重复,连续
 * 两种方案:
 * 1. 遍历有多少个相同的,未优化情况下时间复杂度为O(n2),空间复杂度O(1)
 *    判断相同的数量 sameCount = 0单张-1一对-2两对-3三条-4葫芦-6四条
 * 2. 维护一个以牌面值为key,出现次数为value的map,时间复杂度为O(n),空间复杂度为O(n)
 *    判断map的长度len(map) = 2四条/葫芦-3三条/两对-4一对-5单张,针对2和3再对map中的value进行判断
 * 使用方案1
 */
func (pc *PokerCard) ParseType() error {
	// sort by poker weight and get same times(考虑将该语句调到Parse接口中,但是该接口又耦合该语句)
	pc.sameTimes = sortByPoker(pc.Faces)
	// judge
	faceLen := len(pc.Faces)
	switch pc.sameTimes {
	case 0:
		if (pc.Faces[0] - pc.Faces[faceLen-1]) == 4 {
			// 顺子
			pc.CardType = STRAIGHT
		} else if pc.Faces[0] == 14 && pc.Faces[faceLen-1] == 2 && (pc.Faces[1] - pc.Faces[faceLen-1]) == 3 {
			// 顺子,且将A变成1
			pc.CardType = STRAIGHT
			pc.Faces = append(pc.Faces[1:], 1)
		} else {
			// 单张
			pc.CardType = SIMPLE
		}
	case 1:
		// 一对
		pc.CardType = ONE_PAIRS
	case 2:
		// 两对
		pc.CardType = TWO_PAIRS
	case 3:
		// 三条
		pc.CardType = THREE_SAME
	case 4:
		// 葫芦
		pc.CardType = GOURD
	case 6:
		// 四条
		pc.CardType = FOUR_SAME
	default:
		return errors.New("invalid card type case")
	}
	// 额外判断同花,同花比4条和葫芦小,但是比顺子三条两对一对单张大
	if pc.sameTimes < 4 && pc.ColorCount == 1 {
		if (pc.Faces[0] - pc.Faces[faceLen-1]) == 4 {
			if pc.Faces[0] == 14 {
				// 皇家同花顺
				pc.CardType = ROYAL_FLUSH
			} else {
				// 同花顺
				pc.CardType = STRAIGHT_FLUSH
			}
		} else {
			// 同花
			pc.CardType = FLUSH
		}
	}
	// return
	return nil
}

// 使用计算得分的方案 --更具通用性,允许多副牌一起比较
func (pc *PokerCard) CalculateScore() {
	// todo
}

func parseStringToCard(cardStr string) (faces []int, colorCount int, err error) {
	if len(cardStr) % 2 != 0 {
		err = errors.New("invalid card string length")
		return
	}
	var tmp int
	faces = make([]int, len(cardStr) / 2)
	colorMap := make(map[byte]bool, 4)
	for index, value := range []byte(cardStr) {
		if index % 2 != 0 {
			// input color into map
			if _, ok := colorMap[value]; !ok {
				colorMap[value] = true
			}
			continue
		}
		// transfer face to int and push into face slice
		switch value {
		case 'T':
			faces[index/2] = 10
			//faces = append(faces, 10)
		case 'J':
			faces[index/2] = 11
			//faces = append(faces, 11)
		case 'Q':
			faces[index/2] = 12
			//faces = append(faces, 12)
		case 'K':
			faces[index/2] = 13
			//faces = append(faces, 13)
		case 'A':
			faces[index/2] = 14
			//faces = append(faces, 14)
		default:
			// transfer face to int
			tmp, err = strconv.Atoi(string(value))
			if err != nil {
				return
			} else if tmp < 2 || tmp > 9 {
				err = errors.New("invalid face")
				return
			}
			faces[index/2] = tmp
			//faces = append(faces, tmp)
		}
	}
	// calculate color count
	colorCount = len(colorMap)
	// return
	return
}
