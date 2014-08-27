package faker

import (
         // "fmt"
         "github.com/diebels727/readline"
         "os"
         "math/rand"
        )
type Faker struct {
  FirstNames []string
  firstNamesFile string
  LastNames []string
  lastNamesFile string
  JoinCharacters []string
}

func New() (*Faker) {
  f := Faker{firstNamesFile: "data/name/first_names",lastNamesFile: "data/name/last_names"}
  first_names_file,err := os.Open(f.firstNamesFile)
  if err != nil {
    panic("Error opening first names file")
  }
  reader := readline.Readline{File: first_names_file}
  defer first_names_file.Close()
  reader.Map(f.firstNamesReader)

  last_names_file,err := os.Open(f.lastNamesFile)
  if err != nil {
    panic("Error opening last names file")
  }
  reader = readline.Readline{File: last_names_file}
  defer last_names_file.Close()
  reader.Map(f.lastNamesReader)

  f.JoinCharacters = []string{"",".","_","-"}

  return &f
}

func (f *Faker) firstNamesReader(str string) {
  f.FirstNames = append(f.FirstNames,str)
  return
}

func (f *Faker) lastNamesReader(str string) {
  f.LastNames = append(f.LastNames,str)
  return
}

func (f *Faker) FirstName() string {
  return f.FirstNames[rand.Intn(len(f.FirstNames))]
}

func (f *Faker) LastName() string {
  return f.LastNames[rand.Intn(len(f.LastNames))]
}

func (f *Faker) JoinCharacter() string {
  return f.JoinCharacters[rand.Intn(len(f.JoinCharacters))]
}

func (f *Faker) Username() string {
  return f.FirstName()+f.JoinCharacter()+f.LastName()
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

