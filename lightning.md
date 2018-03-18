# Lightning Network

## Links

- Mainnet LN graph: https://lnmainnet.gaben.win/#
- dx/dt's mainnet dashboard: http://lnstat.ideoflux.com:3000/dashboard/db/lightning-network?refresh=5m&orgId=1
- ACINQ's testnet explorer: https://explorer.acinq.co/#/
- Landing site: http://lightning.network/
- BOLT proposals: https://github.com/lightningnetwork/lightning-rfc/blob/master/00-introduction.md
- Blockstream's c-lightning implementation in C: https://github.com/ElementsProject/lightning
- Lightning Labs lnd implementation in Go: https://github.com/lightningnetwork/lnd
- ACINQ's Eclair implementation in Scala: https://github.com/ACINQ/eclair
- Blockstream's Lightning Charge wordpress plugin: https://blockstream.com/2018/01/16/lightning-charge.html
- Zap desktop UI around lnd: https://github.com/LN-Zap/zap-desktop
- Lightning Labs desktop UI around lnd: https://github.com/lightninglabs/lightning-app
- Instructions for using bitcoind as lnd backend: https://github.com/lightningnetwork/lnd/blob/master/docs/INSTALL.md#running-lnd-using-the-bitcoind-backend
- Lightning app directory at community site: http://dev.lightning.community/lapps/index.html
- Mainnet yalls.org: https://mainnet.yalls.org
- Minimal desktop LN wallet: https://github.com/minecraft2048/ln-pay
- Graph explorer for lnd: https://github.com/altangent/lnd-explorer

In order to maintain the decentralization and security properties of a
proof-of-work based system, each block must be fast to propagate across
the network and fast to verify after it is discovered. The target rate of
block production also needs to be on the order of several minutes in order
to have good thermodynamic guarantees about how expensive it would be to
rewrite history and attack the network. This means that blocks need to be
small and slow, so there's a very limited amount of space available in blocks
for transactions.

The Lightning Network is one proposal for drastically improving the
scalability of Bitcoin-based systems, by routing off-chain transactions
between participants. Idea of LN was introduced in Feb 2015, and works by
opening payment channels between participants. Opening a payment channel
requires a transaction signed by both parties to be mined into the
blockchain. Once opened, the funds bound by the opening transaction can be
used for any number of off-chain transactions. Closing out the channel
requires a second on-chain transaction. LN is a network of payment
channels, to avoid each participant to have a payment channel to each
other participant; all that's needed is that there's some route between
participants through which payments can be routed.

Apart from scalability improvements, privacy can also be improved, since
only sender, recipient, and any nodes routing the payment even need to see
the transaction. Intermediate nodes only see previous and next node as
well as the sum.

There's some restrictions that were added by Rusty Russell when defining
LN protocol in Q1 2017. 0.04 BTC is the largest individual payment
currently allowed in the LN protocol. Routing payments of other nodes
requires each node to have sufficient liquidity to send the transaction
onward, which was the reason for the limit.

Max 20 hops can be used when routing.

Anyone can join as a LN node and provide liquidity, so it's not expected
that collecting routing fees should form a large amount of income, since
more people would be tempted to join at that point.

A percentage-wise fee based on the value of transaction routed through LN
is paid to the routing nodes. This means that over some value, the fee per
byte of transaction of the L1 transactions will be more efficient.

The offchain payment transactions are always valid and could be published
to the Bitcoin blockchain at any point to retrieve your bitcoins, but if
the counterparty cooperates there's no need to publish it.

Resource requirements for running a LN node are quite small, since there's
no requirement to store a large amount of data like the UTXO set.

Funding transaction used to open payment channel refers to transaction id
and is a 2-of-2 multisig transaction, which due to transaction
malleability could end up with different txid. 

Earlier iterations of LN design used Check-Lock-Time-Verify (CLTV) to lock
transactions until block N, which made it a requirement to close out the
channel before that time. Using relative time locks, we can say that the
new transaction will be valid after some other transaction T has been
mined and another 100 blocks have passed.

When closing a channel unilaterally, the user announces their transaction
that gets mined. The output going to that user is timelocked for 100
blocks, to give counterparty a time window to claim the bitcoin instead if
the first user defected. The LN setup has a ahared secret every time the
state of the payment channel changes. If a user tries to spend an earlier
transaction than the correct final state, e.g. a state where the user had
a larger fraction of the channel's funds, the counterparty will know the
shared secret and can claim all the funds within 100 blocks.

Cross-chain atomic swaps also make it possible to run e.g. Bitcoin and
Litecoin LN nodes, receiving BTC and sending out LTC.

The sender builds the route and selects nodes that are part of the path,
including ones with reasonable fees. This means that the sender needs to
store full network topology and information about all channels. Eventually
this will be too large to store for a small node. At 10k users, this is
~1MB, at 1M users it's ~121MB, 100M users is ~12.1 GB. Dynamically updating
the information about fees and channels frequently requires more and more
bandwidth, 10k users of ~1GB / day, 1M users of 121 GB / day, 100M users
will lead to 11 TB / day, or fully utilized 1 GBPS line. Planned solution
is that each node long term shouldn't need to know full topology, but only
close neighborhood. 

Implementations of LN are done by Blockstream, Lightning Labs (lnd), and
Acinq (Eclair). There's also an implementation called Lit by MIT, and
Thunder by Blockchain.info.

As of Dec 2017, Lightning Network is being actively tested on the Bitcoin
test networks, and there's a few trials on the main network.
