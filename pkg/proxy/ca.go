package proxy

import _ "embed"

//go:embed ca.pem
var ca []byte

func GetCert() []byte {
	return ca
}
