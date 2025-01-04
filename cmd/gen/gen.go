package main

import (
	"github.com/Cattle0Horse/url-shortener/internal/model"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/global/query",
		Mode:    gen.WithoutContext, // generate mode
	})

	//g.ApplyBasic(database.Models...)
	g.ApplyBasic(
		model.User{},
		model.Url{},
	)

	g.Execute()
}
