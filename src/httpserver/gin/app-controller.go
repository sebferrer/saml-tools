package httpserver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	payload "samltools/src/payload"
	saml "samltools/src/saml"
)

func Decrypt(c *gin.Context) {
	var payload payload.Payload

	reqBody, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(reqBody, &payload)

	key := saml.MustParsePrivateKey([]byte(payload.PrivateKey))

	decryptedAssertion, err := saml.GetAndDecryptAssertion(payload.SAMLResponse, key)
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusOK, decryptedAssertion)
}
