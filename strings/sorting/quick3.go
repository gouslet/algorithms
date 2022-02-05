package strings

type quick3String struct {
	a []string
}

func NewQuick3String(a []string) *quick3String {
	return &quick3String{a}
}

func (q *quick3String) Sort() {
	q.sort(0, len(q.a)-1, 0)
}

func (q *quick3String) sort(lo, hi, d int) {
	if hi <= lo {
		return
	}

	lt, gt := lo, hi
	v := charAt(q.a[lo], d)
	i := lo + 1
	for i <= gt {
		t := charAt(q.a[i], d)
		if t < v {
			swap(q.a, lt, i)
			lt++
			i++
		} else if t > v {
			swap(q.a, i, gt)
			gt--
		} else {
			i++
		}
	}

	q.sort(lo, lt-1, d)
	if v >= 0 {
		q.sort(lt, gt, d+1)
	}
	q.sort(gt+1, hi, d)
}
