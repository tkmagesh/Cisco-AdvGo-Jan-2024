# Go Advanced

## Magesh Kuppan

## Schedule
- Commence      : 9:00 AM
- Tea Break     : 10:30 AM (20 mins)
- Lunch Break   : 12:30 PM (1 hour)
- Tea Break     : 3:00 PM (20 mins)
- Wind up       : 5:00 PM

## Methodology
- No powerpoint
- Discussion & Code

## Repository
- https://github.com/tkmagesh/cisco-advgo-jan-2024

## Basics
- Data Types, Language Constructs
- Functions, Higher Order Functions
- Collections (Arrays, Slices, Maps)
- Pointer
- Error Handling, Panic & Recovery
- Type Assertion, Interfaces
- Structs, Methods
- Modules & Packages

## Recap
- Interfaces

## Managed Concurrency
- Concurrency
    - Ability to have more than one execution path
    - Concurrency is NOT parallelism
- Builtin scheduler
- Concurrent operations are represented as goroutines (cheap = 2KB)
- Language support for concurrency
    - go keyword
    - channel data type
    - channel operator ( <- )
    - range, select-case constructs
![image](./images/managed-concurrency.png)

### Concurrent Safe State Management
- To detect data race run/build with race detector
    - > go run -race [program.go]
    - > go build -race [program.go]
    - Note: DO NOT create a production build with race detector
### Communication (channel)
- Declaration
    - > var ch chan int
- Initialization
    - > ch = make(chan int)
- Declaration & Initialization (type inference)
    - > ch := make(chan int)
- Operation (using channel operator ( <- ))
    - Send Operation
        - > ch <- [data]
    - Receive Operation
        - > data := <- ch
- Behavior
    - A RECIEVE operation is ALWAYS a blocking operation (blocked until the data is sent through the channel)
    - A SEND operation is blocked until a RECIEVE operation is initiated
![image](./images/channel-behaviors.png)
