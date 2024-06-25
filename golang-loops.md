What are some of the ways you can loop over different data structs/arrays in Golang?
ChatGPT
In Go, there are several ways to loop over data structures like arrays, slices, maps, and channels. Here are the main looping constructs:

For Loop: The basic for loop can be used to iterate over arrays, slices, and strings.
go
Copy code
// Iterating over an array
arr := [3]int{1, 2, 3}
for i := 0; i < len(arr); i++ {
fmt.Println(arr[i])
}

// Iterating over a slice
slice := []int{1, 2, 3}
for i := 0; i < len(slice); i++ {
fmt.Println(slice[i])
}

// Iterating over a string
str := "hello"
for i := 0; i < len(str); i++ {
fmt.Println(string(str[i]))
}
Range-based For Loop: The range keyword is used to iterate over arrays, slices, strings, and maps.
go
Copy code
// Iterating over an array
arr := [3]int{1, 2, 3}
for \_, value := range arr {
fmt.Println(value)
}

// Iterating over a slice
slice := []int{1, 2, 3}
for index, value := range slice {
fmt.Println(index, value)
}

// Iterating over a string
str := "hello"
for index, char := range str {
fmt.Println(index, string(char))
}

// Iterating over a map
m := map[string]int{"a": 1, "b": 2}
for key, value := range m {
fmt.Println(key, value)
}
Looping Over Channels: Using the range keyword with a channel allows iterating over values sent to the channel until it is closed.
go
Copy code
ch := make(chan int)

// Producer goroutine
go func() {
defer close(ch)
for i := 0; i < 5; i++ {
ch <- i
}
}()

// Consumer goroutine
for value := range ch {
fmt.Println(value)
}
These are the primary looping mechanisms in Go, each suitable for different types of data structures and use cases.
