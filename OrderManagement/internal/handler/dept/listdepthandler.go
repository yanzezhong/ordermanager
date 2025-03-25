package dept

import (
	"net/http"

	"OrderManagement/OrderManagement/internal/logic/dept"
	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListDeptHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListDeptReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dept.NewListDeptLogic(r.Context(), svcCtx)
		resp, err := l.ListDept(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
