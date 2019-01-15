package models2

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

type WhatToEncrypt struct {
	PatientDetails bool
}

var _cachedCipher cipher.Block
var _cachedKeyId string

func getCipher() (cipher.Block, string, error) {
	if _cachedCipher != nil && _cachedKeyId != "" {
		return _cachedCipher, _cachedKeyId, nil
	}

	// to set in the fish shell
	// set -x GOFHIR_ENCRYPTION_KEY_BASE64  (dd if=/dev/random bs=32 count=1 | base64)
	// set -x GOFHIR_ENCRYPTION_KEY_ID testKey
	keyB64 := os.Getenv("GOFHIR_ENCRYPTION_KEY_BASE64")
	_cachedKeyId := os.Getenv("GOFHIR_ENCRYPTION_KEY_ID")
	if keyB64 == "" {
		return nil, "", errors.New("missing environment variable: GOFHIR_ENCRYPTION_KEY_BASE64")
	}
	if _cachedKeyId == "" {
		return nil, "", errors.New("missing environment variable: GOFHIR_ENCRYPTION_KEY_ID")
	}
	key, err := base64.StdEncoding.DecodeString(keyB64)
	if err != nil {
		return nil, "", errors.Wrap(err, "invalid environment variable: GOFHIR_ENCRYPTION_BASE64")
	}
	if len(key) != 32 {
		return nil, "", errors.Wrap(err, "environment variable should be 32 bytes: GOFHIR_ENCRYPTION_BASE64")
	}

	_cachedCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, "", errors.Wrap(err, "NewCipher failed")
	}

	return _cachedCipher, _cachedKeyId, nil
}

func shouldEncryptField(name string) bool {
	switch name {
	case "name", "birthDate", "telecom", "address", "photo", "contact", "communication", "text":
		return true
	default:
		return false
	}
}

func encryptBSON(bsonRoot *[]bson.DocElem, resourceType string, whatToEncrypt WhatToEncrypt) error {
	if whatToEncrypt.PatientDetails == false || resourceType != "Patient" {
		return nil
	}

	// will be encrypted
	plaintext := make([]bson.DocElem, 0, 4)

	// new document (with plaintext fields removed)
	newBsonRoot := make([]bson.DocElem, 0, len(*bsonRoot))

	for _, elem := range *bsonRoot {
		if shouldEncryptField(elem.Name) {
			plaintext = append(plaintext, elem)
		} else {
			newBsonRoot = append(newBsonRoot, elem)
		}
	}

	// convert plaintext to bson bytes
	plaintextBytes, err := bson.Marshal(plaintext)
	if err != nil {
		return errors.Wrap(err, "bson.Marshal of plaintext failed")
	}

	// encrypt -- based on https://github.com/gtank/cryptopasta/blob/master/encrypt.go
	block, keyId, err := getCipher()
	if err != nil {
		return errors.Wrap(err, "getCipher failed")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return errors.Wrap(err, "NewGCM failed")
	}

	nonce := make([]byte, gcm.NonceSize()) // random nonce
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return errors.Wrap(err, "failed to generate random nonce for GCM")
	}

	// output is random nonce | GCM ciphertext | GCM tag
	ciphertextBytes := gcm.Seal(nonce, nonce, plaintextBytes, nil)
	ciphertextB64 := base64.StdEncoding.EncodeToString(ciphertextBytes)

	newBsonRoot = append(newBsonRoot, bson.DocElem{Name: "__gofhirEncryptedBSON", Value: ciphertextB64})
	newBsonRoot = append(newBsonRoot, bson.DocElem{Name: "__gofhirEncryptionKeyId", Value: keyId})

	*bsonRoot = newBsonRoot
	return nil
}

func decryptBSON(bsonRoot *[]bson.DocElem) error {

	// new document
	newBsonRoot := make([]bson.DocElem, 0, len(*bsonRoot))
	var ciphertextB64 string
	var expectedKeyId string
	var ok bool

	for _, elem := range *bsonRoot {
		switch elem.Name {
		case "__gofhirEncryptedBSON":
			ciphertextB64, ok = elem.Value.(string)
			if !ok {
				return errors.New("__gofhirEncryptedBSON not a string")
			}

		case "__gofhirEncryptionKeyId":
			expectedKeyId, ok = elem.Value.(string)
			if !ok {
				return errors.New("__gofhirEncryptionKeyId not a string")
			}

		default:
			newBsonRoot = append(newBsonRoot, elem)
		}
	}

	if ciphertextB64 == "" {
		return nil // nothing to decrypt
	}

	// decrypt -- based on https://github.com/gtank/cryptopasta/blob/master/encrypt.go
	block, keyId, err := getCipher()
	if err != nil {
		return errors.Wrap(err, "getCipher failed")
	}
	if keyId != expectedKeyId {
		return errors.Wrapf(err, "encryption keyId (%s) doesn't match expected (%s)", keyId, expectedKeyId)
	}

	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextB64)
	if err != nil {
		return errors.Wrap(err, "failed to decode __gofhirEncryptedBSON")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return errors.Wrap(err, "NewGCM failed")
	}

	if len(ciphertext) < gcm.NonceSize() {
		return errors.New("failed to decode __gofhirEncryptedBSON: too short")
	}

	plaintextBytes, err := gcm.Open(nil, ciphertext[:gcm.NonceSize()], ciphertext[gcm.NonceSize():], nil)

	// convert plaintext to bson
	var plaintextDoc bson.D
	err = bson.Unmarshal(plaintextBytes, &plaintextDoc)
	if err != nil {
		return errors.Wrap(err, "bson.Unmarshal of plaintext failed")
	}

	for _, elem := range plaintextDoc {
		newBsonRoot = append(newBsonRoot, elem)
	}

	*bsonRoot = newBsonRoot
	return nil
}
