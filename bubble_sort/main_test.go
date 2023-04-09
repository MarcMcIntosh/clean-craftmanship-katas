package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	want := []int{}
	got := sort([]int{})
	assert.Equal(t, want, got)

	want = []int{1}
	got = sort([]int{1})
	assert.Equal(t, want, got)

	want = []int{1, 2}
	got = sort([]int{1, 2})
	assert.Equal(t, want, got)

	want = []int{1, 2}
	got = sort([]int{2, 1})
	assert.Equal(t, want, got)

	want = []int{1, 2, 3}
	got = sort([]int{1, 2, 3})
	assert.Equal(t, want, got)

	want = []int{1, 2, 3}
	got = sort([]int{2, 1, 3})
	assert.Equal(t, want, got)

	want = []int{1, 2, 3}
	got = sort([]int{2, 3, 1})
	assert.Equal(t, want, got)

	want = []int{1, 2, 3}
	got = sort([]int{2, 3, 1})
	assert.Equal(t, want, got)

	want = []int{1, 2, 3}
	got = sort([]int{3, 2, 1})
	assert.Equal(t, want, got)

	want = []int{1, 1, 2, 3, 3, 3, 4, 5, 5, 5, 5, 6, 7, 8, 9, 9, 9}
	got = sort([]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 3, 5, 8, 9, 7, 9})
	assert.Equal(t, want, got)
}
