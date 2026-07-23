package tcgapi

type Card struct {
	Category       string     `json:"category"`
	ID             string     `json:"id"`
	Illustrator    string     `json:"illustrator"`
	Image          string     `json:"image"`
	LocalID        string     `json:"localId"`
	Name           string     `json:"name"`
	Rarity         string     `json:"rarity"`
	Set            Set        `json:"set"`
	Variants       Variants   `json:"variants"`
	HP             int        `json:"hp"`
	Types          []string   `json:"types"`
	EvolveFrom     string     `json:"evolveFrom"`
	Description    string     `json:"description"`
	Stage          string     `json:"stage"`
	Attacks        []Attack   `json:"attacks"`
	Weaknesses     []Weakness `json:"weaknesses"`
	Retreat        int        `json:"retreat"`
	RegulationMark string     `json:"regulationMark"`
	Legal          Legal      `json:"legal"`
}

type Set struct {
	CardCount CardCount `json:"cardCount"`
	ID        string    `json:"id"`
	Logo      string    `json:"logo"`
	Name      string    `json:"name"`
	Symbol    string    `json:"symbol"`
}

type CardCount struct {
	Official int `json:"official"`
	Total    int `json:"total"`
}

type Variants struct {
	FirstEdition bool `json:"firstEdition"`
	Holo         bool `json:"holo"`
	Normal       bool `json:"normal"`
	Reverse      bool `json:"reverse"`
	WPromo       bool `json:"wPromo"`
}

type Attack struct {
	Cost   []string `json:"cost"`
	Name   string   `json:"name"`
	Effect string   `json:"effect"`
	Damage int      `json:"damage,omitempty"`
}

type Weakness struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Legal struct {
	Standard bool `json:"standard"`
	Expanded bool `json:"expanded"`
}
