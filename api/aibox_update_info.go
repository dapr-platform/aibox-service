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

func InitAibox_update_infoRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/aibox-update-info/page", Aibox_update_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/aibox-update-info", Aibox_update_infoListHandler)

	r.Post(common.BASE_CONTEXT+"/aibox-update-info", UpsertAibox_update_infoHandler)

	r.Delete(common.BASE_CONTEXT+"/aibox-update-info/{id}", DeleteAibox_update_infoHandler)

	r.Post(common.BASE_CONTEXT+"/aibox-update-info/batch-delete", batchDeleteAibox_update_infoHandler)

	r.Post(common.BASE_CONTEXT+"/aibox-update-info/batch-upsert", batchUpsertAibox_update_infoHandler)

}

// @Summary batch update
// @Description batch update
// @Tags AI盒子软件更新信息视图
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /aibox-update-info/batch-upsert [post]
func batchUpsertAibox_update_infoHandler(w http.ResponseWriter, r *http.Request) {

	var entities []model.Aibox_update_info
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("len of entities is 0"))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Aibox_update_info")
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

	err = common.DbBatchUpsert[model.Aibox_update_info](r.Context(), common.GetDaprClient(), entities, model.Aibox_update_infoTableInfo.Name, model.Aibox_update_info_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
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
// @Param description query string false "description"
// @Param status query string false "status"
// @Param status_name query string false "status_name"
// @Param created_time query string false "created_time"
// @Param updated_time query string false "updated_time"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Aibox_update_info}} "objects array"
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

// @Summary save
// @Description save
// @Tags AI盒子软件更新信息视图
// @Accept       json
// @Param item body model.Aibox_update_info true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Aibox_update_info} "object"
// @Failure 500 {object} common.Response ""
// @Router /aibox-update-info [post]
func UpsertAibox_update_infoHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Aibox_update_info
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Aibox_update_info")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Aibox_update_info)
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

	err = common.DbUpsert[model.Aibox_update_info](r.Context(), common.GetDaprClient(), val, model.Aibox_update_infoTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags AI盒子软件更新信息视图
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Aibox_update_info} "object"
// @Failure 500 {object} common.Response ""
// @Router /aibox-update-info/{id} [delete]
func DeleteAibox_update_infoHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Aibox_update_info")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "v_aibox_update_info", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags AI盒子软件更新信息视图
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /aibox-update-info/batch-delete [post]
func batchDeleteAibox_update_infoHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Aibox_update_info")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "v_aibox_update_info", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
