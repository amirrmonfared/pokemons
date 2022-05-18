package api

import (
	"net/http"

	db "github.com/amirrmonfared/pokemons/db/sqlc"
	"github.com/gin-gonic/gin"
)

type listPokemonRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=10"`
}

func (server *Server) listPokemons(ctx *gin.Context) {
	var req listPokemonRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
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
