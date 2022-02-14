package tries

const R = 256

type tries_arr struct {
	children [R]*tries_arr
	val      any
	size     int
}

func NewTriesArr() *tries_arr {
	return &tries_arr{}
}

func (t *tries_arr) Put(key string, val any) {
	if key == "" {
		t.val = val
		t.size++
	}
	cur := t
	for i, c := range key {
		if cur.children[c] == nil {
			cur.children[c] = new(tries_arr)
		}

		cur = cur.children[c]

		if i == len(key)-1 {
			cur.val = val
			t.size++
		}
	}
}

func (t *tries_arr) Get(key string) any {
	cur := t
	for _, c := range key {
		cur = cur.children[c]

		if cur.val != nil && byte(c) == key[len(key)-1] {
			return cur.val
		}
	}
	return t.val
}

func (t tries_arr) Size() int {
	return t.size
}

func (t *tries_arr) Contains(key string) bool {
	if key == "" {
		return t.val != nil
	}
	cur := t
	for _, c := range key {
		if cur.children[c] == nil {
			return false
		}
		cur = cur.children[c]
		if cur.val != nil {
			return true
		}
	}
	return key == t.val
}

func (t *tries_arr) Keys() (res []string) {
	res = t.KeysWithPrefix("")

	return
}

func (t *tries_arr) KeysWithPrefix(pre string) (res []string) {
	if t == nil {
		return
	}

	for i, c := range pre {
		if b := t.children[c]; b != nil {
			t = b
		}
		if t.val != nil && i == len(pre)-1 {
			res = append(res, pre)
		}
	}
	res = append(res, t.collect(pre)...)

	return
}

func (t *tries_arr) collect(key string) []string {
	res := []string{}

	if t == nil {
		return res
	}

	if t.val != nil {
		res = append(res, key)
	}

	for i, c := range t.children {
		res = append(res, c.collect(key+string(rune(i)))...)
	}

	return res
}
