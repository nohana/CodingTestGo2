package stack

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	s := Stack{}
	s.Push("foo")
	s.Push("bar")

	want := Stack([]string{"foo", "bar"})
	if !reflect.DeepEqual(s, want) {
		// TODO: more detail message
		t.Error("error")
		return
	}
}

func TestPop(t *testing.T) {
	s := Stack([]string{"foo", "bar"})
	t.Run("pop 1", func(t *testing.T) {
		got := s.Pop()
		wantVal := "bar"
		if got != wantVal {
			// TODO: more detail message
			t.Error("error on 1 val")
			return
		}
		want := Stack([]string{"foo"})
		if !reflect.DeepEqual(s, want) {
			// TODO: more detail message
			t.Error("error on 1")
			return
		}
	})

	t.Run("pop 2", func(t *testing.T) {
		got := s.Pop()
		wantVal := "foo"
		if got != wantVal {
			// TODO: more detail message
			t.Error("error on 2 val")
			return
		}
		want := Stack{}
		if !reflect.DeepEqual(s, want) {
			// TODO: more detail message
			t.Error("error on 2")
			return
		}
	})

	t.Run("pop 3", func(t *testing.T) {
		got := s.Pop()
		wantVal := ""
		if got != wantVal {
			// TODO: more detail message
			t.Error("error on 3 val")
			return
		}
		want := Stack{}
		if !reflect.DeepEqual(s, want) {
			// TODO: more detail message
			t.Error("error on 3")
			return
		}
	})
}
