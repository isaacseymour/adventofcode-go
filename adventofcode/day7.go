package adventofcode

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

var (
	wires = make(map[string]wire)
)

type wire struct {
	id        string
	signal    uint16
	filled    bool
	receivers []chan uint16
	lock      sync.Mutex
}

func debug() {
	str := ""
	for id, w := range wires {
		str += fmt.Sprintf("%s has %d receivers, filled: %b, signal: %d\n", id, len(w.receivers), w.filled, w.signal)
	}
	fmt.Println(str)
}

func (w *wire) Provide(signal uint16) {
	w.lock.Lock()
	if w.filled {
		panic(fmt.Sprintf("Value provided twice to wire %s", w.id))
	}

	w.filled = true
	w.signal = signal
	fmt.Printf("Sending %d to %d receivers\n", signal, len(w.receivers))
	for _, receiver := range w.receivers {
		receiver <- signal
	}
	debug()
	w.lock.Unlock()
}

func (w *wire) Receive() chan uint16 {
	c := make(chan uint16, 1)
	w.lock.Lock()
	if w.filled {
		fmt.Printf("Receiving pre-filled value %d of %s\n", w.signal, w.id)
		c <- w.signal
	} else {
		fmt.Printf("Adding receiver to %s\n", w.id)
		w.receivers = append(w.receivers, c)
	}
	debug()
	w.lock.Unlock()
	return c
}

func inputGate(signal uint16, z *wire) {
	fmt.Printf("done: %d -> %s\n", signal, z.id)
	z.Provide(signal)
}

func funnelGate(in *wire, out *wire) {
	fmt.Printf("init: %s -> %s\n", in.id, out.id)
	val := <-in.Receive()
	fmt.Printf("done: %s -> %s\n", in.id, out.id)
	out.Provide(val)
}

func andGate(x *wire, y *wire, z *wire) {
	fmt.Printf("init: %s AND %s -> %s\n", x.id, y.id, z.id)
	xval := <-x.Receive()
	yval := <-y.Receive()
	fmt.Printf("done: %s AND %s -> %s\n", x.id, y.id, z.id)
	z.Provide(xval & yval)
}

func orGate(x *wire, y *wire, z *wire) {
	fmt.Printf("init: %s OR %s -> %s\n", x.id, y.id, z.id)
	xval := <-x.Receive()
	yval := <-y.Receive()
	fmt.Printf("done: %s OR %s -> %s\n", x.id, y.id, z.id)
	z.Provide(xval | yval)
}

func lshiftGate(x *wire, amount uint16, z *wire) {
	fmt.Printf("init: %s LSHIFT %d -> %s\n", x.id, amount, z.id)
	xval := <-x.Receive()
	fmt.Printf("done: %s LSHIFT %d -> %s\n", x.id, amount, z.id)
	z.Provide(xval << amount)
}

func rshiftGate(x *wire, amount uint16, z *wire) {
	fmt.Printf("init: %s RSHIFT %d -> %s\n", x.id, amount, z.id)
	xval := <-x.Receive()
	fmt.Printf("done: %s RSHIFT %d -> %s\n", x.id, amount, z.id)
	z.Provide(xval >> amount)
}

func notGate(x *wire, z *wire) {
	fmt.Printf("init: NOT %s -> %s\n", x.id, z.id)
	xval := <-x.Receive()
	fmt.Printf("done: NOT %s -> %s\n", x.id, z.id)
	z.Provide(^xval)
}

func ensureWire(wires map[string]wire, id string) *wire {
	if signal, err := strconv.ParseUint(id, 10, 16); err == nil {
		// This is a constant-value wire
		return &wire{id: id, filled: true, signal: uint16(signal)}
	}

	w, ok := wires[id]

	if !ok {
		w = wire{id: id, receivers: []chan uint16{}}
		wires[id] = w
	}

	return &w
}

func Day7(input string) uint16 {
	for _, line := range strings.Split(input, "\n") {
		// fmt.Printf("processing line: %s\n", line)
		sides := strings.Split(line, " -> ")

		if len(sides) != 2 {
			continue
		}

		gateString := sides[0]
		targetId := sides[1]

		target := ensureWire(wires, targetId)

		if andIds := strings.Split(gateString, " AND "); len(andIds) == 2 {
			// fmt.Printf("found: %s AND %s -> %s\n", andIds[0], andIds[1], targetId)
			x := ensureWire(wires, andIds[0])
			y := ensureWire(wires, andIds[1])
			go andGate(x, y, target)
		} else if orIds := strings.Split(gateString, " OR "); len(orIds) == 2 {
			// fmt.Printf("found: %s OR %s -> %s\n", orIds[0], orIds[1], targetId)
			x := ensureWire(wires, orIds[0])
			y := ensureWire(wires, orIds[1])
			go orGate(x, y, target)
		} else if lshiftParams := strings.Split(gateString, " LSHIFT "); len(lshiftParams) == 2 {
			// fmt.Printf("found: %s LSHIFT %s -> %s\n", lshiftParams[0], lshiftParams[1], targetId)
			x := ensureWire(wires, lshiftParams[0])
			amount, err := strconv.ParseUint(lshiftParams[1], 10, 16)
			if err != nil {
				panic(err)
			}
			go lshiftGate(x, uint16(amount), target)
		} else if rshiftParams := strings.Split(gateString, " RSHIFT "); len(rshiftParams) == 2 {
			// fmt.Printf("found: %s RSHIFT %s -> %s\n", rshiftParams[0], rshiftParams[1], targetId)
			x := ensureWire(wires, rshiftParams[0])
			amount, err := strconv.ParseUint(rshiftParams[1], 10, 16)
			if err != nil {
				panic(err)
			}
			go rshiftGate(x, uint16(amount), target)
		} else if len(gateString) > 4 && gateString[:4] == "NOT " {
			// fmt.Printf("found: NOT %s -> %s\n", gateString[4:], targetId)
			x := ensureWire(wires, gateString[4:])
			go notGate(x, target)
		} else {
			signal, err := strconv.ParseUint(gateString, 10, 16)
			if err != nil {
				// fmt.Printf("found: (funnel) %s -> %s\n", gateString, targetId)
				inWire := ensureWire(wires, gateString)
				go funnelGate(inWire, target)
			} else {
				// fmt.Printf("found: (input) %s -> %s\n", gateString, targetId)
				go inputGate(uint16(signal), target)
			}
		}

	}

	resultWire, ok := wires["a"]

	if !ok {
		panic("Input doesn't define gate a??")
	}

	result := <-resultWire.Receive()
	return result
}
