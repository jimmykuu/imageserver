package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/pierrre/imageserver"
	imageserver_source "github.com/pierrre/imageserver/source"
)

var (
	formats = map[string]string{
		".jpeg": "jpeg",
		".jpg":  "jpeg",
		".gif":  "gif",
		".png":  "png",
	}

	// Server 使用文件名提供 server 服务
	Server = imageserver.Server(imageserver.ServerFunc(func(params imageserver.Params) (*imageserver.Image, error) {
		source, err := params.GetString(imageserver_source.Param)
		if err != nil {
			return nil, err
		}

		im, err := Get(source)
		if err != nil {
			return nil, &imageserver.ParamError{Param: imageserver_source.Param, Message: err.Error()}
		}
		return im, nil
	}))
)

func Get(name string) (*imageserver.Image, error) {
	format, err := GetFormat(name)
	if err != nil {
		return nil, err
	}

	return loadImage(name, format)
}

func GetFormat(filename string) (string, error) {
	suffix := filepath.Ext(filename)
	if suffix == "" {
		return "", errors.New("invalid filename")
	}

	if format, has := formats[suffix]; has {
		return format, nil
	}

	return "", errors.New("invalid filename")
}

func loadImage(filename string, format string) (*imageserver.Image, error) {
	filePath := filepath.Join(config.ImagePath, filename)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	im := &imageserver.Image{
		Format: format,
		Data:   data,
	}

	return im, nil
}
