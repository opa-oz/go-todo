# go-todo: Prototype it later

Plan now - implement _later_

[![Run test](https://github.com/opa-oz/go-todo/actions/workflows/ci.yaml/badge.svg)](https://github.com/opa-oz/go-todo/actions/workflows/ci.yaml)

## Install
```bash
go get github.com/opa-oz/go-todo
```


## Features

Introducing our prototyping tool designed to streamline your development process!
Say goodbye to lengthy implementation phases and hello to rapid iteration. 
Our tool allows you to seamlessly replace complex functionality with placeholder 'todo' values,
enabling you to focus on the core features of your project without getting bogged down in the details.
With this innovative solution, you'll accelerate your prototyping timeline while maintaining
a clear roadmap for future development.

Get ready to bring your ideas to life faster than ever before!

### 1. "Magic" strings

Imagine you need to prefix all your logs.
It's easy! But what if your function is already in use? Of course, you *need* to refactor it. 

Just todo it and focus on something more important for the moment.

```go
func printFancy(message string) {
    fmt.Println(
        fmt.Sprintf("[%s] -> %s", todo.String("Move prefix to function params", "MyServer"), message),
    )
}

printFancy("My very important message")

// stdout:
// [TODO]: Move prefix to function params
```


### 2. "Magic" numbers

What if your task is to refactor famous [fast inverse square root](https://en.wikipedia.org/wiki/Fast_inverse_square_root)?

_Note:_ Yes, it's not fully translatable to Go, but I'd still like to include it as an example.
```go

func Q_rsqrt(number float32) float32 {
    var x2, y float32

    threehalfs := todo.Float32("Un-magic threehalfs", float32(1.5))
    x2 = number * todo.Float32("Un-magic 1/2", float32(0.5))
    y = number
    // Note: The bit-level manipulation in C is not directly translatable to Go.
    // Instead, we use Go's math.Sqrt function to compute the square root.
    y = 1 / float32(math.Sqrt(float64(y)))
    y = y * (threehalfs - (x2 * y * y)) // 1st iteration
    // y = y * (threehalfs - (x2 * y * y)) // 2nd iteration, this can be removed

    return y
}

Q_rsqrt(4.0)

// stdout:
// [TODO]: Un-magic threehalfs
// [TODO]: Un-magic 1/2
```

Don't forget! You can just mock the whole function!
```go
Q_rsqrt := todo.Float32F("Implement fast inverse square root", float32(0.5))

Q_rsqrt(4.0)

// stdout:
// [TODO]:Implement fast inverse square root 
```


### 3. "I'll do predicate later"

`Bool` is the simplest one to implement
```go
isOdd := todo.BoolF("Implement `isOdd`", false)
isEven := todo.BoolF("Implement `isOdd`", true)

assert.True(t, isEven(3))
assert.False(t, isOdd(3))

```


### 4. Custom errors

What if your code need to return an error, but you aren't sure how specific it should be? Just `todo` it!

```go

func processor(shouldFail bool) (*string, error) {
    if shouldFail {
        return nil, todo.Error("Add custom error here")
    }

    return todo.Ptr[string]("Find string somewhere"), nil
}

r, err := processor(true)

assert.NotNil(t, err)

// stdout:
// [TODO]: Add custom error here
```

### 5. Easily prototype math

_Todo description here+

```go
add := todo.IntF("Implement `add` function", 3)
substract := todo.IntF("Implement `substract` function", 1)

assert.Equal(t, add(1, 2), 3)
assert.Equal(t, substract(add(1, 2), 4), 1)

// stdout:
// [TODO]: "Implement `add` function"
//
// [TODO]: "Implement `add` function"
// [TODO]: "Implement `substract` function"
```

### 6. And more...

Have fun exploring!


----

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/S6S1UZ9P7)

