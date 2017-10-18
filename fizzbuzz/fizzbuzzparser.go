package fizzbuzz

const BUZZ string = "Buzz"
const FIZZ string = "Fizz"
const FIZZ_BUZZ string = "FizzBuzz"
func ParseFizzBuzz(value int) string {
	if isFizz(value) && isBuzz(value) {
		return FIZZ_BUZZ
	}
	if isBuzz(value) {
		return BUZZ
	}
	if isFizz(value){
		return FIZZ
	}
	return ""
}

func isBuzz(value int) bool {
	return value % 5 == 0
}

func isFizz(value int) bool {
	return value % 3 == 0
}
