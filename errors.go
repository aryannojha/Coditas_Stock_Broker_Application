package constants

// Log errors
const (
	OpeningOrderErr               = "error opening orders.csv: %w"
	OpeningFileErr                = "failed opening %s CSV file: %w"
	ReadingContentErr             = "failed reading %s CSV content: %w"
	WatchingStartFail             = "Could not start watching new file: "
	FileTemporarilyUnavailable    = "File temporarily unavailable (may have been replaced): "
	WatcherErr                    = "Watcher error: "
	OrderTruncateFailed           = "failed to truncate orders table for empty file: %w"
	OrderInvalidHeader            = "invalid header in orders.csv expected %v, got %v"
	OrderParsedFailedOrderId      = "ParseOrdersCSV: failed to parse OrderID '%s' at line %d"
	OrderParsedFailedQuantity     = "ParseOrdersCSV: failed to parse Quantity '%s' at line %d"
	OrderParsedFailedPrice        = "ParseOrdersCSV: failed to parse Price '%s' at line %d"
	OrderNoValidOrders            = "ParseOrdersCSV: No valid orders found to save."
	FailedToTruncate              = "failed to truncate and load: %w"
	TradeTruncateFailed           = "failed to truncate trades table for empty file: %w"
	InvalidHeader                 = "invalid %s CSV header. Expected: %v, Got: %v"
	ConInvalidHeader              = "%w: invalid header for %s CSV"
	TradeParseFailedTradeId       = "ParseTradesCSV: failed to parse TradeID '%s' at line %d"
	TradeParseFailedTradePrice    = "ParseTradesCSV: failed to parse TradePrice '%s' at line %d"
	TradeParseFailedTradeQuantity = "ParseTradesCSV: failed to parse TradeQty '%s' at line %d"
	TradeParseFailedTradeTime     = "ParseTradesCSV: failed to parse TradeTime '%s' at line %d"
	TradesNoValidTrades           = "ParseTradesCSV: No valid trades found to save."
	EmptyCSVErr                   = "The %s CSV file is empty."
	ExistingOrdersReadFailed      = "Failed to process existing Orders CSV: "
	ExistingTradesReadFailed      = "Failed to process existing Trades CSV: "
	FailedOrdersJob               = "Failed to add Orders job: %v"
	FailedTradesJob               = "Failed to add Trades job: %v"
)

// Scheduler
const (
	FailedToUpsert                     = "failed to upsert holdings: %v"
	FailedToLoad                       = "failed to truncate and load holdings: %v"
	NoValidRecords                     = "No valid records in file"
	InvalidAvgPrice                    = "invalid average_price"
	InvalidCuspaAvailableQuantity      = "invalid cuspa_available_quantity"
	InvalidHoldingAvailableQuantity    = "invalid holding_available_quantity"
	InvalidT1AvailableQuantity         = "invalid t1_available_quantity"
	InvalidCollateralAvailableQuantity = "invalid collateral_available_quantity"
	InvalidMtfAvailableQuantity        = "invalid mtf_available_quantity"
	GeneralParseFailed                 = "Scheduler: %s parse failed: %v"
	GeneralDatabaseErr                 = "Database error in %s: %v"
	GeneralUnknownErr                  = "Unknown error in %s: %v"
)

// GenericJob Errors
const (
	FileNotFound          = "%s file not found at path: %s"
	FileMissing           = "File Missing"
	JobErr                = "%s Job"
	ParseFailedErr        = "Parsing Failed"
	DataBaseErr           = "DataBase Error"
	UnknownErr            = "Unknown Error"
	EmptyCSV              = "Empty CSV File"
	NoDataRow             = "No data rows found in the file"
	TruncateAndLoadFailed = "Database Truncate and Load Failed"
	UpsertFailed          = "Database Upsert Failed"
)
