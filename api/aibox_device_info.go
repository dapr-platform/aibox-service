package api

import (
	"aibox-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"

	"strings"

	"time"
)

var _ = time.Now()

func InitAibox_device_infoRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/aibox-device-info/page", Aibox_device_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/aibox-device-info", Aibox_device_infoListHandler)

	r.Post(common.BASE_CONTEXT+"/aibox-device-info", UpsertAibox_device_infoHandler)

	r.Delete(common.BASE_CONTEXT+"/aibox-device-info/{id}", DeleteAibox_device_infoHandler)

	r.Post(common.BASE_CONTEXT+"/aibox-device-info/batch-delete", batchDeleteAibox_device_infoHandler)

	r.Post(common.BASE_CONTEXT+"/aibox-device-info/batch-upsert", batchUpsertAibox_device_infoHandler)

}

// @Summary batch update
// @Description batch update
// @Tags AI盒子设备信息视图
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /aibox-device-info/batch-upsert [post]
func batchUpsertAibox_device_infoHandler(w http.ResponseWriter, r *http.Request) {

	var entities []model.Aibox_device_info
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("len of entities is 0"))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Aibox_device_info")
	if exists {
		for _, v := range entities {
			_, err1 := beforeHook(r, v)
			if err1 != nil {
				common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
				return
			}
		}

	}
	for _, v := range entities {
		if v.ID == "" {
			v.ID = common.NanoId()
		}

		if time.Time(v.LatestHeartBeatTime).IsZero() {
			v.LatestHeartBeatTime = common.LocalTime(time.Now())
		}

	}

	err = common.DbBatchUpsert[model.Aibox_device_info](r.Context(), common.GetDaprClient(), entities, model.Aibox_device_infoTableInfo.Name, model.Aibox_device_info_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
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
// @Param latest_heart_beat_time query string false "latest_heart_beat_time"
// @Param status query string false "status"
// @Param status_name query string false "status_name"
// @Param active_event_count query string false "active_event_count"
// @Param critical_event_count query string false "critical_event_count"
// @Param major_event_count query string false "major_event_count"
// @Param minor_event_count query string false "minor_event_count"
// @Param warning_event_count query string false "warning_event_count"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Aibox_device_info}} "objects array"
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
// @Param latest_heart_beat_time query string false "latest_heart_beat_time"
// @Param status query string false "status"
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

// @Summary save
// @Description save
// @Tags AI盒子设备信息视图
// @Accept       json
// @Param item body model.Aibox_device_info true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Aibox_device_info} "object"
// @Failure 500 {object} common.Response ""
// @Router /aibox-device-info [post]
func UpsertAibox_device_infoHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Aibox_device_info
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Aibox_device_info")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Aibox_device_info)
	}
	if val.ID == "" {
		val.ID = common.NanoId()
	}

	if time.Time(val.LatestHeartBeatTime).IsZero() {
		val.LatestHeartBeatTime = common.LocalTime(time.Now())
	}

	err = common.DbUpsert[model.Aibox_device_info](r.Context(), common.GetDaprClient(), val, model.Aibox_device_infoTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags AI盒子设备信息视图
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Aibox_device_info} "object"
// @Failure 500 {object} common.Response ""
// @Router /aibox-device-info/{id} [delete]
func DeleteAibox_device_infoHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Aibox_device_info")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "v_aibox_device_info", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags AI盒子设备信息视图
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /aibox-device-info/batch-delete [post]
func batchDeleteAibox_device_infoHandler(w http.ResponseWriter, r *http.Request) {

	var ids []string
	err := common.ReadRequestBody(r, &ids)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(ids) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("len of ids is 0"))
		return
	}
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Aibox_device_info")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "v_aibox_device_info", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
