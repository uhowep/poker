package poker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type verifyParamStruct struct {
		inSlice []int
		expect []int
	}
	// build verify params
	verifyParams := []*verifyParamStruct{
		&verifyParamStruct{inSlice:[]int{}, expect:[]int{}},
		&verifyParamStruct{inSlice:[]int{1}, expect:[]int{1}},
		&verifyParamStruct{inSlice:[]int{-1, 2, 5}, expect:[]int{5, 2, -1}},
		&verifyParamStruct{inSlice:[]int{8, 6, 4, -2, -3}, expect:[]int{8, 6, 4, -2, -3}},
		&verifyParamStruct{inSlice:[]int{3, -9, 7, -2, 5}, expect:[]int{7, 5, 3, -2, -9}},
		&verifyParamStruct{inSlice:[]int{3, 9, 0, -2, 5}, expect:[]int{9, 5, 3, 0, -2}},
	}
	// verify
	for _, param := range verifyParams {
		quickSort(param.inSlice)
		assert.Equal(t, param.expect, param.inSlice)
	}
	// return
	return
}

func TestBuildStraight(t *testing.T) {
	type inSrt struct {
		inSlice []int
		inGhostNum int
	}
	type verifyParamStruct struct {
		input inSrt
		expectRes []int
		expectSuccess bool
	}
	// build verify params
	verifyParams := []*verifyParamStruct{
		// failure case
		&verifyParamStruct{input:inSrt{inSlice:[]int{5,3,2},inGhostNum:0},expectRes:[]int{},expectSuccess:false,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{5,3,2},inGhostNum:1},expectRes:[]int{},expectSuccess:false,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{9,6,3},inGhostNum:0},expectRes:[]int{},expectSuccess:false,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{9,6,3},inGhostNum:1},expectRes:[]int{},expectSuccess:false,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{14,4,3,2},inGhostNum:0},expectRes:[]int{},expectSuccess:false,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{14,7,6,4,3,2},inGhostNum:0},expectRes:[]int{},expectSuccess:false,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{13,10,10,5,4,3,2},inGhostNum:0},expectRes:[]int{},expectSuccess:false,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{14,13,12,11,5,4,3},inGhostNum:0},expectRes:[]int{},expectSuccess:false,},
		// success case
		&verifyParamStruct{input:inSrt{inSlice:[]int{5,3,2},inGhostNum:2},expectRes:[]int{6,5,4,3,2},expectSuccess:true,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{5,3,2},inGhostNum:3},expectRes:[]int{7,6,5,4,3},expectSuccess:true,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{5,3,2},inGhostNum:4},expectRes:[]int{9,8,7,6,5},expectSuccess:true,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{5,3,2},inGhostNum:5},expectRes:[]int{14,13,12,11,10},expectSuccess:true,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{9,6,3,2},inGhostNum:2},expectRes:[]int{6,5,4,3,2},expectSuccess:true,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{10,7,6},inGhostNum:3},expectRes:[]int{11,10,9,8,7},expectSuccess:true,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{9,6,3,2},inGhostNum:3},expectRes:[]int{10,9,8,7,6},expectSuccess:true,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{9,8,3,2},inGhostNum:3},expectRes:[]int{12,11,10,9,8},expectSuccess:true,},
		// 将14转换成1
		&verifyParamStruct{input:inSrt{inSlice:[]int{14,4,3,2},inGhostNum:1},expectRes:[]int{5,4,3,2,1},expectSuccess:true,},
		// 优先不使用1
		&verifyParamStruct{input:inSrt{inSlice:[]int{14,4,3,2},inGhostNum:2},expectRes:[]int{6,5,4,3,2},expectSuccess:true,},
		// 针对5张
		&verifyParamStruct{input:inSrt{inSlice:[]int{8,7,6,5,4},inGhostNum:0},expectRes:[]int{8,7,6,5,4},expectSuccess:true,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{8,7,6,5,4},inGhostNum:1},expectRes:[]int{9,8,7,6,5},expectSuccess:true,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{8,7,6,5,4},inGhostNum:4},expectRes:[]int{12,11,10,9,8},expectSuccess:true,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{8,7,6,5,4},inGhostNum:5},expectRes:[]int{14,13,12,11,10},expectSuccess:true,},
		// 针对已到顶
		&verifyParamStruct{input:inSrt{inSlice:[]int{12,11,10,9,8},inGhostNum:3},expectRes:[]int{14,13,12,11,10},expectSuccess:true,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{13},inGhostNum:4},expectRes:[]int{14,13,12,11,10},expectSuccess:true,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{14},inGhostNum:5},expectRes:[]int{14,13,12,11,10},expectSuccess:true,},
		// 针对大于5张
		&verifyParamStruct{input:inSrt{inSlice:[]int{10,8,7,6,5,4,3},inGhostNum:0},expectRes:[]int{8,7,6,5,4},expectSuccess:true,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{10,8,7,6,5,4,2},inGhostNum:1},expectRes:[]int{10,9,8,7,6},expectSuccess:true,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{10,10,8,7,6,5,4},inGhostNum:2},expectRes:[]int{11,10,9,8,7},expectSuccess:true,},
		// 针对重复的情况
		&verifyParamStruct{input:inSrt{inSlice:[]int{8,8,7,6,5,4},inGhostNum:2},expectRes:[]int{10,9,8,7,6},expectSuccess:true,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{10,10,8,8,6,5},inGhostNum:3},expectRes:[]int{12,11,10,9,8},expectSuccess:true,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{14,13,13,13},inGhostNum:4},expectRes:[]int{14,13,12,11,10},expectSuccess:true,},
		&verifyParamStruct{input:inSrt{inSlice:[]int{14,13,13,12,12,11,10},inGhostNum:4},expectRes:[]int{14,13,12,11,10},expectSuccess:true,},
	}
	// verify
	for _, param := range verifyParams {
		// deal with
		result, isSuccess := buildStraight(param.input.inSlice, param.input.inGhostNum)
		// build success
		assert.Equal(t, param.expectSuccess, isSuccess)
		// equal expect
		assert.Equal(t, param.expectRes, result, param.input)
	}
	// return
	return
}

