package pokeapi

type ListItemsResp struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type ItemDetails struct {
	Attributes []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"attributes"`
	BabyTriggerFor interface{} `json:"baby_trigger_for"`
	Category       struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"category"`
	Cost          int `json:"cost"`
	EffectEntries []struct {
		Effect   string `json:"effect"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		ShortEffect string `json:"short_effect"`
	} `json:"effect_entries"`
	FlavorTextEntries []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Text         string `json:"text"`
		VersionGroup struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version_group"`
	} `json:"flavor_text_entries"`
	FlingEffect interface{} `json:"fling_effect"`
	FlingPower  interface{} `json:"fling_power"`
	GameIndices []struct {
		GameIndex  int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"game_indices"`
	HeldByPokemon []interface{} `json:"held_by_pokemon"`
	ID            int           `json:"id"`
	Machines      []interface{} `json:"machines"`
	Name          string        `json:"name"`
	Names         []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Sprites struct {
		Default string `json:"default"`
	} `json:"sprites"`
}
