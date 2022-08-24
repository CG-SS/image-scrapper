package config

import (
	"image-scrapper/internal/img_formats"
	"log"
	"net/url"
)

type AppConfig struct {
	OutputDir         string
	URL               *url.URL
	InfoLog           *log.Logger
	ImageFormatFilter []img_formats.ImageFormat
}
