# go-trigger
Go Trigger is a global event trigger for golang. Define an event with a task specified to that
event and Trigger it from anywhere you want.

### Get The Package 
```bash

$ go get github.com/sadlil/go-trigger

```

### How To Use

Import the package into your code. Add the events with `trigger.On` method.
And call that event handler with `trigger.Fire` method.

````go
package main

import (
  "github.com/sadlil/go-trigger"
  "fmt"
)


func main() {
  trigger.On("first-event", func() {
    // Do Some Task Here.
    fmt.Println("Done")
  })
  trigger.Fire("first-event")
}

```


You can define Your events from another package
```go
  trigger.On("second-event", packagename.FunctionName)
  trigger.Fire("second-event")
```


You can Define events with parameteres and return types.
```go
func TestFunc(a, b int) int {
    return a + b
}

// Call Them Using
trigger.On("third-event", TestFunc)
values, err := trigger.Fire("third-event", 5, 6)

// IMPORTANT : You need to type convert Your Returned Values using
// values[0].Int()
// I will try fix this in next version.

```


You can define your event in one package and trigger it another package. Your define and triggers are global.
Define anywhere, fire anywhere. You can define any function in any package u only need to import the function's
package where you defien it. Where You trigger it You do not need to import it there.
```go
//---------------------------------------------
  package a
  
  func AFunction(one, two int) int {
    return one + two
  }
//---------------------------------------------
  package b
  import (
    "yourdirectory/a"
    "github.com/sadlil/go-trigger"
  )
  
  func() {
    trigger.On("new-event", a.AFunction)
  }
//---------------------------------------------
  package c
  import (
    "github.com/sadlil/go-trigger"
  )
  
  func() {
    values, err := trigger.Fire("new-event", 10, 10) // You dont need to import package a here.
    fmt.Println(values[0].Int())
  }
```


### Under Development Feautures
 1. Return already type converted values from Fire.
 2. Add support of Methods on structs events.

### Licence
    Licenced under MIT Licence




Any Suggestions and Bug Report will be gladly appricated.

