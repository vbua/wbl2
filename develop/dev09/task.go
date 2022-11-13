package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"net/url"
	"os"
	"path/filepath"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func htmlHandler(attr string, visitedUrls map[string]struct{}) func(e *colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr(attr))
		if _, ok := visitedUrls[link]; !ok {
			visitedUrls[link] = struct{}{}
			e.Request.Visit(link)
		}
	}
}

func main() {
	// читаем с коммандой строки и вытаскиваем только домен с урла
	input := os.Args[len(os.Args)-1]
	url, err := url.Parse(input)
	if err != nil {
		log.Fatal(err)
	}
	hostname := url.Hostname()

	// создаем главную директорию сайта, куда будем скачивать
	err = os.Mkdir(hostname, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	c := colly.NewCollector(
		// посещаем только заданный домен, на внешние ссылки не идем
		colly.AllowedDomains(hostname),
	)

	// чтобы не скачивать одно и то же несколько раз
	visitedUrls := make(map[string]struct{})

	c.OnHTML("a[href]", htmlHandler("href", visitedUrls))
	c.OnHTML("script[src]", htmlHandler("src", visitedUrls))
	c.OnHTML("img[src]", htmlHandler("src", visitedUrls))
	c.OnHTML("link[href]", htmlHandler("href", visitedUrls))

	c.OnRequest(func(r *colly.Request) {
		if _, ok := visitedUrls[r.URL.RequestURI()]; !ok {
			fmt.Println("Visiting", r.URL)
		}
	})

	// тут уже работаем с респонсом, можем сохранить сожержимое файла
	c.OnResponse(func(r *colly.Response) {
		fullPath := hostname + r.Request.URL.Path
		// если это путь без файла, то создаем папки и туда кидаем index.html
		if filepath.Ext(fullPath) == "" {
			err := os.MkdirAll(fullPath, os.ModePerm)
			if err != nil {
				log.Fatalln(err)
			}
			r.Save(fullPath + "/index.html")
		} else { // если это файл, то так и сохраняем
			err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
			if err != nil {
				log.Fatalln(err)
			}
			r.Save(fullPath)
		}
	})

	err = c.Visit(input)
	if err != nil {
		log.Fatalln(err)
	}
	c.Wait()
	fmt.Println("Successfully downloaded site")
}
