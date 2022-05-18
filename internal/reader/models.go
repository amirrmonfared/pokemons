package reader

type Pokemons struct {
	Id         int32  `csv:"#"`
	Name       string `csv:"Name"`
	Type1      string `csv:"Type 1"`
	Type2      string `csv:"Type 2"`
	Total      int32  `csv:"Total"`
	HP         int32  `csv:"HP"`
	Attack     int32  `csv:"Attack"`
	Defense    int32  `csv:"Defense"`
	SpAtk      int32  `csv:"Sp. Atk"`
	SpDef      int32  `csv:"Sp. Def"`
	Speed      int32  `csv:"Speed"`
	Generation int32  `csv:"Generation"`
	Legendary  bool   `csv:"Legendary"`
}