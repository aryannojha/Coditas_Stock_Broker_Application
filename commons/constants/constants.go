package constants

// search NEST API URL Keys
const (
	ServiceName      = "search"
	PortDefaultValue = 9099
)

// database columns constants
const (
	ScripId            = "id"
	SpreadOrderType    = "SP"
	InstrumentType     = "instrument_type"
	ContractType       = "contract_type"
	ContractId         = "contract_id"
	Exchange           = "exchange"
	ExchangeSegment    = "exchange_segment"
	SymbolName         = "symbol_name"
	ExpiryDate         = "expiry_date"
	OptionType         = "option_type"
	Group              = "group"
	ScripName          = "description"
	TradingSymbol      = "trading_symbol"
	ScripToken         = "scrip_token"
	ISINValue          = "isin"
	Multiplier         = "multiplier"
	DecimalPrecision   = "decimal_precision"
	TickSize           = "tick_size"
	LotSize            = "lot_size"
	UniqueKey          = "unique_key"
	StrikePrice        = "strike_price"
	CombinedScripToken = "combined_scrip_token"
	SegmentIndicator   = "segment_indicator"
	DisplayStrikePrice = "display_strike_price"
	Description        = "description"
	DisplayExpiryDate  = "display_expiry_date"

	UserId   = "user_id"
	ColumnId = "column_id"
)

// contract type constants
const (
	SpreadFuture = "SP-FUTURE"
	Future       = "FUTURE"
	Option       = "OPTION"
)

// query constants
const (
	SearchTextQuery                  = "(\"group\" = ? AND (description ILIKE ? OR trading_symbol ILIKE ?))"
	GetDistinctGroupsByExchangeQuery = `SELECT DISTINCT "group" FROM "scrip_master" WHERE "exchange" = $1`
)

// DB queries
const (
	GetDerivativesInstrumentSelectQuery       = `SELECT DISTINCT "instrument_type","contract_type","contract_id" FROM "scrip_master" WHERE exchange = $1 AND contract_type = $2`
	GetDerivativesScripInformationSelectQuery = `SELECT "id","scrip_token","trading_symbol","lot_size","expiry_date","symbol_name","tick_size","multiplier","decimal_precision","strike_price","combined_scrip_token","unique_key","exchange","instrument_type","option_type","exchange_segment" FROM "scrip_master" WHERE exchange = $1 AND instrument_type = $2 AND symbol_name = $3 AND expiry_date = $4`
	GetDerivativesScripSelectQuery            = `SELECT DISTINCT "symbol_name" FROM "scrip_master" WHERE exchange = $1 AND instrument_type = $2`
	GetDerivativesOptionTypesSelectQuery      = `SELECT DISTINCT "option_type" FROM "scrip_master" WHERE exchange = $1 AND instrument_type = $2 AND symbol_name = $3 AND expiry_date = $4`
	GetDerivativesStrikePriceSelectQuery      = `SELECT DISTINCT "strike_price","display_strike_price" FROM "scrip_master" WHERE exchange = $1 AND instrument_type = $2 AND symbol_name = $3 AND expiry_date = $4 AND option_type = $5`
	SearchEquityScripSelectQuery              = `SELECT "id","exchange","description","trading_symbol","exchange_segment","scrip_token","isin","tick_size","lot_size","unique_key","multiplier","decimal_precision" FROM "scrip_master" WHERE ("group" = $1 AND (description ILIKE $2 OR trading_symbol ILIKE $3))`
	GetDerivativesExpiryDateSelectQuery       = `SELECT DISTINCT "expiry_date","display_expiry_date" FROM "scrip_master" WHERE exchange = $1 AND instrument_type = $2 AND symbol_name = $3 ORDER BY expiry_date ASC`
	GetCircuitLimitsSelectQuery               = `SELECT "multiplier","decimal_precision","expiry_date","combined_scrip_token" FROM "scrip_master" WHERE "exchange_segment" = $1 AND "scrip_token" = $2`
)

const (
	GetDerivativesStrikePriceQueryCondition      = "exchange = ? AND instrument_type = ? AND symbol_name = ? AND expiry_date = ? AND option_type = ?"
	GetDerivativesScripInformationQueryCondition = "exchange = ? AND instrument_type = ? AND symbol_name = ? AND expiry_date = ?"
	GetDerivativesExpiryDateQueryCondition       = "exchange = ? AND instrument_type = ? AND symbol_name = ?"
	GetDerivativesInstrumentQueryCondition       = "exchange = ? AND contract_type = ?"
	GetDerivativesScripQueryCondition            = "exchange = ? AND instrument_type = ?"
	GetDeivativeOptionTypesQueryCondition        = "exchange = ? AND instrument_type = ? AND symbol_name = ? AND expiry_date = ?"
)

const (
	DeleteForOrderBookQuery   = `user_id = ? AND column_id IN (?)`
	DeleteForTradeBookQuery   = `user_id = ? AND column_id IN (?)`
	DeleteRowIfBothZeroQuery  = `user_id = ? AND column_id IN (?) AND order_book_sequence = 0 AND trade_sequence = 0`
	MaximumSequenceQuery      = `COALESCE(MAX(order_book_sequence), 0)`
	UserIdValidation          = `user_id = ?`
	OrderBookSequenceQuery    = `order_book_sequence > 0`
	TradeBookSequenceQuery    = `trade_sequence > 0`
	OrderBookAscSequenceQuery = `order_book_sequence ASC`
	TradeBookAscSequenceQuery = `trade_sequence ASC`
	AddRowBookOrderQuery      = `user_id = ? AND order_book_sequence > 0`
	AddRowTradeOrderQuery     = `user_id = ? AND trade_sequence > 0`
	AddRowQuery               = `user_id = ? AND column_id = ?`
	AddRowConditionOrder      = ` AND order_book_sequence > 0`
	AddRowCConditionTrade     = ` AND trade_sequence > 0`
)

