# go-trigger
Go Trigger is a global event trigger for golang. Define an event with a task specified to that
event and Trigger it from anywhere you want.

### Get The Package 
`
$ go get github.com/sadlil/go-trigger
`


### How To Use

Import the package into your code. Add the events with `trigger.On` method.
And call that event handler with `trigger.Do` method.
