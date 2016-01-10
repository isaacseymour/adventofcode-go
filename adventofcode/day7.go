package adventofcode

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type wire struct {
	id        string
	signal    uint16
	filled    bool
	receivers []chan uint16
	lock      sync.Mutex
}

func (w *wire) Provide(signal uint16) {
	w.lock.Lock()
	if w.filled {
		fmt.Printf("Value provided twice to wire %s. Original %d, new %d\n", w.id, w.signal, signal)
		return
	}

	w.filled = true
	w.signal = signal
	for _, receiver := range w.receivers {
		receiver <- signal
	}
	w.lock.Unlock()
}

func (w *wire) Receive() chan uint16 {
	c := make(chan uint16, 1)
	w.lock.Lock()
	if w.filled {
		c <- w.signal
	} else {
		w.receivers = append(w.receivers, c)
	}
	w.lock.Unlock()
	return c
}

func inputGate(signal uint16, z *wire) {
	z.Provide(signal)
}

func funnelGate(in *wire, out *wire) {
	val := <-in.Receive()
	out.Provide(val)
}

func andGate(x *wire, y *wire, z *wire) {
	xval := <-x.Receive()
	yval := <-y.Receive()
	z.Provide(xval & yval)
}

func orGate(x *wire, y *wire, z *wire) {
	xval := <-x.Receive()
	yval := <-y.Receive()
	z.Provide(xval | yval)
}

func lshiftGate(x *wire, amount uint16, z *wire) {
	xval := <-x.Receive()
	z.Provide(xval << amount)
}

func rshiftGate(x *wire, amount uint16, z *wire) {
	xval := <-x.Receive()
	z.Provide(xval >> amount)
}

func notGate(x *wire, z *wire) {
	xval := <-x.Receive()
	z.Provide(^xval)
}

func ensureWire(wires map[string]*wire, id string) *wire {
	if signal, err := strconv.ParseUint(id, 10, 16); err == nil {
		// This is a constant-value wire
		return &wire{id: id, filled: true, signal: uint16(signal)}
	}

	w, ok := wires[id]

	if !ok {
		w = &wire{id: id, receivers: []chan uint16{}}
		wires[id] = w
	}

	return w
}

func applyRules(wires map[string]*wire, rules []string) {
	for _, line := range rules {
		sides := strings.Split(line, " -> ")

		if len(sides) != 2 {
			continue
		}

		gateString := sides[0]
		targetId := sides[1]

		target := ensureWire(wires, targetId)

		if andIds := strings.Split(gateString, " AND "); len(andIds) == 2 {
			x := ensureWire(wires, andIds[0])
			y := ensureWire(wires, andIds[1])
			go andGate(x, y, target)
		} else if orIds := strings.Split(gateString, " OR "); len(orIds) == 2 {
			x := ensureWire(wires, orIds[0])
			y := ensureWire(wires, orIds[1])
			go orGate(x, y, target)
		} else if lshiftParams := strings.Split(gateString, " LSHIFT "); len(lshiftParams) == 2 {
			x := ensureWire(wires, lshiftParams[0])
			amount, err := strconv.ParseUint(lshiftParams[1], 10, 16)
			if err != nil {
				panic(err)
			}
			go lshiftGate(x, uint16(amount), target)
		} else if rshiftParams := strings.Split(gateString, " RSHIFT "); len(rshiftParams) == 2 {
			x := ensureWire(wires, rshiftParams[0])
			amount, err := strconv.ParseUint(rshiftParams[1], 10, 16)
			if err != nil {
				panic(err)
			}
			go rshiftGate(x, uint16(amount), target)
		} else if len(gateString) > 4 && gateString[:4] == "NOT " {
			x := ensureWire(wires, gateString[4:])
			go notGate(x, target)
		} else {
			signal, err := strconv.ParseUint(gateString, 10, 16)
			if err != nil {
				inWire := ensureWire(wires, gateString)
				go funnelGate(inWire, target)
			} else {
				go inputGate(uint16(signal), target)
			}
		}

	}
}

func Day7(input string) (uint16, uint16) {
	rules := strings.Split(input, "\n")
	wires := make(map[string]*wire)

	applyRules(wires, rules)

	resultWire, ok := wires["a"]

	if !ok {
		panic("Input doesn't define gate a??")
	}

	result1 := <-resultWire.Receive()

	wires = make(map[string]*wire)
	b := ensureWire(wires, "b")
	b.Provide(result1)

	applyRules(wires, rules)

	resultWire, ok = wires["a"]

	if !ok {
		panic("Input doesn't define gate a??")
	}

	result2 := <-resultWire.Receive()
	return result1, result2
}
