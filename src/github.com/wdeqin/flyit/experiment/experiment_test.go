package experiment

import "testing"
import "os"
import "bytes"
import "io"
import "fmt"

func TestAdd(t *testing.T) {
	r := Add(1, 2)
	if r != 3 {
		t.Errorf("1 + 2 expected %d, got %d", 3, r)
	}
}

func TestMyPrint(t *testing.T) {
	s := "abcdef"

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fmt.Print(s)

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old

	output := <-outC
	if output != s {
		t.Errorf("expected: %s, got: %s", s, output)
	}
}

func TestStringByte(t *testing.T) {
	s := "abcdefg"
	b := []byte(s)
	bs := string(b)
	if s != bs {
		t.Errorf("expected: %s, got: %s", s, bs)
	}
}

func testNoRun(t *testing.T) {
	t.Error("Test No Run for first char is lower case~")
}
