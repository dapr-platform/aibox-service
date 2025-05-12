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

func InitAibox_deviceRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/aibox-device/page", Aibox_devicePageListHandler)
	r.Get(common.BASE_CONTEXT+"/aibox-device", Aibox_deviceListHandler)

	r.Post(common.BASE_CONTEXT+"/aibox-device", UpsertAibox_deviceHandler)

	r.Delete(common.BASE_CONTEXT+"/aibox-device/{id}", DeleteAibox_deviceHandler)

	r.Post(common.BASE_CONTEXT+"/aibox-device/batch-delete", batchDeleteAibox_deviceHandler)

	r.Post(common.BASE_CONTEXT+"/aibox-device/batch-upsert", batchUpsertAibox_deviceHandler)

}

// @Summary batch update
// @Description batch update
// @Tags AI盒子设备
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /aibox-device/batch-upsert [post]
func batchUpsertAibox_deviceHandler(w http.ResponseWriter, r *http.Request) {

	var entities []model.Aibox_device
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("len of entities is 0"))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Aibox_device")
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

		if time.Time(v.CreatedTime).IsZero() {
			v.CreatedTime = common.LocalTime(time.Now())
		}

		if time.Time(v.UpdatedTime).IsZero() {
			v.UpdatedTime = common.LocalTime(time.Now())
		}

		if time.Time(v.DeviceTime).IsZero() {
			v.DeviceTime = common.LocalTime(time.Now())
		}

		if time.Time(v.LatestHeartBeatTime).IsZero() {
			v.LatestHeartBeatTime = common.LocalTime(time.Now())
		}

	}

	err = common.DbBatchUpsert[model.Aibox_device](r.Context(), common.GetDaprClient(), entities, model.Aibox_deviceTableInfo.Name, model.Aibox_device_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags AI盒子设备
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param name query string false "name"
// @Param ip query string false "ip"
// @Param build_time_str query string false "build_time_str"
// @Param device_time query string false "device_time"
// @Param latest_heart_beat_time query string false "latest_heart_beat_time"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=common.PageGeneric[model.Aibox_device]} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /aibox-device/page [get]
func Aibox_devicePageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Aibox_device](w, r, common.GetDaprClient(), "o_aibox_device", "id")

}

// @Summary query objects
// @Description query objects
// @Tags AI盒子设备
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param name query string false "name"
// @Param ip query string false "ip"
// @Param build_time_str query string false "build_time_str"
// @Param device_time query string false "device_time"
// @Param latest_heart_beat_time query string false "latest_heart_beat_time"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Aibox_device} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /aibox-device [get]
func Aibox_deviceListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Aibox_device](w, r, common.GetDaprClient(), "o_aibox_device", "id")
}

// @Summary save
// @Description save
// @Tags AI盒子设备
// @Accept       json
// @Param item body model.Aibox_device true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Aibox_device} "object"
// @Failure 500 {object} common.Response ""
// @Router /aibox-device [post]
func UpsertAibox_deviceHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Aibox_device
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Aibox_device")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Aibox_device)
	}
	if val.ID == "" {
		val.ID = common.NanoId()
	}

	if time.Time(val.CreatedTime).IsZero() {
		val.CreatedTime = common.LocalTime(time.Now())
	}

	if time.Time(val.UpdatedTime).IsZero() {
		val.UpdatedTime = common.LocalTime(time.Now())
	}

	if time.Time(val.DeviceTime).IsZero() {
		val.DeviceTime = common.LocalTime(time.Now())
	}

	if time.Time(val.LatestHeartBeatTime).IsZero() {
		val.LatestHeartBeatTime = common.LocalTime(time.Now())
	}

	err = common.DbUpsert[model.Aibox_device](r.Context(), common.GetDaprClient(), val, model.Aibox_deviceTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags AI盒子设备
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Aibox_device} "object"
// @Failure 500 {object} common.Response ""
// @Router /aibox-device/{id} [delete]
func DeleteAibox_deviceHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Aibox_device")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_aibox_device", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags AI盒子设备
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /aibox-device/batch-delete [post]
func batchDeleteAibox_deviceHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Aibox_device")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_aibox_device", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
