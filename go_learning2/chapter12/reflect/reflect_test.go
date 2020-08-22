package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
	"errors"
)

// 检查类型公共函数
func CheckType(v interface{}) {
	// 返回类型
	t := reflect.TypeOf(v)
	// 判断类型
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("不知道是什么类型", t)
	}
}

// 1.测试使用反射检查类型
func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(f)
	CheckType(&f)
}

// 2.测试通过反射获取类型和值
func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	// 通过反射获取类型和值
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	// 通过反射的值反过来获取类型
	t.Log(reflect.ValueOf(f).Type())
}


// 通过反射编写灵活代码
type Employee struct {
	EmployeeID string
	Name       string `haha:"normal"`  // struct tag,是一个key/value结构
	Age        int
}

// 更新年龄的方法
func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

// 3.测试反射获取结构体的成员，执行方法
func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	// 通过反射按名字获取成员的值
	t.Logf("Name: value(%[1]v), Type(%[1]T) ", reflect.ValueOf(*e).FieldByName("Name"))
	
	// 通过反射按名字获取成员的类型
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' field.")
	} else {
		t.Log("nameField", nameField)
		// 通过反射取得struct的tag
		t.Log("Tag:haha", nameField.Tag.Get("haha"))
	}

	// 反射调用方法
	reflect.ValueOf(e).MethodByName("UpdateAge").
		Call([]reflect.Value{reflect.ValueOf(15)})

	t.Log("Updated Age:", e)
}


type Customer struct {
	CookieID string
	Name     string
	Age      int
}

// 4.测试通过反射比较两个map,比较两个切片，比较两个对象
func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 4: "three"}
	
	t.Log("a==b?", reflect.DeepEqual(a, b)) // false

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}

	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2)) // true
	t.Log("s1 == s3?", reflect.DeepEqual(s1, s3)) // false

	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"1", "Mike", 40}
	fmt.Println(c1 == c2) // true
	fmt.Println(reflect.DeepEqual(c1, c2)) // true
}

// 5. 测试 填充空对象
func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 30}
	e := Employee{}

	// 把值放在map中,填充一个空对象,被填充对象中有对应的变量则会填充相应的值
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}

	t.Log(e)
	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}

// 通用的填充设置方法
func fillBySettings(st interface{}, settings map[string]interface{}) error {

	// func (v Value) Elem() Value
	// Elem returns the value that the interface v contains or that the pointer v points to.
	// It panics if v's Kind is not Interface or Ptr.
	// It returns the zero Value if v is nil.

	// 判断必须是一个指针类型
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		return errors.New("the first param should be a pointer to the struct type.")
	}

	// Elem() 获取指针指向的值
	if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
		return errors.New("the first param should be a pointer to the struct type.")
	}

	if settings == nil {
		return errors.New("settings is nil.")
	}

	var (
		field reflect.StructField
		ok    bool
	)

	// 遍历进行填充
	for k, v := range settings {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}

