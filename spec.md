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
struct string {
    unsafeChars *rune
    length int

    func struct(chars *rune, len int) {
        return string{unsafeChars: chars, length: len}
    }

    func [](index int) rune {
        if length > index && index >= 0 {
            return index, unsafeChars[index]
        }

        error "Index out of bounds!"
    }

    func [](index int, value rune) {
        if length > index && index >= 0 {
            unsafeChars[index] = value
        }

        error "Index out of bounds!"
    }

    func struct() *rune {
        return unsafeChars
    }
}
```

### Inheritance

```go
struct A {
    Str string
}

struct B A;

func (struct B) MyFunction() string {
    return struct.Str
}
```

### Union Types

```c
union MyUnion {
    union Foo(a string)
    union Bar {
        union Something
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
struct A {
    Str string
}

func (struct *A) MyFunction() {
    io.println(first := struct.Str, first.length)
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

func defer[T](ptr *T) {
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

struct A {

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
struct A {
    func struct() B {
        return (B)struct
    }
}

struct B {
    func struct() A {
        return (A)struct
    }
}
```

# Freeing Pointers

```go
x := &string{}
*x = "Hello World!"

defer x
free x
```