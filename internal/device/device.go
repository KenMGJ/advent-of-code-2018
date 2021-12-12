package device

var Ops = [...]string{
	"addr", "addi", "mulr", "muli", "banr", "bani", "borr", "bori",
	"setr", "seti", "gtir", "gtri", "gtrr", "eqir", "eqri", "eqrr",
}

type OpAndArgs struct {
	Op int
	A  int
	B  int
	C  int
}

type Device struct {
	Instr []OpAndArgs
	R     []int

	opMap map[int]string
}

func NewDevice(lines []OpAndArgs, registerCount int) *Device {

	opMap := make(map[int]string)
	opMap[0] = "eqri"
	opMap[1] = "bori"
	opMap[2] = "addi"
	opMap[3] = "bani"
	opMap[4] = "seti"
	opMap[5] = "eqrr"
	opMap[6] = "addr"
	opMap[7] = "gtri"
	opMap[8] = "borr"
	opMap[9] = "gtir"
	opMap[10] = "setr"
	opMap[11] = "eqir"
	opMap[12] = "mulr"
	opMap[13] = "muli"
	opMap[14] = "gtrr"
	opMap[15] = "banr"

	r := make([]int, registerCount)
	return &Device{
		Instr: lines,
		R:     r,
		opMap: opMap,
	}
}

func (d *Device) Run() {
	for _, o := range d.Instr {
		d.Exec(o.Op, o.A, o.B, o.C)
	}
}

func (d *Device) Exec(opcode, a, b, c int) {
	d.ExecOpName(d.opMap[opcode], a, b, c)
}

func (d *Device) ExecOpName(op string, a, b, c int) {
	switch op {
	case "addr":
		d.Addr(a, b, c)
	case "addi":
		d.Addi(a, b, c)
	case "mulr":
		d.Mulr(a, b, c)
	case "muli":
		d.Muli(a, b, c)
	case "banr":
		d.Banr(a, b, c)
	case "bani":
		d.Bani(a, b, c)
	case "borr":
		d.Borr(a, b, c)
	case "bori":
		d.Bori(a, b, c)
	case "setr":
		d.Setr(a, b, c)
	case "seti":
		d.Seti(a, b, c)
	case "gtir":
		d.Gtir(a, b, c)
	case "gtri":
		d.Gtri(a, b, c)
	case "gtrr":
		d.Gtrr(a, b, c)
	case "eqir":
		d.Eqir(a, b, c)
	case "eqri":
		d.Eqri(a, b, c)
	case "eqrr":
		d.Eqrr(a, b, c)
	}
}

func (d *Device) Addr(a, b, c int) {
	d.R[c] = d.R[a] + d.R[b]
}

func (d *Device) Addi(a, b, c int) {
	d.R[c] = d.R[a] + b
}

func (d *Device) Mulr(a, b, c int) {
	d.R[c] = d.R[a] * d.R[b]
}

func (d *Device) Muli(a, b, c int) {
	d.R[c] = d.R[a] * b
}

func (d *Device) Banr(a, b, c int) {
	d.R[c] = d.R[a] & d.R[b]
}

func (d *Device) Bani(a, b, c int) {
	d.R[c] = d.R[a] & b
}

func (d *Device) Borr(a, b, c int) {
	d.R[c] = d.R[a] | d.R[b]
}

func (d *Device) Bori(a, b, c int) {
	d.R[c] = d.R[a] | b
}

func (d *Device) Setr(a, b, c int) {
	d.R[c] = d.R[a]
}

func (d *Device) Seti(a, b, c int) {
	d.R[c] = a
}

func (d *Device) Gtir(a, b, c int) {
	if a > d.R[b] {
		d.R[c] = 1
	} else {
		d.R[c] = 0
	}
}

func (d *Device) Gtri(a, b, c int) {
	if d.R[a] > b {
		d.R[c] = 1
	} else {
		d.R[c] = 0
	}
}

func (d *Device) Gtrr(a, b, c int) {
	if d.R[a] > d.R[b] {
		d.R[c] = 1
	} else {
		d.R[c] = 0
	}
}

func (d *Device) Eqir(a, b, c int) {
	if a == d.R[b] {
		d.R[c] = 1
	} else {
		d.R[c] = 0
	}
}

func (d *Device) Eqri(a, b, c int) {
	if d.R[a] == b {
		d.R[c] = 1
	} else {
		d.R[c] = 0
	}
}

func (d *Device) Eqrr(a, b, c int) {
	if d.R[a] == d.R[b] {
		d.R[c] = 1
	} else {
		d.R[c] = 0
	}
}
