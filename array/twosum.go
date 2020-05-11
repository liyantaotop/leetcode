package array

import (
	"sort"
	"strconv"
	"strings"
)

func twoSum(nums []int, target int) []int {
	first, second := 0, 0
	for i:=0;  i<len(nums) ;i++  {
		for k:=(i+1);  k<len(nums) ;k++  {
              if nums[i]+nums[k] == target {
				  first  = i
				  second  = k
				  break
			  }
		}
		if first != 0 || second != 0 {
			break
		}
	}
	return []int{first, second}
}

func maximalSquare(matrix [][]byte) int {
     for k,v :=range matrix {

	}
}