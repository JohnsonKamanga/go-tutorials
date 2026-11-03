# Go Tutorials

A collection of small, standalone Go examples and programming exercises, organized by topic. Each `.go` file with a `main` function is intended to be run independently; this folder is not a single application.

## Requirements

- Go 1.16 or later

There is no Go module or project-wide build configuration in this folder. Run an example by passing its source file to `go run` from the repository root.

## Examples

### Hello world

- [`hello-world/helloworld.go`](hello-world/helloworld.go) — prints a greeting.

### I/O and command-line arguments

- [`io/basic/echo-basic.go`](io/basic/echo-basic.go) — prints command-line arguments on one line.
- [`io/line-by-line/line-by-line-echo.go`](io/line-by-line/line-by-line-echo.go) — prints each argument on its own line.

### Networking

- [`networking/fetching/fetch.go`](networking/fetching/fetch.go) — fetches and prints the response body for each URL argument.
- [`networking/concurrent/fetchall.go`](networking/concurrent/fetchall.go) — fetches multiple URLs concurrently and reports elapsed times and response sizes.

These examples require network access and URL arguments.

### Animation

- [`animations/basic-gif/lissajous.go`](animations/basic-gif/lissajous.go) — generates an animated Lissajous GIF on standard output.

### LeetCode exercises

- [`leetcode/easy/twosum/main.go`](leetcode/easy/twosum/main.go) — implements Two Sum with sample cases in `main`.
- [`leetcode/easy/palindrome-number/main.go`](leetcode/easy/palindrome-number/main.go) — checks whether an integer is a palindrome, with sample cases in `main`.

### System programming exercises

- [`system-programming-exercises/chapter-1/question-2/min-max.go`](system-programming-exercises/chapter-1/question-2/min-max.go) — reads space-separated integers from standard input and prints the minimum and maximum.
- [`system-programming-exercises/chapter-1/question-3/min-max.go`](system-programming-exercises/chapter-1/question-3/min-max.go) — finds the minimum and maximum of integer command-line arguments.
- `system-programming-exercises/chapter-1/question-7/random-number-generator.go` — unfinished stub; it is not currently runnable.

## Running examples

Run these commands from the repository root:

```sh
go run ./hello-world/helloworld.go
go run ./io/basic/echo-basic.go hello go
go run ./io/line-by-line/line-by-line-echo.go hello go
go run ./networking/fetching/fetch.go https://example.com
go run ./networking/concurrent/fetchall.go https://example.com https://go.dev
go run ./animations/basic-gif/lissajous.go > lissajous.gif
go run ./leetcode/easy/twosum/main.go
go run ./leetcode/easy/palindrome-number/main.go
printf '4 1 9\n' | go run ./system-programming-exercises/chapter-1/question-2/min-max.go
go run ./system-programming-exercises/chapter-1/question-3/min-max.go 4 1 9
```

The minimum/maximum examples print `1` and `9` for the sample inputs. The GIF command writes the animation to `lissajous.gif` in the current directory. These are learning examples; there is no project-wide test suite.
