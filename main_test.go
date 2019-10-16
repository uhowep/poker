package poker

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
	"time"
)

func TestCompare(t *testing.T) {
	type validStruct struct {
		Alice string `json:"alice"`
		Bob string `json:"bob"`
		Result int `json:"result"`
	}
	var validData map[string][]validStruct
	// read file
	validJsons, err := ioutil.ReadFile("json/valid.json")
	assert.NoError(t, err)
	// unmarshal json
	assert.NoError(t, json.Unmarshal(validJsons, &validData))
	// get the valid compare slice
	validSlice := validData["matches"]
	if len(validSlice) <= 0 {
		assert.Fail(t, "no matches field in valid.json file")
	}
	// compare
	for _, val := range validSlice {
		lPokerCard := NewPokerCard(val.Alice)
		assert.NoError(t, lPokerCard.Parse())
		rPokerCard := NewPokerCard(val.Bob)
		assert.NoError(t, rPokerCard.Parse())
		assert.Equal(t, val.Result, Compare(lPokerCard, rPokerCard), val.Alice + " " + val.Bob)
	}
	// calculate use time for dealing with
	hands := 0
	startTime := time.Now().Nanosecond()
	for _, val := range validSlice {
		lCard := NewPokerCard(val.Alice)
		_ = lCard.Parse()
		rCard := NewPokerCard(val.Bob)
		_ = rCard.Parse()
		_ = Compare(lCard, rCard)
		hands++
	}
	endTime := time.Now().Nanosecond()
	handleStr := fmt.Sprintf("%s %d %s %d %s ","compare", hands, "hands, spending", endTime - startTime, "nano second")
	fmt.Println(handleStr)
	// return
	return
}

