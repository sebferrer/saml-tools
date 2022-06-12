package main

import (
	// "log"
	httpserver "samltools/src/httpserver/gin"
	// "samltools/src/saml"
	// "samltools/src/util"
)

func main() {
	/*samlResponse := util.Read("../../data/saml_response.xml")
	key := saml.MustParsePrivateKey([]byte(util.Read("../../data/key.pem")))

	decryptedAssertion, err := saml.GetAndDecryptAssertion(samlResponse, key)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(decryptedAssertion)*/

	httpserver.Serve()
}
