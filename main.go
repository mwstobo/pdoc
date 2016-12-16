package main

import (
	"github.com/mwstobo/pdoc/conf"
)

func main() {
	conf.LoadConfig()

	fileWalker := walk.NewFileWalker(*conf.BaseDirectory)
	fuzzyFinder := fuzzy.NewFuzzyFinder()
	fileOpener := open.NewFileOpener(*conf.FileOpenCommand)

	possibleFiles := fuzzyFinder.Find(*conf.SearchQuery, filesystemWalker.Walk())
	if len(possibleFiles) == 1 {
		fileOpener.Open(possibleFiles[0])
	}
}
