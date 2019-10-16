package poker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPokerCard(t *testing.T) {
	type verifyParamStruct struct {
		input string
		expect *PokerCard
	}
	// define verify params
	verifyParams := []verifyParamStruct{
		{
			input:"TsAsQcKhJc",
			expect: &PokerCard{
				Cards: []Card{{Face:10,Color:"s"},{Face:14,Color:"s"},{Face:12,Color:"c"},{Face:13,Color:"h"},{Face:11,Color:"c"}},
				GhostCount:0,
				AuxiliaryData:AuxiliaryData{
					Faces: []int{10,14,12,13,11},
					ColorMap: map[string][]int{"h":[]int{13},"s":[]int{10,14},"c":[]int{12,11}},
				},
			},
		},
		{
			input:"8hTsXn2sXnXnJc",
			expect: &PokerCard{
				Cards: []Card{{Face:8,Color:"h"},{Face:10,Color:"s"},{Face:2,Color:"s"},{Face:11,Color:"c"}},
				GhostCount:3,
				AuxiliaryData:AuxiliaryData{
					Faces: []int{8,10,2,11},
					ColorMap: map[string][]int{"h":[]int{8},"s":[]int{10,2},"c":[]int{11}},
				},
			},
		},
	}
	// verify
	for _, param := range verifyParams {
		pokerCard := NewPokerCard(param.input)
		assert.Equal(t, param.expect, pokerCard, param.input)
	}
	// return
	return
}

