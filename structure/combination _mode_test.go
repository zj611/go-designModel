package structure

// 组合模式
import (
	"fmt"
	"testing"
)

type component interface {
	search(string)
}

type file struct {
	name string
}

func (f *file) search(keyword string) {
	fmt.Println(keyword+" : ", f.name)
}

func (f *file) getName() string {
	return f.name
}

//组合
type folder struct {
	components []component
	name       string
}

func (f *folder) search(keyword string) {
	for _, v := range f.components {
		v.search(keyword)
	}
}

func (f *folder) add(c component) {
	f.components = append(f.components, c)
}

func TestCombinationModel(t *testing.T) {
	f1 := &folder{
		name: "f1",
	}
	file1 := &file{
		name: "file1",
	}
	f1.add(file1)
	f1.search("f100")

	f2 := &folder{
		name: "f2",
	}
	file2 := &file{name: "file2"}
	file3 := &file{name: "file3"}
	f2.add(file2)
	f2.add(file3)
	f2.add(f1)

	f2.search("rose")

}
