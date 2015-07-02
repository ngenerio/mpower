package mpowergo

const (
	BASE_URL_LIVE string = "https://app.mpowerpayments.com/api/v1"
	BASE_URL_TEST string = "https://app.mpowerpayments.com/sandbox-api/v1"
)

type Setup struct {
	MasterKey   string
	PrivateKey  string
	PublicKey   string
	Token       string
	ContentType string
	BASE_URL    string
}

func (setup *Setup) Get(fieldName string) string {
	return get(setup, fieldName)
}

func (setup *Setup) GetHeaders() map[string]string {
	headers := make(map[string]string)

	headers["MP-Master-Key"] = setup.MasterKey
	headers["MP-Private-Key"] = setup.PrivateKey
	headers["MP-Public-Key"] = setup.PublicKey
	headers["MP-Token"] = setup.Token
	headers["Content-Type"] = setup.ContentType

	return headers
}

func NewSetup(setupInfo map[string]string) *Setup {
	setup := &Setup{
		MasterKey:   envOr("MP-Master-Key", setupInfo["MP-Master-Key"]),
		PrivateKey:  envOr("MP-Private-Key", setupInfo["MP-PrivateKey"]),
		PublicKey:   envOr("MP-Public-Key", setupInfo["MP-PublicKey"]),
		Token:       envOr("MP-Token", setupInfo["MP-Token"]),
		ContentType: "application/json",
	}

	if val, ok := setupInfo["mode"]; ok && val == "live" {
		setup.BASE_URL = BASE_URL_LIVE
	} else {
		setup.BASE_URL = BASE_URL_TEST
	}

	return setup
}
