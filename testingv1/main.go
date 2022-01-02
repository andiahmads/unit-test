package testingv1

func Absolute(number int) int {
	if number < 0 {
		return -1 * number
	}
	return number
}

func LuasSegitiga(p int, l int) int {
	var panjang = p
	var luas = l
	return panjang * luas
}

func add(a, b int) int {
	return a + b
}
