package split

import (
	"reflect"
	"testing"
)

// TestSplit: test
// func TestSplit(t *testing.T) {
// 	got := Split("a:b:c", ":")
// 	want := []string{"a", "b", "c"}
// 	if ok := reflect.DeepEqual(got, want); !ok {
// 		t.Fatalf("希望得到%v,实际得到%v", want, got)
// 	}
// }

// func TestNoneSplit(t *testing.T) {
// 	t.Log("test not in s")
// 	got := Split("a:b:c", "*")
// 	want := []string{"a:b:c"}
// 	if ok := reflect.DeepEqual(got, want); !ok {
// 		t.Fatalf("希望得到%v,实际得到%v", want, got)
// 	}
// }

// func TestAdd(t *testing.T) {
// 	got := Add(10, 20)
// 	want := 30
// 	if ok := reflect.DeepEqual(got, want); !ok {
// 		t.Fatalf("want:%v,got:%v", want, got)
// 	}
// }

// func TestMutiSepSplit(t *testing.T) {
// 	got := Split("abcfabcab", "bc")
// 	want := []string{"a", "fa", "ab"}
// 	if ok := reflect.DeepEqual(got, want); !ok {
// 		t.Fatalf("want:%v,got:%v", want, got)
// 	}
// }
func TestMutiSepSplit(t *testing.T) {
	type testItem struct {
		str  string
		sep  string
		want []string
	}
	var tests = map[string]testItem{
		"normal": {"a:b:c", ":", []string{"a", "b", "c"}},
		"none":   {"a:b:c", "*", []string{"a:b:c"}},
		"multi":  {"abcfabcab", "bc", []string{"a", "fa", "ab"}},
	}
	for name, item := range tests {
		t.Log(name)
		// method 1: old
		// ret := Split(item.str, item.sep)
		// if !reflect.DeepEqual(ret, item.want) {
		// 	t.Fatalf("want:%v,got:%v", ret, item.want)
		// }
		// method 2: recommand
		t.Run(name, func(t *testing.T) {
			ret := Split(item.str, item.sep)
			if !reflect.DeepEqual(ret, item.want) {
				t.Fatalf("want:%v,got:%v", ret, item.want)
			}
		})
	}
}
