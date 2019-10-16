package poker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type verifyParamStruct struct {
		inputSlice []int
		expect []int
	}
	// build verify params
	verifyParams := []*verifyParamStruct{
		&verifyParamStruct{inputSlice:[]int{}, expect:[]int{}},
		&verifyParamStruct{inputSlice:[]int{1}, expect:[]int{1}},
		&verifyParamStruct{inputSlice:[]int{-1, 2, 5}, expect:[]int{5, 2, -1}},
		&verifyParamStruct{inputSlice:[]int{8, 6, 4, -2, -3}, expect:[]int{8, 6, 4, -2, -3}},
		&verifyParamStruct{inputSlice:[]int{3, -9, 7, -2, 5}, expect:[]int{7, 5, 3, -2, -9}},
		&verifyParamStruct{inputSlice:[]int{3, 9, 0, -2, 5}, expect:[]int{9, 5, 3, 0, -2}},
	}
	// verify
	for _, param := range verifyParams {
		quickSort(param.inputSlice)
		assert.Equal(t, param.expect, param.inputSlice)
	}
	// return
	return
}

func TestSortByPoker(t *testing.T) {
	type verifyParamStruct struct {
		inputSlice []int
		expectSlice []int
		expectSameTimes int
	}
	// build verify params
	verifyParams := []*verifyParamStruct{
		&verifyParamStruct{inputSlice:[]int{9,9,6,3,3}, expectSlice:[]int{9,9,3,3,6}, expectSameTimes:2},
		&verifyParamStruct{inputSlice:[]int{9,6,6,6,3}, expectSlice:[]int{6,6,6,9,3}, expectSameTimes:3},
		&verifyParamStruct{inputSlice:[]int{9,7,6,5,3}, expectSlice:[]int{9,7,6,5,3}, expectSameTimes:0},
		&verifyParamStruct{inputSlice:[]int{13,3,3,3,3}, expectSlice:[]int{3,3,3,3,13}, expectSameTimes:6},
		&verifyParamStruct{inputSlice:[]int{3,3,2,2,3}, expectSlice:[]int{3,3,3,2,2}, expectSameTimes:4},
		&verifyParamStruct{inputSlice:[]int{13,3,7,4,3}, expectSlice:[]int{3,3,13,7,4}, expectSameTimes:1},
	}
	// verify
	for _, param := range verifyParams {
		// deal with
		sameTimes := sortByPoker(param.inputSlice)
		// 结果符合预期
		assert.Equal(t, param.expectSlice, param.inputSlice)
		assert.Equal(t, param.expectSameTimes, sameTimes)
	}
	// return
	return
}
