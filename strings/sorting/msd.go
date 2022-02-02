package strings

type msd struct {
	w int
	a []string
}

func NewMSD(w int, a []string) *lsd {
	return &lsd{w, a}
}

func (l *lsd) MSDSort() []string {

	return l.a
}
