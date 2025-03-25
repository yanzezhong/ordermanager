package menu

import (
	"net/http"

	"OrderManagement/OrderManagement/internal/logic/menu"
	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetMenuRoutesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMenuRoutesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewGetMenuRoutesLogic(r.Context(), svcCtx)
		resp, err := l.GetMenuRoutes(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
