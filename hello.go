package main
import ( 
	"fmt"
	"errors"
	"math"
)

type person struct {
	name string
	age int
}

func main() {
	fmt.Println("Hello, Worlds")

	p := person{name: "Jake" , age: 23}
	fmt.Println(p.age)

	var x int
	x = 7
	var y int = 8
	t := 6
	var sum int = x+y+t
	fmt.Println(sum)
	fmt.Println(&sum)
	if x<6 {
		fmt.Println("x is lower down 6")	
	} else if x>8 { 
		fmt.Println("else if x is lower down 6")	
	} else {
		fmt.Println("else x is lower down 6")	
	}
	var a [5]int
	a[2] = 7
	fmt.Println(a)
	b := [5]int{1,5,6,7,7}
	fmt.Println(b)
	// slice arrays withot allocate size
	c := []int{5,8,7,8,9}
	c = append(c,13)
	fmt.Println(c)
	vetrics := make(map[string]int)
	vetrics["test"] = 854
	vetrics["Ali"] = 1
	vetrics["Hossein"] = 3
	delete(vetrics, "test")
	fmt.Println(vetrics)
	fmt.Println(vetrics["Ali"])
	for key, value := range vetrics {
		fmt.Println("key:", key, "value:", value)
	}
	// for loop
	for i:=0; i<5 ; i++ {
		fmt.Println(i)	
	}
	//while
	j := 0
	for j<5{
		fmt.Println(j)	
		j++
	}
	// range 
	arr := []string{"a","b","c"}
	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}
	sumresult := summ(5,9)
	fmt.Println(sumresult)
	result, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}


func summ(x int, y int) int {
	return x+y
}

func sqrt (x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for navigate")
	}
	return math.Sqrt(x),nil
}


