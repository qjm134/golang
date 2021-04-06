package binHexOct

import (
	"fmt"
	"strconv"
)

func main() {
	/*num := 14
	k := 8
	*/
	num := 12
	k := 16
	fmt.Println(transfer(num, k))
}

func transfer(num int, k int) string {
	var res string
	for num != 0 {
		quotient := num / k
		remainder := num % k

		if remainder < 10 {
			cur := strconv.Itoa(remainder)
			res = cur + res
		} else {
			cur := 'A' + rune(remainder-10)
			res = string(cur) + res
		}

		num = quotient
	}

	return res
}
