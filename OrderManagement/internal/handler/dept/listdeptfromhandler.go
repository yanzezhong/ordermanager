package dept

import (
	"net/http"

	"OrderManagement/OrderManagement/internal/logic/dept"
	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListDeptFromHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListDeptFromReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dept.NewListDeptFromLogic(r.Context(), svcCtx)
		resp, err := l.ListDeptFrom(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
