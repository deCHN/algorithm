package algorithm

import "testing"

func TestAdd(t *testing.T) {
	want := 2
	if get := Add(1, 1); get != want {
		t.Fatalf("Want %q, but got %q.", want, get)
	}
}

func TestAny(t *testing.T) {
	id := IDImpl{
		ID: "adj",
	}

	t.Log(id, Add(3, 5))

}
