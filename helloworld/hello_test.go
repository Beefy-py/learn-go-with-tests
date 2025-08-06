package helloworld

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	result := HelloWorld()

	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		result := Hello("Kenny", "")
		expected := "Hello, Kenny!"

		assertCorrectMessage(t, result, expected)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		result := Hello("", "")
		expected := "Hello, World!"

		assertCorrectMessage(t, result, expected)
	})

	t.Run("in Spanish", func(t *testing.T) {
		result := Hello("Elodie", "Spanish")
		expected := "Hola, Elodie"
		assertCorrectMessage(t, result, expected)

	})

	t.Run("in French", func(t *testing.T) {
		result := Hello("François", "French")
		expected := "Bonjour, François!"
		assertCorrectMessage(t, result, expected)
	})
}

func assertCorrectMessage(t testing.TB, result, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("Expected %q, but got %q", result, expected)
	}
}
