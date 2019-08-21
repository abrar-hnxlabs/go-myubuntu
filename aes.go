package main

import (
	"log"
	"crypto/aes"
    "crypto/cipher"
    "crypto/rand"
	"io/ioutil"
    "io"
    "path/filepath"
)

func EncryptFile(filename string, password string) int {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("error reading file")
	}
    var key []byte
    key = []byte(password)

	c, err := aes.NewCipher(key)
    if err != nil {
        log.Fatalln(err)
    }

    gcm, err := cipher.NewGCM(c)
    if err != nil {
        log.Fatalln(err)
    }

    // creates a new byte array the size of the nonce
    // which must be passed to Seal
    nonce := make([]byte, gcm.NonceSize())
    // populates our nonce with a cryptographically secure
    // random sequence
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        log.Fatalln(err)
    }

    // here we encrypt our text using the Seal function
    // Seal encrypts and authenticates plaintext, authenticates the
    // additional data and appends the result to dst, returning the updated
    // slice. The nonce must be NonceSize() bytes long and unique for all
    // time, for a given key.
    encryptedContent := gcm.Seal(nonce, nonce, content, nil)
    filebase := filepath.Base(filename)
    ioutil.WriteFile(filebase+".enc", encryptedContent, 0644)
	return 0
}

func Decryptfile(filename string, outputfile string, password string) int {
    content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("error reading file")
	}
    c, err := aes.NewCipher([]byte(password))
    if err != nil {
        log.Fatalln(err)
    }

    gcm, err := cipher.NewGCM(c)
    if err != nil {
        log.Fatalln(err)
    }

    nonceSize := gcm.NonceSize()
    if len(content) < nonceSize {
        log.Fatalln(err)
    }

    nonce, ciphertext := content[:nonceSize], content[nonceSize:]
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        log.Fatalln(err)
    }
    ioutil.WriteFile(outputfile, []byte(plaintext), 0644)
    return 0
}