module github.com/regen-network/regen-ledger/types

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.42.4
	github.com/gogo/protobuf v1.3.3
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.9
	github.com/tendermint/tm-db v0.6.4
	google.golang.org/grpc v1.37.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/regen-network/regen-ledger => ../
