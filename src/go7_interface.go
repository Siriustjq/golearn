package main

import "fmt"

//定义接口
type phone interface {
	call()
}

//利用两个子类型来实现该接口（实现其方法就ok，并无类似于implament等关键字）
type apple struct{}

type nokia struct{}

func (na nokia) call() {
	fmt.Println("this is nokia!")
}

func (ae apple) call() {
	fmt.Println("this is apple!")
}

func main() {
	//多态的体现
	var phone1 phone = nokia{}
	var phone2 phone = apple{}
	//判断该结构体到底是哪种类型的，虽然都是实现了phone接口，但是类型并不一致。
	if value, ok := phone1.(apple); !ok {
		value.call()
	}
	if value, ok := phone2.(apple); ok {
		value.call()
	}
	// switch还将会用来判断一个接口变量的类型
	switch value := phone1.(type) {
	case apple:
		fmt.Println("this is a apple !!!!!", value)
	case nokia:
		fmt.Println("this is a nokia !!!!!", value)
	}
}
