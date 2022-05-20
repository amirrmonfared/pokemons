package api

type userRequest struct {
	Name     string `form:"name"`
	PageID   int32  `form:"page_id"`
	PageSize int32  `form:"page_size"`
	Hp       int32  `form:"hp"`
	Attack   int32  `form:"attack"`
	Defense  int32  `form:"defense"`
}

type listPokemonRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"max=10"`
}

type listPokemonByNameRequest struct {
	Name string `form:"name" binding:"required"`
}

type listPokemonByAbilityRequest struct {
	Hp       int32  `form:"hp"`
	Attack   int32  `form:"attack"`
	Defense  int32  `form:"defense"`
}
