package golist

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math/rand"
	"strings"
	"time"
)

// List represents go slice of generic values.
type List[V comparable] struct {
	values []V
}

func newList[V comparable](values []V) List[V] {
	return List[V]{values: values}
}

// From creates the List[V] from the given slice.
func From[V comparable](values []V) List[V] {
	return newList(values)
}

// FromString split the string by separator, apply the mapper for each value and output the List[V].
func FromString[V comparable](s string, separator string, mapper func(string) (V, bool)) List[V] {
	chunks := strings.Split(s, separator)

	newlist := make([]V, 0, len(chunks))

	for _, chunk := range chunks {
		if newv, mapped := mapper(strings.TrimSpace(chunk)); mapped {
			newlist = append(newlist, newv)
		}
	}

	return newList(newlist)
}

// Range creates the list of given range constrained by Integer.
func Range[V constraints.Integer](min, max V) List[V] {
	newlist := make([]V, max-min+1)
	for i := range newlist {
		newlist[i] = min + V(i)
	}

	return newList(newlist)
}

// Each apply the given func to each element of List[V] and return the new List[E].
func Each[V, E comparable](l List[V], fn func(V) E) List[E] {
	newlist := make([]E, 0, l.Len())

	for _, v := range l.Values() {
		newlist = append(newlist, fn(v))
	}

	return newList(newlist)
}

// L alias to From.
func L[V comparable](values []V) List[V] {
	return From(values)
}

// Var creates the List[V] from variadic.
func Var[V comparable](values ...V) List[V] {
	return L(values)
}

// Values return the builtin slice of V.
func (l List[V]) Values() []V {
	return l.values
}

// First return the first element of List[V]
func (l List[V]) First() V {
	var v V

	if l.Len() > 0 {
		v = l.values[0]
	}

	return v
}

// Last return the last element of List[V].
func (l List[V]) Last() V {
	var v V

	if l.Len() > 0 {
		v = l.values[l.Len()-1]
	}

	return v
}

// Len return actual slice len.
func (l List[V]) Len() int {
	return len(l.values)
}

// Add allow to add element to the List[V].
func (l List[V]) Add(v V) List[V] {
	l.values = append(l.values, v)
	return l
}

// Delete deletes the element from slice by index.
func (l List[V]) Delete(index uint) List[V] {
	if l.Len() <= int(index) {
		return l
	}

	l.values = append(l.values[:index], l.values[index+1:]...)
	return l
}

// Filter filters element of the List[V].
func (l List[V]) Filter(filter func(V) bool) List[V] {
	newlist := make([]V, 0, l.Len())

	for _, v := range l.values {
		if filter(v) {
			newlist = append(newlist, v)
		}
	}

	return newList(newlist)
}

// Each apply the mapper function for each value and return the List[V].
func (l List[V]) Each(mapper func(V) V) List[V] {
	newlist := make([]V, 0, l.Len())

	for _, v := range l.values {
		newlist = append(newlist, mapper(v))
	}

	return newList(newlist)
}

// Chunk creates the slice of List[V] for the given slice size.
func (l List[V]) Chunk(size int) []List[V] {
	var chunks []List[V]

	slice := l.values
	sliceLen := len(slice)
	for i := 0; i < sliceLen; i += size {
		end := i + size

		if end > sliceLen {
			end = sliceLen
		}

		chunks = append(chunks, newList(slice[i:end]))
	}

	return chunks
}

// Join other lists with the target list.
func (l List[V]) Join(lists ...List[V]) List[V] {
	outputLen := l.Len()

	for _, list := range lists {
		outputLen += list.Len()
	}

	newlist := make([]V, 0, outputLen)
	newlist = append(newlist, l.values...)

	for _, list := range lists {
		newlist = append(newlist, list.values...)
	}

	return newList(newlist)
}

// Nth return each nth element of the List[V].
func (l List[V]) Nth(nth int) List[V] {
	newlist := make([]V, 0, l.Len())

	for i, v := range l.values {
		if i%nth == 0 {
			newlist = append(newlist, v)
		}
	}

	return newList(newlist)
}

// Random return random element from List[V].
func (l List[V]) Random() V {
	rand.Seed(time.Now().Unix())
	return l.values[rand.Intn(l.Len())]
}

// Contains check that V exists in List[V].
func (l List[V]) Contains(v V) bool {
	for _, lv := range l.values {
		if lv == v {
			return true
		}
	}

	return false
}

// Reverse return reversed List[V].
func (l List[V]) Reverse() List[V] {
	newlist := make([]V, l.Len())
	copy(newlist, l.values)

	for i, j := 0, len(newlist)-1; i < j; i, j = i+1, j-1 {
		newlist[i], newlist[j] = newlist[j], newlist[i]
	}

	return newList(newlist)
}

// Shuffle the target List[V].
func (l List[V]) Shuffle() List[V] {
	newlist := make([]V, l.Len())
	copy(newlist, l.values)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(newlist), func(i, j int) { newlist[i], newlist[j] = newlist[j], newlist[i] })

	return newList(newlist)
}

// Unique return only unique values of the List[V]
func (l List[V]) Unique() List[V] {
	unique, visited := make([]V, 0, l.Len()), make(map[V]bool, l.Len())

	for _, v := range l.values {
		if _, exists := visited[v]; !exists {
			unique = append(unique, v)
			visited[v] = true
		}
	}

	return newList(unique)
}

// Zip creates List[V] together with other List[V].
func (l List[V]) Zip(other List[V]) ([]List[V], error) {
	if l.Len() != other.Len() {
		return []List[V]{}, fmt.Errorf("list must be of the same length")
	}

	zipped := make([]List[V], l.Len())

	for i, v := range l.values {
		zipped[i] = Var(v, other.values[i])
	}

	return zipped, nil
}
