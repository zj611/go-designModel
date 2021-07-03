package testing

import (
	"fmt"
	"testing"
)
//优点:实现了解耦
//缺点:违背 "开闭原则"
//适合:创建的对象比较少

type Product interface {
	SetName(name string)
	GetName() string
}

type Product1 struct {
	name string
}

func (p *Product1)SetName(name string)  {
	p.name = name
}

func (p *Product1)GetName()string  {
	return "Product1's name is " + p.name
}

type Product2 struct {
	name string
}

func (p *Product2)SetName(name string)  {
	p.name = name
}

func (p *Product2)GetName()string  {
	return "Product2's name is " + p.name
}

type ProductType int

//工厂
type ProductFactory struct {

}

const  (
	p1 ProductType = iota
	p2
)

func (pf ProductFactory)Create(productType ProductType)Product  {
	switch productType {
	case p1:
		return &Product1{}
	case p2:
		return &Product2{}
	default:
		return nil
	}

}

func TestSimpleFactoryModel(t *testing.T){
	var pf ProductFactory
	p := pf.Create(p1)
	p.SetName("i am p1")
	fmt.Println(p.GetName())

	p = pf.Create(p2)
	p.SetName("i am p2")
	fmt.Println(p.GetName())
}

