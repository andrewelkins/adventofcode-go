package mathy

import (	
	"fmt"
	"strconv"
)

func MaxInt(nums ...int) int {
	maxNum := nums[0]
	for _, v := range nums {
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum
}
func MinInt(nums ...int) int {
	minNum := nums[0]
	for _, v := range nums {
		if v < minNum {
			minNum = v
		}
	}
	return minNum
}

func AbsInt(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

func SumIntSlice(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
}

func MultiplyIntSlice(nums []int) int {
	product := 1
	for _, n := range nums {
		product *= n
	}
	return product
}

func Sum(arr []any) int {
    n := 0
    for _, v := range arr {
        switch v := v.(type) {
        case int:
            n += v
        case string:
            i, err := strconv.Atoi(v)
            if err != nil {
                panic(err)
            }
            n += i
        default:
            panic(fmt.Sprintf("unsupported type %T", v))
        }
    }
    return n
}


func Reverse(s string) (result string) {
	for _,v := range s {
	  result = string(v) + result
	}
	return 
  }