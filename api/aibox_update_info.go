package api

import (
	"aibox-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"

	"time"
)

var _ = time.Now()

func InitAibox_update_infoRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/aibox-update-info/page", Aibox_update_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/aibox-update-info", Aibox_update_infoListHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags AI盒子软件更新信息视图
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param version query string false "version"
// @Param type query string false "type"
// @Param type_name query string false "type_name"
// @Param file_path query string false "file_path"
// @Param file_name query string false "file_name"
// @Param file_key query string false "file_key"
// @Param description query string false "description"
// @Param status query string false "status"
// @Param status_name query string false "status_name"
// @Param created_time query string false "created_time"
// @Param updated_time query string false "updated_time"
// @Produce  json
// @Success 200 {object} common.Response{data=common.PageGeneric[model.Aibox_update_info]} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /aibox-update-info/page [get]
func Aibox_update_infoPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Aibox_update_info](w, r, common.GetDaprClient(), "v_aibox_update_info", "id")

}

// @Summary query objects
// @Description query objects
// @Tags AI盒子软件更新信息视图
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param version query string false "version"
// @Param type query string false "type"
// @Param type_name query string false "type_name"
// @Param file_path query string false "file_path"
// @Param file_name query string false "file_name"
// @Param file_key query string false "file_key"
// @Param description query string false "description"
// @Param status query string false "status"
// @Param status_name query string false "status_name"
// @Param created_time query string false "created_time"
// @Param updated_time query string false "updated_time"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Aibox_update_info} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /aibox-update-info [get]
func Aibox_update_infoListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Aibox_update_info](w, r, common.GetDaprClient(), "v_aibox_update_info", "id")
}
