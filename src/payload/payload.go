package payload

type Payload struct {
	SAMLResponse string `json:"saml_response"`
	PrivateKey   string `json:"private_key"`
}
