package behavior

import (
	"container/list"
	"fmt"
	"testing"
)

//观察者模式

//抽象主题
type Subject interface {
	Add(o Observer)
	Send(str string)
}

//-------------天气主题------------------
type WeatherSubject struct {
	title string
	l     *list.List
}

//天气主题非侵入式实现抽象主题
func (sub *WeatherSubject) Add(o Observer) {
	sub.l.PushBack(o)
}

//天气主题非侵入式实现抽象主题
func (sub *WeatherSubject) Send(str string) {
	for i := sub.l.Front(); i != nil; i = i.Next() {
		(i.Value).(Observer).Receive(sub.title + "发送的：" + str)
	}
}

//-------------新闻主题------------------
type NewsSubject struct {
	title string
	l     *list.List
}

//新闻主题非侵入式实现抽象主题
func (sub *NewsSubject) Add(o Observer) {
	sub.l.PushBack(o)
}

//新闻主题非侵入式实现抽象主题
func (sub *NewsSubject) Send(str string) {
	for i := sub.l.Front(); i != nil; i = i.Next() {
		(i.Value).(Observer).Receive(sub.title + "发送的：" + str)
	}
}

//抽象观察者
type Observer interface {
	Receive(str string)
}

//-------------a观察者------------------
type aObserver struct {
	name string
}

//a观察者非侵入式实现抽象观察者
func (o *aObserver) Receive(str string) {
	fmt.Println("A类观察者【" + o.name + "】接收" + str)
}

//-------------b观察者------------------
type bObserver struct {
	name string
}

//b观察者非侵入式实现抽象观察者
func (o *bObserver) Receive(str string) {
	fmt.Println("B类观察者【" + o.name + "】接收" + str)
}

func TestObserverModel(t *testing.T) {
	a := &aObserver{
		name: "张三",
	}

	b := &bObserver{
		name: "李四",
	}

	//天气类允许a和b类型观察者订阅
	weather := WeatherSubject{
		l:     list.New(),
		title: "武汉新闻",
	}
	weather.Add(a)
	weather.Add(b)

	//新闻类只允许b类型观察者订阅
	news := NewsSubject{
		l:     list.New(),
		title: "孝感热点",
	}
	news.Add(b)

	weather.Send("全省暴雨红色警报")
	news.Send("全市停工停课")
}
