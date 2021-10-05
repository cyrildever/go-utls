package crypto

import (
	"encoding/base64"

	"golang.org/x/crypto/ssh"
)

// SSHPublicKey2String returns the stringified version of the passed SSH public key
func SSHPublicKey2String(pk ssh.PublicKey) string {
	return base64.StdEncoding.EncodeToString(pk.Marshal())
}
