module github.com/regen-network/regen-ledger/types/testutil

go 1.15

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

require (
	github.com/CosmWasm/wasmd v0.15.0
	github.com/cosmos/cosmos-sdk v0.42.4
	github.com/gorilla/mux v1.8.0
	github.com/rakyll/statik v0.1.7
	github.com/regen-network/regen-ledger v1.0.0
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.10
	github.com/tendermint/tm-db v0.6.4
	google.golang.org/grpc v1.37.0
)
