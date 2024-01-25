package cyoa

import "encoding/json"

type Adventure map[string]Arc

type Arc struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	NextArc string `json:"arc"`
}

func GetAdventure(data []byte) (Adventure, error) {
	a := Adventure{}
	err := json.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}
	return a, nil
}
