package day6

type Visit struct {
	r, c int
	dir  int32
}

type VisitLog map[Visit]struct{}

func (vl VisitLog) DejaVu(r int, c int, dir int32) bool {
	_, ok := vl[Visit{r, c, dir}]
	return ok
}

func (vl VisitLog) Log(r int, c int, dir int32) {
	visit := Visit{r, c, dir}
	vl[visit] = struct{}{}
}

func NewVisitLog() VisitLog {
	return make(VisitLog)
}
