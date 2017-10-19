package ll

import "testing"
import "fmt"

func TestAppending(t *testing.T) {
	var list Sll
	list.Append(&Node{Data: 88})
	list.Append(&Node{Data: 89})
	list.Append(&Node{Data: 11})
	for i := 0; i <= 22; i++ {
		list.Append(&Node{Data: i})
	}

	list.Printll()
}

func TestFind(t *testing.T) {
	var list Sll
	list.Append(&Node{Data: 88})
	list.Append(&Node{Data: 89})
	list.Append(&Node{Data: 11})
	for i := 0; i <= 22; i++ {
		list.Append(&Node{Data: i})
	}

	prev, d := list.Find(4)
	fmt.Println(prev, d)

	prev, d = list.Find(88)
	fmt.Println(prev, d)

	prev, d = list.Find(22)
	fmt.Println(prev, d)

	prev, d = list.Find(188)
	fmt.Println(prev, d)
}

func TestToSlice(t *testing.T) {
	var list Sll
	list.Append(&Node{Data: 88})
	list.Append(&Node{Data: 89})
	list.Append(&Node{Data: 11})
	for i := 0; i <= 22; i++ {
		list.Append(&Node{Data: i})
	}

	s := list.Lltoas()
	fmt.Println(s)
}
