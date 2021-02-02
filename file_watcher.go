package main
import (
    "io/ioutil"
    "fmt"
    "os"
    "time"
    "strings"
    "strconv"
    "encoding/csv"
)


const watchedpath = "./source"

func main() {
    for {
        d, _ := os.Open(watchedpath)
        files, _ := d.Readdir(-1)
        for _, fi := range files {
            filePath := watchedpath + "/" + fi.Name()
            f, _ := os.Open(filePath)
             data, _ := ioutil.ReadAll(f)
             f.Close()
             os.Remove(filePath)

             go func(data string) {
                reader := csv.NewReader(strings.NewReader(data))
                records , _ := reader.ReadAll()

                for _, r := range records {
                    invoice := new(Invoice)
                    invoice.Number = r[0]
                    invoice.Amount, _ = strconv.ParseFloat(r[1], 64)
                    invoice.PurchaseOrderNumber, _ = strconv.Atoi(r[2])
                    unixtime, _ := strconv.ParseInt(r[3], 10, 64)

                    invoice.InvoiceDate = time.Unix(unixtime, 0)

                    fmt.Printf("Received Invoice '%v' for $%.2f and submitted for processing\n", invoice.Number, invoice.Amount)
                }
             }(string(data))
        }

        d.Close()
        time.Sleep(100 * time.Millisecond)
    }
}

type Invoice struct {
    Number string
    Amount float64
    PurchaseOrderNumber int
    InvoiceDate time.Time
}


