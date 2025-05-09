package api

import (
	"aibox-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"

	"time"
)

var _ = time.Now()

func InitAibox_active_event_statsRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/aibox-active-event-stats/page", Aibox_active_event_statsPageListHandler)
	r.Get(common.BASE_CONTEXT+"/aibox-active-event-stats", Aibox_active_event_statsListHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags AI盒子活动事件统计视图
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param level query string false "level"
// @Param level_name query string false "level_name"
// @Param event_count query string false "event_count"
// @Produce  json
// @Success 200 {object} common.Response{data=common.PageGeneric[model.Aibox_active_event_stats]} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /aibox-active-event-stats/page [get]
func Aibox_active_event_statsPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Aibox_active_event_stats](w, r, common.GetDaprClient(), "v_aibox_active_event_stats", "level")

}

// @Summary query objects
// @Description query objects
// @Tags AI盒子活动事件统计视图
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param level query string false "level"
// @Param level_name query string false "level_name"
// @Param event_count query string false "event_count"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Aibox_active_event_stats} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /aibox-active-event-stats [get]
func Aibox_active_event_statsListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Aibox_active_event_stats](w, r, common.GetDaprClient(), "v_aibox_active_event_stats", "level")
}
