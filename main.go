package main
import ("fmt";"runtime")
// test to check the operating system
func main() {
os := runtime.GOOS
arch := runtime.GOARCH
fmt.Println("Architecture:", arch)
fmt.Println("Operating system:", os)
whatos(os)
}
const (
	//  defining constants for different operating systems
	Windows = "windows"
	Linux   = "linux"
	MacOS   = "darwin"
)
// whatos checks the operating system and print a message accordingly
func whatos(os string) {
if os == Windows {
	fmt.Println("This is a Windows operating system.")
} else if os == Linux {
	fmt.Println("This is a Linux operating system.")
} else if os == MacOS {
	fmt.Println("This is a macOS operating system.")
}
}