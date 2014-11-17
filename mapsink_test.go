package mapsink_test

import (
	"testing"

	"github.com/bakins/go-metrics-map"
	h "github.com/bakins/test-helpers"
)

func TestNew(t *testing.T) {
	m := mapsink.New()
	h.Assert(t, m != nil, "sink is nil")
}

func TestSetGauge(t *testing.T) {
	m := mapsink.New()

	m.SetGauge([]string{"foo", "bar", "baz"}, 100)
	v, ok := m.Get("foo.bar.baz")
	h.Assert(t, ok, "value not found")
	h.Assert(t, 100 == v, "value is not 100")

	m.SetGauge([]string{"foo", "bar", "baz"}, 200)
	v, ok = m.Get("foo.bar.baz")
	h.Assert(t, ok, "value not found")
	h.Assert(t, 200 == v, "value is not 200")
}

func TestEmitKey(t *testing.T) {
	m := mapsink.New()

	m.EmitKey([]string{"foo", "bar", "baz"}, 100)
	v, ok := m.Get("foo.bar.baz")
	h.Assert(t, ok, "value not found")
	h.Assert(t, 100 == v, "value is not 100")
}

func TestIncrCounter(t *testing.T) {
	m := mapsink.New()

	key := []string{"foo", "bar", "baz"}
	m.SetGauge(key, 100)

	m.IncrCounter(key, 5)

	v, ok := m.Get(m.FlattenKey(key))
	h.Assert(t, ok, "value not found")
	h.Assert(t, 105 == v, "value is not 105")

}
