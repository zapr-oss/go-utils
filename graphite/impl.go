package graphite

import (
	"log"
	"net"
	"strconv"
	"time"
)

type connection interface {
	Close() error
	Write(p []byte) (int, error)
	connect() error
}

type tcpConnection struct {
	host string
	conn net.Conn
}

func newConnection(host string) *tcpConnection {
	c := new(tcpConnection)
	c.host = host

	return c
}

func (c *tcpConnection) Close() (err error) {
	if c.conn != nil {
		err = c.conn.Close()
		c.conn = nil
	}
	return
}

func (c *tcpConnection) Write(p []byte) (int, error) {
	if c.conn == nil {
		err := c.connect()
		if err != nil {
			log.Printf("Graphite.connect: %v", err)
			return 0, err
		}
	}

	c.conn.SetWriteDeadline(time.Now().Add(writeTimeout))
	return c.conn.Write(p)
}

func (c *tcpConnection) connect() (err error) {
	c.conn, err = net.DialTimeout("tcp", c.host, connectTimeout)
	if err != nil {
		c.conn = nil
	}
	return
}

func (gr *Graphite) fillBuffer(currentTime time.Time) {
	current_time := strconv.Itoa(int(currentTime.Unix()))

	if gr.buffer.Len() > maxBufSize {
		log.Printf("Graphite.sendMetrics: buffer size > %d. Reset buffer.", maxBufSize)
		gr.buffer.Reset()
	}

	gr.metricsMap.Range(func(key, val interface{}) bool {
		metric, ok := val.(*Counter)
		if ok {

			name := metric.GetName()
			value := metric.GetValue()

			gr.buffer.WriteString(gr.prefix)
			gr.buffer.WriteString(name + ".count")
			gr.buffer.WriteString(" ")
			gr.buffer.WriteString(strconv.FormatFloat(value, 'f', 12, 64))
			gr.buffer.WriteString(" ")
			gr.buffer.WriteString(current_time)
			gr.buffer.WriteString("\n")
		}
		return true
	})
}

func (gr *Graphite) sendMetrics(currentTime time.Time) {
	gr.fillBuffer(currentTime)

	if gr.buffer.Len() > 0 {
		_, err := gr.conn.Write(gr.buffer.Bytes())
		if err != nil {
			gr.conn.Close()
			return
		}
		gr.buffer.Reset()
	}
}

func (gr *Graphite) handleChans() {
	for {
		select {
		case t := <-gr.tickerChan:
			gr.sendMetrics(t)
		case _, _ = <-gr.stopChan:
			return
		}
	}
}
