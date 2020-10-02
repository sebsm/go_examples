package main

import (
	"fmt"
)

var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Smith Sade"
	doctorNumber int    = 15
	season       int    = 20
)

func main() {
	var i int
	i = 10
	var j int = 30
	k := 40.
	fmt.Println(i)
	i = 10          //1010
	var div int = 3 //0011
	var vid int8 = 8
	fmt.Println(i)
	var stat bool = true
	fmt.Printf("%v, %T\n", j, j)

	fmt.Printf("%v, %T\n", k, k)

	fmt.Printf("%v, %T\n", stat, stat)

	fmt.Println(i + div)
	fmt.Println(i - div)
	fmt.Println(i * div)
	fmt.Println(i / div)
	fmt.Println(i % div)

	fmt.Println(i + int(vid))

	fmt.Println(i & div)  //0010
	fmt.Println(i | div)  //1011
	fmt.Println(i ^ div)  //1001
	fmt.Println(i &^ div) //0100

	a := 8              //2^3 * 2^3 = 2^6
	fmt.Println(a << 3) //2^3 / 2^3 = 2^6
	fmt.Println(a >> 3) //2^3 * 2^3 = 2^0

	n := 3.14
	n = 13.7e72
	n = 2.1e14
	fmt.Printf("%v, %T\n", n, n)

	var o complex64 = 2i
	fmt.Printf("%v, %T\n", o, o)

	g := 10.2
	h := 3.7
	fmt.Println(g + h)
	fmt.Println(g - h)
	fmt.Println(g * h)
	fmt.Println(g / h)

	z := 5 + 7i
	x := 9 - 4i
	fmt.Println(z + x)
	fmt.Println(z - x)
	fmt.Println(z * x)
	fmt.Println(z / x)

	var l complex128 = 3 + 6i
	fmt.Printf("%v, %T\n", real(l), real(l))
	fmt.Printf("%v, %T\n", imag(l), imag(l))

	var d complex128 = complex(1, 2)
	fmt.Printf("%v, %T\n", d, d)

	///////////////////////////////////////////////STRINGS///////////////////////////////////

	s := "Ok, thanks, yes, no"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", string(s[2]), s[2])

	s2 := "No, yes, goodbye"
	fmt.Printf("%v, %T\n", s+s2, s+s2)

	s3 := []byte(s2)
	fmt.Printf("%v, %T\n", s3, s3)

	s4 := 'b'
	fmt.Printf("%v, %T\n", s4, s4)

	/////////////ARRAYS/SLICES//////////////////////
	g1 := 87
	g2 := 90
	g3 := 93
	fmt.Printf("G's: %v, %v, %v\n", g1, g2, g3)

	g4 := []int{0, 2, 3, 4, 5}
	fmt.Printf("G's: %v\n", g4)

	var student [3]string
	student[0] = "Bob"
	student[1] = "Lila"
	student[2] = "Luna"
	fmt.Printf("Student: %v\n", student)
	fmt.Printf("Number of students: %v\n", len(student))
	fmt.Printf("Capacity: %v", cap(student))

	v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	q := v[:]
	w := v[3:]
	r := v[:6]
	t := v[3:6]

	fmt.Println(v)
	fmt.Println(q)
	fmt.Println(w)
	fmt.Println(r)
	fmt.Println(t)

	qw := make([]int, 3, 100)
	fmt.Println(qw)
	fmt.Printf("Length: %v\n", len(qw))
	fmt.Printf("Capacity: %v\n", cap(qw))

	qe := []int{}
	fmt.Println(qe)
	fmt.Printf("Length: %v\n", len(qe))
	fmt.Printf("Capacity: %v\n", cap(qe))

	qe = append(qe, 2)
	fmt.Println(qe)
	fmt.Printf("Length: %v\n", len(qe))
	fmt.Printf("Capacity: %v\n", cap(qe))

	///////////////////////MAPS/STRUCTS////////////////////////
	statePop := make(map[string]int)
	statePop = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
	}
	fmt.Println(statePop)
	statePop["c"] = 9
	fmt.Println(statePop["b"])

	pop, err := statePop["z"]
	fmt.Println(pop, err)
	_, erd := statePop["a"]
	fmt.Println(erd)

	aDoc := struct{ name string }{name: "John Peer"}
	nextDoc := &aDoc
	nextDoc.name = "Tom Been"
	fmt.Println(aDoc)
	fmt.Println(nextDoc)

	type bird struct {
		name   string
		weight float64
		height float64
	}
	///////IF
	number := 50
	quess := -10
	if quess < 1 || quess > 100 {
		fmt.Println("The quess must be between 1 and 100")
	}
	if quess > 1 && quess < 100 {
		if number > quess {
			fmt.Println("Too low")
		}
		if number == quess {
			fmt.Println("Exactly this one")
		}
		if number < quess {
			fmt.Println("Too high")
		}
	}
	//////SWITCH
	switch 3 {
	case 3:
		fmt.Println("one")

	case 2:
		fmt.Println("two")
	default:
		fmt.Println("Neither of them")
	}

	var qt interface{} = "8"
	switch qt.(type) {
	case int:
		fmt.Println("qt is an integer")
	case float64:
		fmt.Println("qt is a float64")
	case string:
		fmt.Println("qt is a string")
	default:
		fmt.Println("qt is something else")
	}
	////FOR
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			i /= 2
		} else {
			i = 2*i + 1
		}

	}
	gh := 0
	for gh < 5 {
		fmt.Println(gh)
		gh++
		if gh == 5 {
			break
		}

	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)

	}

	sk := [4]int{0, 1, 2, 3}
	for k, v := range sk {
		fmt.Println(k, v)
	}

	sv := "Hello"
	for k, v := range sv {
		fmt.Println(k, string(v))
	}

	///POINTERS
	var af int = 42
	var bf *int = &af
	fmt.Println(af, *bf)
	af = 27
	fmt.Println(af, *bf)
	*bf = 14

	type str struct {
		foo int
	}
	var ms *str
	ms = new(str)
	(*ms).foo = 45
	fmt.Println((*ms).foo)
	dvd := divide(5.4, 3.7)
	fmt.Println(dvd)

	////FUNCS///////
	var divi func(float64, float64) (float64, error)
	divi = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by 0")
		}
		return a / b, nil
	}
	der, errr := divi(5.0, 0.0)
	if errr != nil {
		fmt.Println(errr)
		return
	}
	fmt.Println(der)

	////INTERFACES

	var wa writer = consoleWriter{}
	wa.Write([]byte("Hello guys!"))

	myInt := intCounter(0)
	var inc incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

type incrementer interface {
	Increment() int
}

type intCounter int

func (ic *intCounter) Increment() int {
	*ic++
	return int(*ic)
}

type writer interface {
	Write([]byte) (int, error)
}
type consoleWriter struct {
}

func (cwk consoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func divide(a, b float64) float64 {
	if b == 0 {
		panic("Can't divide by zero)")
	}
	return a / b
}
