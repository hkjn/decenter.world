# Softforks

- Preferred
- Bitcoin scripts with `OP_TRUE` are "anyone can spend" output, since they
  will immediately fulfill spending requirements
- But since `OP_TRUE` can have parameters afterwards, that provides
  backwards-compatible upgrade mechanism
- So miners supporting new feature (e.g. Schnorr signatures) will have
  more narrow definition than "anyone can spend", such as validating
  that the signature is correct, so a naive user trying to spend the
  outputs would not get mined by upgraded miners, leading to a softfork
  activation mechanism if majority of miners have upgraded
- Used by BIP16, BIP30, BIP34, ..
- During upgrade, there's three types of miners: upgraded miners,
  non-upgraded miners, and users of both categories
- Non-upgraded users will during this time be vulnerable to a 51% attack,
  since they could grab transactions that the users wouldn't be able to
  validate
- Sidechains are just the same; they look like something that was taken
  by an arbitrary spender
- The subset of users who have upgraded can be viewed as an in-chain
  sidechain, they understand an extended feature
- So sidechains can be seen as a generalization and modularization of
  existing upgrade behavior
- By making the Bitcoin network effectively an SPV client of the sidechain
  network, there's loose coupling; problems on sidechain network don't
  leak into the Bitcoin network
- The cost is different security model, hashrate on sidechain network
  can steal coins; if there's mergemining and transaction fees this
  shouldn't be much of an issue since they're at similar level
- Bitcoin network could enforce the validity of sidechain blocks, so main
  network hashrate prevents sidechain recovery transaction unless sidechain
  agreed; but this makes sidechain rules part of Bitcoin consensus
- If 85% of miners were doing mergemining with bitcoind and sidechaind
  daemons side-by-side, this means that the miner can check that the valid
  bitcoin transactions also are valid according to the sidechain
- And if the sidechain rules were enforced by Bitcoin consensus rules,
  it could call out to sidechaind to verify validity of the transaction,
  although this risks splitting the network if the valid/invalid response
  were to be nondeterministic
- Enforcing determinism could be done by abstract virtual machines like
  Moxie / MoxieBox
- I.e. very simple CPU similator could run the sidechain code, similarly
  to Bitcoin Script
- This provides nice onramp, where you start out with weaker-than-Bitcoin
  security, but there's no coupling
- If sidechain becomes widely used, we can then make it a safe upgrade
  to Bitcoin
- Blocksize upgrades could somewhat surprisingly be softforked in, by
  introducing extension block of larger size, which in larger
  merkle tree can be "linked in" in original tree in a way that they're
  ignored by people who don't understand the link
