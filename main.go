package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/badugisoft/go-template/funcs"
	"github.com/badugisoft/xson"

	cli "gopkg.in/urfave/cli.v2"
)

func main() {
	app := cli.App{
		Name:    "go-template",
		Usage:   "go template commandline tool",
		Version: "0.0.1",
		Action:  run,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "path",
				Aliases: []string{"p"},
				Usage:   "input data file path\n\tif this is missing, stdin is used",
			},
			&cli.StringFlag{
				Name:    "type",
				Aliases: []string{"t"},
				Value:   "autodetect",
				Usage:   "input data file or stdin data type.\n\tif path is missing, default type is yaml\n\tsupported types : yaml, xml, json\n\t",
			},
			&cli.StringFlag{
				Name:    "main",
				Aliases: []string{"m"},
				Value:   "main",
				Usage:   "main template name\n\t",
			},
		},
		ArgsUsage: "[template files]",
	}

	app.Run(os.Args)
}

func run(c *cli.Context) error {
	dataPath := c.String("path")
	dataType := c.String("type")
	var realType xson.Type

	if dataType == "autodetect" {
		if dataPath == "" {
			realType = xson.YAML
		} else {
			realType = xson.GetType(dataPath)
			if realType == xson.UNKNOWN {
				return cli.Exit("unsupported data type : "+dataType, -1)
			}
		}
	} else {
		realType = xson.GetType(dataType)
		if realType == xson.UNKNOWN {
			return cli.Exit("unsupported data type : "+dataType, -1)
		}
	}

	if c.Args().Len() == 0 {
		return cli.Exit("at least one template file is required", -2)
	}

	var fileData []byte
	var err error
	if dataPath != "" {
		fileData, err = ioutil.ReadFile(dataPath)
		if err != nil {
			return cli.Exit("read data file failed : "+dataType+", "+err.Error(), -3)
		}
	} else {
		fileData, err = ioutil.ReadAll(bufio.NewReader(os.Stdin))
		if err != nil {
			return cli.Exit("read data from stdin : "+dataType+", "+err.Error(), -3)
		}
	}

	data := map[string]interface{}{}
	err = xson.Unmarshal(realType, fileData, &data)
	if err != nil {
		return cli.Exit("read data failed : "+err.Error(), -4)
	}

	t, err := template.New("").Funcs(funcs.FuncMap).ParseFiles(c.Args().Slice()...)
	if err != nil {
		return cli.Exit("parsing template failed : "+err.Error(), -6)
	}

	if len(t.Templates()) == 1 {
		err = t.ExecuteTemplate(os.Stdout, t.Templates()[0].Name(), data)
	} else {
		err = t.ExecuteTemplate(os.Stdout, c.String("main"), data)
	}

	if err != nil {
		return cli.Exit("executing template failed : "+err.Error(), -7)
	}

	return nil
}
