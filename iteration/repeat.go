package iteration

//const repeatCount = 5

func Repeat(letter string, repeatCount int) (word string) {
	for i := 0; i < repeatCount; i++ {
		word += letter
	}
	return
}
