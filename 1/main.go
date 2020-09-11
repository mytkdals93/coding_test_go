package main

import "fmt"

func calc(arr []int, index int) ([]int, error) {
	result := []int{}
	if index < 0 {
		return result, fmt.Errorf("index는 0보다 작을 수 없습니다.: %d", index)
	}
	if index >= len(arr) {
		return result, fmt.Errorf("index는 배열의 lenth보다 작아야합니다.: %d", index)
	}
	if index == 0 {
		result = arr[1:]
		return result, nil
	}
	if index == len(arr) {
		result = arr[0 : len(arr)-1]
		return result, nil
	}

	result = append(result, arr[0:index]...)
	result = append(result, arr[index+1:len(arr)]...)

	return result, nil
}
func main() {

}
