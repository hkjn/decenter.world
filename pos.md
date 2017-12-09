# Proof-of-stake

- PoS proves that a group of people came to consensus that a certain
  timeline is acceptable, but it does not prove that's the only timeline
  they claimed was acceptable; nothing stops them from signing alternate
  histories to e.g. doublespend
- gmaxwell and others were excited when first proposals came out in 2011
- But deeper analysis showed fundamental problems, which only can be
  handled by abandoning sybil resistance, or requiring developer-signed
  blocks to prevent reorgs
- What happens if you own some coins / stake at the beginning of the
  network, which you later sell, then at a later time you create
  a simulated network where you didn't sell the coins, that you show
  to new people to the network
- Only known fixes is to not use PoS but instead "ask a friend" or similar
  (not sybil resistant)
- Often degrades into alternate PoW by grinding through many alternate
  histories to find ones where you own all the coins, which is more
  vulnerable to temporary attacks where you rent VMs from Amazon instead
  of needing to build ASICs
