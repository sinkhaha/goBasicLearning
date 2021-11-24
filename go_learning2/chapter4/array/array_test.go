package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int // 声明占3个空间的数组，并初始化为默认0
	t.Log(arr[1], arr[2])

	arr1 := [4]int{1, 2, 3, 4}   // 4个空间的数组，并初始化
	arr3 := [...]int{1, 3, 4, 5} // 不定义多少个元素，由初始化值决定
	arr1[1] = 5

	t.Log(arr1, arr3)
}

// 测试数组的遍历
func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	// 普通遍历
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	// 方式二：go特有的遍历,idx(索引)和e(元素)必须同时使用，不然过不了编译
	for idx, e := range arr3 {
		t.Log(idx, e)
	}

	// _ 表示忽略index值，即索引不需要用到
	for _, e := range arr3 {
		t.Log(e)
	}
}

// 测是数组的截取
func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3_sec := arr3[:] // 截取全部，第一个数不写默认从0开始，第二个数不写默认到末尾,不支持负数
	t.Log(arr3_sec)
}
