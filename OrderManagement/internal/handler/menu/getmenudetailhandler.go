package menu

import (
	"net/http"

	"OrderManagement/OrderManagement/internal/logic/menu"
	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetMenuDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMenuDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewGetMenuDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetMenuDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
