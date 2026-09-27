package run

func FillSlice(slice []rune, value rune) {
	slice[0] = value

	for j := 1; j < len(slice); j *= 2 {
		copy(slice[j:], slice[:j])
	}
}
