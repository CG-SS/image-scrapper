package scrapper

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"image-scrapper/internal/config"
	"io"
	"mvdan.cc/xurls/v2"
	"net/http"
	url2 "net/url"
	"os"
	"path/filepath"
	"strings"
)

var appConfig *config.AppConfig

func Init(config *config.AppConfig) {
	appConfig = config
}

func Run() {
	initialMsg := "Starting scrapper"
	if len(appConfig.ImageFormatFilter) > 0 {
		var filterStr []string
		for _, filter := range appConfig.ImageFormatFilter {
			filterStr = append(filterStr, filter.String())
		}

		initialMsg = initialMsg + " with filters (" + strings.Join(filterStr, ",") + ")"
	}
	log(initialMsg + ".")

	exploreLink()
}

func log(msg string) {
	if appConfig.InfoLog != nil {
		appConfig.InfoLog.Println(msg)
	}
}

func er(e error) {
	if appConfig.InfoLog != nil {
		appConfig.ErrorLog.Println(e.Error())
	}
}

func isValidImgSuffix(s string) bool {
	for _, format := range appConfig.ImageFormatFilter {
		if strings.HasSuffix(strings.ToUpper(s), "."+format.String()) {
			return true
		}
	}

	return false
}

func getImgLinks(doc *goquery.Document) []*url2.URL {
	var imgLinks []*url2.URL

	doc.Each(func(index int, item *goquery.Selection) {
		matches := xurls.Strict().FindAllString(item.Text(), -1)
		if matches != nil {
			for _, m := range matches {
				u, err := url2.Parse(m)

				if err == nil && isValidImgSuffix(u.Path) {
					imgLinks = append(imgLinks, u)
				}
			}
		}
	})

	return imgLinks
}

func getFileName(link *url2.URL) (string, error) {
	pathSplit := strings.Split(link.Path, "/")
	pathSplitLen := len(pathSplit)

	if pathSplitLen <= 0 {
		return "", errors.New("failed to split link")
	}

	return pathSplit[pathSplitLen-1], nil
}

func downloadImage(link *url2.URL) error {
	linkStr := link.String()

	log("Downloading " + linkStr)

	res, err := http.Get(linkStr)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	if res.StatusCode != 200 {
		return errors.New(fmt.Sprintf("received status code %d", res.StatusCode))
	}

	fileName, err := getFileName(link)
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(appConfig.OutputDir, fileName))
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}

	return nil
}

func downloadAllImages(links []*url2.URL) {
	for _, imgLink := range links {
		err := downloadImage(imgLink)
		if err != nil {
			er(err)
		}
	}
}

func exploreLink() {
	url := appConfig.URL.String()
	log("Vising " + url)

	res, err := http.Get(url)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	if err != nil {
		er(err)
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		er(err)
		return
	}

	imgLinks := getImgLinks(doc)

	for _, img := range imgLinks {
		log("Got " + img.EscapedPath())
	}

	downloadAllImages(imgLinks)
}
