package gocontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

type key string

const (
	a key = "a"
	b key = "b"
	c key = "c"
	d key = "d"
	e key = "e"
	f key = "f"
)

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, a, "B")
	contextC := context.WithValue(contextA, c, "C")

	contextD := context.WithValue(contextB, d, "D")
	contextE := context.WithValue(contextB, e, "E")

	contextF := context.WithValue(contextC, f, "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value("b")) // key tidak ada di dia sendiri dan parent2 nya
	fmt.Println(contextF.Value("c")) // bisa ambil parent
	fmt.Println(contextF.Value("f")) // bisa ambil dia sendiri
	fmt.Println(contextA.Value("c")) // ga bisa ambil punya child nya
}

// ##########################################################################CANCEL##########################################################################################

// Goroutine Leak = goroutine yang masih menyala, padahal sudah tidak dibutuhkan lagi
func CreateCounter() chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			destination <- counter
			counter++
		}
	}()
	return destination
}

func TestGoroutineLeak(t *testing.T) {
	fmt.Println("Total goroutine :", runtime.NumGoroutine())

	destination := CreateCounter()
	for n := range destination {
		fmt.Println("Counter ", n)
		if n == 10 {
			break
		}
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total goroutine :", runtime.NumGoroutine())
}

func CreateCounterSolution(ctx context.Context) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // Simulation slow
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total goroutine :", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounterSolution(ctx)
	for n := range destination {
		fmt.Println("Counter ", n)
		if n == 10 {
			break
		}
	}

	cancel()

	time.Sleep(2 * time.Second)

	fmt.Println("Total goroutine :", runtime.NumGoroutine())
}

// ##############################################################################TIMEOUT######################################################################################

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total goroutine :", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounterSolution(ctx)
	for n := range destination {
		fmt.Println("Counter ", n)
	}

	fmt.Println("Total goroutine :", runtime.NumGoroutine())
}

// ##############################################################################DEADLINE######################################################################################

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total goroutine :", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounterSolution(ctx)
	for n := range destination {
		fmt.Println("Counter ", n)
	}

	fmt.Println("Total goroutine :", runtime.NumGoroutine())
}
