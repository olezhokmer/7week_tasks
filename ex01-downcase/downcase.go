package downcase

func Downcase(s string) (string, error) {
	u := ""
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			c += 32
		}
		u += string(c)
	}
	return u, nil
}
