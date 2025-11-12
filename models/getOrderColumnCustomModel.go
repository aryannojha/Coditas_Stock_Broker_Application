package models

type BFFOrderColumnCustomGetRequest struct {
	UserId           *uint64 `json:"userId" validate:"required" example:"1"`
	AdvanceOrderType string  `json:"advanceOrderType" validate:"required,oneof=OrderBook TradeBook" example:"OrderBook/TradeBook"`
}

type BFFOrderColumnCustomGetResponse struct {
	ColumnId           uint64 `json:"columnId"`
	ResponseColumnName string `json:"responseColumnName"`
	DisplayColumnName  string `json:"displayColumnName"`
}
