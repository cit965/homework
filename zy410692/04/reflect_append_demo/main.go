package main

import (
	"fmt"
	"reflect"
)

// reflect append方法使用 主要是将元素添加到切片里面
func main() {

	a := []int{2, 4}
	// 这里一般都传入a的指针，因为反射有个可写的地址，如果传入reflect.ValueOf(a) 会panic
	// 一般是通过传入reflect.ValueOf(&a),再通过Elem()函数去掉指针，算作固定套路吧

	//这也就是if pv.Type().Kind() != reflect.Ptr || pv.Type().Elem().Kind() != reflect.Slice
	//这个条件判断应该就是这么来的
	var b reflect.Value = reflect.ValueOf(&a)
	b = b.Elem()
	fmt.Println("Slice:", a)
	b = reflect.Append(b, reflect.ValueOf(80))
	fmt.Println("Slice after appending data", b)
}
