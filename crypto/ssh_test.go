package crypto_test

import (
	"io/ioutil"
	"testing"

	"github.com/cyrildever/go-utls/crypto"
	"golang.org/x/crypto/ssh"
	"gotest.tools/assert"
)

// TestSSH ...
//
// NB: Instruction used to create PEM key:
//	ssh-keygen -t rsa -b 2048 -C test
func TestSSH(t *testing.T) {
	pemBytes, err := ioutil.ReadFile("test/ssh_rsa.key")
	if err != nil {
		t.Fatal(err)
	}
	signer, err := ssh.ParsePrivateKey(pemBytes)
	assert.NilError(t, err)
	pk := signer.PublicKey()
	pkStr := crypto.SSHPublicKey2String(pk)
	assert.Equal(t, pkStr, "AAAAB3NzaC1yc2EAAAADAQABAAABAQDF6Hb06N9dse1rdPruUb2F9i4escDb0FbqXhYgKt+lmLJKfI6rM5CWmoHmYaUs6ysZcmyE3LbVYv14kcho3bKPJJwdvprgrAbq3NXNNlwBeawPtyOfRC59ubH3Uxq7V1qPUivFQjY5FzZSainRTxJI4/faM494BiXTr9AnoBZIin69Fo7Jd3YbhKPeZOJDlp73uTwN8WG6AgVF9Nzvcx6boIhJiY8vLbSbNnDNzN+CmO1yW98FzK+whaR5KcNSka99MBtj80zx1QJJ+QCIr36KUrtQ3uwGsT20JIJ5ozqgLPuvg7CpcUEuGLvn862Bmcxv56R67/XVZ6ytfZLFv3N3")
}
