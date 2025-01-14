```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        ch := make(chan int, 10) // Buffered channel to avoid blocking

        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        fmt.Printf("Goroutine %d: ", i)
                        ch <- i // Send i to the channel
                }(i)
        }

        go func() {
                wg.Add(1) // Wait for receiving goroutine as well
                defer wg.Done()
                for i := 0; i < 10; i++ {
                        val := <-ch // Receive from the channel
                        fmt.Printf("Received %d\n", val)
                }
        }()

        wg.Wait()
        close(ch) // Close channel after all goroutines are done
}
```