func TestBuildFlush(t *testing.T) {
	type verifyParamStruct struct {
		inMap map[string][]int
		inGhostNum int
		expectRes []int
		expectSuccess bool
	}
	// build verify params
	verifyParams := []*verifyParamStruct{
		&verifyParamStruct{
			inMap:map[string][]int{"s":[]int{3,2},"h":[]int{9,5},"d":[]int{12,12},"c":[]int{}},
			inGhostNum:2,
			expectRes:[]int{},
			expectSuccess:false,
		},
		&verifyParamStruct{
			inMap:map[string][]int{"s":[]int{3,2},"h":[]int{9,5},"d":[]int{12,12},"c":[]int{}},
			inGhostNum:3,
			expectRes:[]int{14,14,14,12,12},
			expectSuccess:true,
		},
		&verifyParamStruct{
			inMap:map[string][]int{"s":[]int{3,2},"h":[]int{9,5},"d":[]int{12,12},"c":[]int{13}},
			inGhostNum:4,
			expectRes:[]int{14,14,14,14,13},
			expectSuccess:true,
		},
		&verifyParamStruct{
			inMap:map[string][]int{"s":[]int{3,2},"h":[]int{9,8,5,4,2},"d":[]int{},"c":[]int{}},
			inGhostNum:0,
			expectRes:[]int{9,8,5,4,2},
			expectSuccess:true,
		},
		&verifyParamStruct{
			inMap:map[string][]int{"s":[]int{13,9,3,2},"h":[]int{12,9,8,4,2},"d":[]int{14},"c":[]int{}},
			inGhostNum:2,
			expectRes:[]int{14,14,13,9,3},
			expectSuccess:true,
		},
		&verifyParamStruct{
			inMap:map[string][]int{"s":[]int{12,11,8,4,2},"h":[]int{12,11,9,5,3},"d":[]int{14},"c":[]int{}},
			inGhostNum:2,
			expectRes:[]int{14,14,12,11,9},
			expectSuccess:true,
		},
		&verifyParamStruct{
			inMap:map[string][]int{"s":[]int{13,12,10,7,6},"h":[]int{13,12,10,9,2},"d":[]int{14},"c":[]int{}},
			inGhostNum:0,
			expectRes:[]int{13,12,10,9,2},
			expectSuccess:true,
		},
	}
	// verify
	for _, param := range verifyParams {
		maxResultVal, isSuccess := buildFlush(param.inMap, param.inGhostNum)
		// build success
		assert.Equal(t, param.expectSuccess, isSuccess)
		// equal expect
		assert.Equal(t, param.expectRes, maxResultVal, param.inMap)
	}
	// return
	return
}

