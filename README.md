# Go Data Race Example

This repository demonstrates a common data race in Go programs involving concurrent channel access. The `bug.go` file contains the buggy code, and `bugSolution.go` shows the corrected version.

## Bug Description

The original code launches 10 goroutines, each sending a value to a shared channel.  Another goroutine receives from this channel. Without proper synchronization, the program exhibits a data race because multiple goroutines try to access the channel concurrently, leading to unpredictable behavior.

## Solution

The solution uses a `sync.WaitGroup` to ensure all sending goroutines complete before the receiving goroutine finishes, preventing the data race.  The `close(ch)` function is called after the `wg.Wait()` to safely close the channel once all goroutines have finished.