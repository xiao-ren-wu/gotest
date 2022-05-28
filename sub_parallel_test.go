package gotest

import (
	"testing"
	"time"
)

func parallelTest1(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second)
	t.Log("parallel test1 finish")
}

func parallelTest2(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
	t.Log("parallel test2 finish")
}

func parallelTest3(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
	t.Log("parallel test1 finish")
}

func TestSubParallel(t *testing.T) {
	t.Run("sub parallel", func(t *testing.T) {
		t.Run("pt1", parallelTest1)
		t.Run("pt2", parallelTest2)
		t.Run("pt3", parallelTest3)
	})
	t.Log("finish sub parallel test")
}
