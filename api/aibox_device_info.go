package api

import (
	"aibox-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"

	"time"
)

var _ = time.Now()

func InitAibox_device_infoRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/aibox-device-info/page", Aibox_device_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/aibox-device-info", Aibox_device_infoListHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags AI盒子设备信息视图
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param ip query string false "ip"
// @Param build_time_str query string false "build_time_str"
// @Param device_time query string false "device_time"
// @Param latest_heart_beat_time query string false "latest_heart_beat_time"
// @Param status query string false "status"
// @Param upgrade_tasks query string false "upgrade_tasks"
// @Param status_name query string false "status_name"
// @Param active_event_count query string false "active_event_count"
// @Param critical_event_count query string false "critical_event_count"
// @Param major_event_count query string false "major_event_count"
// @Param minor_event_count query string false "minor_event_count"
// @Param warning_event_count query string false "warning_event_count"
// @Produce  json
// @Success 200 {object} common.Response{data=common.PageGeneric[model.Aibox_device_info]} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /aibox-device-info/page [get]
func Aibox_device_infoPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Aibox_device_info](w, r, common.GetDaprClient(), "v_aibox_device_info", "id")

}

// @Summary query objects
// @Description query objects
// @Tags AI盒子设备信息视图
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param ip query string false "ip"
// @Param build_time_str query string false "build_time_str"
// @Param device_time query string false "device_time"
// @Param latest_heart_beat_time query string false "latest_heart_beat_time"
// @Param status query string false "status"
// @Param upgrade_tasks query string false "upgrade_tasks"
// @Param status_name query string false "status_name"
// @Param active_event_count query string false "active_event_count"
// @Param critical_event_count query string false "critical_event_count"
// @Param major_event_count query string false "major_event_count"
// @Param minor_event_count query string false "minor_event_count"
// @Param warning_event_count query string false "warning_event_count"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Aibox_device_info} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /aibox-device-info [get]
func Aibox_device_infoListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Aibox_device_info](w, r, common.GetDaprClient(), "v_aibox_device_info", "id")
}
