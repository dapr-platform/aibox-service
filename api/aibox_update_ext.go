package api

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
)

func InitAiboxUpdateExtRoute(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/upload", uploadHandler)
	r.Post(common.BASE_CONTEXT+"/download", downloadHandler)
}

// @Summary 上传文件
// @Description 上传文件
// @Tags 软件更新文件管理
// @Accept multipart/form-data
// @Param version formData string true "版本号"
// @Param type formData string true "类型"
// @Param file formData file true "文件"
// @Success 200 {string} string "上传成功"
// @Failure 400 {string} string "上传失败"
// @Router /upload [post]
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// 解析表单
	err := r.ParseMultipartForm(100 << 20) // 100MB 限制
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("解析表单失败: "+err.Error()))
		return
	}

	// 获取表单参数
	version := r.FormValue("version")
	typeStr := r.FormValue("type")
	if version == "" || typeStr == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("版本号或类型不能为空"))
		return
	}

	// 获取上传的文件
	file, header, err := r.FormFile("file")
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("获取文件失败: "+err.Error()))
		return
	}
	defer file.Close()

	// 创建目录结构 uploads/<type>/<version>/
	uploadDir := filepath.Join("uploads", typeStr, version)
	err = os.MkdirAll(uploadDir, 0755)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg("创建目录失败: "+err.Error()))
		return
	}

	// 创建目标文件
	filePath := filepath.Join(uploadDir, header.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg("创建文件失败: "+err.Error()))
		return
	}
	defer dst.Close()

	// 计算MD5
	hash := md5.New()
	tee := io.TeeReader(file, hash)

	// 复制文件内容
	_, err = io.Copy(dst, tee)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg("保存文件失败: "+err.Error()))
		return
	}

	// 获取MD5值
	md5sum := hex.EncodeToString(hash.Sum(nil))

	common.HttpResult(w, common.OK.AppendMsg("上传成功").WithData(map[string]string{
		"filePath": filePath,
		"md5":      md5sum,
	}))
}

// @Summary 下载文件
// @Description 下载文件
// @Tags 软件更新文件管理
// @Param version formData string true "版本号"
// @Param type formData string true "类型"
// @Param filename query string true "文件名"
// @Success 200 {string} string "下载成功"
// @Failure 400 {string} string "下载失败"
// @Router /download [post]
func downloadHandler(w http.ResponseWriter, r *http.Request) {
	// 解析表单
	err := r.ParseForm()
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("解析表单失败: "+err.Error()))
		return
	}

	// 获取参数
	version := r.FormValue("version")
	typeStr := r.FormValue("type")
	filename := r.URL.Query().Get("filename")

	if version == "" || typeStr == "" || filename == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("版本号、类型或文件名不能为空"))
		return
	}

	// 构建文件路径
	filePath := filepath.Join("uploads", typeStr, version, filename)

	// 检查文件是否存在
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		common.HttpResult(w, common.ErrNotFound.AppendMsg("文件不存在"))
		return
	}
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg("获取文件信息失败: "+err.Error()))
		return
	}

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg("打开文件失败: "+err.Error()))
		return
	}
	defer file.Close()

	// 设置响应头
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	// 发送文件内容
	_, err = io.Copy(w, file)
	if err != nil {
		// 此时已经开始发送响应，无法再发送错误信息
		// 只能记录日志
		fmt.Println("发送文件失败:", err)
	}
}
