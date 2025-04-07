package coverindex

// TODO: Append Where Clause on Statement limitOf,orderBy curosr page
// AND (created_at, id) < (lastCreateTime, lastId)
type CursorPage struct {
	LastAutoPkValueTxKey     string
	LastCreateTimeValueTxKey string
	LastUpdateTimeValueTxKey string

	UseOrderAutoIncPk  bool
	UseOrderCreateTime bool
	UseOrderUpdateTime bool
}
