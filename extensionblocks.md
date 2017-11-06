# Extension Blocks

- Each extension block would be a SHA256 hash of bytecode of Moxie
  script or similar
- Goal can be to disallow further softforks by just having changes
  via sidechains or extension blocks
- In endgoal, new rules could be written, tested, validated, then
  published to blockchain as a transaction, after which sidechain
  would be live and others could opt in by transferring coins to
  start using it
- Cryptosystem rules can be published as a special type of transaction,
  in very succinct bytecode that just needs to fit within a block
- Bitcoin being mutable is risky, since even coin cap could be softforked
- But being too immutable is also risky is also bad, and core parts
  could be locked down while allow extensions

