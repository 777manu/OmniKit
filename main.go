package main
import ("fmt";"runtime")
// test to check the operating system
func main() {
os := runtime.GOOS
fmt.Println("Operating system:", os)
}