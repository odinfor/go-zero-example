package handler

import (
	"net/http"

	"go-zero-example/service/search/cmd/api/internal/logic"
	"go-zero-example/service/search/cmd/api/internal/svc"
	"go-zero-example/service/search/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func searchHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSearchLogic(r.Context(), ctx)
		resp, err := l.Search(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
