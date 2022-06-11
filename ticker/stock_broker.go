package ticker

type broker struct {
	ID         string
	name       string
	priceChart map[string]float64
}

type Broker interface {
	Update(s *stock)
}

func NewBroker(p map[string]float64) Broker {
	return &broker{
		priceChart: p,
	}
}

func (b *broker) Update(s *stock) {
	if _, ok := b.priceChart[s.name]; ok {
		b.priceChart[s.name] = s.price
	}
}
