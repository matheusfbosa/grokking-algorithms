package chapter2_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/matheusfbosa/grokking-algorithms/chapter2"
)

func TestNewLinkedList(t *testing.T) {
	ll := chapter2.NewLinkedList()

	if ll.Head != nil {
		t.Error("Expected head to be nil for a new linked list")
	}
}

func TestLinkedList_Append(t *testing.T) {
	t.Run("AppendEmptyList", func(t *testing.T) {
		ll := chapter2.NewLinkedList()

		ll.Append(1)

		if ll.Head == nil || ll.Head.Data != 1 || ll.Head.Next != nil {
			t.Error("Append failed for an empty list")
		}
	})

	t.Run("AppendNonEmptyList", func(t *testing.T) {
		ll := chapter2.NewLinkedList()

		ll.Append(1)
		ll.Append(2)

		if ll.Head == nil || ll.Head.Data != 1 || ll.Head.Next == nil || ll.Head.Next.Data != 2 {
			t.Error("Append failed for a non-empty list")
		}
	})
}

func TestLinkedList_Display(t *testing.T) {
	t.Run("DisplayEmptyList", func(t *testing.T) {
		ll := chapter2.NewLinkedList()
		expectedOutput := "nil\n"

		output := captureOutput(func() {
			ll.Display()
		})

		if output != expectedOutput {
			t.Errorf("Expected output: %s, Got: %s", expectedOutput, output)
		}
	})

	t.Run("DisplayWithElements", func(t *testing.T) {
		ll := chapter2.NewLinkedList()
		ll.Append(1)
		ll.Append(2)
		expectedOutput := "1 -> 2 -> nil\n"

		output := captureOutput(func() {
			ll.Display()
		})

		if output != expectedOutput {
			t.Errorf("Expected output: %s, Got: %s", expectedOutput, output)
		}
	})
}

// captureOutput is a helper function to capture stdout during testing.
func captureOutput(f func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	os.Stdout = rescueStdout
	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
