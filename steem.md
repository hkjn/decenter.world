# Steem concerns

- Q1: How decentralized is the governance of Steem blockchain? If
  something happened to Steemit Inc, would the network continue
  just the same? How many external developers are there? If
  there was a change which Steemit Inc strongly opposed, but the
  community had consensus was a good idea, would it be deployed?
- Q2: How many full nodes are there? I.e. how many individual computers
  would need to be subverted before existing data on the Steem blockchain
  couldn't be accessed?
- Q3: How many witnesses are there? I.e. how many individual computers / people
  would need to be attacked before new blocks no longer could be created?

## Consensus

How does DPoS solve nothing-at-stake problems?

I.e, if an attacker ever has sufficient (50%+?) stake in the system
at any block, couldn't they vote in their own witnesses at that point,
and so a newly starting node would believe their false history from
that point onward (even if they had lower stake later in the "true"
chain)? In consensus research, it seems that the security properties
of PoS schemes are known to be weaker than PoW systems, with common
"fixes" such as having known third parties signing blocks or checkpoints
baked into source code much less ideal.

## Centralized

The governance of the system is very important to make sure that it's
not controlled by a single company or group of people. It's not clear
how many independent developers are working on Steem, or whether witnesses
actually would reject consensus changes proposed by Steemit Inc.

## Extrinsic motivation

When someone shares content with us, we'd like them to do that
because they care about the content, not because they're expecting
a payout. Extrinsic explicit incentivization supresses authentic motives.

## Value of posts should be few cents for almost all users

Actual value of facebook users are on order of dozens of dollars / year,
so payouts of hundreds of dollars or more per post can't really make sense
economically, even if authors / curators are allowed to capture most value
they create.

## Consensus mechanism may privilege large holders

Why are only 1 / 20 blocks PoW?

PoS tends to centralize wealth in existing owners.

Is this just to be able to call the consensum mechanism
"hybrid PoW / PoS"?

## Uneven distribution of tokens

Initial distribution mechanism was to quickly mine blocks with large
inflation rate [ todo: more details ], leading few people [ 45? ]
controlling > 50% of all Steem Power.

Under Pos / DPoS, existing SP owners will receive more SP both in absolute
numbers and in proportion of all SP.

Inflation was 100% new STEEM per year earlier, but lowered to 9.5%
with 0.5% less per year in 2016:

- https://steemit.com/steem/@steemitblog/final-review-of-steem-economic-changes

## Low supply of liquid STEEM causes high prices

Less than 10% of actual STEEM tokens in circulation [ todo: more details ]
are actually traded (the rest locked up in SP), which pushes up price.
