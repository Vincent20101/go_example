package main

import "testing"

func TestCheckDir(t *testing.T) {
	testDatas := []struct {
		dir  string
		want bool
	}{
		{"C", true},
		{"./", true},
		//{"*sdf", false},
		//{"V:/abc", false},
	}

	for _, v := range testDatas {
		res := checkDir(v.dir)
		if res != v.want {
			t.Error("error")
		}
	}
}
