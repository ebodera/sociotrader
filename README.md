# sociotrader

![icon](https://storage.googleapis.com/indie-hackers.appspot.com/product-avatars/stock-trading-bot/ibTLPyjwVebnZjMGKvz6ztarnuV2)

## Architecture
- Trade Client
- Data Server
- Algorithms
- Sentiment Analysis

### Trade Client

[alpaca-trade-api-go](https://github.com/alpacahq/alpaca-trade-api-go)
> `alpaca-trade-api-go` is a Go library for the Alpaca trade API. It allows rapid trading algo development easily, with support for the both REST and streaming interfaces. For details of each API behavior, please see the online API document.

### Data Server

[MarketStore](https://github.com/alpacahq/marketstore)
> MarketStore is a database server optimized for financial time-series data. You can think of it as an extensible DataFrame service that is accessible from anywhere in your system, at higher scalability.

### Algorithms

[A Buy-on-dip algo for Alpaca API](https://github.com/alpacahq/samplealgo01)
> This is a simple algo that trades every day refreshing portfolio based on the EMA ranking. Among the universe (e.g. SP500 stocks), it ranks by daily (price - EMA) percentage as of trading time and keep positions in sync with lowest ranked stocks.

[S&P100 replication algo for Alpaca API](https://github.com/alpacahq/sp100algo)
> This is an algorithm that manage your portfolio by simply replication S&P100 index by buying the underlying stocks.


### Sentiment Analysis

[StockTwits?](https://stocktwits.com)


## Resources

[NOTES](NOTES.md)

### Alpaca
- [Alpaca API](https://alpaca.markets/) - Market data to test out bots
- [Alpaca Repo](https://github.com/alpacahq)
  - [roboadvisor](https://github.com/alpacahq/roboadvisor)
- [Alpaca Documentation](https://docs.alpaca.markets)
- [Alpaca Insomnia Workspace](https://github.com/alpacahq/insomnia-workspace)

### IEXTrading
- [IEX Trading](https://iextrading.com/developer/) - Broker firm to execute trades
