package main
import ( 
	"fmt"
	"time"
	"sync"
)

func main() {
	go count("ship")
	go count("cat")
	fmt.Scanln()
}

func count(thing string) {
	for i:= 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
