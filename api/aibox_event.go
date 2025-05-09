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

func InitAibox_eventRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/aibox-event/page", Aibox_eventPageListHandler)
	r.Get(common.BASE_CONTEXT+"/aibox-event", Aibox_eventListHandler)

	r.Post(common.BASE_CONTEXT+"/aibox-event", UpsertAibox_eventHandler)

	r.Delete(common.BASE_CONTEXT+"/aibox-event/{id}", DeleteAibox_eventHandler)

	r.Post(common.BASE_CONTEXT+"/aibox-event/batch-delete", batchDeleteAibox_eventHandler)

	r.Post(common.BASE_CONTEXT+"/aibox-event/batch-upsert", batchUpsertAibox_eventHandler)

}

// @Summary batch update
// @Description batch update
// @Tags AI盒子事件
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /aibox-event/batch-upsert [post]
func batchUpsertAibox_eventHandler(w http.ResponseWriter, r *http.Request) {

	var entities []model.Aibox_event
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("len of entities is 0"))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Aibox_event")
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

	}

	err = common.DbBatchUpsert[model.Aibox_event](r.Context(), common.GetDaprClient(), entities, model.Aibox_eventTableInfo.Name, model.Aibox_event_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags AI盒子事件
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param dn query string false "dn"
// @Param title query string false "title"
// @Param device_id query string false "device_id"
// @Param content query string false "content"
// @Param picstr query string false "picstr"
// @Param level query string false "level"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=common.PageGeneric[model.Aibox_event]} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /aibox-event/page [get]
func Aibox_eventPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Aibox_event](w, r, common.GetDaprClient(), "o_aibox_event", "id")

}

// @Summary query objects
// @Description query objects
// @Tags AI盒子事件
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param dn query string false "dn"
// @Param title query string false "title"
// @Param device_id query string false "device_id"
// @Param content query string false "content"
// @Param picstr query string false "picstr"
// @Param level query string false "level"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Aibox_event} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /aibox-event [get]
func Aibox_eventListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Aibox_event](w, r, common.GetDaprClient(), "o_aibox_event", "id")
}

// @Summary save
// @Description save
// @Tags AI盒子事件
// @Accept       json
// @Param item body model.Aibox_event true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Aibox_event} "object"
// @Failure 500 {object} common.Response ""
// @Router /aibox-event [post]
func UpsertAibox_eventHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Aibox_event
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Aibox_event")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Aibox_event)
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

	err = common.DbUpsert[model.Aibox_event](r.Context(), common.GetDaprClient(), val, model.Aibox_eventTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags AI盒子事件
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Aibox_event} "object"
// @Failure 500 {object} common.Response ""
// @Router /aibox-event/{id} [delete]
func DeleteAibox_eventHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Aibox_event")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_aibox_event", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags AI盒子事件
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /aibox-event/batch-delete [post]
func batchDeleteAibox_eventHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Aibox_event")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_aibox_event", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
