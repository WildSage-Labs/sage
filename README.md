# SAGE cosmos blockchain monitoring tool
## How this works:
1) Upon start, sage will query the configured chains, to determine latest block height
2) Works then start querying the chain for new blocks 
3) Upon new block, wallet balances are queried

TODO:
1) Create a condition that can be set (balance increase, decrease or a wallet transaction happened)
2) Notifications to discord,tg (whatever is enabled)