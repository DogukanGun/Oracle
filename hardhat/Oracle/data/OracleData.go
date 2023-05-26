package data

type SetLendingAddress struct {
	LendingPoolAddress string `json:"lendingPoolAddress"`
	OracleAddress      string `json:"oracleAddress"`
}

type GetLendingAddress struct {
	OracleAddress string `json:"oracleAddress"`
}

type SetInterestRate struct {
	OracleAddress string `json:"oracleAddress"`
	InterestRate  string `json:"interestRate"`
}

type IncreaseTotalBarrowedLimit struct {
	OracleAddress string `json:"oracleAddress"`
	Amount        string `json:"amount"`
}

type DecreaseTotalBarrowedLimit struct {
	OracleAddress string `json:"oracleAddress"`
	Amount        string `json:"amount"`
}
