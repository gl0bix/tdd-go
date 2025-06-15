package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello World' when an empty string is supplied", func(*testing.T) {
		got := Hello("", "")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in spanish", func(*testing.T) {
		got := Hello("Javier", "spanish")
		want := "Hola Javier"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
