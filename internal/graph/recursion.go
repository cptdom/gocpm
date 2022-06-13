package graph

func DoEarly(starts []*Node) {
	for _, n := range starts {
		var max uint64
		for _, d := range n.Prev {
			if d.ef > max {
				max = d.ef
			}
		}
		n.es = max
		n.ef = max + n.Duration

		DoEarly(n.Next)
	}
}

func DoLatest(ends []*Node) {
	for _, n := range ends {
		for _, d := range n.Prev {
			if d.lf > n.ls || d.lf == 0 {
				d.lf = n.ls
				d.ls = d.lf - d.Duration
				d.Float = d.lf - d.ef
			}
		}
		DoLatest(n.Prev)
	}
}
