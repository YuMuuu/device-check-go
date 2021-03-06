package devicecheck

import (
	"io/ioutil"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestNewCredentialFile(t *testing.T) {
	filename := "revoked_private_key.p8"

	cred := NewCredentialFile(filename).(credentialFile)

	assert.Equal(t, filename, cred.filename)
}

func TestCredentialFile_key(t *testing.T) {
	cred := NewCredentialFile("revoked_private_key.p8")

	key, err := cred.key()

	assert := assert.New(t)
	assert.NotNil(key)
	assert.Nil(err)
}

func TestCredentialFile_key_UnknownFile(t *testing.T) {
	cred := NewCredentialFile("unknown_file.p8")

	key, err := cred.key()

	assert := assert.New(t)
	assert.Nil(key)
	assert.NotNil(err)
}

func TestNewCredentialBytes(t *testing.T) {
	raw, err := ioutil.ReadFile("revoked_private_key.p8")
	assert.Nil(t, err)

	cred := NewCredentialBytes(raw).(credentialBytes)

	assert.Equal(t, raw, cred.raw)
}

func TestCredentialBytes_key(t *testing.T) {
	raw, err := ioutil.ReadFile("revoked_private_key.p8")
	assert.Nil(t, err)
	cred := NewCredentialBytes(raw)

	key, err := cred.key()

	assert := assert.New(t)
	assert.NotNil(key)
	assert.Nil(err)
}

func TestNewCredentialString(t *testing.T) {
	raw, err := ioutil.ReadFile("revoked_private_key.p8")
	assert.Nil(t, err)
	str := *(*string)(unsafe.Pointer(&raw))

	cred := NewCredentialString(str).(credentialString)

	assert.Equal(t, str, cred.str)
}

func TestCredentialString_key(t *testing.T) {
	raw, err := ioutil.ReadFile("revoked_private_key.p8")
	assert.Nil(t, err)
	str := *(*string)(unsafe.Pointer(&raw))
	cred := NewCredentialString(str)

	key, err := cred.key()

	assert := assert.New(t)
	assert.NotNil(key)
	assert.Nil(err)
}