func TestBuildStraightFlush(t *testing.T) {
	type inputStruct struct {
		inSlice []int
		inColorMap map[string][]int
		inGhostNum int
	}
	type verifyParamStruct struct {
		input inputStruct
		expectRes []int
		expectSuccess bool
	}
	// define test data
	verifyParams := []*verifyParamStruct{
		// failure case
		&verifyParamStruct{input:inputStruct{
			inSlice:    []int{5,3,2},
			inColorMap: map[string][]int{"h":[]int{5,3,2}},
			inGhostNum: 0,
		}, expectRes: []int{}, expectSuccess: false},
		&verifyParamStruct{input:inputStruct{
			inSlice:    []int{5,3,2},
			inColorMap: map[string][]int{"h":[]int{5,3,2}},
			inGhostNum: 1,
		}, expectRes: []int{}, expectSuccess: false},
		&verifyParamStruct{input:inputStruct{
			inSlice:    []int{14,4,3,2},
			inColorMap: map[string][]int{"h":[]int{14,4,3,2}},
			inGhostNum: 0,
		}, expectRes: []int{}, expectSuccess: false},
		&verifyParamStruct{input:inputStruct{
			inSlice:    []int{14,7,6,4,3,2},
			inColorMap: map[string][]int{"h":[]int{14,7,6,4,3,2}},
			inGhostNum: 0,
		}, expectRes: []int{}, expectSuccess: false},
		&verifyParamStruct{input:inputStruct{
			inSlice:    []int{14,13,12,11,10,9,8},
			inColorMap: map[string][]int{"h":[]int{13,12},"s":[]int{14,10,8},"c":[]int{11,9}},
			inGhostNum: 2,
		}, expectRes: []int{}, expectSuccess: false},
		// success case
		&verifyParamStruct{input:inputStruct{
			inSlice:    []int{5,3,2},
			inColorMap: map[string][]int{"h":[]int{5,3,2}},
			inGhostNum: 3,
		}, expectRes: []int{7,6,5,4,3}, expectSuccess: true},
		&verifyParamStruct{input:inputStruct{
			inSlice:    []int{14,10,5,3,2},
			inColorMap: map[string][]int{"h":[]int{14,10,5,3,2}},
			inGhostNum: 1,
		}, expectRes: []int{5,4,3,2,1}, expectSuccess: true},
		&verifyParamStruct{input:inputStruct{
			inSlice:    []int{5,3,2},
			inColorMap: map[string][]int{"h":[]int{5,3,2}},
			inGhostNum: 4,
		}, expectRes: []int{9,8,7,6,5}, expectSuccess: true},
		&verifyParamStruct{input:inputStruct{
			inSlice:    []int{5,3,2},
			inColorMap: map[string][]int{"h":[]int{5,3,2}},
			inGhostNum: 5,
		}, expectRes: []int{14,13,12,11,10}, expectSuccess: true},
		// max select case
		&verifyParamStruct{input:inputStruct{
			inSlice:    []int{14,13,12,11,10,9,8},
			inColorMap: map[string][]int{"h":[]int{12,13},"s":[]int{14,8},"c":[]int{9,11,10}},
			inGhostNum: 2,
		}, expectRes: []int{13,12,11,10,9}, expectSuccess: true},
		&verifyParamStruct{input:inputStruct{
			inSlice:    []int{14,13,12,11,10,9,8},
			inColorMap: map[string][]int{"h":[]int{12,13},"s":[]int{14,8},"c":[]int{9,11,10}},
			inGhostNum: 3,
		}, expectRes: []int{14,13,12,11,10}, expectSuccess: true},
		&verifyParamStruct{input:inputStruct{
			inSlice:    []int{14,13,12,8,6,4,3,2},
			inColorMap: map[string][]int{"h":[]int{4,2,3,14},"s":[]int{13,8},"c":[]int{6,12}},
			inGhostNum: 1,
		}, expectRes: []int{5,4,3,2,1}, expectSuccess: true},
		// same
		&verifyParamStruct{input:inputStruct{
			inSlice:    []int{14,13,13,6,6,4,3,2},
			inColorMap: map[string][]int{"h":[]int{4,2,3,14},"s":[]int{13,6},"c":[]int{6,13}},
			inGhostNum: 1,
		}, expectRes: []int{5,4,3,2,1}, expectSuccess: true},
		&verifyParamStruct{input:inputStruct{
			inSlice:    []int{14,12,12,11,10,9,9,4,3,2},
			inColorMap: map[string][]int{"h":[]int{4,2,3,14},"s":[]int{12,11,10,9},"c":[]int{9,12}},
			inGhostNum: 1,
		}, expectRes: []int{13,12,11,10,9}, expectSuccess: true},
	}
	// verify
	for _, param := range verifyParams {
		// deal with
		result, isSuccess := buildStraightFlush(param.input.inSlice, param.input.inColorMap, param.input.inGhostNum)
		// build success
		assert.Equal(t, param.expectSuccess, isSuccess)
		// equal expect
		assert.Equal(t, param.expectRes, result, param.input)
	}
}

