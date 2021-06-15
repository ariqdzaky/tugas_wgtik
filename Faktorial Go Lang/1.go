package main
import "fmt"

func main() {
var k, n, i int
	
	fmt.Scan(&n)

	i = 1
	k = 1
	for k<=n {
		i = i * k
		k = k + 1
	}
	fmt.Print(i)
}