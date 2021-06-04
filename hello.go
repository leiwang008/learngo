package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"example.com/morestrings"
	"github.com/google/go-cmp/cmp"
	"github.com/leiwang008/utils"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
	"rsc.io/quote"
)

func main() {
	debugmsg := utils.Debugmsg(true)

	file, err := initLog()
	if err != nil {
		fmt.Printf("When initializing the log, met error:\n%v\n", err)
	} else {
		defer file.Close()
	}

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
	fmt.Printf("morestrings.Add(%v, %v) = %v \n\n", a, b, morestrings.Add(a, b))

	x := "World"
	y := "Hello"
	m, n := utils.Swap(x, y)
	fmt.Printf("utils.Swap(%v, %v)= %v, %v \n\n", x, y, m, n)

	const Pi = 3.14
	const Wd = "Word"
	const T = true
	fmt.Printf("Pi %v is type %T\nWd %v is type %T\nT %v is type %T\n\n", Pi, Pi, Wd, Wd, T, T)

	learnConstant()
	learnLoop()
	learnFor()
	learnSwitch()
	learnDefer()
	learnDeferStackCalls()
	learnPointer()
	learnStruct()
	learnArray()
	learnSlice()
	learnSlice2()
	learnpic()
	learnCountWord()

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
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

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
	fmt.Printf("float Small=%v, the type is %T\n\n", needFloat(Small), needFloat(Small))
}
func learnLoop() {
	sum := 0
	for i := 1; i < 10; i++ {
		sum = sum + i
	}
	fmt.Printf("the addition is %d\n", sum)
}

