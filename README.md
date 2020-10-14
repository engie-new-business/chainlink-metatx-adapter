# chainlink-metatx-adapter

Use a gnosis safe to fulfill request. Replace `EthTx` from your job.

This is just a poc from [rockside.io](https://rockside.io/) team. The main idea behind this poc is to see if a relayer can relay transactions for a data provider and accept a reimbursement in link. This poc doesn't include the reimbursement.

## Configuration

First set these env variable

```
RPC_URL                   # node url
GNOSIS_ADDRESS            # gnosis safe address
GNOSIS_OWNER_PRIVATE_KEY  # gnosis safe owner private key
RELAYER_PRIVATE_KEY       # relayer private
```

The relayer need to have some ether to make the transactions. 

Then:

1. install a node-operator

2. [configure the node-operator to use this external adapter](https://docs.chain.link/docs/node-operators)

3. Configure the job to use the external adapter instead of `EthTx`

4. Serve data using the node-operator

## Cloud

Compatible with google cloud function.

