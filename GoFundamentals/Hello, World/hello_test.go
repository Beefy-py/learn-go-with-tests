package main

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	result := HelloWorld()

	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}

func TestHello(t *testing.T){
	expected := "Hello, Kenny!";
	result := Hello("Kenny");

	if result != expected {
		t.Errorf("Expected %q, but got %q", expected, result)	;
	}
}