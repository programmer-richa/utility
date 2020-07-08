package testing

import (
	"fmt"
	"github.com/programmer-richa/utility/encrypt"
	"testing"
)

// TestEncrypt runs several test cases to check the correctness of
//the Encrypt function defined in encrypt package.
func TestEncrypt(t *testing.T) {
	tests := []struct {
		name string
		passphrase   string
		data  string
	}{
		{
			name: "Testing with key 1",
			passphrase :  "123456789012345678901234",
			data: "This is a test message.",
		},
		{
			name: "Testing with key 2",
			passphrase:  "9988998899880099887788",
			data: "This is a test message.",
		},
		{
			name: "Testing with key 2",
			passphrase:  "God is One",
			data: "This is a test message.",
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			err1,encryptedText := encrypt.Encrypt(c.data, c.passphrase)
			err2,decryptedText :=encrypt.Decrypt(encryptedText, c.passphrase)
			if  err1 !=nil || err2!=nil || decryptedText!= c.data {
				t.Fatal("Encrypt Validator Failed", c.name)
			} else {
				fmt.Println("Encrypt Validator-", c.name, "Pass")
			}
		})
	}
}
