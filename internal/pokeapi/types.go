package pokeapi

type LocationArea struct {
	ID 		int 		`json:"id"`
	Name	string  `json:"name"`
	URL 	string  `json:"url"`
}

type LocationResponse struct {
	Count 		int							`json:"count"`
	Next 			*string					`json:"next"`
	Previous  *string 				`json:"previous"`
	Results   []LocationArea	`json:"results"`
}

type Config struct {
	NextURL 		string
	PreviousURL	string
}