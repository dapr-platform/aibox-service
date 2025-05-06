package api

import (
	"aibox-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"

	"time"
)

var _ = time.Now()

func InitAibox_event_infoRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/aibox-event-info/page", Aibox_event_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/aibox-event-info", Aibox_event_infoListHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags AI盒子事件详情视图
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param dn query string false "dn"
// @Param title query string false "title"
// @Param device_id query string false "device_id"
// @Param content query string false "content"
// @Param picstr query string false "picstr"
// @Param level query string false "level"
// @Param level_name query string false "level_name"
// @Param status query string false "status"
// @Param status_name query string false "status_name"
// @Param created_time query string false "created_time"
// @Param updated_time query string false "updated_time"
// @Param device_name query string false "device_name"
// @Param device_ip query string false "device_ip"
// @Param device_status query string false "device_status"
// @Param device_status_name query string false "device_status_name"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Aibox_event_info}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /aibox-event-info/page [get]
func Aibox_event_infoPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Aibox_event_info](w, r, common.GetDaprClient(), "v_aibox_event_info", "id")

}

// @Summary query objects
// @Description query objects
// @Tags AI盒子事件详情视图
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param dn query string false "dn"
// @Param title query string false "title"
// @Param device_id query string false "device_id"
// @Param content query string false "content"
// @Param picstr query string false "picstr"
// @Param level query string false "level"
// @Param level_name query string false "level_name"
// @Param status query string false "status"
// @Param status_name query string false "status_name"
// @Param created_time query string false "created_time"
// @Param updated_time query string false "updated_time"
// @Param device_name query string false "device_name"
// @Param device_ip query string false "device_ip"
// @Param device_status query string false "device_status"
// @Param device_status_name query string false "device_status_name"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Aibox_event_info} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /aibox-event-info [get]
func Aibox_event_infoListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Aibox_event_info](w, r, common.GetDaprClient(), "v_aibox_event_info", "id")
}
