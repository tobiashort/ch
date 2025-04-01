//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	. "github.com/tobiashort/ch/canton"
	. "github.com/tobiashort/ch/coord"
)

type Geometry struct {
	Coordinates [][][]float32
}

type Properties struct {
	NAME string
}

type Feature struct {
	Properties Properties
	Geometry   Geometry
}

type FeatureCollection struct {
	Features []Feature
}

func assertNil(val any) {
	if val != nil {
		panic(val)
	}
}

func nameToAbbr(name string) string {
	switch name {
	case "Aargau":
		return "AG"
	case "Appenzell Ausserrhoden":
		return "AR"
	case "Appenzell Innerrhoden":
		return "AI"
	case "Basel-Landschaft":
		return "BL"
	case "Basel-Stadt":
		return "BS"
	case "Bern":
		return "BE"
	case "Fribourg":
		return "FR"
	case "Genève":
		return "GE"
	case "Glarus":
		return "GL"
	case "Graubünden":
		return "GR"
	case "Jura":
		return "JU"
	case "Luzern":
		return "LU"
	case "Neuchâtel":
		return "NE"
	case "Nidwalden":
		return "NW"
	case "Obwalden":
		return "OW"
	case "Schaffhausen":
		return "SH"
	case "Schwyz":
		return "SZ"
	case "Solothurn":
		return "SO"
	case "St. Gallen":
		return "SG"
	case "Thurgau":
		return "TG"
	case "Ticino":
		return "TI"
	case "Uri":
		return "UR"
	case "Vaud":
		return "VD"
	case "Valais":
		return "VS"
	case "Zug":
		return "ZG"
	case "Zürich":
		return "ZH"
	default:
		panic(name)
	}
}

func main() {
	data, err := os.ReadFile("canton/gen/swissBOUNDARIES3D_1_3_TLM_KANTONSGEBIET.geojson")
	assertNil(err)
	fc := FeatureCollection{}
	err = json.Unmarshal(data, &fc)
	assertNil(err)

	cantons := make(map[string]Canton)
	for _, feature := range fc.Features {
		canton, ok := cantons[feature.Properties.NAME]
		if !ok {
			canton = Canton{}
			canton.Name = feature.Properties.NAME
			canton.Abbr = nameToAbbr(canton.Name)
			canton.Polygons = make([][]Coord, 0)
		}
		for _, geometry := range feature.Geometry.Coordinates {
			polygon := make([]Coord, 0)
			for _, northWest := range geometry {
				north := northWest[0]
				west := northWest[1]
				polygon = append(polygon, Coord{North: north, West: west})
			}
			canton.Polygons = append(canton.Polygons, polygon)
			cantons[canton.Name] = canton
		}
	}

	tmpl, err := template.ParseFiles("canton/gen/canton.gotmpl")
	assertNil(err)
	for _, canton := range cantons {
		fileName := fmt.Sprintf("canton/%s.go", strings.ToLower(canton.Abbr))
		fmt.Println(fileName)
		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		assertNil(err)
		err = tmpl.Execute(file, canton)
		assertNil(err)
		cmd := exec.Command("gofmt", "-w", fileName)
		cmdOut, err := cmd.CombinedOutput()
		if err != nil {
			panic(string(cmdOut))
		}
	}

	tmpl, err = template.ParseFiles("canton/gen/draw_all.gotmpl")
	assertNil(err)
	fileName := fmt.Sprintf("canton/draw_all.go")
	fmt.Println(fileName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	assertNil(err)
	err = tmpl.Execute(file, cantons)
	assertNil(err)
	cmd := exec.Command("gofmt", "-w", fileName)
	cmdOut, err := cmd.CombinedOutput()
	if err != nil {
		panic(string(cmdOut))
	}
}