const (
	OrderBook                 = "OrderBook"
	TradeBook                 = "TradeBook"
	OrderBookSequence         = "order_book_sequence"
	TradeBookSequence         = "trade_sequence"
	TotalColumns              = "totalColumns"
	AdvanceOrderType          = "advanceOrderType"
	ColumnsCount              = "columnsCount"
	UsersId                   = "userId"
	ResponseColumnName        = "response_column_name"
	DisplayColumnName         = "display_column_name"
	UserFetchedSuccess        = "Fetched user-specific columns successfully"
	UserFetchedFailed         = "Failed to fetch user-specific columns"
	DefaultColumnsSuccess     = "Inserted default columns successfully for first-time user"
	UserNotFoundInsertDefault = "No user records found, inserting default display columns"
	CountFailed               = "Failed to count user records in order_columns"
)

const (
	OrderBookAscSequenceTestQuery = `SELECT user_id, order_book_sequence FROM order_columns WHERE user_id = ? ORDER BY order_book_sequence ASC`
	UpdateForOrderBookTestQuery   = `UPDATE "order_columns" SET "order_book_sequence"=$1 WHERE user_id = $2 AND column_id IN ($3)`
	DeleteRowIfBothZeroTestQuery  = `DELETE FROM order_columns WHERE user_id = ? AND column_id IN (?) AND order_book_sequence = 0 AND trade_book_sequence = 0`
	DeleteForTradeBookTestQuery   = `UPDATE "order_columns" SET "trade_sequence"=$1 WHERE user_id = $2 AND column_id IN (?)`
)

const (
	ToDisplayQuery                  = "is_display = ?"
	ExistingUserFetchingSelectQuery = "order_columns.column_id, order_column_details.response_column_name, order_column_details.display_column_name"
	ExistingUserFetchingJoinQuery   = "JOIN order_column_details ON order_columns.column_id = order_column_details.column_id"
	ExistingUserFetchingWhereQuery  = "order_columns.user_id = ?"
	SwitchOrderBookWhereQuery       = "order_columns.order_book_sequence > 0"
	SwitchTradeBookWhereQuery       = "order_columns.trade_sequence > 0"
	SwitchOrderBookOrderQuery       = "order_columns.order_book_sequence ASC"
	SwitchTradeBookOrderQuery       = "order_columns.trade_sequence ASC"
)

const (
	GetDefaultColumnsTestQuery       = `SELECT column_id, response_column_name, display_column_name FROM "order_column_details" WHERE is_display = $1`
	GetUserColumnsTradeBookTestQuery = `SELECT order_columns.column_id, order_column_details.response_column_name, order_column_details.display_column_name FROM "order_columns" JOIN order_column_details ON order_columns.column_id = order_column_details.column_id WHERE order_columns.user_id = $1 AND order_columns.trade_sequence > 0 ORDER BY order_columns.trade_sequence ASC WHERE order_columns.user_id = $1 AND order_columns.trade_sequence > 0 ORDER BY order_columns.trade_sequence ASC`
	GetOrderBookColumnsTestQuery     = `SELECT order_columns.column_id, order_column_details.response_column_name, order_column_details.display_column_name FROM "order_columns" JOIN order_column_details ON order_columns.column_id = order_column_details.column_id WHERE order_columns.user_id = $1 AND order_columns.order_book_sequence > 0 ORDER BY order_columns.order_book_sequence ASC`
	CountUserOrderColumnsTestQuery   = `SELECT count(*) FROM "order_columns" WHERE user_id = $1`
	GetUserOrderBookColumnsTestQuery = `SELECT order_columns.column_id, order_column_details.response_column_name, order_column_details.display_column_name FROM "order_columns" JOIN order_column_details ON order_columns.column_id = order_column_details.column_id WHERE order_columns.user_id = $1 AND order_columns.order_book_sequence > 0 ORDER BY order_columns.order_book_sequence ASC`
	GetUserTradeBookColumnsTestQuery = `SELECT order_columns.column_id, order_column_details.response_column_name, order_column_details.display_column_name FROM "order_columns" JOIN order_column_details ON order_columns.column_id = order_column_details.column_id WHERE order_columns.user_id = $1 AND order_columns.trade_sequence > 0 ORDER BY order_columns.trade_sequence ASC`
    // CountUserOrderColumnsTestQuery2   = "SELECT count(*) FROM order_columns WHERE user_id = ?"
    // DefaultColumnsFetchQuery         = "SELECT column_id, response_column_name, display_column_name FROM order_column_details"
    // InsertUserOrderColumnsQuery      = "INSERT INTO order_columns (user_id, column_id, order_book_sequence, trade_book_sequence, created_at, last_updated_at) VALUES ($1, $2, $3, $4, $5, $6)"

	ToDisplayQueryTestQuery = `SELECT * FROM "order_column_details" WHERE is_display = $1`
	DefaultColumnsSelectQuery      = `SELECT column_id, response_column_name, display_column_name FROM "order_column_details" WHERE is_display = \$1`
	InsertDefaultColumnsTestQuery  = `INSERT INTO order_columns`
	ABC = `SELECT * FROM "order_column_details" WHERE is_display = $1`

	
)
