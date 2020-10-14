module github.com/rocksideio/chainlink-metatx-adapter

go 1.13

require (
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/ethereum/go-ethereum v1.9.21
	github.com/linkpoolio/bridges v0.0.0-20200226172010-aa6f132d477e
	github.com/rocksideio/chainlink-metatx-adapter/app v0.0.0-00010101000000-000000000000
)

replace github.com/rocksideio/chainlink-metatx-adapter/app => ./app
