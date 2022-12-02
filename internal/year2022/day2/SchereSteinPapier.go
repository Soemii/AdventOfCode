package day2

type SchereSteinPapier struct {
	schere bool
	stein  bool
	papier bool
}

func (s SchereSteinPapier) GetPoints() int {
	if s.stein {
		return 1
	} else if s.papier {
		return 2
	} else {
		return 3
	}
}

func (s SchereSteinPapier) win(p SchereSteinPapier) bool {
	if s.schere && p.papier {
		return true
	}
	if s.stein && p.schere {
		return true
	}
	if s.papier && p.stein {
		return true
	}
	return false
}

func (s SchereSteinPapier) draw(p SchereSteinPapier) bool {
	return s.stein && p.stein || s.papier && p.papier || s.schere && p.schere
}

func (s SchereSteinPapier) GetWin() SchereSteinPapier {
	if s.schere {
		return SchereSteinPapier{stein: true}
	}
	if s.stein {
		return SchereSteinPapier{papier: true}
	}
	if s.papier {
		return SchereSteinPapier{schere: true}
	}
	return SchereSteinPapier{stein: true}
}

func (s SchereSteinPapier) GetLose() SchereSteinPapier {
	if s.schere {
		return SchereSteinPapier{papier: true}
	}
	if s.stein {
		return SchereSteinPapier{schere: true}
	}
	if s.papier {
		return SchereSteinPapier{stein: true}
	}
	return SchereSteinPapier{stein: true}
}

func (s SchereSteinPapier) GetDraw() SchereSteinPapier {
	if s.schere {
		return SchereSteinPapier{schere: true}
	}
	if s.stein {
		return SchereSteinPapier{stein: true}
	}
	if s.papier {
		return SchereSteinPapier{papier: true}
	}
	return SchereSteinPapier{stein: true}
}
