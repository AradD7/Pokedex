package pokeapi


type Encounters struct {
	PokemonEncounters 	[]PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon        Pokemon          `json:"pokemon"`
	VersionDetails []VersionDetails `json:"version_details"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type VersionDetails struct {
	EncounterDetails []EncounterDetails `json:"encounter_details"`
	MaxChance        int                `json:"max_chance"`
	Version          Version            `json:"version"`
}

type EncounterDetails struct {
	Chance          int    `json:"chance"`
	ConditionValues []any  `json:"condition_values"`
	MaxLevel        int    `json:"max_level"`
	Method          Method `json:"method"`
	MinLevel        int    `json:"min_level"`
}

type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Method struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
