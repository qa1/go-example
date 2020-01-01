package main
import ( 
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("ship", c)	
	//for {
	//	msg,open := <- c
	//	if !open {
	//		break		
	//	}
	//	fmt.Println(msg)
	//}
	// buffer channel: send channel block until some one ready to recive
	// because of that we have to get and recive in 2 diffrent goroutin
	// but with buffer we can do that
	// c := make(chan string, 2)
	// that means we can have 2 channel read and not lock until some pick it up in same block
	for msg := range c {
		fmt.Println(msg)	
	}
}

func count(thing string, c chan string) {
	for i:= 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
