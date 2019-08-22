package main

import (
	"log"
	"crypto/aes"
    "crypto/cipher"
    "crypto/rand"
	"io/ioutil"
    "io"
    "path/filepath"
    "encoding/base64"
    "strings"
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

    encryptedContent := gcm.Seal(nonce, nonce, content, nil)
    base64 := base64.StdEncoding.EncodeToString(encryptedContent)
    filebase := filepath.Base(filename)
    ioutil.WriteFile(filebase+".enc", []byte(base64), 0644)
	return 0
}

func Decryptfile(filename string, password string) int {
    if strings.Contains(filename, ".enc") == false {
        log.Fatalln("we need an encrypted file with .enc extension")
        return -1
    }
    
    base64Content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("error reading file")
    }
    content, err := base64.StdEncoding.DecodeString(string(base64Content))
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
    outputfile := strings.Replace(filename, ".enc", "",1)
    ioutil.WriteFile(outputfile, []byte(plaintext), 0644)
    return 0
}