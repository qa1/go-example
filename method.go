package main

import "fmt"

type rect struct {
    width, height int
}
func (r *rect) area() int {
    return r.width * r.height
}
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}
    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim()) 
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
    fmt.Println("گو به صورت خودکار تبدیل بین مقادیر و اشاره گرها را برای فراخوانی متدها انجام می دهد. شما ممکنه از یک نوع اشاره گر گیرنده برای جلوگیری از کپی کردن روی فراخوانی متدها استفاده کنید یا به متدها برای تغییر دادن سازه های  دریافتی اجازه دهید.")
}
