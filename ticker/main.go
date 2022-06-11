package ticker

import "fmt"

func main() {

	appl := &stock{
		name: "AMZN",
		price: 100.0,
	}

	googl := &stock{
		name: "GOOGL",
		price: 120.0,
	}

	stockBroker := &broker{
		ID: "BR102",
		name: "zerodha",
		priceChart: map[string]float64{"AMZN": 163.4, "GOOGL": 2024},
	}

	if err := appl.Add(stockBroker); err != nil {
		fmt.Printf("could not add broker: %s", stockBroker.name)
	}

	if err := googl.Add(stockBroker); err != nil {
		fmt.Printf("could not add broker: %s", stockBroker.name)
	}

	appl.UpdatePrice(200.0)
	googl.UpdatePrice(2241)
}
