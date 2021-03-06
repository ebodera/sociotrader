package main

import (
	"log"
	"sort"
	"time"

	"github.com/VividCortex/ewma"
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	strftime "github.com/jehiah/go-strftime"
	"github.com/shopspring/decimal"
)

type Order struct {
	symbol string
	qty    int
	side   string
}

type Diff struct {
	symbol string
	diff   float64
}

var (
	alpacaClient = alpaca.NewClient(common.Credentials())
	NY           = "America/New_York"
	sp100        = []string{"AAL", "AAPL", "ADBE", "ADI", "ADP", "ADSK", "ALGN", "ALXN", "AMAT", "AMD", "AMGN", "AMZN", "ASML", "ATVI", "AVGO", "BIDU", "BIIB", "BKNG", "BMRN", "CDNS", "CELG", "CERN", "CHKP", "CHTR", "CMCSA", "COST", "CSCO", "CSX", "CTAS", "CTRP", "CTSH", "CTXS", "DLTR", "EA", "EBAY", "EXPE", "FAST", "FB", "FISV", "FOX", "FOXA", "GILD", "GOOG", "GOOGL", "HAS", "HSIC", "IDXX", "ILMN", "INCY", "INTC", "INTU", "ISRG", "JBHT", "JD", "KHC", "KLAC", "LBTYA", "LBTYK", "LRCX", "LULU", "MAR", "MCHP", "MDLZ", "MELI", "MNST", "MSFT", "MU", "MXIM", "MYL", "NFLX", "NTAP", "NTES", "NVDA", "NXPI", "ORLY", "PAYX", "PCAR", "PEP", "PYPL", "QCOM", "REGN", "ROST", "SBUX", "SIRI", "SNPS", "SWKS", "SYMC", "TMUS", "TSLA", "TTWO", "TXN", "UAL", "ULTA", "VRSK", "VRSN", "VRTX", "WBA", "WDAY", "WDC", "WLTW", "WYNN", "XEL", "XLNX"}
)

func init() {
	log.Printf("Running w/ credentials [%v %v]\n", common.Credentials().ID, common.Credentials().Secret)

	acct, err := alpacaClient.GetAccount()
	if err != nil {
		log.Fatalln("Unable to get account data: ", err)
	}

	log.Println(*acct)
}

func getPrices(symbols []string, endDt time.Time) map[string][]alpaca.Bar {
	var i, j int
	var e error
	barset := make(map[string][]alpaca.Bar, len(symbols))
	opts := alpaca.ListBarParams{Timeframe: "1D"}

	for i <= len(symbols)-1 {

		if (i + 200) > len(symbols) {
			j = len(symbols)
		} else {
			j = i + 200
		}

		if i == 0 {
			barset, e = alpaca.ListBars(symbols[i:j], opts)
			if e != nil {
				log.Println("Unable to get prices: ", e)
			}
		} else {
			a, e := alpaca.ListBars(symbols[i:j], opts)
			if e != nil {
				log.Println("Unable to get prices: ", e)
			}

			for k, v := range a {
				barset[k] = v
			}
		}

		i += 200
	}

	return barset
}

func prices(symbols []string) map[string][]alpaca.Bar {
	loc, _ := time.LoadLocation(NY)
	now := time.Now().In(loc)
	endDt := now
	mktOpen, _ := time.ParseInLocation("03:04PM", "09:30AM", loc)

	if now.After(mktOpen) {
		endDt = endDt.Add(-1 * time.Minute)
	}

	return getPrices(symbols, endDt)
}

func calcScores(priceDf map[string][]alpaca.Bar, dayindex int) []Diff {
	var diffs []Diff
	param := 10
	ema := ewma.NewMovingAverage(float64(param)) //=> returns a VariableEWMA with a decay of 2 / (5 + 1)

	for symbol, df := range priceDf {

		if len(df) <= param {
			continue
		}

		for _, bar := range df {
			ema.Add(float64(bar.Close))
		}

		last := float64(df[len(df)-1].Close)
		diff := (last - ema.Value()) / last
		diffs = append(diffs,
			Diff{
				symbol: symbol,
				diff:   diff,
			})
	}

	sort.Slice(diffs[:], func(i, j int) bool {
		return diffs[i].diff < diffs[j].diff
	})

	return diffs
}

func getOrders(c *alpaca.Client, priceDf map[string][]alpaca.Bar, size int, max int) []Order {
	var orders []Order

	acct, err := alpacaClient.GetAccount()
	if err != nil {
		log.Fatalln("Unable to get account data: ", err)
	}

	log.Println("Calculating scores...")
	ranked := calcScores(priceDf, -1)
	log.Printf("%v", ranked)

	// toBuy := mapset.NewSet()
	// toSell := mapset.NewSet()

	for i := range ranked[:len(ranked)-1/20] {
		symbol := ranked[i].symbol
		last := len(priceDf[symbol]) - 1
		price := decimal.NewFromFloat32(priceDf[symbol][last].Close)
		if price.LessThan(acct.Cash) {
			continue
		}
		toBuy.Add(symbol)
	}

	return orders
}

func main() {
	var done string

	for {
		clock, err := alpaca.GetClock()
		if err != nil {
			log.Fatalln("Unable to get clock: ", err)
		}

		now := clock.Timestamp

		//clock.IsOpen
		if done != strftime.Format("%Y-%m-%d", now) {
			log.Println("Getting prices data frame...")
			priceDf := prices(sp100)

			log.Println("Getting list of orders...")
			orders := getOrders(alpacaClient, priceDf, 100, 5)
			//trade(orders)

			done = strftime.Format("%Y-%m-%d", now)
			log.Printf("Done for %s", done)
			log.Printf("%v", orders)
			break
		}
	}
}
