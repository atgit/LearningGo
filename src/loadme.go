/*
- Support http verbs
- Fully saturate the bandwidth or cpu
- Configurable current connections, payload, duration
- Time of execution
*/
package main

import "net/http"
import "io/ioutil"
import "time"
import "fmt"
import "flag"

func main() {
    var purl = flag.String("url", "http://google.com", "url")
    var currentConnections int = 10
    // var sizeOfPayload int = 1024 // Bytes
    // var duration int = 10        // Seconds
    var pverb = flag.String("verb", "get", "HTTP verb")
    var poutput = flag.Bool("output", false, "Output http response")

    var high, low, total time.Duration

    flag.Parse()
    fmt.Println(*purl)
    
    for i := 0; i < currentConnections; i++ {
        var resp *http.Response
        var err  error

        switch *pverb {
        case "get":
            var start = time.Now()
            resp, err = http.Get(*purl)
            elapsed := time.Since(start)
            if high < elapsed {
                high = elapsed
            }
            if low == 0 {
                low = elapsed
            } else if low > elapsed {
                low = elapsed
            }
            total += elapsed
            
        default:
            fmt.Println("Unsupport http verb!")
            return 
        }
        if err != nil {
            fmt.Println("Error occured ...")
            return
        }
        defer resp.Body.Close()
        if *poutput {
            body, _ := ioutil.ReadAll(resp.Body)
            fmt.Println(body)
        }
    }
    fmt.Println("Low: ", low)
    fmt.Println("High: ", high)
    fmt.Println("Avg: ", time.Duration(float64(total)/float64(currentConnections)))
}
