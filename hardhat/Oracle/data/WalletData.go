package data

type CreateUserAddress struct {
	UserHash string `json:"userHash"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type AddNewAsset struct {
	UserAddress string `json:"userAddress"`
	Sign        string `json:"sign"`
}

type Deposit struct {
	UserAddress string `json:"userAddress"`
	Sign        string `json:"sign"`
	Amount      string `json:"amount"`
}

type Withdraw struct {
	UserAddress string `json:"userAddress"`
	Sign        string `json:"sign"`
	Amount      string `json:"amount"`
}
