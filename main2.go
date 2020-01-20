package main
import "fmt"
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
type person struct {
    name string
    age  int
}
func main() {
    fmt.Println(fact(7))
    
    intseq := intSeq()
    fmt.Println(intseq())

    fmt.Println(&person{name: "Ann", age: 40})
    p := person{name: "Sean", age: 50}
    sp := &p
    fmt.Println(sp.age)
    
    //slices
    s := make([]string, 3)
    fmt.Println("emp:", s)
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])
    fmt.Println("len:", len(s))
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)
    l := s[2:5]
    fmt.Println("sl1:", l)
    l = s[:5]
    fmt.Println("sl2:", l)
    l = s[2:]
    fmt.Println("sl3:", l)
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
    
    //map
    m := make(map[string]int)
    m["k1"] = 7
    m["k2"] = 13
    fmt.Println("map:", m)
    v1 := m["k1"]
    fmt.Println("v1: ", v1)
    fmt.Println("len:", len(m))
    delete(m, "k2")
    fmt.Println("map:", m)
    _, prs := m["k2"]
    fmt.Println("prs:", prs)
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

}
