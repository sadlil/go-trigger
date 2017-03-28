# go-trigger
Go Trigger is a global event trigger for golang. Define an event with a task specified to that
event and Trigger it from anywhere you want.

### Get The Package 
```bash
$ go get -u github.com/sadlil/go-trigger
```

### How to switch to a specific version
`go get` the package. Go to the package directory in your $GOPATH/src.
Change the tag using git.
`go install` the package.
 
```bash
$ go get -u github.com/sadlil/go-trigger
$ cd $GOPATH/src/github.com/sadlil/go-trigger
$ git checkout tags/<tag_name>
$ go install
```
#### Currently [available Tags](https://github.com/sadlil/go-trigger/releases)    
 - [v0.01](https://github.com/sadlil/go-trigger/releases/tag/v0.01)
     - Global event with unique key Id.
     - Trigger Event,
     - List Events,
     - Clear and Delete Events,
     - Trigger event in background,
     - Local Event,
     - Thread Safe Mutex Lock.

### How To Use
#### 1. Global Events
Import the package into your code. Add events with `trigger.On` method.
And call that event handler with `trigger.Fire` method. All the events added
like this will be global events. You can call `Fire` from anywhere.

```go
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


You can define your events from another package
```go
  trigger.On("second-event", packagename.FunctionName)
  trigger.Fire("second-event")
```


You can define events with parameteres and return types.
```go
func TestFunc(a, b int) int {
    return a + b
}

// Call them using
trigger.On("third-event", TestFunc)
values, err := trigger.Fire("third-event", 5, 6)

// IMPORTANT : You need to type convert Your Returned Values using
// values[0].Int()
```


You can define your event in one package and trigger it another package. Your event and trigger are global.
Define anywhere, fire anywhere. You can define any function in any package as event you only need to
import the function's specified package where you define the event. Where you trigger the event, you do not
need to import it there.
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
    values, err := trigger.Fire("new-event", 10, 10) 
    // You don't need to import package a here.
    fmt.Println(values[0].Int())
  }
```

You can run events in background with `FireBackground()`
```go
func main() {
  trigger.On("first-event", func() {
    for i := 1; i <= 1000; i++ {
      fmt.Println(i)
    }
  })
  channel, err := trigger.FireBackground("first-event")
  fmt.Println("Event runs")
  //read the returned channel
  values := <- channel
  
  trigger.FireBackground("first-event")
  fmt.Println("Running 2nd Event")
}

```

#### 2. Local Events
Trigger instance that will not effecct the global event. All event added to
an local event instace can call only via this trigger instance. This is
implementation of plugable `Trigger` interface.

Create a local trigger instance,

```go
package main

import (
  "github.com/sadlil/go-trigger"
  "fmt"
)


func main() {
  t := trigger.New()
  t.On("first-event", func() {
    // Do Some Task Here.
    fmt.Println("Done")
  })
  t.Fire("first-event")
  
  // t2 is another trigger instance that will be separate from t1.
  t2 := trigger.New()
  t2.On("first-event", func() {
    // Do Some Task Here.
    fmt.Println("Done")
  })
  t2.Fire("first-event")
}

```
**All other methods are availabe on any local trigger instance**


### Available methods
```go
On(event string, task interface{}) error
  - Add a Event. task must be function. Throws an error if the event is duplicated.
   
Fire(event string, params ...interface{}) ([]reflect.Value, error)
  - Fires the task specified with the event key. params are the parameter and
  [] is the returned values of task. Fire Triggers the event and wait for it to
  end until it goes to execute the following codes.
  
FireBackground(event string, params ...interface{}) (chan []reflect.Value, error)
  - Fires the task specified with the event key. Unlike Fire it runs the event in
  background in go routine. It triggers the event but does not wait for the event
  to end. It writes the returned values of the event in a channel and returns the
  channel of reflect.Values. You can get the returned values by reading the
  channel (IE. ret := <- returned channel).
  
  - As FireBackground does not wait for the event to end first, if your program 
  exits it will stop any running event that did not finishes. So make sure your
  background events exits before ending the program.   


Clear(event string) error
  - Delete a event from the event list. throws an error if event not found.
  
ClearEvents()
  - Deletes all event from the event list.
  
HasEvent(event string) bool
  - Checks if a event exists or not. Return true if the event list have a event with 
  that key.  false otherwise.
  
Events() []string
  - Returns all the events added.
  
EventCount() int
  - Returns count of the events. If none found return 0;
  
```


### Under Development Features
 1. Multiple event handler for a event.

### Licence
    Licenced under MIT Licence


##### Any Suggestions and Bug Report will be gladly appreciated.
