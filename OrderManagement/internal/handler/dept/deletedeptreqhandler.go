package dept

import (
	"net/http"

	"OrderManagement/OrderManagement/internal/logic/dept"
	"OrderManagement/OrderManagement/internal/svc"
	"OrderManagement/OrderManagement/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteDeptReqHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListDeptOptionsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dept.NewDeleteDeptReqLogic(r.Context(), svcCtx)
		resp, err := l.DeleteDeptReq(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
