package ll

import "testing"

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
