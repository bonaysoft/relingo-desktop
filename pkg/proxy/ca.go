package proxy

import _ "embed"

//go:embed ca.pem
var ca []byte

//go:embed ca.key
var caKey []byte

func GetCert() []byte {
	return ca
}
