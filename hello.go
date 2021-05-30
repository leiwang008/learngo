package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"

	"example.com/morestrings"
	"github.com/google/go-cmp/cmp"
	"github.com/leiwang008/utils"
	"rsc.io/quote"
)

func main() {
	debugmsg := debugmsg(true)

	file, err := initLog()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	log.Println(debugmsg + "main started ... ")

	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Hello())

	//define a slice of names
	names := []string{"Tom", "Sandy", "Jack"}
	msgs, err := utils.Hellos(names)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(msgs)
		for _, name := range names {
			fmt.Println(morestrings.ReverseRunes(msgs[name]))
		}
	}

	fmt.Println(cmp.Diff("hello world", "Hello World"))

	a := 12
	b := 4
	fmt.Printf("morestrings.Add(%v, %v) = %v \n", a, b, morestrings.Add(a, b))

	x := "World"
	y := "Hello"
	m, n := utils.Swap(x, y)
	fmt.Printf("utils.Swap(%v, %v)= %v, %v \n", x, y, m, n)

	const Pi = 3.14
	const Wd = "Word"
	const T = true
	fmt.Printf("Pi %v is type %T\nWd %v is type %T\nT %v is type %T\n", Pi, Pi, Wd, Wd, T, T)

	learnConstant()
	learnLoop()
	learnFor()
	learnSwitch()
}

func initLog() (*os.File, error) {
	//create a log file
	filename := "./log/hello" + strconv.FormatInt(time.Now().Unix(), 10) + ".log"
	file, err := os.Create(filename)
	if err != nil || file == nil {
		return nil, errors.New(fmt.Sprintf("Cannot open log file %v, due to %v", filename, err))
	}
	fmt.Printf("open file %v ", file)
	// defer file.Close()
	log.SetOutput(file)
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	log.Println("started ")

	return file, nil
}

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(i int) int {
	return i * 1
}
func needFloat(f float64) float64 {
	return f * 1.0
}
func learnConstant() {
	// fmt.Println(needInt(Big))
	fmt.Printf("int Small=%v, the type is %T\n", needInt(Small), needInt(Small))
	fmt.Printf("float Big=%v, the type is %T\n", needFloat(Big), needFloat(Big))
	fmt.Printf("float Small=%v, the type is %T\n", needFloat(Small), needFloat(Small))
}
func learnLoop() {
	sum := 0
	for i := 1; i < 10; i++ {
		sum = sum + i
	}
	fmt.Printf("the addition is %d", sum)
}

func learnSwitch() {

	os, found := os.LookupEnv("GOOS")
	if !found || os == "" {
		log.Printf("found %v, os %v\n", found, os)
		os = runtime.GOOS
	}
	fmt.Print("go runs on ")
	switch os {
	case "linux":
		fmt.Println("Linux.")
	case "darwin":
		fmt.Println("MAC os")
	default:
		fmt.Printf("%v os\n", os)
	}
}

func learnFor() {
	fmt.Println("sqrt(2)=" + sqrt(2) + "\nsqrt(-4)=" + sqrt(-4))

	f := 569.0
	fmt.Printf("sqrt2(%v)=%v\n", f, sqrt2(f))
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

//calculate the square root of a value
func sqrt2(x float64) float64 {
	debugmsg := debugmsg(true)
	log.Printf(debugmsg+"calculating sqare root of %v\n", x)
	z := x / 2
	dif := 10.0
	for i := 0; dif > 0.01; i++ {
		pz := z
		z -= (z*z - x) / (2 * z)
		dif = math.Abs(z - pz)
		log.Printf(debugmsg+"diff = %v, iterate %v\n", dif, i)
	}

	return z
}

func debugmsg(fullname bool) string {
	pc, _, _, _ := runtime.Caller(1)

	funcp := runtime.FuncForPC(pc)
	fmt.Println(funcp.Name())
	//length := len(path.Base(currentFile)) + len(fmt.Sprint(line)) + len(path.Ext(funcp.Name())) + len("   :: (, ) = ")

	if fullname {
		return funcp.Name() + "() "
	}

	return path.Ext(funcp.Name() + "() ")

}
