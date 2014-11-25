// XOR encryption key manager for zcryp.
// Copyright 2014 John Ehringer <jhe@5khz.com>.
// Provided under the terms of the MIT license in the included LICENSE file.

package main

// KeyState maintains the state of the XOR key.
type KeyState struct {
	key string
	keylen, keyidx int
}

// Create a new KeyState based on an input key of arbitrary length.
func NewKeyState(key string) *KeyState {
	keystate := new(KeyState)
	keystate.key = key
	keystate.keylen = len(keystate.key)
	keystate.keyidx = 0

	return keystate
}

// Fetch the next byte of the KeyState and increment the internal cursor.
func (keystate *KeyState) NextByte() byte {
	nextbyte := keystate.key[keystate.keyidx]

	keystate.keyidx++

	if keystate.keyidx >= keystate.keylen {
		keystate.keyidx = 0
	}

	return nextbyte
}
