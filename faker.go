package faker

import (
         // "fmt"
         "github.com/diebels727/readline"
         "os"
        )
type Faker struct {
  FirstNames []string
  firstNamesFile string
}

func New() (*Faker) {
  f := Faker{firstNamesFile: "data/name/first_names"}
  file,err := os.Open(f.firstNamesFile)
  if err != nil {
    panic("Error opening file")
  }
  reader := readline.Readline{File: file}
  defer file.Close()

  reader.Map(f.firstNamesReader)
  return &f
}

func (f *Faker) firstNamesReader(str string) {
  f.FirstNames = append(f.FirstNames,str)
  return
}

// func init() {
//   *first_names = (make([]string,0))
//   // first_names_file_name := "data/name/first_names"
//   // file,err := os.Open(first_names_file_name)
//   // if err != nil {
//   //   panic("Error opening file")
//   // }
//   // reader := readline.Readline{File: file}
//   // defer file.Close()
//   //
//   // reader.Map(firstNamesReader)
//   fmt.Println(first_names)
// }

