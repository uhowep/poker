package poker

// 保持一定间隔以方便适应未来增加牌型
const (
	ROYAL_FLUSH = 100	// 皇家同花顺
	STRAIGHT_FLUSH = 90	// 同花顺
	FOUR_SAME = 80		// 四条
	GOURD = 70			// 葫芦
	FLUSH = 60			// 同花
	STRAIGHT = 50		// 顺子
	THREE_SAME = 40		// 三张
	TWO_PAIRS = 30		// 两对
	ONE_PAIRS = 20		// 一对
	SIMPLE = 10			// 单张
)

const CARD_MAX_LENGTH = 5	// 暂无用,原意构建可支持扩展的规则(如最终非5张等)


type Card struct {
	Face int		// 牌面
	Color string	// 花色
}

type AuxiliaryData struct {
	Faces []int		// 除癞子外的牌面数组
	ColorMap map[string][]int		// 花色对应的牌面
}

type PokerCard struct {
	Cards []Card		// 除癞子外
	GhostCount int		// 癞子数量
	Type int			// 牌型
	ResultFaces []int	// 解析后的最终结果
	AuxiliaryData AuxiliaryData		// 中间数据
}