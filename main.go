package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/urfave/cli/v2"
)

func main() {
	var contentDir string
	var templateFile string
	var outFile string

	var address string
	var port int
	var siteDir string
	var file string

	app := &cli.App{
		Name: "help",
		Commands: []*cli.Command{
			{
				Name: "render",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "contentDir",
						Aliases:     []string{"c"},
						Value:       "content",
						Destination: &contentDir,
					},
					&cli.StringFlag{
						Name:        "templateFile",
						Aliases:     []string{"t"},
						Value:       "template.html",
						Destination: &templateFile,
					},
					&cli.StringFlag{
						Name:        "outFile",
						Aliases:     []string{"o"},
						Value:       "site/index.html",
						Destination: &outFile,
					},
				},
				Action: func(c *cli.Context) error {
					return renderTemplate(contentDir, templateFile, outFile)
				},
			},
			{
				Name: "serve",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "address",
						Aliases:     []string{"a"},
						Value:       "",
						Destination: &address,
					},
					&cli.IntFlag{
						Name:        "port",
						Aliases:     []string{"p"},
						Value:       9001,
						Destination: &port,
					},
					&cli.StringFlag{
						Name:        "siteDir",
						Value:       "site",
						Destination: &siteDir,
					},
					&cli.StringFlag{
						Name:        "file",
						Aliases:     []string{"f"},
						Value:       "site/index.html",
						Destination: &file,
					},
				},
				Action: func(c *cli.Context) error {
					return runServer(address, port, siteDir, file)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func renderTemplate(contentDir string, templateFile string, outFile string) error {
	files, err := parseFiles(contentDir)
	if err != nil {
		return err
	}

	fileNames := make([]string, len(files))
	contents := make([]string, len(files))
	for i, f := range files {
		fileNames[i] = f.name
		c, err := readContent(f.path)
		if err != nil {
			return err
		}
		contents[i] = c
	}

	templateData := TemplateData{Files: fileNames, Contents: contents}

	templateName := filepath.Base(templateFile)
	t := template.New(templateName)
	t, err = t.ParseFiles(templateFile)
	if err != nil {
		return err
	}

	w, err := os.Create(outFile)
	if err != nil {
		return err
	}

	return t.ExecuteTemplate(w, templateName, templateData)
}

type TemplateData struct {
	Files    []string
	Contents []string
}

func readContent(f string) (string, error) {
	content, err := os.ReadFile(f)
	if err != nil {
		return "", err
	}
	return string(content), err
}

type File struct {
	path string
	name string
}

func parseFiles(contentDir string) ([]File, error) {
	file, err := os.Open(contentDir)
	if err != nil {
		return nil, err
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if !fileInfo.IsDir() {
		return nil, errors.New(fmt.Sprintf("%v is not a directory", contentDir))
	}

	entries, err := file.ReadDir(0)
	if err != nil {
		return nil, err
	}

	result := make([]File, 0)
	for _, entry := range entries {
		fileName := entry.Name()
		path := filepath.Join(contentDir, fileName)
		name := strings.TrimSuffix(fileName, filepath.Ext(fileName))
		result = append(result, File{path: path, name: name})
	}

	if len(result) == 0 {
		return nil, errors.New("no files in content directory")
	}

	return result, nil
}

func runServer(address string, port int, siteDir string, file string) error {
	e := echo.New()

	e.File("/", file)
	e.Static("/", siteDir)

	if err := e.Start(fmt.Sprintf("%v:%v", address, port)); err != http.ErrServerClosed {
		return err
	}

	return nil
}
