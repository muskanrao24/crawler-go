package config

const (
	RpcProtocol = "tcp"

	ItemSaverPort = 1234
	WorkerPort0   = 9000

	ElasticIndex = "dating_profile"

	ItemSaverRpcMethod    = "ItemSaverService.Save"
	CrawlServiceRpcMethod = "CrawlService.Process"

	// Parser Names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"
)
