package resources

type Kit struct {
	Id          int    `json:"id"`
	Grade_id    int    `json:"grade_id"`
	Grade       string `json:"grade"`
	Name        string `json:"name"`
	Series      string `json:"series"`
	Price       int    `json:"price"`
	Release     string `json:"release"`
	Description string `json:"description"`
}
