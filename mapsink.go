// Package mapsink provides a simple map-based sink for go-metrics.  It does
// not support intervals, only totals since creation.
package mapsink

import (
	"encoding/json"
	"strings"
	"sync"
)

// MapSink is a simple sink that stores metrics in a map.
type MapSink struct {
	sync.RWMutex
	Data map[string]float32
}

// New initializes a new MapSink
func New() *MapSink {
	s := &MapSink{
		Data: make(map[string]float32),
	}
	return s
}

// FlattenKey flattens the key for formatting, removes spaces
func (s *MapSink) FlattenKey(parts []string) string {
	joined := strings.Join(parts, ".")
	return strings.Map(func(r rune) rune {
		switch r {
		case ':':
			fallthrough
		case ' ':
			return '_'
		default:
			return r
		}
	}, joined)
}

// SetGauge sets the flattend key to the value.
func (s *MapSink) SetGauge(key []string, val float32) {
	flatKey := s.FlattenKey(key)
	s.Lock()
	s.Data[flatKey] = val
	s.Unlock()
}

// EmitKey sets the flattend key to the value.
func (s *MapSink) EmitKey(key []string, val float32) {
	flatKey := s.FlattenKey(key)
	s.Lock()
	s.Data[flatKey] = val
	s.Unlock()
}

// IncrCounter increments the flattened key by val. If the key does not currently, exist
// then it is set to val
func (s *MapSink) IncrCounter(key []string, val float32) {
	flatKey := s.FlattenKey(key)
	s.Lock()
	s.Data[flatKey] += val
	s.Unlock()
}

// AddSample does nothing.  Timings are not supported.
func (s *MapSink) AddSample(key []string, val float32) {

}

// MarshalJSON marshals the sinks data into json
func (s *MapSink) MarshalJSON() ([]byte, error) {
	s.Lock()
	defer s.Unlock()
	return json.Marshal(s.Data)
}

// Get retrieves a single value.
func (s *MapSink) Get(key string) (float32, bool) {
	s.Lock()
	defer s.Unlock()
	val, ok := s.Data[key]
	return val, ok
}
