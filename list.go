package golist

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

// List represents go slice of generic values as concurrency safe List[V].
type List[V comparable] struct {
	mutex  *sync.RWMutex
	values []V
}

func newList[V comparable](values []V) List[V] {
	return List[V]{
		mutex:  &sync.RWMutex{},
		values: values,
	}
}

// From creates the List[V] from the given slice.
func From[V comparable](values []V) List[V] { return newList(values) }

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
func L[V comparable](values []V) List[V] { return From(values) }

// Var creates the List[V] from variadic.
func Var[V comparable](values ...V) List[V] { return L(values) }

// Values return the builtin slice of V.
func (l List[V]) Values() []V { return l.values }

// First return the first element of List[V]
func (l List[V]) First() V {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	var v V

	if l.Len() > 0 {
		v = l.values[0]
	}

	return v
}

// Last return the last element of List[V].
func (l List[V]) Last() V {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	var v V

	if l.Len() > 0 {
		v = l.values[l.Len()-1]
	}

	return v
}

// Len return actual slice len.
func (l List[V]) Len() int { return len(l.values) }

// Empty return true if len is zero.
func (l List[V]) Empty() bool { return l.Len() == 0 }

// Add allow to add element to the List[V].
func (l List[V]) Add(v V) List[V] {
	l.mutex.Lock()
	l.values = append(l.values, v)
	l.mutex.Unlock()
	return l
}

// Delete deletes the element from slice by index.
func (l List[V]) Delete(index uint) List[V] {
	l.mutex.Lock()
	defer l.mutex.Unlock()

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

// JoinToString joins the list elements if V is type of string. Otherwise, skip the value.
func (l List[V]) JoinToString(separator string) string {
	sparts := make([]string, 0, l.Len())

	for _, v := range l.values {
		switch vv := any(v).(type) {
		case string:
			sparts = append(sparts, vv)
		default:
			return ""
		}
	}

	return strings.Join(sparts, separator)
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
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	rand.Seed(time.Now().Unix())
	return l.values[rand.Intn(l.Len())]
}

// Contains check that V exists in List[V].
func (l List[V]) Contains(v V) bool {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

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

// Interface return List[V] as slice of interfaces.
func (l List[V]) Interface() []interface{} {
	newlist := make([]interface{}, 0, l.Len())

	for _, v := range l.values {
		newlist = append(newlist, v)
	}

	return newlist
}

// Fill create the List[V] from given V by the specified count.
func Fill[V comparable](symbol V, count int) List[V] {
	newlist := make([]V, 0, count)

	for i := 0; i < count; i++ {
		newlist = append(newlist, symbol)
	}

	return newList(newlist)
}

// Sequence generate the List[string] of the given symbol concatenated with index. Useful for SQL query generation.
// In example: golist.Sequence("$", 3, 1).Values() [$1, $2, $3]
func Sequence(symbol string, count int, start int) List[string] {
	newlist := make([]string, 0, count)
	end := count + start - 1

	for i := start; i <= end; i++ {
		newlist = append(newlist, symbol+strconv.Itoa(i))
	}

	return newList(newlist)
}

// Every apply the given func to each element of List[V] and return the new slice of E.
func Every[V comparable, E any](l List[V], fn func(V) E) []E {
	newlist := make([]E, 0, l.Len())

	for _, v := range l.Values() {
		newlist = append(newlist, fn(v))
	}

	return newlist
}

func Between[V comparable](min, max, chunk int, f func(start, count int) V) List[V] {
	newlist := make([]V, 0, max)

	for i := min; i <= max; i++ {
		if i%chunk == 0 {
			newlist = append(newlist, f(i-chunk+1, chunk))
		}
	}

	return newList(newlist)
}
