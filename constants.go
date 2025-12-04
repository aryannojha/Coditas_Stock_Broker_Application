package constants

const (
	WatchDir = "./data"
)

// Names of columns Orders and Trades
const (
	OrdersCSV     = "orders.csv"
	TradesCSV     = "trades.csv"
	TradeId       = "trade_id"
	OrderId       = "order_id"
	TradePrice    = "trade_price"
	TradeQuantity = "trade_qty"
	TradeTime     = "trade_time"
	Symbol        = "symbol"
	Quantity      = "quantity"
	Price         = "price"
	Status        = "status"
	ConOrders     = "orders"
	Orders        = "Orders"
	ConTrades     = "trades"
	Trades        = "Trades"
	WebHoldings   = "Holdings"
	CNC           = "CNC"
	MTF           = "MTF"
)

const (
	OrdersAddress = WatchDir + "/" + OrdersCSV
	TradesAddress = WatchDir + "/" + TradesCSV
	HoldingsPath  = "./data/HoldingsEDIS.csv"
	OrdersPath    = "./data/orders.csv"
	TradesPath    = "./data/trades.csv"
)

// Queries
const (
	OrderTruncateAndLoadQuery = `TRUNCATE TABLE orders RESTART IDENTITY`
	TradeTruncateAndLoadQuery = `TRUNCATE TABLE trades RESTART IDENTITY`
	HoldingsEDISTruncateQuery = `TRUNCATE TABLE holdings_edis`
	TruncateFailedTable       = `TRUNCATE TABLE holdings_failed_message`
)

// Logs Info or Warnings
const (
	InitializingWatcher         = "Initializing CSV file watcher..."
	FailedCreatingWatcher       = "Failed to create fsnotify watcher: "
	FailedWatchingDirectory     = "Could not watch directory: "
	WatchingDirectory           = "Watching directory: "
	WatchingExistingFiles       = "Watching existing file: "
	ExistingOrdersRead          = "Processing existing Orders CSV on startup: "
	ExistingTradesRead          = "Processing existing Trades CSV on startup: "
	FileNotFoundWait            = "File not found yet, will watch for creation: "
	FileCreated                 = "File created: "
	UpdateDetectionOrders       = "Detected update in orders.csv"
	UpdateDetectionTrades       = "Detected update in trades.csv"
	ChangeInUnrelatedFile       = "Change detected in unrelated file: "
	FileReappear                = "File reappeared: "
	StoppingCSV                 = "Context cancelled, stopping CSV watcher gracefully."
	OrderBulkSave               = "ParseOrdersCSV: Attempting to bulk save %d valid orders."
	TradeBulkSave               = "ParseTradesCSV: Attempting to bulk save %d valid orders."
	OrderandTradeWatcherService = "Starting Orders & Trades CSV Watcher Service"
	ManualTriggerScheduler      = "manually triggering the Scheduler"
	ReceivedSignal              = "Received signal: %v. Shutting down gracefully"
	UnknownJobType              = "Unknown job type: %s. Use 'orders', 'trades', or 'holdings'."
	OrdersSchedulerStart        = "Starting Orders Scheduler"
	OrdersSchedulerStop         = "Stopping Orders Scheduler"
	TradesSchedulerStart        = "Starting Trades Scheduler"
	TradesSchedulerStop         = "Stopping Trades Scheduler"
)

// Names of columns HoldingsEDIS
const (
	Holdings                    = "holdings"
	AccountID                   = "account_id"
	Isin                        = "isin"
	AveragePrice                = "average_price"
	CuspaAvailableQuantity      = "cuspa_available_quantity"
	HoldingAvailableQuantity    = "holding_available_quantity"
	T1AvailableQuantity         = "t1_available_quantity"
	CollateralAvailableQuantity = "collateral_available_quantity"
	MtfAvailableQuantity        = "mtf_available_quantity"
	ProductCode                 = "product_code"
	TriggerTime                 = "Scheduler triggered at %v"
	GeneralParseSuccess         = "Scheduler: %s parse success"
	Batchsize                   = 400
)

// Logs Info
const (
	SavingRecordsUpsert          = "Saving %d Holdings EDIS records via UPSERT"
	SavingRecordsTruncateandLoad = "Saving %d Holdings EDIS records via UPSERT"
	StartScheduler               = "Starting Daily Scheduler"
	FailedAddCronJob             = "Failed to add cron job: %v"
	StopScheduler                = "Stopping daily scheduler gracefully"
	RunHoldingsJobNow            = "Running Holdings job immediately"
	SkippingLine                 = "Skipping line %d: expected %d columns, got %d"
	SkipLineAllEmpty             = "Skipping line %d: all fields empty"
	ProductCodeLineSkip          = "Skipping line %d: Invalid ProductCode '%s'"
)
