package structure

import (
	"fmt"
	"testing"
)

// 装饰器模式

//定义抽象的组建
type Component interface {
	Operate()
}

type Component1 struct {
}

func (c Component1) Operate() {
	fmt.Println("c1 operate")
}

//定义一个抽象的装饰器
type Decorator interface {
	Component
	Do() //装饰行为，可以多个
}

//定义一个具体的装饰器

type Decorator1 struct {
	c Component
}

func (d Decorator1) Do() {
	fmt.Println("c1 发生了装饰行为")
}

//重新实现Component的方法
func (d Decorator1) Operate() {
	d.Do()
	d.c.Operate()
	d.Do()
}

func TestDecorator(t *testing.T) {
	//无装饰模式
	c1 := &Component1{}
	c1.Operate()
	fmt.Println("------")
	//装饰模式
	c2 := &Decorator1{
		c: c1,
	}
	c2.Operate()

}
