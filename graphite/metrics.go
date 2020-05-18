package graphite

type Counter struct {
	value float64
	name  string
}

func newCounter(name string) *Counter {
	return &Counter{0, name}
}

func (counter *Counter) GetValue() float64 {
	return counter.value
}

func (counter *Counter) GetName() string {
	return counter.name
}

func (counter *Counter) Inc() {
	counter.value += 1
}

type GraphiteMetric interface {
	getValue() float64
	getName() string
}