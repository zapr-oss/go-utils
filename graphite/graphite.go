// Package graphite implements a graphite client with aggregation of metrics in a short period of time and sending the result to graphite.
// The package was written for cases when an application is running on thousands of instances and each of the instances generates hundreds of thousands of events per second.
//
package graphite

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"sync"
	"time"
)

const (
	connectTimeout = 400 * time.Millisecond
	writeTimeout   = 1 * time.Second
	maxBufSize     = 5 * 1 << 20 // 5 MiB
)

var defaultGraphite *Graphite

func InitDefaultGraphite(graphiteConfig GraphiteConfig) {
	dG, err := NewGraphiteUsingConfig(graphiteConfig)

	if err != nil {
		log.Fatal("Error in initializing default graphite")
	}

	defaultGraphite = dG
}

func Start() error {
	return defaultGraphite.Start()
}

func Stop() error {
	return defaultGraphite.Stop()
}

func GetCounter(name string) (*Counter) {

	if defaultGraphite == nil {
		log.Println("Error in creating counter. Call InitDefaultGraphite func to initialize default graphite")
		return newCounter(name)
	}

	if err := defaultGraphite.isEverythingOk(); err != nil {
		log.Println("Warn: graphite not started. Call Start() before getting counter")
	}

	counter := newCounter(name)

	if val, ok := defaultGraphite.metricsMap.Load(name); ok {
		counter, ok = val.(*Counter)
		if ok {
			return counter
		} else {
			log.Fatalf("Metric with same NAME and different TYPE exists, Name: %v, Type: %v\n", name, reflect.ValueOf(val).Type().String())
		}
	}

	defaultGraphite.metricsMap.Store(name, counter)

	return counter
}

// Graphite encapsulates an API that allows you to handle metric values and send them to graphite.
// Multiple goroutines may invoke methods on a Graphite simultaneously.
type Graphite struct {
	host          string
	prefix        string
	flushInterval time.Duration

	conn       connection
	ticker     *time.Ticker
	tickerChan <-chan time.Time

	metricsMap sync.Map

	stopChan chan struct{}

	buffer   bytes.Buffer
	disabled bool
	started  bool
}

// NewGraphite creates a new Graphite with connection to host:port. Application only needs one Graphite to work with all metrics.
// The prefix parameter is assigned to each metric name when it is sent to the graphite server.
// Graphite sends aggregated metrics to the server each flushInterval period. The flushInterval can't be less than a one second.
// Sending metrics to the server is easy to disable from the application config without changing the code. Use the disabled option to do this.
func NewGraphite(host string, port uint16, environment string, prefix string, flushInterval time.Duration, disabled bool) (*Graphite, error) {
	if disabled == true {
		graph := new(Graphite)
		graph.disabled = true
		return graph, nil
	}

	if flushInterval < time.Second {
		return nil, fmt.Errorf("NewGraphite: Flush interval (%v) < 1s", flushInterval)
	}
	graph := new(Graphite)
	graph.host = host + ":" + strconv.Itoa(int(port))
	graph.conn = newConnection(graph.host)
	if len(prefix) > 0 {
		graph.prefix = environment + "." + prefix
		if graph.prefix[len(graph.prefix)-1] != '.' {
			graph.prefix = graph.prefix + "."
		}
	}
	graph.flushInterval = flushInterval

	graph.metricsMap = sync.Map{}
	return graph, nil
}

func NewGraphiteUsingConfig(graphiteConfig GraphiteConfig) (*Graphite, error) {
	if graphiteConfig.Disabled == true {
		graph := new(Graphite)
		graph.disabled = true
		return graph, nil
	}

	if time.Duration(graphiteConfig.FlushIntervalInSec)*time.Second < time.Second {
		return nil, fmt.Errorf("NewGraphite: Flush interval (%v) < 1s", graphiteConfig.FlushIntervalInSec)
	}
	graph := new(Graphite)
	graph.host = graphiteConfig.Host + ":" + strconv.Itoa(int(graphiteConfig.Port))
	graph.conn = newConnection(graph.host)
	if len(graphiteConfig.Prefix) > 0 {
		graph.prefix = graphiteConfig.Environment + "." + graphiteConfig.Prefix
		if graph.prefix[len(graph.prefix)-1] != '.' {
			graph.prefix = graph.prefix + "."
		}
	}
	graph.flushInterval = time.Duration(graphiteConfig.FlushIntervalInSec) * time.Second

	graph.metricsMap = sync.Map{}
	return graph, nil
}

// Start creates a goroutine, which sends the aggregated metrics to graphite.
// Start should be called once when the application is initialized as soon as all metrics are registered with functions Register*
func (graphite *Graphite) Start() error {
	if graphite == nil {
		return fmt.Errorf("Start: Call NewGraphite() before Start()")
	}

	if graphite.disabled == true {
		return nil
	}

	if graphite.started == true {
		return fmt.Errorf("Graphite already started")
	}

	graphite.stopChan = make(chan struct{})
	graphite.ticker = time.NewTicker(graphite.flushInterval)
	graphite.tickerChan = graphite.ticker.C
	go graphite.handleChans()
	graphite.started = true

	return nil
}

// Stop completes the goroutine of sending metrics to graphite. Typically, in a real application, this is not required, but only for tests.
func (graphite *Graphite) Stop() error {
	if graphite == nil {
		return fmt.Errorf("Stop: Call NewGraphite() before Stop()")
	}

	if graphite.disabled == true {
		return nil
	}

	if graphite.started != true {
		return fmt.Errorf("Stop: Call Start() before Stop()")
	}

	close(graphite.stopChan)

	return nil
}

func (graphite *Graphite) isEverythingOk() error {
	if graphite == nil {
		return fmt.Errorf("Call NewGraphite() before using any metric")
	}

	if graphite.disabled == true {
		return nil
	}

	if graphite.started != true {
		return fmt.Errorf("Call Start() before using any metric")
	}

	return nil
}

func (graphite *Graphite) GetCounter(name string) (*Counter) {

	if err := graphite.isEverythingOk(); err != nil {
		log.Println("Warn: graphite not started. Call Start() before getting counter")
	}

	counter := newCounter(name)

	if val, ok := graphite.metricsMap.Load(name); ok {
		counter, ok = val.(*Counter)
		if ok {
			return counter
		} else {
			log.Printf("Metric with same NAME and different TYPE exists, Name: %v, Type: %v\n", name, reflect.ValueOf(val).Type().String())
			return counter
		}
	}

	graphite.metricsMap.Store(name, counter)

	return counter
}

type GraphiteConfig struct {
	Host               string `json:"host"`
	Port               uint16 `json:"port"`
	Prefix             string `json:"prefix"`
	Environment        string `json:"environment"`
	FlushIntervalInSec int64  `json:"flushIntervalInSec"`
	Disabled           bool   `json:"disabled"`
}
