package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/urfave/cli/v2"
)

type Args struct {
	contentDir string
	template   string
	address    string
	port       int
}

func main() {
	var contentDir string
	var template string
	var address string
	var port int

	app := &cli.App{
		Name: "help",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "contentDir",
				Value:       "content",
				Destination: &contentDir,
			},
			&cli.StringFlag{
				Name:        "template",
				Value:       "template.html",
				Destination: &template,
			},
			&cli.StringFlag{
				Name:        "address",
				Value:       "",
				Destination: &address,
			},
			&cli.IntFlag{
				Name:        "port",
				Value:       9001,
				Destination: &port,
			},
		},
		Action: func(*cli.Context) error {
			args := Args{
				contentDir: contentDir,
				template:   template,
				address:    address,
				port:       port,
			}
			return run(args)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args Args) error {
	files, err := parseFiles(args.contentDir)
	if err != nil {
		return err
	}

	template, err := NewTemplate(args.template)
	if err != nil {
		return err
	}

	runServer(files, template, args.address, args.port)

	return nil
}

func runServer(files []File, template *Template, address string, port int) error {
	e := echo.New()

	h := NewHandler(files)
	e.Static("/static", "static")
	e.GET("/", h.handle)
	e.GET("/:name", h.handle)

	e.Renderer = template

	if err := e.Start(fmt.Sprintf("%v:%v", address, port)); err != http.ErrServerClosed {
		return err
	}

	return nil
}

type Handler struct {
	files     []File
	fileNames []string
}

func NewHandler(files []File) Handler {
	fileNames := make([]string, len(files))
	for i, f := range files {
		fileNames[i] = f.name
	}
	return Handler{
		files:     files,
		fileNames: fileNames,
	}
}

type TemplateData struct {
	Files    []string
	Selected int
	Content  string
}

func (h Handler) handle(c echo.Context) error {
	name := c.Param("name")
	index, file := h.matchFile(name)
	content, err := readContent(file.path)
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, templateName, TemplateData{h.fileNames, index, content})
}

func (h Handler) matchFile(name string) (int, File) {
	if name == "" {
		return 0, h.files[0]
	}

	for i, f := range h.files {
		if f.name == name {
			return i, f
		}
	}

	return 0, h.files[0]
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

const templateName = "template.html"

type Template struct {
	templates *template.Template
}

func NewTemplate(templateFile string) (*Template, error) {
	t := template.New(templateName)
	t, err := t.ParseFiles(templateFile)
	if err != nil {
		return nil, err
	}
	return &Template{t}, err
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
