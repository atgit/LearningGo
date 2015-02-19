// Philosophy
// Go is about composition rather than type hierachies
// Built-in asynchronous and synchronization primitives
// No generic


// Naming Convention
// file name     - lower case - hello.go
// pakcage name  - lower case, single word underscore is not recommended
//                 the package name is the base name of its source directory
// getter/setter - Owner()/SetOwner(user); var owner User
// 

package main

import (
    "fmt"
    "runtime"
    "math"
    "time"
    "reflect"
    "mypackage"
)


//// Variable
var i, j int = 1, 2   // Variables with initializers
var k = 3             // Type inference - based on the precision of the constant

// Group declaration can indicate relations between items
var (
    ii = 42           // int
    ff = 3.142        // float64
    gg = 0.867 + 0.5i // complex128
)

// Type conversion must be explicit (compared with C)
var f float64 = float64(i)


//// Constant
// - can be character, string, bool, or numeric values
// - created at compile time
const Number = 1      // untyped constant that can be used in any expression and converted implicitly
const True bool = true// typed constant

// const group
const (
    Big   = 1 << 100
    Small = Big >> 99
)

// iota in const group
const (
    One = 1<<iota   // 1, 1<<0
    Two             // 2, 1<<1
    Four            // 4, 1<<2
)


//// Pointer - hold the memory address of a variable
func pointers() {
    var p *int
    p = &i          // & operator generates a pointer to its operand
    fmt.Println("Type of *int: ", reflect.TypeOf(p));

    fmt.Println(*p) // * operator denotes the pointer's underlying value
    *p = 11         // Set value through a pointer

    p = nil         // zero value
}


//// Struct
type Vertex struct {
    X float64
    Y float64
}
func structs() {
    // Short variable declarations - commonly used for declaraing temparary local variables
    // - only in functions
    v := Vertex{1, 2}
    v.X = 3

    // Access through pointer
    p := &v
    p.X = 4
    fmt.Println(p.X)

    // Struct literals
    v1 := Vertex{1, 2}  // has type Vertex
    v2 := Vertex{X: 1}  // Y:0 is implicit
    v3 := Vertex{}      // X:0 and Y:0
    p  = &Vertex{1, 2} // has type *Vertex
    fmt.Println(v1, v2, v3)
}


//// Array
// - Arrays are values. Assigning one array to another copies all the elements - cautious to pass an array to a function
// - fixed size as the length is the part of the type
func arrays() {
    var arr [2]string
    arr[0] = "Hello, "
    arr[1] = "world!"

    // Or you can have the compiler count
    arr1 := [...]string{"Hello, ", "world!"}
    fmt.Println(arr1[0] + arr1[1])


    // Slice
    // - A slice is a descriptor of an array segment
    // - A slice consists of a pointer to an array, the length of the segment, and its capacity (the maximum length of the segment)
    // - Slicing does NOT copy data. Therefore, modifying data via a slice is essentially changing the data in its underlying array
    p := []int{2, 3, 5, 7, 11, 13}
    fmt.Printf("len: %v; cap: %v;\n", len(p), cap(p))
    fmt.Printf("p[1:4] == %v; len: %v; cap: %v;\n", p[1:4], len(p[1:4]), cap(p[1:4]))
    fmt.Printf("p[:3] == %v; len: %v; cap: %v;\n", p[:3], len(p[:3]), cap(p[:3]))  // missing low index implies 0
    fmt.Println("p[4:] == ", p[4:])  // Missing high index implies len(s)

    // Making slices
    a := make([]int, 0, 5) // len(b)=0, cap(b)=5
    fmt.Println(len(a), cap(a))

    // The range form of the for loop iterates over a slice or map
    // or 'for i := range p'
    // or 'for _, v := range p'
    for i, v := range p {
        fmt.Println(i, v)
    }

    // Growing slics
    s := p[1:4]
    ss := append(s, 15)
    fmt.Println(ss)     // [3 5 7 15]
    fmt.Println(p)      // [2 3 5 7 15 13] - If the append operation does not exceed the capacity of the underlying array, changes will apply to the current underlying array

    pp := append(p, 17) // Create an other array because the old one does not have more capacity
    fmt.Println(p)
    fmt.Println(pp)
}


//// Map
func maps() {
    var m map[string]Vertex

    // Must use make to create map
    m = make(map[string]Vertex)
    m["Home"] = Vertex {
        -31.817337,
        115.792259,
    }

    delete(m, "Home")
    elem, exist  := m["Home"]   // elem is zero and exist is false if not found
    fmt.Println(elem, exist)
}


