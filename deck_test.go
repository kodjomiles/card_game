package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	lengthOfDeck := len(d)
	actualExpectedLen := 56
	if lengthOfDeck != actualExpectedLen {
		t.Errorf("Expected %d, got %d", actualExpectedLen, lengthOfDeck)
	}

	firstDeckCardValue := "Ace of Diamond"
	lastDeckCardValue := "King of Spades"
	if d[0] != firstDeckCardValue {
		t.Errorf("Expected first element in deck to be %v but got %v", firstDeckCardValue, d[0])
	}

	if d[len(d)-1] != lastDeckCardValue {
		t.Errorf("Expected last element in deck to be %v but got %v", lastDeckCardValue, d[len(d)-1])
	}
}

func TestSaveToFile(t *testing.T) {
	d := newDeck()

	_deckSaveFileTesting := "_deskSaveFileTesting"

	err := d.saveToFile(_deckSaveFileTesting)

	if err != nil {
		t.Errorf("Something went wrong saving the file: %v", err)
	}

	checkFileExistenceAndRemove(t, d, _deckSaveFileTesting)

}

func checkFileExistenceAndRemove(t *testing.T, d Deck, filename string) {
	err := d.deleteFile(filename)
	if err != nil {
		t.Fatalf("Failed to delete file, %v", err)
	}
}
