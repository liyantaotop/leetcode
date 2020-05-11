package array

import (
	"sort"
	"strconv"
	"strings"
)


//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。

//两层循环  获取a+b 在从set 里寻找-c  去重
func ThreeSum1(nums []int) [][]int {
	s := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, exist := s[nums[i]]; exist {
			temp := s[nums[i]] + 1
			s[nums[i]] = temp
			continue
		}
		s[nums[i]] = 1
	}
	result := make(map[string][]int)
	arrenLth := len(nums)
	for i := 0; i < arrenLth; i++ {
		for k := (i + 1); k < arrenLth; k++ {
			target := -(nums[i] + nums[k])
			//移除set中元素
			temp := s[nums[i]] - 1
			s[nums[i]] = temp
			temp2 := s[nums[k]] - 1
			s[nums[k]] = temp2
			v, exist := s[target]
			if exist && v > 0 { //去重
				dd := []int{nums[i], nums[k], target}
				sort.Ints(dd)
				sInt := make([]string, len(dd))
				for _, v := range dd {
					sInt = append(sInt, strconv.Itoa(v))
				}
				cc := strings.Join(sInt, "-")
				result[cc] = dd
				temp3 := s[target] - 1
				s[target] = temp3
			} else {
				//添加移除的数据
				temp := s[nums[i]] + 1
				s[nums[i]] = temp
				temp2 := s[nums[k]] + 1
				s[nums[k]] = temp2
			}
		}
	}
	ss := [][]int{}
	for _, v := range result {
		ss = append(ss, v)
	}
	return ss
}

//夹逼法
func threeSum2(nums []int) [][]int {
	s := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		s[nums[i]] = i
	}
	result := make(map[string][]int)
	for i := 0; i < len(nums); i++ {
		for k := (i + 1); k < len(nums); k++ {
			target := nums[i] + nums[k]
			_, exist := s[target]
			if exist {
				dd := []int{nums[i], nums[k], target}
				sort.Ints(dd)
				sInt := make([]string, len(dd))
				for _, v := range dd {
					sInt = append(sInt, strconv.Itoa(v))
				}
				cc := strings.Join(sInt, "-")
				result[cc] = dd
			}
		}
	}
	ss := [][]int{}
	for _, v := range result {
		ss = append(ss, v)
	}
	return ss
}
