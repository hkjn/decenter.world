# Sidechains 

Security provided by sidechains is different than security of Bitcoin itself

## Security

- Miners attacking main Bitcoin chain lose out on transaction fees, and
  can only doublespend (mine their own transaction, then reorg by
  mining a longer chain)
- Sidechains are SPV wallets from main chain's perspective, which means
  that attacks can steal coins on sidechain, not just double-spend them

## sipa notes 

01:23 < sipa> sidechains, in my view, are a way to experiment with new blockchain technology without first needing to create a new currency
01:23 < sipa> but they don't really give you any scaling advantages
01:24 < sipa> there are still chains involved that need publishing and validation
01:25 < sipa> payment channels are different in that there are no published transactions for all the payments - they're kept private between the participants, and they only publish an aggregated
              transaction that represents multiple transactions
01:27 < sipa> there is some similarity with private federated sidechains
01:27 < sipa> except those are just between a small number of fixed participants
01:28 < sipa> it's like this... hey puff and i are going to be transferring money between eachother very often, let's lock up some coins on the main chain that both of us always have to sign off on
01:28 < sipa> and we can create our own "chain" with updates to our balances
01:30 < sipa> the disadvantage is: it's just between the two of us (or 3, or 4, or 20, but not with 10000), and there is no way to settle in case one of the parties goes offline... so you really only want to do it
              with entities you know in real life and can go after if they hold your coins hostage
01:30 < sipa> the advantage is: you're not restricted to bitcoin-like technology in the sidechain - it can be anything you come up, as long as all participants agree on it
01:30 < sipa> which is great if you want to experiment with new technology
01:31 < sipa> lightning doesn't let you do that... lightning payments _are_ bitcoin transactions - just most of them aren't published
