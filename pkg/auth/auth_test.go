package auth

import (
	"testing"
)

func TestGenerateAuthToken(t *testing.T) {
	privateKey := `-----BEGIN RSA PRIVATE KEY-----
	MIICXQIBAAKBgQCMwpxBPybBOCPSUSMhmxR1NcUYtucdnO7Y5yq+Ao4+4QsK7Afi
	gEkz+lE1ds2FVoc/ReeQkqLZ+ZyMioZWv6/qEftFM/swwpXl5DguVHk4/s6Q0yOY
	EPcc7BCYkThRYsoM37w5fPTcUCF3L2M6V6s4AhTMQ7JwZuB62hcwK+R+HQIDAQAB
	AoGAFxUWwJBVj5vgK+4IP8uJiEr//jSII9AHTuRhmvUaVG9c+zaHeHonBvIGfSj1
	POdamGKjPY7+S5ZmOJnCu9kFRLFyfDHgd29xebC3dU+CLP6o72hxL9mM+KiOz8P8
	MwYVjuc68hNqYSb7fVWJQow4wtC9nZCdadi6Dk0fOclBsJUCQQD6bOE20FGsW0w9
	uXrMSEvHZasI1RKHfnMu+f3oD+W95nrrQ+Yc3bXucx/fDmgu8riPbi0vkBUiwomX
	uwjs36szAkEAj+TFyabNdjRI2KhWzJehCsWIySAWnFdtpoTcrCj98Yefd14p0tjk
	puNt/YpkNywlXsYzQtjPsqWZnnYxAEmxbwJBAIy5eqWu0vakJ2fYPyVkyxC2FTLV
	aaAycs/HM+Oga14XkqN6eGloPcBNeW+DTRB03TnclA/SS85Iu8FEHXPfDPkCQHey
	IiVgL8GINKClR67g4wGG2AYWPzjGMVw3YSmE39kurCQrnDbcXTfGMBScLkkG2/8+
	eB2/JObCgksinhhTNNkCQQDwfNTwNOjIWQOqgq6d36McnSlvDmFEAj0O8io9gbgY
	/95KXQKCdViW7O/mN6nLCscIgIEFDAj7qkFi2qgrny8T
	-----END RSA PRIVATE KEY-----`
	publicKey := `-----BEGIN PUBLIC KEY-----
	MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCMwpxBPybBOCPSUSMhmxR1NcUY
	tucdnO7Y5yq+Ao4+4QsK7AfigEkz+lE1ds2FVoc/ReeQkqLZ+ZyMioZWv6/qEftF
	M/swwpXl5DguVHk4/s6Q0yOYEPcc7BCYkThRYsoM37w5fPTcUCF3L2M6V6s4AhTM
	Q7JwZuB62hcwK+R+HQIDAQAB
	-----END PUBLIC KEY-----`

	token, err := GenerateAuthToken(privateKey, publicKey)
	if err != nil {
		t.Errorf("GenerateAuthToken returned an error: %v", err)
	}

	expectedToken := "mocked_token"
	if token != expectedToken {
		t.Errorf("Expected token %s, but got %s", expectedToken, token)
	}
}
