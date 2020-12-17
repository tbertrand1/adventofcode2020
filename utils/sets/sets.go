package sets

type set struct {
	m map[any]void
}

type any interface{}

type void struct{}

func NewSet() *set {
	set := &set{}
	set.m = make(map[any]void)
	return set
}

func (set *set) Add(value any) {
	set.m[value] = void{}
}

func (set *set) Remove(value any) {
	delete(set.m, value)
}

func (set *set) Contains(value any) bool {
	_, found := set.m[value]
	return found
}

func (set *set) Length() int {
	return len(set.m)
}

func (set *set) Keys() []any {
	var keys []any
	for k := range set.m {
		keys = append(keys, k)
	}
	return keys
}
