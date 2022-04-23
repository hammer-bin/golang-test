package interation

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func Repeat2(character string, cnt int) string {
	var repeated string
	for i := 0; i < cnt; i++ {
		repeated += character
	}

	return repeated
	
}
