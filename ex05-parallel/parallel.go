package letter

type Letters map[rune]int

func New() *Letters {
	return &Letters{}
}

func Frequency(s string) Letters {
	l := *New()
	for _, ch := range s {
		l[ch]++
	}
	return l
}

func ConcurrentFrequency(s []string) Letters {
	sm := Letters{}
	count := len(s)
	r := make(chan Letters, count)
	for _, s := range s {
		go func(s string) {
			r <- Frequency(s)
		}(s)
	}

	for i := 0; i < count; i++ {
		for r, f := range <-r {
			sm[r] += f
		}
	}
	return sm
}
