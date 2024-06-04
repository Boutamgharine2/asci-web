package ascii




func Checkinput(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] < 32 || r[i] > 127 { 
			
			if r[i] == 10 {
				continue
			} else {
				return false
			}
		}
	}
	return true
}
