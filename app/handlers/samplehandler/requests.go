package samplehandler

import (
	"net/http"

	"github.com/daison12006013/gorvel/pkg/engines"
	"github.com/daison12006013/gorvel/pkg/errors"
	"github.com/daison12006013/gorvel/pkg/facade/logger"
	"github.com/daison12006013/gorvel/pkg/helpers"
	"github.com/daison12006013/gorvel/pkg/storage"
)

func Requests(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.MuxEngine)
	req := engine.Request
	res := engine.Response

	// prepare the data
	data := helpers.MP{
		"req.All()":          req.All(),
		"req.IsForm()":       req.IsForm(),
		"req.IsJson()":       req.IsJson(),
		"req.IsMultipart()":  req.IsMultipart(),
		"req.WantsJson()":    req.WantsJson(),
		"req.GetIp()":        req.GetIp(),
		"req.GetUserAgent()": req.GetUserAgent(),
	}

	return res.Json(data, http.StatusOK)
}

func FileStorage(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.MuxEngine)
	req := engine.Request
	res := engine.Response

	files, err := req.GetFiles()

	if err != nil {
		return res.Json(helpers.MP{
			"error": err.Error(),
		}, http.StatusOK)
	}

	images := files["files"]
	logger.Info("Image Length", len(images))

	// initialize local storage
	store := storage.NewLocalStorage()

	for _, image := range images {
		err := store.Put(image.Filename, image)
		if err != nil {
			return &errors.AppError{Code: 400, Error: err}
		}

		size, _ := store.Size(image.Filename)
		go logger.Info("Storage Size: ", size)
		go logger.Info("File Exist: ", store.Exists(image.Filename))
		go logger.Info("File Missing: ", store.Missing(image.Filename))
		store.Delete(image.Filename)
	}

	if err != nil {
		return res.Json(helpers.MP{
			"error": err.Error(),
		}, http.StatusOK)
	}

	// prepare the data
	data := helpers.MP{
		"file": len(images),
	}

	return res.Json(data, http.StatusOK)
}
