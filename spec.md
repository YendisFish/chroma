# Functions

```go
func main() {
    io.println(first := "Hello, World!", first.length)
}

func add(a int, b int) int {
    return a + b
}
```

There are no private functions! They're for LOSERS!!!

# Variables

```go
var foo string = "Foo"
var foo = "Foo"
foo := "Foo"
```

# Types

```go
type string struct {
    unsafeChars *rune
    length int
}

func (s *string) struct *rune {
    return s.unsafeChars
}

func (s *string) [](index int) rune {
    if s.length > index && index >= 0 {
        return s.unsafeChars[index]
    }

    panic("Index out of bounds!")
}

func (s *string) [](index int, value rune) {
    if s.length > index && index >= 0 {
        s.unsafeChars[index] = value
    }

    panic("Index out of bounds!")
}
```

### Inheritance

```go
type A struct {
    prop string
}

type B A

func (b *B) PrintProp() {
    io.println(b.prop, b.prop.length)
}
```

### Interfaces

```go
type IList[T] interface {
    elements *T
    length int
    capacity int

    Add()
    Remove()
}

type List[T] struct {
    elements *T
    length int
    capacity int
}

func (l *List) Add() {

}

func (l *List) Remove() {

}

// List implements IList!!!
func main() {
    var lst *IList[int] = &List[int]{} //interfaces always take pointers as they are reference types!
}
```

### Union Types

```c
union MyUnion {
    union Foo(a string)
    union Bar {
        union Somethingcharacters
    }
}
```

You can implement functions for union types as well!

```go
func (a *Something) UnionFunc() { }
```

Union types MUST be by reference!

### Implementing Instance Functions

```go
type A struct {
    Str string
}

func (a *A) MyFunction() {
    io.println(first := a.Str, first.length)
}

```

# Loops/Statements

### If

```go
if x < 5 {
    //do something
}
```

### Error

```go
x := error err := someFunc() {
    //handle errors
}

//use x
```

### Switch

```go
switch x {
    case 1 => {
        //run some code
    }
    case 2 => x = x + 7
    default => x = 0
}

//switching on types
switch type x {
    type int => {
        //run code
    }
}

```

### For

```go
for i, num range x {
    io.println(num)
}

//I do not need num
for i, _ range x {
    io.println(num)
}

//range for loops always need 'i' to be defined since they use it under the hood

//traditional for loops
for i := 0; i < x.length; i++ {
    err := error x[i]
    io.println(x[i])
}
```

**PLEASE NOTE THAT MOST INDEXERS DO NOT RETURN ERRORS BUT JUST PANIC!!!!**

### While

```go
while true {
    //do something
}
```

# Allocation Control

```go
//stack allocated
var x string := string{}

//heap allocated
var x *string = &string{5 /*number of elements in buffer*/} // requires package allocator to be defined
```

# Allocators

```go
package myallocator type allocator

import func malloc(size int) *interface {}
import func free(ptr *interface) {}

func &[T](size int) *T {
    return malloc(sizeof T * size)
}

func delete[T](ptr *T) {
    free(ptr)
}
```

Using allocators

```go
select myallocator package main
select myallocator func someFunction()
select myallocator struct A
select myallocator default //default allocator

import {
    "myproject/myallocator"
}

func someFunction() {

}

type A struct {

}
```

# Type Alias

```go
import {
    "corelib/string"
    string string.string
}
```

# Conversion Operators

```go
type A struct {
}

func (a *A) struct B {
    return (B)a
}

type B struct {
}

func (b *B) struct A {
    return (A)b
}
```

# Freeing Pointers

```go
x := &string{}
*x = "Hello World!"

defer x
delete x
```

# From operator

The from operator will copy by reference *into* a new type

```go
a := &string{5}
var b [5]string = from a // the reference to a is still stored in b, we have just copied the reference into our array
var c []string = from b[5] // we can also copy by reference into a slice (however this functionality is more shorthand than anything)
```

# Stack Buffers

```go
func Foo() *byte {
    const buff [5]byte // we have inlined a 5 byte buffer onto the stack
    //these buffers cannot be returned, and are not allowed to escape their enclosing scope
    //well... they are... but you need to be explicit to do so, for instance:
    return &buff[0] //this will cause UB
}

//we can also inline buffers in structs
type A struct {
    const buff [5]byte
}
```
