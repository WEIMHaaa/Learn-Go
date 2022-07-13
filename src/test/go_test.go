package src

import "testing"

func Add(a, b int) int {
	return a + b
}

// 成功用例
func TestAdd1(t *testing.T) {
	r := Add(1, 2)
	if r != 3 {
		t.Errorf("Add(1,2) failed. Got %d, expected 3.", r)
	}
}

// 跳过用例 go test -short
func TestAdd2(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	r := Add(1, 2)
	if r != 2 {
		t.Errorf("Add(1,2) failed. Got %d, expected 3.", r)
	}
}

// 失败用例
func TestAdd3(t *testing.T) {
	r := Add(1, 2)
	if r != 1 {
		t.Errorf("Add(1,2) failed. Got %d, expected 3.", r)
	}
}
