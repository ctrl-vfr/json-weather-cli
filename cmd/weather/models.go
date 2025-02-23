package weather

type Response struct {
	Date        string `json:"date"`
	Temperature int    `json:"temperature"`
	Weather     string `json:"weather"`
	Cycle       string `json:"cycle"`
}
