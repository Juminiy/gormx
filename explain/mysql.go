package explain

type MySQLPlan struct {
	ID           int64
	SelectType   string
	Table        string
	Partitions   string
	Type         string
	PossibleKeys string
	Key          string
	KeyLen       int64
	Ref          string
	Rows         int64
	Filtered     int64
	Extra        string
}