func TestBuildSame(t *testing.T) {
	type inputStruct struct {
		inSlice []int
		inGhostNum int
	}
	type verifyParamStruct struct {
		input inputStruct
		expectRes []int
		expectSuccess bool
	}
	// define verify params
	verifyParams := []*verifyParamStruct{
		// failure case
		&verifyParamStruct{input:inputStruct{inSlice:[]int{8,5,4,4}, inGhostNum:0},expectRes:[]int{},expectSuccess:false},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{8,4,4}, inGhostNum:1},expectRes:[]int{},expectSuccess:false},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{4,4}, inGhostNum:2},expectRes:[]int{},expectSuccess:false},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{11,9,8,7,5,2}, inGhostNum:0},expectRes:[]int{},expectSuccess:false},
		// 5 ghost
		&verifyParamStruct{input:inputStruct{inSlice:[]int{8,7,7,4,}, inGhostNum:5},expectRes:[]int{14,14,14,14,13},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{4,2}, inGhostNum:5},expectRes:[]int{14,14,14,14,13},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{14,13,12,11,10}, inGhostNum:5},expectRes:[]int{14,14,14,14,13},expectSuccess:true},
		// 4 ghost
		&verifyParamStruct{input:inputStruct{inSlice:[]int{8,7,7,4,4}, inGhostNum:4},expectRes:[]int{14,14,14,14,8},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{14,14,13,13,2}, inGhostNum:4},expectRes:[]int{14,14,14,14,13},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{13,13,13,3,2}, inGhostNum:4},expectRes:[]int{14,14,14,14,13},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{11,10,10,10,10}, inGhostNum:4},expectRes:[]int{14,14,14,14,11},expectSuccess:true},
		// 3 ghost
		&verifyParamStruct{input:inputStruct{inSlice:[]int{11,10,10,10,10}, inGhostNum:3},expectRes:[]int{11,11,11,11,10},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{14,13,12,11}, inGhostNum:3},expectRes:[]int{14,14,14,14,13},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{14,14,12,11}, inGhostNum:3},expectRes:[]int{14,14,14,14,13},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{13,13,12}, inGhostNum:3},expectRes:[]int{13,13,13,13,14},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{11,11}, inGhostNum:3},expectRes:[]int{11,11,11,11,14},expectSuccess:true},
		// 2 ghost
		&verifyParamStruct{input:inputStruct{inSlice:[]int{14,13,12,11}, inGhostNum:2},expectRes:[]int{14,14,14,13,12},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{9,9,3,2}, inGhostNum:2},expectRes:[]int{9,9,9,9,3},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{9,3,3,3,2,2}, inGhostNum:2},expectRes:[]int{3,3,3,3,14},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{9,3,3,2,2,2}, inGhostNum:2},expectRes:[]int{3,3,3,3,9},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{9,9,9,9,2}, inGhostNum:2},expectRes:[]int{9,9,9,9,14},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{9,9,8,8,2}, inGhostNum:2},expectRes:[]int{9,9,9,9,8},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{9,8,8,2}, inGhostNum:2},expectRes:[]int{8,8,8,8,9},expectSuccess:true},
		// 1 ghost
		&verifyParamStruct{input:inputStruct{inSlice:[]int{9,8,7,2}, inGhostNum:1},expectRes:[]int{9,9,8,7,2},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{9,8,8,2}, inGhostNum:1},expectRes:[]int{8,8,8,9,2},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{9,8,8,4,4,2}, inGhostNum:1},expectRes:[]int{8,8,8,4,4},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{9,8,8,4,2,2}, inGhostNum:1},expectRes:[]int{8,8,8,2,2},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{8,4,4,4,2,2}, inGhostNum:1},expectRes:[]int{4,4,4,4,8},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{8,4,4,4,2,2,2}, inGhostNum:1},expectRes:[]int{4,4,4,4,8},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{8,4,4,4,4,2,2}, inGhostNum:1},expectRes:[]int{4,4,4,4,14},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{8,8,8,4,4,4,4,2}, inGhostNum:1},expectRes:[]int{8,8,8,8,4},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{8,8,8,8,4,4,4,4}, inGhostNum:1},expectRes:[]int{8,8,8,8,14},expectSuccess:true},
		// 0 ghost
		&verifyParamStruct{input:inputStruct{inSlice:[]int{11,9,8,8,5,2}, inGhostNum:0},expectRes:[]int{8,8,11,9,5},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{11,9,8,8,8,2}, inGhostNum:0},expectRes:[]int{8,8,8,11,9},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{11,9,9,8,8,8,2}, inGhostNum:0},expectRes:[]int{8,8,8,9,9},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{9,8,8,8,8,7,7,2}, inGhostNum:0},expectRes:[]int{8,8,8,8,9},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{9,8,8,6,3,3,2}, inGhostNum:0},expectRes:[]int{8,8,3,3,9},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{9,8,8,6,6,3,3,2}, inGhostNum:0},expectRes:[]int{8,8,6,6,9},expectSuccess:true},
		&verifyParamStruct{input:inputStruct{inSlice:[]int{8,8,8,6,6,3,3,3}, inGhostNum:0},expectRes:[]int{8,8,8,6,6},expectSuccess:true},
	}
	// verify
	for _, param := range verifyParams {
		result, isSuccess := buildSame(param.input.inSlice, param.input.inGhostNum)
		// check is success
		assert.Equal(t, param.expectSuccess, isSuccess, param.input)
		// check result
		assert.Equal(t, param.expectRes, result, param.input)
	}
	// return
	return
}