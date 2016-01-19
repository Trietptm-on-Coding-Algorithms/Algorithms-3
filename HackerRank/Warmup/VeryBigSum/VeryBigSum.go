Vpackage main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N string
	fmt.Scanln(&N)
	bytes, _ := ioutil.ReadAll(os.Stdin)

	var sum int64 = 0
	for _, value := range strings.Split(string(bytes), " ") {
		parse, _ := strconv.ParseInt(value, 10, 64)
		sum += parse
	}

	fmt.Println(sum)
}
