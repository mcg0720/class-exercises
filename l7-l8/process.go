package main

type Process struct {
	clock []int
	pid   int
	N     int
}

func (p *Process) Start(N int, PID int) {

	var list []int

	p.pid = PID
	p.N = N
	p.clock = append(list, 0, 0, 0) //initalize the clock to [0, 0, 0]

}

func (p *Process) Internal() {
	p.clock[p.pid] = p.clock[p.pid] + 1 //increment the process posiiton by 1
}

func (p *Process) Send() []int {
	//p.clock[p.pid] = p.clock[p.pid] + 1
	return p.clock
}

func (p *Process) Receive(ts []int) {
	for n := range ts {
		if ts[n] > p.clock[n] {
			p.clock[n] = ts[n]
		}
	}
}

func Compare(ts1 []int, ts2 []int) int {

	var comp int

	if ts1[0] < ts2[0] { //find the first comparison to see if it holds throughout
		comp = -1
	} else if ts1[0] > ts2[0] {
		comp = 1
	} else {
		comp = 0
	}

	for n := range ts1 {
		if ts1[n] < ts2[n] {
			if comp == -1 {
				continue
			} else {
				comp = 0
				break
			}
		} else {
			if comp == 1 {
				continue
			} else {
				comp = 0
				break
			}
		}
	}
	return comp
}
