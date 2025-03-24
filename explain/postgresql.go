package explain

type PgExplainFormatJSON struct {
	NodeType     string  `json:"Node Type"`
	RelationName string  `json:"Relation Name"`
	Alias        string  `json:"Alias"`
	StartupCost  float64 `json:"Startup Cost"`
	TotalCost    float64 `json:"Total Cost"`
	PlanRows     int64   `json:"Plan Rows"`
	PlanWidth    int64   `json:"Plan Width"`
	Filter       string  `json:"Filter"`
}
