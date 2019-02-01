package isUnique

import "testing"

func Test_EmptyArray(t *testing.T) {
	str := s{}
	_, err := str.isUnique()
	if err == nil {
		t.Fail()
	}
}

func Test_NonEmpty_NoDuplicates(t *testing.T) {
	str := s{1, 4, 7, 5, 3, 99, 11}
	unique, err := str.isUnique()
	if err != nil {
		t.Fail()
	}
	if !unique {
		t.Fail()
	}
}

func Test_NonEmpty_Duplicates(t *testing.T) {
	str := s{1, 4, 7, 5, 4, 3, 99, 11}
	unique, err := str.isUnique()
	if err != nil {
		t.Fail()
	}
	if unique {
		t.Fail()
	}
}

func Test_NonEmpty_DuplicatesAtEdges(t *testing.T) {
	str := s{1, 4, 7, 5, 4, 3, 99, 11, 1}
	unique, err := str.isUnique()
	if err != nil {
		t.Fail()
	}
	if unique {
		t.Fail()
	}
}
