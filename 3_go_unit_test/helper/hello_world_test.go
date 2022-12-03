package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Benchmark Test
// go test -v -bench=.
// go test -v -run=NotMatchUnitTest -bench=.
// go test -v -run=NotMatchUnitTest -bench=. ./...
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Zulfikar")
	}
}

// go test -v -bench=BenchmarkHelloWorldSub/Zulfikar
func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Zulfikar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Zulfikar")
		}
	})

	b.Run("Isnaen", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Isnaen")
		}
	})
}

func BenchmarkHelloWorldTable(b *testing.B) {
	bechmarks := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Zulfikar)",
			request:  "Zulfikar",
			expected: "Hello Zulfikar",
		},
		{
			name:     "HelloWorld(Isnaen)",
			request:  "Isnaen",
			expected: "Hello Isnaen",
		},
	}

	for _, benchmark := range bechmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.name)
			}
		})
	}
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Zulfikar)",
			request:  "Zulfikar",
			expected: "Hello Zulfikar",
		},
		{
			name:     "HelloWorld(Isnaen)",
			request:  "Isnaen",
			expected: "Hello Isnaen",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

// hanya jalan di 1 package, format harus TestMain
func TestMain(m *testing.M) {
	fmt.Println("Sebelum testing")

	m.Run()

	fmt.Println("Sesudah testing")
}

// go test -v -run TestSubTest
// go test -v -run TestSubTest/Zulfikar
// go test -v -run /Zulfikar
// go test -v -run /Isnaen
func TestSubTest(t *testing.T) {
	t.Run("Zulfikar", func(t *testing.T) {
		result := HelloWorld("Zulfikar")
		assert.Equal(t, "Hello Zulfikar", result)
	})

	t.Run("Isnaen", func(t *testing.T) {
		result := HelloWorld("Isnaen")
		assert.Equal(t, "Hello Isnaen", result)
	})
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can't tun on MacOS")
	}

	result := HelloWorld("Zulfikar")
	assert.Equal(t, "Hello Zulfikar", result)
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorldFail("Zulfikar")
	assert.Equal(t, "Hello Zulfikar", result)
	fmt.Println("Test Done Assert!!!!")
}

func TestHelloWorldRequired(t *testing.T) {
	result := HelloWorldFail("Zulfikar")
	require.Equal(t, "Hello Zulfikar", result)
	fmt.Println("Test Done Reuire!!!!")
}

func TestHelloWorldZul(t *testing.T) {
	result := HelloWorld("Zulfikar")

	if result != "Hello Zulfikar" {
		panic("Result bukan Zulfikar")
	}
}

func TestHelloWorldArip(t *testing.T) {
	result := HelloWorld("Arip")

	if result != "Hello Arip" {
		panic("Result bukan Arip")
	}
}

func TestHelloWorldNopal(t *testing.T) {
	result := HelloWorldFail("Nopal")

	if result != "Hello Nopal" {
		t.Fail()
	}

	fmt.Println("Test Nopal Done")
}

func TestHelloWorldAkib(t *testing.T) {
	result := HelloWorldFail("Akib")

	if result != "Hello Akib" {
		t.FailNow()
	}

	fmt.Println("Test Nopal Done")
}

func TestHelloWorldAndi(t *testing.T) {
	result := HelloWorldFail("Andi")

	if result != "Hello Andi" {
		t.Error("Result bukan 'Hello Andi'")
	}

	fmt.Println("Test Andi Done")
}

func TestHelloWorldJuki(t *testing.T) {
	result := HelloWorldFail("Juki")

	if result != "Hello Juki" {
		t.Fatal("Result bukan 'Hello Juki'")
	}

	fmt.Println("Test Juki Done")
}
