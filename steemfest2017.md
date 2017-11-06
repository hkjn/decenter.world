# The Scale of steemd

Michael Vandeberg: The Scale of steemd, modularity to scale SMTs and beyond

"If something has to happen to the blockchain, I am in some form responsible for it."

Scheduled hardforks, had #16 to reduce inflation drastically in 2016.

Appbase

Next HF: Steem Velocity:
- Mined accounts
- Subsidized account creation

Transaction rate is ~700k / day, i.e. 8 / sec

Consensus state is 16GB+, i.e. requirement for full node that has all state

Running witness nodes will become too expensive:
- Independent Reindex
  - I.e. sharding
- Parallel Reindex
- Distributed Node
  - I.e. cluster of backing nodes serving up one API

Hard forks: Transactions that earlier were invalid now would be valid
Soft forks: Transactions that earlier were valid now (but not historically) would be invalid

Attacks requires 51% of SP (to vote in your own witnesses) or 11 witnesses colluding, claimed

Full witness node seems like it will refer to actually storing the blockchain state.

"Trusted mode" seems to be some suggestion to trust witnesses even more than before

Claims that since largest four Bitcoin mining pools are > 50% of hash power, only four
entities would need to be subverted - but if miners started following different rules,
users / HODLers with full nodes would reject those coins (like Segwit2x).

# My life with the Steempunks by @sneak, Jeffrey Paul

2017 scaling roadmap

6x traffic since last year

Up to 12 TPS at peak

430k accounts

~100 instances on EC2 for Steemit

Will focus more on i18n, mike@steemit.com for helping out

Auto-scaling, self-healing, containerized, microservices, dev-driven deploys

Different FOSS projects:
- sbds
- hivemind
- yo
- jussi
- overseer
- gatekeeper
- boulton

"If Steemit.com is the only gatekeeper into this ecosystem, we've failed at our job."

New features: Community support, multi-transport notifications

38 people currently and "doesn't suck", sneak@steemit.com.

2018 plans: "our community is everything to us". Adding Steemit iOS app.

# Disruptive shifts revolutionizing our economy, by @mrssteemit

APPICS, SMT on Steem

"Next-Generation Social Media App"

"Global Movement"

"What most people don't know is that there's over 800 different blockchains, with
different sets of advantages."

"So don't worry if you don't understand the concept of how the monetary system works,
just start using it; best time is right now."

"We want to give back to Steemit and give 5% back to the community before the ICO."

# ICOs and the democratization of finance

- jeff7
- Jeff Mazer
- SteemIt CFO
- Current ICO market
- What to look for in ICOs
- $3B+ in 2017 so far
- Institutional investors are starting to contribute, 65 - 70 VC funds
- Others (family offices, foundations, pension funds, insurance companies)
- Really took off in May 2017
- Toys R' Us had $7B of debt
- Speculation is driving market, FOMO
- If you invested in every ICO you'd have 1320% return
- Bad or irrelevant whitepapers, and people don't read them before investing
- Things will get worse before they get better
- Private solutions as "self-regulating"
- Benefits company sponsoring solution
- Protocol Labs says they'll verify everything


# dollarvigilante

- "It's hard to look back to 1.5 years ago, which is like 50 years in cryptocurrency space."
- "Average person doesn't care it it's on a blockchain, if it's decentralized, they
  just want it to work."

# SteemIt inc panel

- Scaling Steemit
- Blockchain, backend, frontend
- HTTP + JSON-RPC > WebSocket for analytics, apparently
- steemd is memory hog
- Trying to keep up with transaction growth
- ramdisk for steemd state
- multi-threading AppBase + sharding
- Johan built distributed tracing system, which can shut down DDOS DNS lookups; runs their own resolvers?
- Javascript + Typescript + Python, looking at Golang
- If there's many blockchain explorers, how would users avoid entering their private keys everywhere?
- Account creation fee is solving equivalent problem as transaction fees for UTXO systems
- Looking into setups to subsidize account creation by existing users
- Steemd throughput is aim, reducing latency as temporary improvement but AppBase is big project to make more efficient
- What is the plan when doing further on-chain scaling is infeasible? Sharding? Off-chain?
- Incentivizes users to solve their own problems

# Steem apps

## feruz / good-karma

Built eSteem, alternative blockchain viewer.

## ?? / dtube creator

IPFS is somewhat flaky, sometimes doesn't load

## prc

Distributed music hopsting on IPFS:
- https://dsound.audio/#/feed

## ??

busy.org: Alternative blockchain viewer / social network
steemconnect.org: Show apps connected to account

## ??

utopian.io creator

Diego / https://steemit.com/@elear


# Governance in blockchains by @chris420

- Airbnb does have problems: misrepresentations, damage by guests, nonexisting apartments
- Bitcoin scaling debate: Many proposals and some "agreements", but slow progress
- 8 Bitcoin miners "control the system", but Ethereum has "proof of Vitalik"
- 2 Ether mining pools control 50%+ mining power
- Change from PoW to PoS is controversial, ice age was delayed
- DPoS systems claimed to be better in governance due to being "more flexible"
- Bitshares consist of votes on governance changes
- Bitshares have Workers, where voters put tokens towards proposals

# Golos project

- @serejandmyself
- Golos
- CEO of GolosFund

# @wmogoyar

- Criteria for ICO evaluation:
  - Startup characteristics
  - Crypto sale resiliency
  - Operational transparency
  - Business model
- Be clear about token functionality
- Work with an experienced lawyer
- Be transparent in token operations and performance
- Don't be greedy, it's about decentralization
- Steem is about attention economy
- Average person spends ~1 hr on facebook / day
- wmougayar@gmail.com
- "The less money you have, the better decisions you make."
