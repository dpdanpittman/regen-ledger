module github.com/regen-network/regen-ledger/x/ecocredit

go 1.15

require (
	github.com/cockroachdb/apd/v2 v2.0.2
	github.com/cosmos/cosmos-sdk v0.42.4
	github.com/gogo/protobuf v1.3.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/regen-network/regen-ledger v0.6.2
	github.com/regen-network/regen-ledger/orm v0.0.0-00010101000000-000000000000
	github.com/regen-network/regen-ledger/types v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.37.0
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/regen-network/regen-ledger => ../../

replace github.com/regen-network/regen-ledger/orm => ../../orm

replace github.com/regen-network/regen-ledger/types => ../../types

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