func TestPokerCard_Parse(t *testing.T) {
	type verifyParamStruct struct {
		input string
		expectRes []int
		expectType int
		occurErr bool
	}
	// define test data
	verifyParams := []*verifyParamStruct{
		// no ghost
		&verifyParamStruct{input:"6s7s3cAcKdTd4h", expectRes:[]int{14,13,10,7,6}, expectType:SIMPLE, occurErr: false},
		&verifyParamStruct{input:"6s7s3cAcKdTdAh", expectRes:[]int{14,14,13,10,7}, expectType:ONE_PAIRS, occurErr: false},
		&verifyParamStruct{input:"6sAs3cAcKd3d4h", expectRes:[]int{14,14,3,3,13}, expectType:TWO_PAIRS, occurErr: false},
		&verifyParamStruct{input:"6s7s3c6cKd6d4h", expectRes:[]int{6,6,6,13,7}, expectType:THREE_SAME, occurErr: false},
		&verifyParamStruct{input:"6s7s3c9c8dTd4h", expectRes:[]int{10,9,8,7,6}, expectType:STRAIGHT, occurErr: false},
		&verifyParamStruct{input:"Qs7sJcAc7dTdKh", expectRes:[]int{14,13,12,11,10}, expectType:STRAIGHT, occurErr: false},
		&verifyParamStruct{input:"2s2s3c4c5d6d9h", expectRes:[]int{6,5,4,3,2}, expectType:STRAIGHT, occurErr: false},
		&verifyParamStruct{input:"2s2s3c4c5d6d6h", expectRes:[]int{6,5,4,3,2}, expectType:STRAIGHT, occurErr: false},
		&verifyParamStruct{input:"2s2s3c4c5d6d2h", expectRes:[]int{6,5,4,3,2}, expectType:STRAIGHT, occurErr: false},
		&verifyParamStruct{input:"2sKs3cAc5dTd4h", expectRes:[]int{5,4,3,2,1}, expectType:STRAIGHT, occurErr: false},
		&verifyParamStruct{input:"6s7s8c9cTdJdQh", expectRes:[]int{12,11,10,9,8}, expectType:STRAIGHT, occurErr: false},
		&verifyParamStruct{input:"6s7s3cAsKdTs4s", expectRes:[]int{14,10,7,6,4}, expectType:FLUSH, occurErr: false},
		&verifyParamStruct{input:"6s6s3cAsKdTs4s", expectRes:[]int{14,10,6,6,4}, expectType:FLUSH, occurErr: false},
		&verifyParamStruct{input:"6s6s3c3sKdTs4s", expectRes:[]int{10,6,6,4,3}, expectType:FLUSH, occurErr: false},
		&verifyParamStruct{input:"6s6s3c6sKdTs4s", expectRes:[]int{10,6,6,6,4}, expectType:FLUSH, occurErr: false},
		&verifyParamStruct{input:"6s6s3s6sKsTs4s", expectRes:[]int{13,10,6,6,6}, expectType:FLUSH, occurErr: false},
		&verifyParamStruct{input:"6s4s3c6c6dTd4h", expectRes:[]int{6,6,6,4,4}, expectType:GOURD, occurErr: false},
		&verifyParamStruct{input:"6s4s3s6s6s5d4h", expectRes:[]int{6,6,6,4,4}, expectType:GOURD, occurErr: false},
		&verifyParamStruct{input:"6s4s3s6s6s3d4h", expectRes:[]int{6,6,6,4,4}, expectType:GOURD, occurErr: false},
		&verifyParamStruct{input:"6s7s7cAc7d6d7h", expectRes:[]int{7,7,7,7,14}, expectType:FOUR_SAME, occurErr: false},
		&verifyParamStruct{input:"6s7s7sAs7s6s7s", expectRes:[]int{7,7,7,7,14}, expectType:FOUR_SAME, occurErr: false},
		&verifyParamStruct{input:"6s7s3sQc5sTs4s", expectRes:[]int{7,6,5,4,3}, expectType:STRAIGHT_FLUSH, occurErr: false},
		&verifyParamStruct{input:"6s7s3sAs5s8s4s", expectRes:[]int{8,7,6,5,4}, expectType:STRAIGHT_FLUSH, occurErr: false},
		&verifyParamStruct{input:"TsQs3sAs5s2s4s", expectRes:[]int{5,4,3,2,1}, expectType:STRAIGHT_FLUSH, occurErr: false},
		&verifyParamStruct{input:"6s7s7sAs5s8s4s", expectRes:[]int{8,7,6,5,4}, expectType:STRAIGHT_FLUSH, occurErr: false},
		&verifyParamStruct{input:"6s7s7s8s5s8s4s", expectRes:[]int{8,7,6,5,4}, expectType:STRAIGHT_FLUSH, occurErr: false},
		&verifyParamStruct{input:"6s7s7s7s5s8s4s", expectRes:[]int{8,7,6,5,4}, expectType:STRAIGHT_FLUSH, occurErr: false},
		&verifyParamStruct{input:"TsQsKsAs7sJs4s", expectRes:[]int{14,13,12,11,10}, expectType:ROYAL_FLUSH, occurErr: false},
		&verifyParamStruct{input:"TsQsKsAsAsJs4s", expectRes:[]int{14,13,12,11,10}, expectType:ROYAL_FLUSH, occurErr: false},
		&verifyParamStruct{input:"TsQsKsAsAsJsKs", expectRes:[]int{14,13,12,11,10}, expectType:ROYAL_FLUSH, occurErr: false},
		&verifyParamStruct{input:"TsQsKsAsAsJsAs", expectRes:[]int{14,13,12,11,10}, expectType:ROYAL_FLUSH, occurErr: false},
		// 1 ghost
		&verifyParamStruct{input:"6s9s3cXnKdTd4h", expectRes:[]int{13,13,10,9,6}, expectType:ONE_PAIRS, occurErr: false},
		&verifyParamStruct{input:"6s7s3cXnKdTdKh", expectRes:[]int{13,13,13,10,7}, expectType:THREE_SAME, occurErr: false},
		&verifyParamStruct{input:"6s6s3cXnKdTdKh", expectRes:[]int{13,13,13,6,6}, expectType:GOURD, occurErr: false},
		&verifyParamStruct{input:"6s6s3cXnKd3dKh", expectRes:[]int{13,13,13,6,6}, expectType:GOURD, occurErr: false},
		&verifyParamStruct{input:"6s7s3cXnKdTd4h", expectRes:[]int{7,6,5,4,3}, expectType:STRAIGHT, occurErr: false},
		&verifyParamStruct{input:"As5s3cXnKdTd4h", expectRes:[]int{5,4,3,2,1}, expectType:STRAIGHT, occurErr: false},
		&verifyParamStruct{input:"JsQs3cXnKdTd4h", expectRes:[]int{14,13,12,11,10}, expectType:STRAIGHT, occurErr: false},
		&verifyParamStruct{input:"JsQsAcXnKdTd4h", expectRes:[]int{14,13,12,11,10}, expectType:STRAIGHT, occurErr: false},
		&verifyParamStruct{input:"5sAs3sXnKsTs4h", expectRes:[]int{14,14,13,10,5}, expectType:FLUSH, occurErr: false},
		&verifyParamStruct{input:"Js8s7sXnKdTh4s", expectRes:[]int{14,11,8,7,4}, expectType:FLUSH, occurErr: false},
		&verifyParamStruct{input:"Js8s3sXnKsTd4h", expectRes:[]int{14,13,11,8,3}, expectType:FLUSH, occurErr: false},
		&verifyParamStruct{input:"3s3s3cXnKdTd4h", expectRes:[]int{3,3,3,3,13}, expectType:FOUR_SAME, occurErr: false},
		&verifyParamStruct{input:"3s3s4cXn3d4d4h", expectRes:[]int{4,4,4,4,3}, expectType:FOUR_SAME, occurErr: false},
		&verifyParamStruct{input:"3s3s3cXnKdTd3h", expectRes:[]int{3,3,3,3,14}, expectType:FOUR_SAME, occurErr: false},
		&verifyParamStruct{input:"5sAs3sXnAsAs4h", expectRes:[]int{14,14,14,14,5}, expectType:FOUR_SAME, occurErr: false},
		&verifyParamStruct{input:"Js9s9sXnKsTs4h", expectRes:[]int{13,12,11,10,9}, expectType:STRAIGHT_FLUSH, occurErr: false},
		&verifyParamStruct{input:"5sAs3sXnKsTs4s", expectRes:[]int{5,4,3,2,1}, expectType:STRAIGHT_FLUSH, occurErr: false},
		&verifyParamStruct{input:"KsAs3sXnJsQs4h", expectRes:[]int{14,13,12,11,10}, expectType:ROYAL_FLUSH, occurErr: false},
		// 2 ghost
		&verifyParamStruct{input:"7c9s3cXn2dXnQh", expectRes:[]int{12,12,12,9,7}, expectType:THREE_SAME, occurErr: false},
		&verifyParamStruct{input:"6c9s3cXnKdXn4h", expectRes:[]int{7,6,5,4,3}, expectType:STRAIGHT, occurErr: false},
		&verifyParamStruct{input:"6c9s3cXn8dXnQh", expectRes:[]int{12,11,10,9,8}, expectType:STRAIGHT, occurErr: false},
		&verifyParamStruct{input:"6cTs3cXn8dXnQh", expectRes:[]int{12,11,10,9,8}, expectType:STRAIGHT, occurErr: false},
		&verifyParamStruct{input:"6cTs3cXn8dXnQc", expectRes:[]int{14,14,12,6,3}, expectType:FLUSH, occurErr: false},
		&verifyParamStruct{input:"4c9s3cXnKdXn4c", expectRes:[]int{4,4,4,4,13}, expectType:FOUR_SAME, occurErr: false},
		&verifyParamStruct{input:"4c9s3cXnAdXn4c", expectRes:[]int{4,4,4,4,14}, expectType:FOUR_SAME, occurErr: false},
		&verifyParamStruct{input:"4cAsAcXnQdXn4c", expectRes:[]int{14,14,14,14,12}, expectType:FOUR_SAME, occurErr: false},
		&verifyParamStruct{input:"4c4sAcXnQdXn4c", expectRes:[]int{4,4,4,4,14}, expectType:FOUR_SAME, occurErr: false},
		&verifyParamStruct{input:"AcAsAcXnQdXn4c", expectRes:[]int{14,14,14,14,13}, expectType:FOUR_SAME, occurErr: false},
		&verifyParamStruct{input:"4c5c3cXnAcXn4c", expectRes:[]int{7,6,5,4,3}, expectType:STRAIGHT_FLUSH, occurErr: false},
		&verifyParamStruct{input:"4c5h3cXnAcXn4c", expectRes:[]int{5,4,3,2,1}, expectType:STRAIGHT_FLUSH, occurErr: false},
		&verifyParamStruct{input:"6c9s3cXnKdXn4c", expectRes:[]int{7,6,5,4,3}, expectType:STRAIGHT_FLUSH, occurErr: false},
		&verifyParamStruct{input:"QcKhTcXnAcXnQc", expectRes:[]int{14,13,12,11,10}, expectType:ROYAL_FLUSH, occurErr: false},
		&verifyParamStruct{input:"QcKcTcXnAcXnQc", expectRes:[]int{14,13,12,11,10}, expectType:ROYAL_FLUSH, occurErr: false},
		// 3 ghost
		&verifyParamStruct{input:"6sXn3cXnXnTd4h", expectRes:[]int{10,10,10,10,6}, expectType:FOUR_SAME, occurErr: false},
		&verifyParamStruct{input:"6sXn6cXnXnTd4h", expectRes:[]int{10,10,10,10,6}, expectType:FOUR_SAME, occurErr: false},
		&verifyParamStruct{input:"6sXn6cXnXnTd6h", expectRes:[]int{10,10,10,10,6}, expectType:FOUR_SAME, occurErr: false},
		&verifyParamStruct{input:"6sXn6cXnXnAd4h", expectRes:[]int{14,14,14,14,6}, expectType:FOUR_SAME, occurErr: false},
		&verifyParamStruct{input:"6sXn6cXnXnAdAh", expectRes:[]int{14,14,14,14,13}, expectType:FOUR_SAME, occurErr: false},
		&verifyParamStruct{input:"9sXn3cXnXnTd4c", expectRes:[]int{7,6,5,4,3}, expectType:STRAIGHT_FLUSH, occurErr: false},
		&verifyParamStruct{input:"6cXn3cXnXnTd4c", expectRes:[]int{8,7,6,5,4}, expectType:STRAIGHT_FLUSH, occurErr: false},
		&verifyParamStruct{input:"8cXn3cXnXnTc4c", expectRes:[]int{12,11,10,9,8}, expectType:STRAIGHT_FLUSH, occurErr: false},
		&verifyParamStruct{input:"AcXn3cXnXn4h9c", expectRes:[]int{5,4,3,2,1}, expectType:STRAIGHT_FLUSH, occurErr: false},
		&verifyParamStruct{input:"AcXn3cXnXnTc4c", expectRes:[]int{14,13,12,11,10}, expectType:ROYAL_FLUSH, occurErr: false},
		// 4 ghost
		&verifyParamStruct{input:"6cXn3cXnXn4hXn", expectRes:[]int{10,9,8,7,6}, expectType:STRAIGHT_FLUSH, occurErr: false},
		&verifyParamStruct{input:"QcXn3cXnXn4hXn", expectRes:[]int{14,13,12,11,10}, expectType:ROYAL_FLUSH, occurErr: false},
		&verifyParamStruct{input:"AcXn3cXnXn4hXn", expectRes:[]int{14,13,12,11,10}, expectType:ROYAL_FLUSH, occurErr: false},
		// 5 ghost
		&verifyParamStruct{input:"XnXn3cXnXn4hXn", expectRes:[]int{14,13,12,11,10}, expectType:ROYAL_FLUSH, occurErr: false},
		// 6 ghost
		&verifyParamStruct{input:"XnXn3cXnXnXnXn", expectRes:[]int{14,13,12,11,10}, expectType:ROYAL_FLUSH, occurErr: false},
		// length large than 7 case
		&verifyParamStruct{input:"6s7s3cAcKdTd4h2c9s", expectRes:[]int{14,13,10,9,7}, expectType:SIMPLE, occurErr: false},
		&verifyParamStruct{input:"6s7s3cAcKdTd4h2c3s", expectRes:[]int{3,3,14,13,10}, expectType:ONE_PAIRS, occurErr: false},
		&verifyParamStruct{input:"6s7s3c2cKdTd4h2c3s", expectRes:[]int{3,3,2,2,13}, expectType:TWO_PAIRS, occurErr: false},
		&verifyParamStruct{input:"6s7s3c3cKdTd4h2c3s", expectRes:[]int{3,3,3,13,10}, expectType:THREE_SAME, occurErr: false},
		&verifyParamStruct{input:"6s5s3cAcAdAd4h2c9s", expectRes:[]int{6,5,4,3,2}, expectType:STRAIGHT, occurErr: false},
		&verifyParamStruct{input:"6d5d3cAcAdAd4h2c9d", expectRes:[]int{14,14,9,6,5}, expectType:FLUSH, occurErr: false},
		&verifyParamStruct{input:"6d2d3cAcAdAd4h2c9d", expectRes:[]int{14,14,14,2,2}, expectType:GOURD, occurErr: false},
		&verifyParamStruct{input:"4c2d3cAcAdAd4h2c5c", expectRes:[]int{5,4,3,2,1}, expectType:STRAIGHT_FLUSH, occurErr: false},
		&verifyParamStruct{input:"4cQc3cAcJcKcTc2c5c", expectRes:[]int{14,13,12,11,10}, expectType:ROYAL_FLUSH, occurErr: false},
		&verifyParamStruct{input:"6s7s3cXn3d3dXn2c3s", expectRes:[]int{7,6,5,4,3}, expectType:STRAIGHT_FLUSH, occurErr: false},
	}
	// verify
	for _, param := range verifyParams {
		pokerCard := NewPokerCard(param.input)
		err := pokerCard.Parse()
		// occur error --目前没想到error的case
		if param.occurErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		// expect type
		assert.Equal(t, param.expectType, pokerCard.Type, param.input)
		// expect result
		assert.Equal(t, param.expectRes, pokerCard.ResultFaces, param.input)
	}
}