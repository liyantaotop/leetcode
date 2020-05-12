package array

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestThreeSum(t *testing.T) {
	paramsIn := []int{-1,0,1,2,-1,-4}
	result := threeSumJiaBi(paramsIn)
	assert := assert.New(t)
	assert.Equal(result, [][]int{{0, 0, 0}}, "they should be equal")
}

//两层循环  获取a+b 在从set 里寻找-c  去重
func ThreeSum(nums []int) [][]int {
	//生成set
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
			first, second := nums[i], nums[k]
			target := -(first + second)
			//标记 移除set中元素
			s[first] = s[first] - 1
			s[second] = s[second] - 1
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
			}
			//添加移除的数据
			s[first] = s[first] + 1
			s[second] = s[second] + 1
		}
	}
	//返回结果
	ss := [][]int{}
	for _, v := range result {
		ss = append(ss, v)
	}
	return ss
}

//夹逼法  1.排序
//2.遍历  // TODO 未完成
func threeSumJiaBi(nums []int) [][]int {
	sort.Ints(nums)
	result := make(map[string][]int)
	fmt.Printf("first :%+v \n ",nums)

	arrLength := len(nums)
	for i := 0; i < arrLength -2; i++ {
		firstItem := nums[i]
		leftIndex := i+1
		startItem := nums[leftIndex]
		rightIndex :=arrLength-1
		endItem := nums[rightIndex]
		for {
			fmt.Printf("first :%d start: %d end :%d \n ",firstItem,startItem,endItem)

			if firstItem == -(startItem + endItem){
				//找到一组数据
				dd := []int{firstItem, startItem, endItem}
				sort.Ints(dd)
				sInt := make([]string, len(dd))
				for _, v := range dd {
					sInt = append(sInt, strconv.Itoa(v))
				}
				cc := strings.Join(sInt, "-")
				result[cc] = dd
				break
			}else if firstItem > -(startItem + endItem) {

				rightIndex = rightIndex - 1
				endItem = nums[leftIndex]
				continue
			}else if firstItem < -(startItem + endItem) {
				leftIndex = leftIndex+1
				startItem = nums[leftIndex]
				continue
			}
		}
	}
	//返回结果
	ss := [][]int{}
	for _, v := range result {
		ss = append(ss, v)
	}
	return ss
}
