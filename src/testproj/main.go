package main

import (
	"fmt"
	"reflect"
	"time"
	"unicode/utf8"
	"unsafe"
)

func main() {
	//var str string
	//str = "Hello, world!"
	//fmt.Println(string(str[0]))

	var a string
	a = `Превет`
	type btype string

	fmt.Println("Code number = ", a[0])
	fmt.Println("bite = ", len(a))
	fmt.Println("RUne = ", utf8.RuneCountInString(a))
	fmt.Println(reflect.TypeOf(a))

	if a == "" || len(a) == 0 {

	}

	var (
		nowTime         = time.Now()
		pi      float32 = 3.14159
	)
	const (
		Black = iota + 1
		Gray
		White
		Yellow
		Blue
	)

	fmt.Println(nowTime, pi, " --- ")
	fmt.Println(Black == Gray, " ---- ", White == Yellow, " ---- ", Black, Gray, White, Yellow, Blue)

	type Weekday int
	const (
		Monday Weekday = iota + 1
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	var today Weekday = Tuesday
	tomorrow := today
	fmt.Println("today =", today, "tomorrow =", tomorrow)

	const (
		one = 2*iota + 1
		three
		five
		seven
		nine
		eleven
	)

	fmt.Println("\n", one, "\t", three, five, seven, nine, eleven)

	// import variables from package Contacts.go
	//contacts.SetSupport("Служба поддержки")
	//fmt.Println(contacts.GetContact())
	//fmt.Println("Email:", contacts.Email)

	date := time.Date(1982, 04, 12, 0, 0, 0, 0, time.UTC)

	switch y := date.Year(); {
	case y >= 1946 && y < 1965:
		fmt.Println(date.Year())
	case y >= 1965 && y < 1981:
		fmt.Println("Привет, представитель X!")
	case y >= 1981 && y < 1997:
		fmt.Println("Привет, миллениал!")
		fallthrough
	case y >= 1997 && y < 2013:
		fmt.Println("Привет, зумер!")
		fallthrough
	case y >= 2013:
		fmt.Println("Привет, альфа!")
	default:
		fmt.Println("Здравствуйте!")
	}
	fmt.Println("--------------------- \n\n")

	var v uint = 10
	for i := 0; i < 10; i++ {
		v++
		fmt.Println(v)
	}
	fmt.Println(" inn ", v)

	fmt.Println("--------------------- \n\n")

	array := [3]int{1, 3, 5}
	for k, v := range array {
		fmt.Printf("array[%d]: %d\n", k, v)
	}

	fmt.Println("--------------------- \n\n")

	sum, limit := 0, 10
	for i := 1; true; i++ {
		if i%2 != 0 {
			continue // переход к следующему числу, так как i — нечётное
		}

		if sum+i > limit {
			break // выход из цикла, так как сумма превысит заданный предел
		}

		sum += i
		fmt.Println("--", i)
		fmt.Println(sum)
	}

	fmt.Println("--------------------- \n\n")

outerLoopLabel:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("[%d, %d]\n", i, j)
			continue outerLoopLabel

		}
		fmt.Printf("[out %d]\n", i)

	}
	fmt.Println("End")

	fmt.Println("--------------------- \n\n")

	for i := 1; i <= 20; i++ {
		switch {
		case i%3 == 0:
			fmt.Println("Fizz")
			continue
		case i%5 == 0:
			fmt.Println("Buzz")
			continue
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
			continue
		}
		fmt.Printf("%d \n", i)
	}

	fmt.Println("--------------------- \n\n")

	for i := 1; i <= 16; i++ {
		found := false

		if i%3 == 0 {
			fmt.Printf("Fizz")
			found = true
		}
		if i%5 == 0 {
			fmt.Printf("Buzz")
			found = true
		}

		if !found {
			fmt.Println(i)
			continue
		}

		fmt.Println("-")
	}

	fmt.Println("--------------------- \n\n")

	var a22 *int
	p22 := a22

	fmt.Printf("%d - %T \n", a22, p22)
	fmt.Println(a22, p22)
	fmt.Println(reflect.TypeOf(p22))

	fmt.Println("--------------------- \n\n")

	//const c = 5
	//p1 := &"abc" // ошибка компиляции
	//p2 := &с // ошибка компиляции

	type A struct {
		IntField int
	}
	p := &A{
		IntField: 10,
	}
	fmt.Println(p)

	fmt.Println("--------------------- \n\n")

	a23 := 1
	p23 := &a23
	b23 := &p23
	v23 := &b23
	*p23 = 3
	**b23 = 5

	a23 += 4 + *p23 + **b23

	fmt.Println(*p23, **b23)

	fmt.Printf("%T", a23, p23, b23, v23, "\n")
	fmt.Println(a23, p23, b23, v23)

	fmt.Println("--------------------- \n\n")
	/*
		считает, сколько строк пользователь ввёл в консоль, и после ввода каждой новой строки выводит общее количество на экран. Напишите реализацию функции f.
	*/
	/*reader := bufio.NewReader(os.Stdin)
	fmt.Println("Interaction counter")
	cnt := 0

	for {
		fmt.Print("->")
		_, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		f(&cnt)

		fmt.Printf("User input %d lines\n", cnt)
	}*/

	fmt.Println("---------- arrays ----------- \n")
	//arrays

	arr := [...]int{2, 255, 225}
	fmt.Println("all: ", arr, "along ", arr[0])

	thisWeekDay := [7]int{0: 99, 5: -44, 6: 888}
	fmt.Println("SIZE bit`s:", unsafe.Sizeof(thisWeekDay))
	fmt.Println("type :", reflect.TypeOf(thisWeekDay)) // unique type

	//если я хочу сделать массив динамическим то нужно использовать только констануту
	const a01 = 10
	var arr01 = [a01]int{2, 5, 6, 9}
	fmt.Print(arr01, "\n")

	var weekTemp = [7]int{15, 4, 6, 8, 11, 9, 5}
	if thisWeekDay != weekTemp {
		fmt.Println("dont equally arrays")
	}

	// method 1 sequential
	sumTemp := 0
	for i := 0; i < len(weekTemp); i++ {
		sumTemp += weekTemp[i]
	}
	average := sumTemp / len(weekTemp)
	fmt.Println("average sum: ", average)

	// method 2
	sumTemp2 := 0
	for _, temp := range &weekTemp {
		sumTemp2 += temp
	}
	average2 := sumTemp2 / len(weekTemp)
	fmt.Println("average sum2: ", average2)
	fmt.Println("type :", reflect.TypeOf(weekTemp))

	//Многомерные массивы
	var thisMonthDay [3][7]int
	fmt.Println(thisMonthDay)
	fmt.Println(thisMonthDay[2][6])

	fmt.Println("---------- slice ----------- \n")
	//slice

	var slice []int = make([]int, 5, 9)
	fmt.Print("type :", reflect.TypeOf(slice), "\n")
	fmt.Println(slice)

}

func f(cnt *int) {
	*cnt++
}
