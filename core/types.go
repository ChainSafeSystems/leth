package core

type Config struct {
	Networks map[string]Network 	`json:"networks"`
}

type Network struct {
	Url string						`json:"url,omitempty"`
	From string						`json:"from,omitempty"`
	Keystore string					`json:"keystore,omitempty"`
	Password string					`json:"password,omitempty"`
	Gas int64 						`json:"gas,omitempty"`
	GasPrice int64  				`json:"gasPrice,omitempty"`
	Id string						`json:"id,omitempty"`
}

type Transaction struct {
	From string 					`json:"from"`	
	To string 						`json:"to,omitempty"`
	Gas string 						`json:"gas,omitempty"`
	GasPrice string 				`json:"gasPrice,omitempty"`
	Value string 					`json:"value,omitempty"`
	Data string 					`json:"data,omitempty"`
}