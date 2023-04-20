package fastpass

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"io"
	"os"

	"io/ioutil"

	"github.com/pkg/errors"
)

//FastPass is a fastpass service
type FastPass struct {
	Key     [32]byte
	Entries Entries
	fi      *os.File
}

//New returns an instantiated fastpass
func New() *FastPass {
	return &FastPass{
		Entries: make(Entries, 0, 100),
	}
}

//Open opens a database
func (fp *FastPass) Open(path string) (err error) {
	fp.fi, err = os.OpenFile(path, os.O_RDWR, 0600)
	if err != nil {
		return errors.Wrap(err, "failed to open fi")
	}
	dat, err := ioutil.ReadAll(fp.fi)
	if err != nil {
		return errors.Wrap(err, "failed to read fi")
	}
	plaintext, err := fp.decrypt(dat)
	if err != nil {
		return errors.Wrap(err, "failed to decrypt fi")
	}
	err = json.Unmarshal(plaintext, &fp.Entries)
	return errors.Wrap(err, "failed to unmarshal")
}

//Create a new fastpass database.
//it sets fp to the new database.
func (fp *FastPass) Create(path string) (err error) {
	fp.fi, err = os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
	if err != nil {
		return errors.Wrap(err, "failed to open database")
	}

	var byt []byte
	if byt, err = json.Marshal(fp.Entries); err != nil {
		panic(err)
	}
	_, err = fp.fi.Write(fp.encrypt(byt))
	return errors.Wrap(err, "failed to write")
}

//key must be 32 bytes
func (fp *FastPass) encrypt(plaintext []byte) (ciphertext []byte) {
	block, err := aes.NewCipher(fp.Key[:])
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	ciphertext = gcm.Seal(ciphertext, nonce, plaintext, nil)
	return append(nonce, ciphertext...)
}

//key must be 32 bytes
func (fp *FastPass) decrypt(ciphertext []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(fp.Key[:])
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonce := ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]

	plaintext, err = gcm.Open(plaintext, nonce, ciphertext, nil)

	return plaintext, errors.Wrap(err, "failed to decrypt")
}

//Close flushes and invalidates fp
func (fp *FastPass) Close() error {
	fp.Flush()
	err := fp.fi.Close()
	return errors.Wrap(err, "failed to write")
}

//Flush flushes the database to disk
func (fp *FastPass) Flush() error {
	byt, err := json.Marshal(fp.Entries)
	if err != nil {
		panic(err)
	}
	fp.fi.Truncate(0)
	fp.fi.Seek(0, 0)
	_, err = fp.fi.Write(fp.encrypt(byt))
	return err
}
