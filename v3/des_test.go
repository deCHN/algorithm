package algorithm

import "testing"

func TestAdd(t *testing.T) {
   want := 10
   if get := Add(1,2,3,4); get != want {
      t.Fatalf("Want %d, but got %d.", want, get)
   }
}
