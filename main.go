package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"image-scrapper/internal/config"
	"image-scrapper/internal/img_formats"
	"image-scrapper/internal/scrapper"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	parser := argparse.NewParser(
		"image-scrapper",
		"Crawls and scraps a given website for images, and outputs them in the given directory.",
	)
	rawOutputDir := parser.String("o", "output-dir", &argparse.Options{Required: true, Help: "Where to output the images."})
	rawWebsiteUrl := parser.String("u", "url", &argparse.Options{Required: true, Help: "URL for what website to crawl."})
	rawImgFormatFilter := parser.StringList(
		"f",
		"image-filter",
		&argparse.Options{Required: false, Help: "What formats to ignore. Default none. Options: " + strings.Join(img_formats.AllFormatsString(), ",")},
	)
	verbose := parser.Flag("v", "verbose", &argparse.Options{Required: false, Help: "If present, will enable logging."})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	fileInfo, err := os.Stat(*rawOutputDir)
	if err != nil || !fileInfo.IsDir() {
		fmt.Print(parser.Usage(err))
		return
	}

	websiteUrl, err := url.Parse(*rawWebsiteUrl)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	var imgFormatFilter []img_formats.ImageFormat
	if len(*rawImgFormatFilter) > 0 {
		for _, format := range *rawImgFormatFilter {
			parsedFormat := img_formats.ParseImageFormat(format)
			if parsedFormat != img_formats.UNKNOWN {
				imgFormatFilter = append(imgFormatFilter, parsedFormat)
			}
		}
	}

	var infoLog *log.Logger
	if *verbose {
		infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	}

	scrapper.Init(&config.AppConfig{
		OutputDir:         *rawOutputDir,
		URL:               websiteUrl,
		InfoLog:           infoLog,
		ImageFormatFilter: imgFormatFilter,
	})
	scrapper.Run()
}
