package iteration

const repeatCount = 5

func Repeat2(str string) string {
	var res string
	for i := 0; i < repeatCount; i++ {
		res += str
	}
	return res
}
