# Goroutines

`Golang` was a language created with a focus on **performance on computers with more than one processing core**. It is a language with high-level language syntax with performance close to low-level languages. One of the amazing features that the language offers is the facility to manage concurrency


In practice, this means we can jointly perform several tasks (the number is conditioned by the number of cores) . 


repository we have two tasks that run in parallel.

> [calculator](calculator/calculator.go)  <br/>

> [fake-download-file](fake-download-file/fake-download-file.go)

Both are treated by Go as routines.

* In the main function we invoke both functions and instantiate the WaitGroup method from the [sync package](https://golang.org/pkg/sync/)  (the package that will help us manage the routines) . 

* Using the [Add method](https://golang.org/pkg/sync/#WaitGroup.Add) we inform the total number of routines to be executed

* We pass the WaitGroup method instance as an argument to our routines and we use the keyword `go` to characterize that it is a Goroutine
.

* And lastly we use the [defer method](https://tour.golang.org/flowcontrol/12) in our routines to inform the sync package that our functions have completed

