// XOR encryption key manager for zcryp.
// Copyright 2014 John Ehringer <jhe@5khz.com>.
// Provided under the terms of the MIT license in the included LICENSE file.

package main

type KeyState struct {
	key string
	keylen, keyidx int
}

func NewKeyState(key string) *KeyState {
	keystate := new(KeyState)
	keystate.key = key
	keystate.keylen = len(keystate.key)
	keystate.keyidx = 0

	return keystate
}

func (keystate *KeyState) NextByte() byte {
	nextbyte := keystate.key[keystate.keyidx]

	keystate.keyidx++

	if keystate.keyidx >= keystate.keylen {
		keystate.keyidx = 0
	}

	return nextbyte
}
