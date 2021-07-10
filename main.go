package main

import "fmt"

type product11 interface {
	setName(name string)
	getName() string
}

type shoe struct {
	name string
}

func (p *shoe) setName(name string) {
	p.name = name
}
func (p *shoe) getName() string {
	return p.name
}

type shirt struct {
	name string
}

func (p *shirt) setName(name string) {
	p.name = name
}
func (p *shirt) getName() string {
	return p.name
}

type factory struct {
}

func (f *factory) getProduct(pName string) product11 {
	switch pName {
	case "shoe":
		return &shoe{name: "shoe"}
	case "shirt":
		return &shirt{name: "shirt"}
	default:
		return nil
	}
}

func main() {
	var f factory
	p := f.getProduct("shoe")
	fmt.Println(p.getName())
}
