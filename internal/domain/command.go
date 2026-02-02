package domain

type Command struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Usage       string `json:"usage"`
	HelpText    string `json:"helpText"`
}
