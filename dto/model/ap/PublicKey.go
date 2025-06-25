package ap

type PublicKey struct {
	Id           string `json:"id" validate:"required"`
	Owner        string `json:"owner" validate:"required"`
	PublicKeyPem string `json:"publicKeyPem" validate:"required"`
}