func learnSwitch() {
	debugmsg := utils.Debugmsg(true)

	os, found := os.LookupEnv("GOOS")
	if !found || os == "" {
		log.Printf(debugmsg+"found %v, os %v\n", found, os)
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

	//Switch match from top to bottom
	day := time.Tuesday
	fmt.Printf("\nWhen is %v?", day)
	switch day {
	case time.Now().Weekday():
		fmt.Println("It is today.")
	case time.Now().Weekday() + 1:
		fmt.Println("It is tomorrow.")
	case time.Now().Weekday() + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away")
	}
	fmt.Println()

	//switch without condition: might be a clean way to replace a long if-else clause
	hour := time.Now().Hour()
	fmt.Printf("It is %v now, what should we say?\n", hour)
	switch {
	case hour < 12:
		fmt.Println("Good morning.")
	case hour < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	fmt.Println()

}

func learnFor() {
	fmt.Println("sqrt(2)=" + sqrt(2) + "\nsqrt(-4)=" + sqrt(-4))

	f := 569.0
	fmt.Printf("sqrt2(%v)=%v\n\n", f, sqrt2(f))
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

//calculate the square root of a value
func sqrt2(x float64) float64 {
	debugmsg := utils.Debugmsg(true)
	log.Printf(debugmsg+"calculating sqare root of %v\n", x)
	z := x / 2
	dif := 10.0
	for i := 0; dif > 0.01; i++ {
		pz := z
		z -= (z*z - x) / (2 * z)
		dif = math.Abs(z - pz)
		log.Printf(debugmsg+"diff = %v, iterate %v\n", dif, i)
	}
	log.Printf(debugmsg+"The sqare root is %v", z)
	return z
}

//defer will postpone the execution until the surrounding function returns
//the deferred call's argruents are evaluated immediately, the the execution will not be performed until the surrounding function returns
func learnDefer() {
	defer fmt.Print(" World")
	fmt.Print("Hello")
	//the output should be "Hello World"
}

//deferred calls are pushed into a stack. When the surrounding function returns, the deferred call are executed in the FILO order
func learnDeferStackCalls() {
	fmt.Println("\nstart defer stack call")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
	//output
	//start defer stack call
	//done
	//9
	//8
	//...
	//0
}

func learnPointer() {
	i, j := 42, 2701

	fmt.Println()
	p := &i
	fmt.Printf("the pointer p %v is pointing at variable i, it's holding the value %v\n", p, *p)
	*p += 8
	fmt.Printf("change the variable i through its pointer p %v, now the i is %v\n", p, i)

	p = &j
	fmt.Printf("Now the pointer p %v is pointing at variable j, it's holding the value %v\n", p, *p)
	*p = *p / 37
	fmt.Printf("change the variable j through its pointer p %v, now the j is %v\n", p, j)

}

func learnStruct() {
	fmt.Println()
	v := utils.Vertex{X: 23, Y: 56}
	fmt.Printf("The structure Vertex v is %v, its address is %v\n", v, &v)
	v.X += -2
	v.Y -= 6
	fmt.Printf("The structure Vertex v is %v, its address is %v\n", v, &v)

	// var p *utils.Vertex
	// p = &v
	p := &v
	fmt.Printf("The pointer p %v is pointing at structure variable v, it is holding the value %v\n", p, *p)
	fmt.Printf("Changing the structure variable v %v through its pointer %v\n", v, p)
	p.X = 89
	p.Y = 90
	fmt.Printf("Now the structure variable v is %v\n", v)
	fmt.Printf("The pointer v == the pointer p? %v\n", p == &v)

	//struct literals
	s1 := utils.Vertex{X: 1, Y: 2}
	s2 := utils.Vertex{X: 5}
	s3 := utils.Vertex{}
	sp := &utils.Vertex{X: 1, Y: 2}

	fmt.Printf("s1=%v, s2=%v, s3=%v, p=%v, the sp is of Type %T\n", s1, s2, s3, sp, sp)

}

func learnArray() {

	fmt.Println()
	//assign array by each element
	var greetings [2]string
	greetings[0] = "Hello"
	greetings[1] = "World"
	fmt.Printf("%v %v\n", greetings[0], greetings[1])
	fmt.Printf("%v\n", greetings)

	//assing array with an array literal
	primes := [8]int{2, 3, 5, 7, 11, 13, 17, 19}
	fmt.Printf("Primes are %v\n", primes)

	//copy part of the array into a slice which is a dynamically-sized, flexible view into the elements of an array
	slice := primes[1:5]
	fmt.Printf("The slice is %v\n", slice)

	//Slices are like references to arrays, Changing the elements of a slice modifies the corresponding elements of its underlying array.
	slice[1] = 90
	fmt.Printf("Changing slice[1] to 90, now the original array also changed as %v\n", primes)

}

func learnSlice() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19}

	printSlice(primes) // 8, 8

	primes = primes[3:] //5, 5
	printSlice(primes)

	primes = primes[:0] //0, 5 ? why the capacity is still 5? should be 0!
	printSlice(primes)

	primes = primes[3:5] //2, 2
	printSlice(primes)

	primes = primes[0:] //2, 2
	printSlice(primes)

}

func printSlice(slice []int) {
	fmt.Printf("slice=%v length=%v, capacity=%v\n", slice, len(slice), cap(slice))
}

func learnSlice2() {
	fmt.Println()
	pow := make([]int, 10) // make() will initialize the range with its default value
	fmt.Printf("pow = %v\n", pow)

	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for i, value := range pow {
		fmt.Printf("pow[%v]=%v\n", i, value)
	}
}

func learnpic() {
	pic.Show(Pic)
}

func Pic(dx, dy int) [][]uint8 {

	fmt.Println()
	result := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		res := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			res[j] = uint8(dx * dy) //(dx+dy)/2, dx^dy
		}
		result[i] = res
	}

	return result
}

func learnCountWord() {
	wc.Test(wordCount)
}

func wordCount(text string) map[string]int {
	counts := make(map[string]int)

	tokens := strings.Fields(text)
	for _, token := range tokens {
		count, ok := counts[token]
		if ok {
			counts[token] = count + 1
		} else {
			counts[token] = 1
		}
	}
	return counts
}
