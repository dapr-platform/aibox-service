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

func InitAibox_updateRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/aibox-update/page", Aibox_updatePageListHandler)
	r.Get(common.BASE_CONTEXT+"/aibox-update", Aibox_updateListHandler)

	r.Post(common.BASE_CONTEXT+"/aibox-update", UpsertAibox_updateHandler)

	r.Delete(common.BASE_CONTEXT+"/aibox-update/{id}", DeleteAibox_updateHandler)

	r.Post(common.BASE_CONTEXT+"/aibox-update/batch-delete", batchDeleteAibox_updateHandler)

	r.Post(common.BASE_CONTEXT+"/aibox-update/batch-upsert", batchUpsertAibox_updateHandler)

}

// @Summary batch update
// @Description batch update
// @Tags AI盒子软件更新
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /aibox-update/batch-upsert [post]
func batchUpsertAibox_updateHandler(w http.ResponseWriter, r *http.Request) {

	var entities []model.Aibox_update
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("len of entities is 0"))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Aibox_update")
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

	err = common.DbBatchUpsert[model.Aibox_update](r.Context(), common.GetDaprClient(), entities, model.Aibox_updateTableInfo.Name, model.Aibox_update_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags AI盒子软件更新
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param version query string false "version"
// @Param type query string false "type"
// @Param file_path query string false "file_path"
// @Param file_name query string false "file_name"
// @Param file_key query string false "file_key"
// @Param description query string false "description"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=common.PageGeneric[model.Aibox_update]} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /aibox-update/page [get]
func Aibox_updatePageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Aibox_update](w, r, common.GetDaprClient(), "o_aibox_update", "id")

}

// @Summary query objects
// @Description query objects
// @Tags AI盒子软件更新
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param version query string false "version"
// @Param type query string false "type"
// @Param file_path query string false "file_path"
// @Param file_name query string false "file_name"
// @Param file_key query string false "file_key"
// @Param description query string false "description"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Aibox_update} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /aibox-update [get]
func Aibox_updateListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Aibox_update](w, r, common.GetDaprClient(), "o_aibox_update", "id")
}

// @Summary save
// @Description save
// @Tags AI盒子软件更新
// @Accept       json
// @Param item body model.Aibox_update true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Aibox_update} "object"
// @Failure 500 {object} common.Response ""
// @Router /aibox-update [post]
func UpsertAibox_updateHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Aibox_update
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Aibox_update")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Aibox_update)
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

	err = common.DbUpsert[model.Aibox_update](r.Context(), common.GetDaprClient(), val, model.Aibox_updateTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags AI盒子软件更新
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Aibox_update} "object"
// @Failure 500 {object} common.Response ""
// @Router /aibox-update/{id} [delete]
func DeleteAibox_updateHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Aibox_update")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_aibox_update", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags AI盒子软件更新
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /aibox-update/batch-delete [post]
func batchDeleteAibox_updateHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Aibox_update")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_aibox_update", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
