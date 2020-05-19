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

// Increase count given
func (counter *Counter) IncCount(count int) {
	counter.value += float64(count)
}

// Set the count value
func (counter *Counter) SetValue(count int) {
	counter.value = float64(count)
}

type GraphiteMetric interface {
	getValue() float64
	getName() string
}
