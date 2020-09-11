package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	assert := assert.New(t)
	tc := []int{1, 2, 3, 4, 5, 6, 10, 100, 200, 300}
	// 음수값 삭제
	result, err := calc(tc, -3)
	assert.Error(err)
	// 길이보다 큰 값 삭제
	result, err = calc(tc, len(tc)+2)
	assert.Error(err)
	//0 삭제
	result, err = calc(tc, 0)
	assert.NoError(err)
	assert.Equal([]int{2, 3, 4, 5, 6, 10, 100, 200, 300}, result)
	//2 삭제
	result, err = calc(tc, 2)
	assert.NoError(err)
	assert.Equal([]int{1, 2, 4, 5, 6, 10, 100, 200, 300}, result)
	//마지막 앞 삭제
	result, err = calc(tc, len(tc)-2)
	assert.NoError(err)
	assert.Equal([]int{1, 2, 3, 4, 5, 6, 10, 100, 300}, result)
	//마지막 삭제
	result, err = calc(tc, len(tc)-1)
	assert.NoError(err)
	assert.Equal([]int{1, 2, 3, 4, 5, 6, 10, 100, 200}, result)

}
