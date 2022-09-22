package golist

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func BenchmarkListJoin(b *testing.B) {
	b.ReportAllocs()

	list := From([]int{1, 2, 3})
	for i := 0; i < b.N; i++ {
		list.Join(
			From([]int{i}),
		)
	}
}

func TestListFrom(t *testing.T) {
	l := From([]int{1, 2, 3})
	assert.Equal(t, []int{1, 2, 3}, l.Values())
}

func TestListFromString(t *testing.T) {
	l := FromString("1, 2, 3", ",", func(v string) (int, bool) {
		if n, err := strconv.Atoi(v); err == nil {
			return n, true
		}

		return 0, false
	})

	assert.Equal(t, []int{1, 2, 3}, l.Values())
}

func TestListAdd(t *testing.T) {
	l := L([]int{1, 2})
	l = l.Add(3)

	assert.Equal(t, []int{1, 2, 3}, l.Values())
}

func TestListDelete(t *testing.T) {
	l := L([]int{1, 2})
	l = l.Delete(0)
	assert.Equal(t, []int{2}, l.Values())
}

func TestListFilter(t *testing.T) {
	l := L([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, []int{2, 4, 6}, l.Filter(func(v int) bool { return v%2 == 0 }).Values())
}

func TestListEach(t *testing.T) {
	l := L([]int{1, 2, 3})
	assert.Equal(t, []int{2, 4, 6}, l.Each(func(v int) int { return v * 2 }).Values())
}

func TestListChunk(t *testing.T) {
	l := L([]int{1, 2, 3, 4, 5})
	chunks := l.Chunk(2)
	assert.Equal(t, 3, len(chunks))
	assert.Equal(t, []int{1, 2}, chunks[0].Values())
	assert.Equal(t, []int{3, 4}, chunks[1].Values())
	assert.Equal(t, []int{5}, chunks[2].Values())
}

func TestListJoin(t *testing.T) {
	l := L([]int{1, 2, 3})
	l = l.Join(L([]int{4, 5}), L([]int{6, 7}))
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, l.Values())
}

func TestRange(t *testing.T) {
	l := Range(1, 10)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, l.Values())
}

func TestListNth(t *testing.T) {
	l := Range(1, 10)
	assert.Equal(t, []int{1, 4, 7, 10}, l.Nth(3).Values())
}

func TestListFirst(t *testing.T) {
	assert.Equal(t, 2, L([]int{2, 3, 10}).First())
	assert.Equal(t, 0, Var[int]().First())
}

func TestListLast(t *testing.T) {
	assert.Equal(t, 10, L([]int{2, 3, 10}).Last())
	assert.Equal(t, "", Var[string]().Last())
}

func TestListRandom(t *testing.T) {
	l := L([]int{1, 2, 3, 4, 5})
	assert.True(t, l.Contains(l.Random()))
}

func TestListReverse(t *testing.T) {
	assert.Equal(t, []int{5, 4, 3, 2, 1}, L([]int{1, 2, 3, 4, 5}).Reverse().Values())
}

func TestListUnique(t *testing.T) {
	assert.Equal(t, []int{1, 2}, Var(1, 1, 1, 1, 2).Unique().Values())
}

func TestListZip(t *testing.T) {
	lists, err := Var(1, 2, 3).Zip(Var(4, 5, 6))
	assert.Nil(t, err)
	assert.Equal(t, 3, len(lists))
	assert.Equal(t, []int{1, 4}, lists[0].Values())
	assert.Equal(t, []int{2, 5}, lists[1].Values())
	assert.Equal(t, []int{3, 6}, lists[2].Values())
}

func TestEach(t *testing.T) {
	assert.Equal(t, []string{"1", "2", "3"}, Each(Var(1, 2, 3), func(v int) string {
		return strconv.Itoa(v)
	}).Values())
}
