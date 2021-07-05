package testing

import (
	"errors"
	"fmt"
	"testing"
)

//优点:实现了解耦
//缺点:违背 "开闭原则"
//适合:创建的对象比较少

type Products interface {
	SetName(name string)
	GetName() string
}

//产品
type Drink struct {
	name string
	code int
}

func (p *Drink) SetName(name string) {
	p.name = name
}

func (p *Drink) GetName() string {
	return "Product's name is " + p.name
}

//具体产品 牛奶
type milk struct {
	Drink
}

func newMilk() *milk {
	return &milk{Drink{
		name: "milk",
		code: 1,
	}}
}

//具体产品 果汁
type orange struct {
	Drink
}

func newOrange() *orange {
	return &orange{Drink{
		name: "orange",
		code: 1,
	}}
}

//工厂
type ProductFactory_ struct {
}

const (
	product1 = "milk"
	product2 = "orange"
)

func (pf ProductFactory_) Create(productType string) (Products, error) {
	switch productType {
	case product1:
		return newMilk(), nil
	case product2:
		return newOrange(), nil
	default:
		return nil, errors.New("not found")
	}

}

func TestSimpleFactoryModel2(t *testing.T) {
	var pf ProductFactory_
	p, err := pf.Create("milk")
	if err == nil {
		fmt.Println(p.GetName())
	}

	p, err = pf.Create("orange")
	if err == nil {
		fmt.Println(p.GetName())
	}

}
