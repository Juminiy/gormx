package explain

type SqliteVMInstruction struct {
	Addr    int64
	OpCode  string
	P1      int64
	P2      int64
	P3      int64
	P4      int64
	P5      int64
	Comment string
}
