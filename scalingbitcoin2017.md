# Pedro Moreno-Sanches

- < 10 GB transactions / sec
- > 135 GB disk required
- Payment channels:
  - Opening channel:
    - Alice deposits 5 BTC with Bob
    - Time-locked contract that allows her to get them back afterwards
  - Arbitrarily many off-chain transactions
  - Closing channel:
    - Coins in last state of channel settled on chain
  - Impractical to have N^2 channels between each two users
- Lightning network:
  - Hash Time-Lock Contract: enables conditional payment between users
  - Pay Bob 1 BTC if he provides x so that H(x) = y, within time T 
  - Multiple chained HTLC enables multi-hop payments: HTLC(Alice, Bob, 1, y, 30) -> HTLC(Bob, Cat, 1, y, 29)
  - Bob (between Alice and Cat) does not gain or lose coins in total
  - Want Balance Security and Serializability
  - Privacy Properties:
    - Off-path value privacy: Someone not in payment route should not see sender and receiver
    - On-path relationship privacy: Someone on the payment route should not see more than hops before / after them
  - Issue that same condition y appears in both HTLC steps
  - Their proposal is a P2P network called Fulgor, with multi-hop HTLC
  - Builds on non-interactive zero knowledge (ZKBoo[GMO16]):
  - Miners trivially can solve serializability requirement
  - Impossibility result in getting both concurrency and privacy (no global identifiers)
  - Ralyo: Full concurrency, but weaker privacy
  - Other privacy concerns are amount and timing of transaction

# BOLT Anonymous Payment Channels for Decentralized Currencies by Ian Miers

- Founding scientist of ZCash
- Batching payments together is familiar from bar tabs
  - If you trust the bar, you can hand them your card
  - If the bar trusts you, you give them the card at the end
- Payment channels are methods to swap IOUs
  - Pay Moe $100 with credit cards
  - Moe gives you an IOU for $95 and one beer
- Individual payment channel compromise privacy
- Payment channel networks help, since you can do onion routing and only entrypoint is seen
- Decentralized / secure / private is a "Zooko's triangle" type thing
- They have a zero-knowledge setup to add more privacy
- 1-2 sec to set up channel, less than 100 ms per hop
- Uses zero-knowledge proofs (but not zkSNARKs, i.e. no trusted setup)
- Tumblebit works on current network
- Their new work needs an op code in Bitcoin
- Channel node operators need to store ~30 bytes / client and to monitor network

# ValueShuffle: Mixing Confidential Transactions by Tim Ruffing

- Layer 1 has bad privacy properties
- Many services clusters and deanonymizes users
- CoinJoin is a way to send their own coins to their own fresh UTXOs
  by mixing it with other users
- Users need to find others to mix transactions with
- P2P mixing proposal ValueShuffle, together with Pedro
- Peers agree on the output and sign result
- DiceMix + CoinJoin => CoinShuffle++
- Combines different properties from Stealth addresses and Confidential Transactions well

# Graphene: A New Protocol for Block Propagation Using Set Reconciliation by Brian N. Levin

- Graphene is way to send block info in 1/10th of compact block size
- Full blocks: Send TXIDs instead of full transactions
- Actually just send prefix of TXIDs, first 5-6 bytes give 1/10**12 odds of mistake
- BlueMatt and maku on IRC claim that paper assumes tree-structured network, and
  otherwise is "useless apart from as an academic result"
  - But others point out that this is not required by current results,
    although there's a wider paper that might require this
- Assumes ordering of transactions in "canonical order", which roasbeef points out
  requires extra data to be sent

# BlockSci: a Platform for Blockchain Science and Exploration by Harry Kalodner

- Tries to analyze outputs that moved recently, discounting change addresses etc
- Correlates trade volume from exchanges with moving outputs on the blockchain
- Builds FOSS tool for blockchain analytics

 
