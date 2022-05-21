package api

import (
	"net/http"

	db "github.com/amirrmonfared/pokemons/db/sqlc"
	"github.com/gin-gonic/gin"
)

// broker checking for the request and send the request to right handler
func (server *Server) broker(ctx *gin.Context) {
	var req userRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if req.Name != "" { // in case of filter pokemon by name
		server.getPokemonByName(ctx)
	}
	if req.PageID != 0 { // in case of listing pokemons
		server.listPokemons(ctx)
	}
	if req.Hp != 0 || req.Attack != 0 || req.Defense != 0 { 
		server.getPokemonByAbility(ctx)
	}
}

func (server *Server) listPokemons(ctx *gin.Context) {
	var req listPokemonRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if req.PageSize == 0 {
		req.PageSize = 10
	}

	arg := db.ListPokemonsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	pokemons, err := server.store.ListPokemons(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, pokemons)
}

func (server *Server) getPokemonByName(ctx *gin.Context) {
	var req listPokemonByNameRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	pokemons, err := server.store.GetPokemonByName(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, pokemons)
}

func (server *Server) getPokemonByAbility(ctx *gin.Context) {
	var req listPokemonByAbilityRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListPokemonsByAbilityParams{
		Hp:      req.Hp,
		Defense: req.Defense,
		Attack:  req.Attack,
	}

	pokemons, err := server.store.ListPokemonsByAbility(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, pokemons)
}
