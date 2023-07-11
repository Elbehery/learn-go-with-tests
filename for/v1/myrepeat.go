package iteration

const repeatCount = 5

func Repeat2(str string, n int) string {
	var res string
	for i := 0; i < n; i++ {
		res += str
	}
	return res
}
