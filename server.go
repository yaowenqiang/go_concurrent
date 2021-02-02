package main
import (
    "net/http"
    "io/ioutil"
    "encoding/xml"
    "fmt"
    "time"
    "runtime"
)


type QuoteResponse struct {
    Status string
    Name string
    LastPrice float32
    Change float32
    ChangePercent float32
    timeStamp string
    MSDate float32
    MarketCap int
    Volume int
    ChangeYTD float32
    ChangePercentYTD float32
    High float32
    Low float32
    Open float32
}


func main() {

    runtime.GOMAXPROCS(4)
    stocksymbols := []string{
        "googl",
        "msft",
        "aapl",
        "bbry",
        "hpq",
        "vz",
        "t",
        "tmus",
        "s",
    }
    start := time.Now()

    numcomplete := 0
    for _, symbol := range stocksymbols {
        go func(symbol string) {
            resp, _ := http.Get("http://dev.markitondemand.com/Api/v2/Quote?symbol=" + symbol)
            defer resp.Body.Close()
            body, _ := ioutil.ReadAll(resp.Body)

            quote := new(QuoteResponse)
            xml.Unmarshal(body, &quote)
            fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)
            numcomplete++
        }(symbol)
    }
    for numcomplete < len(stocksymbols) {
        time.Sleep(10 * time.Millisecond)
    }

    elapsed := time.Since(start)

    fmt.Printf("Execution time: %s", elapsed)
}
