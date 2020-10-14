# chainlink-metatx-adapter

## Configuration

```
RPC_URL                   # node url
GNOSIS_ADDRESS            # gnosis safe address
GNOSIS_OWNER_PRIVATE_KEY  # gnosis safe owner private key
RELAYER_PRIVATE_KEY       # relayer private
```

## Configuration

1. install a node-operator

2. [configure the node-operator to use this external adapter](https://docs.chain.link/docs/node-operators)

3. Configure the node-operator to use the external adapter instead of `EthTx`

4. Serve data using the node-operator

## Cloud

Compatible with google cloud function.
