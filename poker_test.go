package poker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseStringToCard(t *testing.T) {
	// check error
	var err error
	cardParam1 := "3h9cKs"
	cardParam2 := "3h9cKs1d8h"
	_, _, err = parseStringToCard(cardParam1)
	_, _, err = parseStringToCard(cardParam2)
	assert.Error(t, err)
	// check transfer
	cardParam3 := "3hKsJsQhTc"
	faces, colorCount, err := parseStringToCard(cardParam3)
	assert.NoError(t, err)
	assert.Equal(t, []int{3,13,11,12,10}, faces)
	assert.Equal(t, 3, colorCount)
}

func TestPokerCard_ParseType(t *testing.T) {
	type verifyParamStruct struct {
		inputCard string
		expectType int
	}
	// define verify params
	verifyParams := []verifyParamStruct{
		{inputCard:"8hQcTh4s2s", expectType:SIMPLE},
		{inputCard:"TsKsAsQsJs", expectType:ROYAL_FLUSH},
		{inputCard:"TsKsAhQsJs", expectType:STRAIGHT},
		{inputCard:"4s5s6s7s8s", expectType:STRAIGHT_FLUSH},
		{inputCard:"8hQhTh4h2h", expectType:FLUSH},
		{inputCard:"ThQhTh4h2h", expectType:FLUSH},		// FLUSH large than PAIR
		{inputCard:"ThQhThTh2h", expectType:FLUSH},		// FLUSH large than THREE_SAME
		{inputCard:"8s8h8c8d6s", expectType:FOUR_SAME},
		{inputCard:"8s8s8s8s6s", expectType:FOUR_SAME},	// FOUR_SAME large than FLUSH
		{inputCard:"8h8c8d6s6c", expectType:GOURD},
		{inputCard:"JhJdQc3dJc", expectType:THREE_SAME},
		{inputCard:"3hKs3cKc9d", expectType:TWO_PAIRS},
		{inputCard:"QcQh4h5h6c", expectType:ONE_PAIRS},
	}
	// verify
	for _, param := range verifyParams {
		pokerCard := NewPokerCard(param.inputCard)
		// no error to execute parse type
		assert.NoError(t, pokerCard.ParseType())
		// equal card type
		assert.Equal(t, param.expectType, pokerCard.CardType, param.inputCard)
	}
	// return
	return
}

