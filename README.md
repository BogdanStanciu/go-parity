This Go program processes a list of integers, classifies them as odd or even, and stores them in separate slices (odds and evens). It achieves this using goroutines for concurrent execution, channels for communication between goroutines, and a WaitGroup for synchronization.

## Step-by-Step Breakdown

### 1. Type Definition:

```go
type in struct {
    value    int32
    oddChan  chan int32
    evenChan chan int32
}
```

This struct in is designed to hold:

- value: An integer to be classified as odd or even.
- oddChan: A channel where odd numbers will be sent.
- evenChan: A channel where even numbers will be sent.

The struct allows each number to be paired with the necessary channels to send it to the correct destination (odd or even).

### 2. Global Channel Declaration:

```go
var serverChan chan in
```

serverChan is a global channel of type in. It is used to send data to a central server (goroutine) for processing.

### 3. Server Function (Goroutine):

```go
func Server() {
    for v := range serverChan {
        if v.value%2 == 0 {
            v.evenChan <- v.value
        } else {
            v.oddChan <- v.value
        }
    }
}
```

This function listens on serverChan in an infinite loop (or until the channel is closed).

For each message (v) received:
It checks if the value (v.value) is even (v.value % 2 == 0):
If even, it sends the value to the evenChan.
If odd, it sends the value to the oddChan.

The Server function is a central dispatcher that routes numbers to either oddChan or evenChan depending on their classification.

### 4. Main Function:

This is the entry point of the program, where all goroutines are set up.

Channels for Odd and Even Numbers:

```go
oddChan := make(chan int32)
evenChan := make(chan int32)
```

Two channels are created, oddChan for odd numbers and evenChan for even numbers.

Slicing Arrays for Results:

```go
odds, evens := []int32{}, []int32{}
```

Two slices odds and evens are initialized to store classified numbers.

WaitGroup Setup:

```go
wg := &sync.WaitGroup{}
wg.Add(2)

```

A sync.WaitGroup is initialized with a count of 2 to track the two goroutines that will process numbers from oddChan and evenChan.

The program will wait for both goroutines to finish before proceeding.

### 5. Goroutines for Receiving and Storing Odd and Even Numbers:

Odd Number Goroutine:

```go
go func() {
    for v := range oddChan {
        odds = append(odds, v)
    }
    wg.Done()
}()
```

A goroutine listens on oddChan. For each value received, it appends the value to the odds slice.
Once oddChan is closed (by the main goroutine), this goroutine finishes, and wg.Done() is called to signal that it’s complete.

Even Number Goroutine:

```go
go func() {
    for v := range evenChan {
        evens = append(evens, v)
    }
    wg.Done()
}()
```

Similar to the odd number goroutine, this one listens on evenChan, appends values to the evens slice, and signals completion using wg.Done() once evenChan is closed.

### 6. Launching the Server:

```go
go Server()
```

### 7. Sending Data to the Server:

```go
go func() {
    for _, v := range list {
        serverChan <- in{v, oddChan, evenChan}
    }
    close(oddChan)
    close(evenChan)
}()
```

Another goroutine iterates over the list of numbers (list := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
For each number v in the list, it creates an in struct, containing:

- v: The value itself.
- oddChan: The channel to send odd numbers.
- evenChan: The channel to send even numbers.

It sends the in struct to the serverChan to be processed by the Server() goroutine.
After all numbers have been sent, it closes both oddChan and evenChan to signal that no more data will be sent, which will eventually terminate the goroutines listening on these channels.

### 8. Waiting for Completion:

```go
wg.Wait()
```

The program waits until both goroutines (that are appending numbers to odds and evens) have finished their tasks. This is ensured by the WaitGroup.

### 9. Printing the Results:

```go
fmt.Println("odds")
for _, result := range odds {
    fmt.Printf("%d\n", result)
}
fmt.Println("evens")
for _, result := range evens {
    fmt.Printf("%d\n", result)
}
```

After all goroutines have finished, the program prints the contents of the odds and evens slices.

## Key Concepts

Concurrency and Goroutines: The program uses multiple goroutines to handle tasks concurrently:
One goroutine classifies numbers as odd or even (Server()).
Two additional goroutines append odd and even numbers to separate slices.
Another goroutine sends data to the server.

Channels: Channels (serverChan, oddChan, and evenChan) are used to communicate between goroutines safely without using shared memory.
serverChan dispatches the numbers to be processed.
oddChan and evenChan are used to collect odd and even numbers, respectively.

Synchronization: The sync.WaitGroup ensures that the main function waits for the two goroutines appending numbers to finish before printing the results.

Avoiding Deadlocks:

The channels are closed (close(oddChan) and close(evenChan)) once all the data has been sent. This prevents deadlocks, ensuring that the goroutines listening on these channels terminate gracefully after processing all incoming values.

The program will print the odd and even numbers from the list:

```bash
odds
1
3
5
7
9
evens
2
4
6
8
10
```

## Conclusion

The program effectively demonstrates how to utilize Go’s concurrency model, channels, and goroutines to process data in parallel. It classifies a list of integers into odd and even categories, and it handles communication and synchronization between concurrent tasks using channels and WaitGroup.
