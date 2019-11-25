package writer

import (
	"ejercicios-de-prueba/src/writer"
	"testing"
)

func TestHelloToNothing(t *testing.T) {
	name := ""
	expected := "Hello Go"
	result := writer.Hello(name)

	if expected != result {
		t.Error("The result wasn`t expect")
	}
}
