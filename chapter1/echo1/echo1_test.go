package echo1

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func BenchmarkStringConcantenation(b *testing.B) {
	commands := generateCommands()

	for _, cmd := range commands {
		b.Run("Quadratic-Len-"+strconv.Itoa(len(cmd)), func(b *testing.B) {
			setupArgs(cmd)
			for i := 0; i < b.N; i++ {
				quadraticConcatenation()
			}
		})
	}

	for _, cmd := range commands {
		b.Run("Linear-Len-"+strconv.Itoa(len(cmd)), func(b *testing.B) {
			setupArgs(cmd)
			for i := 0; i < b.N; i++ {
				linearConcatenation()
			}
		})
	}
}

func BenchmarkStringConcantenationParallel(b *testing.B) {
	commands := generateCommands()

	for _, cmd := range commands {
		fmt.Println("")
		setupArgs(cmd)
		b.Run("Quadratic-Len-"+strconv.Itoa(len(cmd)), func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					quadraticConcatenation()
				}
			})
		})
	}

	for _, cmd := range commands {
		setupArgs(cmd)
		b.Run("Linear-Len-"+strconv.Itoa(len(cmd)), func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					linearConcatenation()
				}
			})
		})
	}
}

func setupArgs(cmd []string) {
	os.Args = cmd
}

func generateCommands() [][]string {
	length := []int{100, 1000, 1000_0, 1000_00}
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")

	var commands [][]string
	for _, l := range length {
		var cmd []string
		for range l {
			var word string
			b := make([]rune, 1)
			for j := range b {
				b[j] = chars[rand.Intn(len(chars))]
			}
			word += string(b)       // random word
			cmd = append(cmd, word) // build random sentence/command
		}
		commands = append(commands, cmd)
	}
	return commands
}
