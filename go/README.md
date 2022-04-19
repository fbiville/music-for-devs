# Sound experiments in Go

## Prereqs

 - [Go](https://go.dev/dl/)
 - [ffmpeg](https://ffmpeg.org/download.html)

## Sine City

Play a sine wave for a specific duration at a given frequency.

```go
go run ./cmd/01-simple -duration 5 -frequency 550
```

## In Your Phase

Play 2 sine waves offset by a specific phase in radians.

```go
go run ./cmd/02-phase -phase 1.57079632679
```

## The Shape of U(t)

Listen to the same melody with different note shapes!

```go
go run ./cmd/03-shape
```
