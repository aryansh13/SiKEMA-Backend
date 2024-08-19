package saveQRCode

type InputSaveQRCode struct {
	EventID  string
	SSID     string `json:"ssid"`
	Password string `json:"password"`
}
