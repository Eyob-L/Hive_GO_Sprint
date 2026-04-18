package sprint

func IntVsFloat(i int, f float32) string {
	fi := float32(i)
	if fi < f {
		return "Float"
	}else if fi > f {
		return "Integer"
	}else {
		return "Same"
	}
}
