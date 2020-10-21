package core

import "github.com/Elfsilon/proc_emulator/proc/gates"

func Mux4(x0, x1, x2, x3, d0, d1 bool) bool {
	notd0, notd1 := gates.NOT(d0), gates.NOT(d1)
	sig1 := gates.AND3(x0, d0, d1)
	sig2 := gates.AND3(x1, notd0, d1)
	sig3 := gates.AND3(x2, d0, notd1)
	sig4 := gates.AND3(x3, notd0, notd1)
	return gates.OR4(sig1, sig2, sig3, sig4)
}

func Demux4(x, d0, d1 bool) []bool {
	notd0, notd1 := gates.NOT(d0), gates.NOT(d1)
	sig1 := gates.AND3(x, d0, d1)
	sig2 := gates.AND3(x, notd0, d1)
	sig3 := gates.AND3(x, d0, notd1)
	sig4 := gates.AND3(x, notd0, notd1)
	return []bool{sig1, sig2, sig3, sig4}
}
