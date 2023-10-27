# Go study notes

## Get started

### Initialise a Go project

```bash
go mod init github.com/user/repo
```

### Build a Go project

```bash
go build main.go
```

### Run a Go project

```bash
go run main.go
```

### Run tests

```bash
go test main_test.go
```

### Install a package

```bash
go get github.com/author/repo@version
```

### Install third-party tools

```bash
go install github.com/user/repo@version
```

---

## Variable declaration

Multiple ways to declare variables.

> **Note** > **Zero value**  
> Any declared variable that doesn't have a value assigned is automatically assigned the _zero value_ for the corresponding type.

### `var`

Declare a variable using `var`

```go
var foo int // declared and assigned the zero value, `0` for the `int` type
```

Declare and assign a value to a variable using `var`

```go
var foo int = 23

// since the default type for integer literals is `int`, the type annotation can be omitted
var foo = 23
```

Declare multiple variables using `var`

```go

var foo, bar int = 23, 42

// again, type annotation can be omitted
var foo, bar = 23, 42

// different types can be infered and assigned
var baz, qux = 69, "foo"
```

Declare multiple variables in a _declaration list_

```go
var (
    foo int
    bar string = "hello"
    baz = true
)
```

### `:=`

Short declaration format can be used in place of a `var` declaration where type inference is being used _and_ where the variable is being declared in a block scope (i.e. not package level).

```go
foo := 23 // equivalent to `var foo = 23`

baz, qux = 69, "foo" // equivalent to `var baz, qux = 69, "foo"`
```

### `const`

`const` is used to declare an immutable variable.

> **Note**  
> Constants can only be assigned values that can be determined at compile-time.  
> In practice, this means they can only be assigned literals.

```go
const foo int = 23

const bar = "hello"

const (
    baz = 69
    qux = "foo"
)

```

---

## Literals, default and inferred types

Go has five kinds of literals.

### 1. String literal

Interpreted string literal

```go
"Hello, world!"
```

Raw string literal

```go
`hello
world`
```

### 2. Integer literal

```go
23 // base ten
42_000_000 // base ten with underscores for readability (no different to just `42000000`)

0b11100 // binary

0o660 // octal

0x123af // hexadecimal
```

### 3. Floating point literal

```go
23.42 // floating point literal

6.42e23 // floating point literal with exponent
```

### 4. Rune literal

```go
'a' // unicode rune literal
```

### 5. Imaginary value literals

_Not interested in imaginary numbers at this point; come back to subject in future_

---

## Primitive types

### Booleans

`bool` represents boolean variables.

They can be assigned a value of `true` or `false`.

The zero value for the `boolean` type is `false`.

```go
var toggle bool // automatically assigned the zero value: `false`

var toggle = true

toggle := true
```

### Strings

`string` represents a string of one or more Unicode characters.

The zero value for the `string` type is an empty string.

The `+` operator is used to concatenate strings.

```go
var message string // automatically assigned the zero value: `""`
message = "Hi, there!"

var welcome = "Hello, world!"

hello := "Hello"
world := "world"

greeting := hello + ", " + world + "!"
```

### Numerics

Go has 12 numeric types across three categories.

| Category | Type      | Value                                       | Description                      |
| -------- | --------- | ------------------------------------------- | -------------------------------- |
| Integer  | `int8`    | -128 to 127                                 | 8 bit integer                    |
| Integer  | `int16`   | -32768 to 32767                             | 16 bit integer                   |
| Integer  | `int32`   | -2147483648 to 2147483647                   | 32 bit integer                   |
| Integer  | `int64`   | -9223372036854775808 to 9223372036854775807 | 64 bit integer                   |
| Integer  | `uint8`   | 0 to 255                                    | 8 bit unsigned integer           |
| Integer  | `uint16`  | 0 to 65535                                  | 16 bit unsigned integer          |
| Integer  | `uint32`  | 0 to 4294967295                             | 32 bit unsigned integer          |
| Integer  | `uint64`  | 0 to 18446744073709551615                   | 64 bit unsigned integer          |
| Floating | `float32` | See IEEE 754 spec                           | 32 bit floating point (IEEE 754) |
| Floating | `float64` | See IEEE 754 spec                           | 64 bit floating point (IEEE 754) |
| Complex  | -         | -                                           | -                                |

Aliases and special integer types

| Type   | Description                                                             |
| ------ | ----------------------------------------------------------------------- |
| `int`  | `int32` on 32 bit platform; `int64` on 64 bit platform                  |
| `unit` | Same rules as `int`, just unsigned (i.e. values are all 0 and positive) |
| `byte` | Alias for `uint8`                                                       |

#### Arithmetic operators

The usual arithmetic operators are supported for performing math.

| Operator | Operation      |
| -------- | -------------- |
| `+`      | Addition       |
| `-`      | Subtraction    |
| `/`      | Division       |
| `*`      | Multiplication |
| `%`      | Modulus        |

#### Comparison operators

Comparison operators work on numerics and return a boolean `true` or `false`.

| Operator | Comparision                 |
| -------- | --------------------------- |
| `==`     | Is equal to                 |
| `!=`     | Is not equal to             |
| `>`      | Is greater than             |
| `<`      | Is less than                |
| `>=`     | Is greater than or equal to |
| `<=`     | Is less than or equal to    |

#### Bitwise operators

The usual bitwise operators are supported.

| Operator | Operation      |
| -------- | -------------- |
| `<<`     | Bitshift left  |
| `>>`     | Bitshift right |
| `&`      | Logical AND    |
| `\|`     | Logical OR     |
| `^`      | Logical XOR    |

#### Modifying numeric variables in place

All arithmetic and bitwise operators can be combined with `=` to perform an operation and assign back to the variable.

```go
x, y := 23, 69

x += 1 // 24
y -= 27 // 42

x <<= 1 // 48
y >>= 1 // 21
```

#### Complex numbers

_Not interested in this topic currently; revisit at a later point_

### Runes

TODO:

---

## Composite types

### Arrays

Arrays are rarely actually used in Go. Usually, it's a slice that is wanted.

Arrays are zero-indexed.

Declare an empty array of ints with a length of 5. All of the positions are initialised to the empty value for int, which is `0`.

```go
var arr = [5]int

arr[0] // 0
arr[1] // 0
arr[2] // 0
arr[3] // 0
arr[4] // 0

```

Declare an array of ints with length of 5 and initialise with an array literal

```go
arr1 := [5]int{5, 10, 15, 20, 25}

arr1[0] // 5
arr1[1] // 10
arr1[2] // 15
arr1[3] // 20
arr1[4] // 25

arr2 := [...]int{30, 40, 50}

arr2[0] // 30
arr2[1] // 40
arr2[2] // 50
```

Declare a sparse array; where only some elements are initialised with a value are others are initialised with their zero value.

```go
arr := [10]int{5, 3: 20, 25, 8: 45}
// [5, 0, 0, 20, 25, 0, 0, 0, 45, 0]

```

Simulate a multi-dimensional array

```go
arr := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}

// [[1, 2, 3], [4, 5, 6]]
```

### Slices

### Maps

#### Sets

### Structs

- include struct embedding

### `make`, `len`, `append`, `copy` and `delete`

---

## Generics

---

## Control flow

---

## Blocks and scope

---

## Methods

---

## Interfaces

---

## Errors

---

## Concurrency

---

## I/O

---

## Time

---

## JSON

---

## HTTP

---

## Context

---

## Testing

---

## Reflection
