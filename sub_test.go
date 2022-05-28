package gotest

import "testing"

func sub1(t *testing.T) {
	t.Log("sub1")
	a := 1
	b := 2
	expect := 3
	actual := Add(a, b)
	if actual != expect {
		t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expect)
	}
}

func sub2(t *testing.T) {
	t.Log("sub2")
	a := 1
	b := 2
	expect := 3
	actual := Add(a, b)
	if actual != expect {
		t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expect)
	}
}

func sub3(t *testing.T) {
	t.Log("sub3")
	a := 1
	b := 2
	expect := 3
	actual := Add(a, b)
	if actual != expect {
		t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expect)
	}
}

func TestSub(t *testing.T) {
	t.Run("1", sub1)
	t.Run("2", sub2)
	t.Run("3", sub3)

}
func Add(a, b int) int {
	return a + b
}

func TestAddWithCase(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"a=3,b=1",
			args{3, 1},
			4,
		},
		{
			"a=3,b=-1",
			args{3, -1},
			4,
		},
		{
			"a=-2,b=-1",
			args{-2, -1},
			-3,
		},
		{
			"a=-2,b=0",
			args{-2, -1},
			-3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
