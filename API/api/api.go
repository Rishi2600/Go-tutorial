package api

// import (
// 	"encoding/json"
// 	"net/http"
// )

/*few structs to represent parameters for endpoints and responses*/

//coin balance params
type CoinBalanceParams struct {
	Username string
}

//coin balance response
type CoinBalanceResponse struct {
	//status code
	Code int

	//account balance
	Balance int64
}

//error response
type Error struct {
	//status code
	Code int

	//error message
	Message string
}
