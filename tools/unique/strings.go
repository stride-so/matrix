package unique

type StringSlice struct {
	set map[string]struct{}
}

func (u *StringSlice) Add(val string) {
	if u.set == nil {
		u.set = make(map[string]struct{})
	}
	if !u.Has(val) {
		u.set[val] = struct{}{}
	}
}

func (u StringSlice) Has(val string) bool {
	_, ok := u.set[val]
	return ok
}

func (u *StringSlice) Values() []string {
	unique := []string{}
	for k := range u.set {
		unique = append(unique, k)
	}
	return unique
}

func NewStringSlice(vals []string) *StringSlice {
	u := &StringSlice{
		set: make(map[string]struct{}),
	}
	for _, v := range vals {
		u.Add(v)
	}
	return u
}