//// Flow control
func flowcontrol() {
    //// Flow-control
    sum := 0
    for i =0; i<5; i++ {    // pre and post statements can be empty
        sum +=i
    }

    if v := math.Pow(10, 2); v < 100 {  // pre statements are allowed, like for
        fmt.Println(v)
    } else {
        // Variables declared inside an if short statement are also available inside any of the else blocks.
        fmt.Printf("%g >= %g\n", v, 100)
    }
    // Can't use v here, though

    switch os := runtime.GOOS; os {
        case "darwin":
        fmt.Println("OS X.")
        case "linux":
        fmt.Println("Linux.")
        default:
        fmt.Printf("%s.", os)
    }
}


//// Function
func add(x int, y int) int {    // The scope of function's params and return value is the same as the function body
    fmt.Println("add() entering ...")
    z := x + y    // Short assignment statement is not available outside of a func
    return z
}

// Return multiple values
func swap(x, y string) (string, string) {
    return y, x
}

// Named result parameters
func multiply(x, y int) (ret int) {
    ret = x * y
    return
}

// Closure
func counter() func(int) int{
    c := 0
    fmt.Println("Initial value: ", c)
    return func(v int) int {
        c += v
        return c
    }
}
func closures() {
    c := counter()

    fmt.Println(c(2))
    fmt.Println(c(3))
}


// Methods - Go does not have classes. Instead, you can define methods on struct types
// a method can be declared on any types in your package (not from another package)
// Why receiving a pointer?
// - Avoid copying the value
// - Allow modification
func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func methods() {
    v := &Vertex{3, 4}      // The type here must match with the method definition
    fmt.Println(v.Abs())
}


//// Interface
// - Defined by a set of methods
// - Interfaces are satisfied implicity - a type implements an interface by implementing the methods
//   - Benefits: decouple interface definition packages and implementation packages
type Abser interface {  // Naming convention - suffix 'er' is used when the interface only has one method
    Abs() float64
}
func interfaces() {
    var abser Abser
    v := &Vertex{3, 4}

    abser = v
    fmt.Println(abser.Abs())
}


//// Allocation
func allocation() {
    // new
    // - zero the memory
    // - return *T
    v := new(Vertex)
    fmt.Println("Zeroed by new: ", v.X, v.Y)

    // composite literal
    v1 := Vertex{1, 2}
    fmt.Println("Constrcted by composite literal: ", v1.X, v1.Y)
    // or by label the fields - the order of fields can be changed
    v2 := Vertex{Y:4, X:3 }
    fmt.Println("Constrcted by composite literal: ", v2.X, v2.Y)

    // make(T, args)
    // - create slices, maps and channels only
    // - initialized instead of zeroed object as those three object need to be initialized before use
    // - does not return a pointer
}


//// Goroutine
// - Lightweight thread managed by Go runtime
// - Run in the same address space
func goroutines() {
    go add(1, 2)
}


//// Channel
// - Typed
// - Send and receive operations are blocked until the other side is ready - synchronization without explicit locks
func sum(a []int, c chan int) {
    sum := 0
    for _, v := range a {
          sum += v
    }
    c <- sum        // Send a value to a channel
}
func channels() {
    a := []int{7, 2, 8, -9, 4, 0}
    ch := make(chan int, 100)     // Buffered channel - Send operation will be blocked when the buffer is full
go sum(a[:len(a)/2], ch)
    go sum(a[len(a)/2:], ch)
    v1st, v2nd := <-ch, <-ch    // Receive values from a channel
    fmt.Println(v1st, v2nd, v1st + v2nd)

    // Close a channel - only when the receiver must be told there are no more values; and should be done by the sender
    close(ch)
    for i := range ch {
          fmt.Println("This will not be printed as ch is closed", i)
    }

    // Select - wait and run randomly if multiple channels are ready
    ch_input := make(chan int)
    ch_quit_cmd := make(chan int)
    go func(ch chan int) {
          time.Sleep(500 * time.Millisecond)
          ch_input <- 1
            ch_quit_cmd <- 1
    }(ch_input)
    for {
        select {
        case i := <- ch_input:
            fmt.Println(i)
        case <- ch_quit_cmd:
            fmt.Println("Quitting ...")
            return
        default:        // The default case will run if no other case is ready
            fmt.Println("    .")
            time.Sleep(50 * time.Millisecond)
        }
    }
}



func main() {
    flowcontrol()
    pointers()
    structs()
    arrays()
    maps()
    closures()
    methods()
    interfaces()
    allocation()
    goroutines()
    channels()

    mypackage.SayHello("Yu");


    //// Defer the execution of a function until the surrounding function returns
    // The deferred call's arguments are evaluated immediately - when defer executs 
    // Stacking defers - called in last-in-first-out order
    defer fmt.Println("1st deferred call in main(): ", 1)
    defer fmt.Println("2nd deferred call in main(): ", 2)

    fmt.Println("main() leaving ...")
}
