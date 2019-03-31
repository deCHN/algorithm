package algorithm

import "testing"

func TestAdd(t *testing.T) {
   want := 7
   if get := Add(3,4); get != want {
      t.Fatalf("Want %d, but got %d.", want, get)
   }
}
