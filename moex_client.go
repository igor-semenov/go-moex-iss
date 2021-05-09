package moex_iss

import "fmt"
import "go-moex-iss"
import "net/http"

func main() {
    client := &http.Client{}
    moex := moex_iss.NewClient(client)
    sber, err := moex.GetSecInfo("SBER")
    if err != nil {
        fmt.Printf("Error: %s", err.Error())
	return
    }
    if sber == nil {
	fmt.Print("SBER not found")
	return
    }
    fmt.Printf("SBER: %v", *sber)
}

