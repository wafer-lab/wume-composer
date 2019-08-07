package controllers

import (
	"net/http"

	"wume-composer/internal/pkg/common/files"
	"wume-composer/internal/pkg/common/logger"
	"wume-composer/internal/pkg/db"
	"wume-composer/internal/pkg/models"
)

func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	if !requireAuth(w, r) {
		return
	}

	file, header, err := getFormFile(w, r, int64(5*files.MByte), "avatar")
	if err != nil {
		return
	}

	url, err := files.UploadAvatar(file, header)
	if !handleCommonErrors(w, err, []string{"avatar"}) {
		return
	}

	if err = file.Close(); err != nil {
		logger.Error("Impossible to close file: " + err.Error())
	}

	oldAvatar, err := db.UserUpdateAvatar(jwtData(r).Id, url)
	if err != nil {
		sendJson(w, http.StatusInternalServerError, models.GetDeveloperErrorAnswer(err.Error()))
		logger.Error("Developer error: " + err.Error())
		return
	}

	err = files.RemoveAvatar(oldAvatar)
	if err != nil {
		sendJson(w, http.StatusInternalServerError, models.GetDeveloperErrorAnswer(err.Error()))
		logger.Error("Developer error: " + err.Error())
		return
	}

	sendJson(w, http.StatusOK, models.GetUploadAvatarAnswer(url))
}
