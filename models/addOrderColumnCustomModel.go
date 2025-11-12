package models

type BFFOrderColumnAddRequest struct {
	UserId           uint64 `json:"userId" validate:"required" example:"12"`
	ColumnId         []uint64  `json:"columnId" validate:"required"`
	AdvanceOrderType string `json:"advanceOrderType" validate:"required,oneof=OrderBook TradeBook" example:"OrderBook/TradeBook"`
}

type BFFOrderColumnAddResponse struct {
	Message string `json:"message" example:"success"`
}
