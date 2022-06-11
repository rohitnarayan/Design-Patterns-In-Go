package ticker

import "errors"

type stock struct {
	name    string
	price   float64
	brokers []*broker
}

type Stock interface {
	Add(b *broker) error
	Remove(b *broker) error
	Notify()
}

func NewStock(name string, listPrice float64) Stock {
	return &stock{
		name:  name,
		price: listPrice,
	}
}

func (s *stock) Add(b *broker) error {
	for _, a := range s.brokers {
		if a.ID == b.ID {
			return errors.New("broker already added as watcher")
		}
	}

	s.brokers = append(s.brokers, b)
	return nil
}

func (s *stock) Remove(b *broker) error {
	for i, a := range s.brokers {
		if a.ID == b.ID {
			s.brokers = append(s.brokers[0:i], s.brokers[i+1:]...)
			return nil
		}
	}

	return errors.New("broker not found")
}

func (s *stock) Notify() {
	for _, b := range s.brokers {
		b.Update(s);
	}
}

func (s *stock) UpdatePrice(price float64) {
	s.price = price
	s.Notify()
}
