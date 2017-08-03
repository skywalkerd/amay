package fileutil

import "strings"

type ITemplate interface {
  Apply() error
}

// Holds all information to take a templated file `from` and create a new one
type Template struct {
  File
  from File
  patterns map[string]string
}

// Initiates
func (t *Template) Init(from File) {
  t.patterns = make(map[string]string)
  t.from = from
}

func (t *Template) Apply() () {
  content := ""
  for key, value := range t.patterns {
    content += strings.Replace(t.from.Content(), key, value, -1)
  }
  t.SetContent(content)
}

func (t *Template) AddPattern(key string, value string) {
  t.patterns[key] = value
}