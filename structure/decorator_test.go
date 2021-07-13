package structure

import (
	"fmt"
	"testing"
)

// 装饰器模式

//定义抽象的组建
type Food interface {
	getPrice() int
}

type milk struct {
}

func (c milk) getPrice() int {
	return 15
}

//定义一个抽象的装饰器(增值服务)
type DecoratorService interface {
	Food
	Do() //装饰行为，可以多个
}

//定义一个具体的装饰器(食物打包增值服务)
type DecoratorServicePackage struct {
	c Food
}

func (d DecoratorServicePackage) Do() {
	fmt.Println("c1 发生了装饰行为")
}

//重新实现Food的getPrice方法
func (d DecoratorServicePackage) getPrice() int {
	price := d.c.getPrice()
	return price + 10
}

func TestDecorator(t *testing.T) {
	//无装饰模式
	c1 := &milk{}
	fmt.Println(c1.getPrice())

	fmt.Println("------")

	//装饰模式
	c2 := &DecoratorServicePackage{
		c: c1,
	}
	c2.Do()
	fmt.Println(c2.getPrice())
}
