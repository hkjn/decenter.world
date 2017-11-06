# Bitshares

Originated from Protoshares project, which turned into BitsharesX.

Graphene was split out from Bitshares 2.0

`bitshares/bitshares-core,bitshares-ui`

`witness_node`: Witness node has JSON-RPC over WebSockets

`cli_wallet`: Communicates with `witness_node

Several web wallets:

wallet.bitshares.eu
openledger.io
www.freedomledger.com

Bitshares mobile wallet from BitUniverse, open source

cryptofresh.com: block explorer

PyBitShares library for backend deevelopment

bts_tools: Monitoring for backend

`stakemachine`, `Bitcrab/transbot`: Trading bot

## Graphene features

- Human-readable names
- Hierarchical, multi-sig access control
- "Smart coins", user-issued assets
- Witnessess, DPoS block producers
- About 25 witnesses currently
- Witnesses publish prices
- Committe members vote for blockchain parameters changes like block size
- Workers: Produce work to be done, get votes
- Proxies: Vote for proposals on behalf for users
