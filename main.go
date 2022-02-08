package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Lookup struct {
	PriceRoute struct {
		BlockNumber  int    `json:"blockNumber"`
		Network      int    `json:"network"`
		SrcToken     string `json:"srcToken"`
		SrcDecimals  int    `json:"srcDecimals"`
		SrcAmount    string `json:"srcAmount"`
		DestToken    string `json:"destToken"`
		DestDecimals int    `json:"destDecimals"`
		DestAmount   string `json:"destAmount"`
		BestRoute    []struct {
			Percent int `json:"percent"`
			Swaps   []struct {
				SrcToken      string `json:"srcToken"`
				SrcDecimals   int    `json:"srcDecimals"`
				DestToken     string `json:"destToken"`
				DestDecimals  int    `json:"destDecimals"`
				SwapExchanges []struct {
					Exchange      string   `json:"exchange"`
					SrcAmount     string   `json:"srcAmount"`
					DestAmount    string   `json:"destAmount"`
					Percent       int      `json:"percent"`
					PoolAddresses []string `json:"poolAddresses"`
					Data          struct {
						Version int    `json:"version"`
						GasUSD  string `json:"gasUSD"`
					} `json:"data"`
				} `json:"swapExchanges"`
			} `json:"swaps"`
		} `json:"bestRoute"`
		GasCostUSD         string `json:"gasCostUSD"`
		GasCost            string `json:"gasCost"`
		Side               string `json:"side"`
		TokenTransferProxy string `json:"tokenTransferProxy"`
		ContractAddress    string `json:"contractAddress"`
		ContractMethod     string `json:"contractMethod"`
		PartnerFee         int    `json:"partnerFee"`
		SrcUSD             string `json:"srcUSD"`
		DestUSD            string `json:"destUSD"`
		Partner            string `json:"partner"`
		MaxImpactReached   bool   `json:"maxImpactReached"`
		Hmac               string `json:"hmac"`
	} `json:"priceRoute"`
}

func main() {
	// ============ Network ============
	fmt.Print("Network ID (Mainnet - 1, Ropsten - 3, Polygon - 137, BSC - 56, Avalanche - 43114): ")
	var network string
	_, nwkEx := fmt.Scanln(&network)
	if nwkEx != nil {
		log.Fatal(nwkEx)
	}

	// ============ Src Token ============
	fmt.Print("Address of the first token: ")
	var srcToken string
	_, srcEx := fmt.Scanln(&srcToken)
	if srcEx != nil {
		log.Fatal(srcEx)
	}

	// ============ Src Decimals ============
	fmt.Print("Decimals of the first token: ")
	var srcDecimals string
	_, srcDecEx := fmt.Scanln(&srcDecimals)
	if srcDecEx != nil {
		log.Fatal(srcDecEx)
	}

	// ============ Dest Token ============
	fmt.Print("Address of the second token: ")
	var destToken string
	_, destEx := fmt.Scanln(&destToken)
	if destEx != nil {
		log.Fatal(destEx)
	}

	// ============ Dest Decimals ============
	fmt.Print("Decimals of the second token: ")
	var destDecimals string
	_, destDecEx := fmt.Scanln(&destDecimals)
	if destDecEx != nil {
		log.Fatal(destDecEx)
	}

	// =============================
	const amount = "1000000000000000000"
	url := "https://apiv5.paraswap.io/prices/?srcToken=" + srcToken + "&destToken=" + destToken + "&amount=" + amount + "&srcDecimals=" + srcDecimals + "&destDecimals=" + destDecimals + "&side=SELL&network=" + network

	req, reqEx := http.NewRequest(http.MethodGet, url, nil)
	if reqEx != nil {
		log.Fatal(reqEx)
	}

	res, resEx := http.DefaultClient.Do(req)
	if resEx != nil {
		log.Fatal(resEx)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readEx := ioutil.ReadAll(res.Body)
	if readEx != nil {
		log.Fatal(readEx)
	}

	lookup := Lookup{}
	jsonEx := json.Unmarshal(body, &lookup)
	if jsonEx != nil {
		log.Fatal(jsonEx)
	}

	// https://apiv5.paraswap.io/prices/
	// ?srcToken=0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE
	// &destToken=0x6b175474e89094c44da98b954eedeac495271d0f
	// &amount=10000000000000000000
	// &srcDecimals=18
	// &destDecimals=18
	// &side=SELL
	// &network=1
	fmt.Println(lookup.PriceRoute.BlockNumber)
}
