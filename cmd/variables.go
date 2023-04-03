package cmd

import (
	"os"

	"github.com/flashbots/mev-boost-relay/common"
)

var (
	defaultNetwork          = common.GetEnv("NETWORK", "")
	defaultBeaconURIs       = common.GetSliceEnv("BEACON_URIS", []string{"http://localhost:3500"})
	defaultBuilderWhitelist = common.GetSliceEnv("WHITELISTED_BUILDERS", []string{"0xa18fdb012ecadc1d9af521917959d5915f7d17ca1a88953a28cd182c000237e49a1c3a7050ee923459d477041ad0ae45,0xa5eec32c40cc3737d643c24982c7f097354150aac1612d4089e2e8af44dbeefaec08a11c76bd57e7d58697ad8b2bbef5,0x8e39849ceabc8710de49b2ca7053813de18b1c12d9ee22149dac4b90b634dd7e6d1e7d3c2b4df806ce32c6228eb70a8b,0x8931ae674d7b9b0165b784a13301bcc102f70faf4576ec3dc4a31949dde831ec60e79444ed911742fa7b4720691d1e45,0xb1d229d9c21298a87846c7022ebeef277dfc321fe674fa45312e20b5b6c400bfde9383f801848d7837ed5fc449083a12,0xa412007971217a42ca2ced9a90e7ca0ddfc922a1482ee6adf812c4a307e5fb7d6e668a7c86e53663ddd53c689aa3d350"})
	defaultPriorityBuilders = common.GetSliceEnv("PRIORITY_BUILDERS", []string{""})
	defaultRedisURI         = common.GetEnv("REDIS_URI", "localhost:6379")
	defaultPostgresDSN      = common.GetEnv("POSTGRES_DSN", "")
	defaultMemcachedURIs    = common.GetSliceEnv("MEMCACHED_URIS", nil)
	defaultLogJSON          = os.Getenv("LOG_JSON") != ""
	defaultLogLevel         = common.GetEnv("LOG_LEVEL", "info")

	beaconNodeURIs            []string
	whitelistedBuilderPubKeys []string
	redisURI                  string
	postgresDSN               string

	priorityBuilders []string
	memcachedURIs    []string

	logJSON  bool
	logLevel string

	network string
)
