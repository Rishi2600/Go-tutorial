package api

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	code    int
	balance int64
}

type Error struct {
	code    int
	message string
}
