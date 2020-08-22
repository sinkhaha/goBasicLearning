package array_test

import "testing"

// 测试数组的初始化
func TestArrayInit(t *testing.T) {
   // 方式一： 声明3个空间的数组，并初始化为默认0
   var a [3]int 
   a[0] = 1
   t.Log(a[0], a[1], a[2], len(a))

   // 方式二：声明同时初始化
   b := [3]int{ 1, 2, 3 }
   t.Log(b[0], len(b))

   // 方式三： 二维数组初始化
   c := [2][2]int{ {1,2}, {3,4} }
   t.Log(c[0], len(c))

   // 方式四：不定义多少个元素，由初始化值决定
   d := [...]int{ 1, 2, 3, 4 }
   t.Log(d[0], len(d))
}

// 测试数组的遍历
func TestArrayTravel(t *testing.T) {
	arr := [...]int{ 1, 2, 3, 4 }
	// 方式一：普通遍历
	for i := 0; i < len(arr); i++ {
		t.Log(i, arr[i])
	}
	
	// (推荐)方式二：go特有的遍历,idx(索引)和e(元素)必须同时使用，不然过不了编译
	for idx, e := range arr {
		t.Log(idx, e)
	}

	// 方式二：go特有的遍历,_表示不关心idx索引的值
	for _, e := range arr {
		t.Log(e)
	}
	for idx, _ := range arr {
		t.Log(idx)
	}
 }

// 测试数组的截取
// arr[开始索引(包含):结束索引(不包含)]
// 第一个数不写默认从0开始，第二个数不写默认到末尾,不支持负数
func TestArraySection(t *testing.T) {
	arr := [...]int{ 1, 2, 3, 4 }
	t.Log(arr[1:3]) // [2, 3]
	t.Log(arr[:len(arr)]) // 可省略第一个数，默认从0开始，[1,2,3,4]
	t.Log(arr[1:]) // 可省略第二个数，默认到末尾[1,2,3,4]
	
	arr2 := arr[0:4]
	t.Log(arr2) // [1,2,3,4]
}
