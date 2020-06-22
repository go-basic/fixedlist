package fixedlist

import "testing"

func TestFixedList_Add(t *testing.T) {
	f := NewFixedList(2)
	f.Add("a")
	if f.Data()[0] != "a" {
		t.Error("add err")
	}
}

func TestFixedList_Len(t *testing.T) {
	var length int
	f := NewFixedList(length)
	f.Add("a")
	f.Add("b")
	f.Add("c")
	if f.Len() != length {
		t.Error("len err")
	}
}

func TestFixedList_Data(t *testing.T) {
	var length int
	f := NewFixedList(length)
	f.Add("a")
	f.Add("b")
	f.Add("c")
	if len(f.Data()) != length {
		t.Error("data err")
	}
}
