package slice_test

import "testing"

// 切片声明测试
// 切片是一个结构体，[]type,len,cap ,其中len个元素会被初始化为默认值0，未初始化的元素不可访问
func TestSliceInit(t *testing.T) {
   // 方式一：跟数组的区别就在于不必元素的个数
   var a []int 
   // len是可访问的元素个数，cap是代表容量 
   t.Log(len(a), cap(a)) // 0 0
   a = append(a, 1) // 添加元素
   t.Log(len(a), cap(a), a[0]) // 1 1 1

   // 方式二
   b := []int{}
   t.Log(len(b), cap(b)) // 0 0

   // 方式三
   c := []int{ 1, 2, 3 }
   t.Log(len(c), cap(c)) // 3 3 

   // 方式四：长度3，容量5
   d := make([]int, 3, 5)
   t.Log(len(d), cap(d)) // 3 5
}

// 测试切片cap的扩容，当容量不够时，切片会乘以2的自增长
func TestSliceGrowing(t *testing.T) {
   b := []int{}
   for i := 0;i < 10;i++ {
      b = append(b, i)
      t.Log(len(b), cap(b))
   }
}


// 测试切片共享存储空间
func TestSliceShareMemory(t *testing.T) {
   year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug",
      "Sep", "Oct", "Nov", "Dec"}
   Q2 := year[3:6]
   // [Apr May Jun] 3 9   容量9是从Apr开始到末尾
   t.Log(Q2, len(Q2), cap(Q2))  

   summer := year[5:8]
   // [Jun Jul Aug] 3 7   容量7是从Jun开始到末尾
   t.Log(summer, len(summer), cap(summer)) 

   // summber[0]和Q2[2]指向同一个空间，此时修改了同一个存储空间
   summer[0] = "unknow" 
   t.Log(summer[0], Q2[2]) // unknow unknow
}
