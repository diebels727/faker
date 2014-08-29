package faker

import (
         "github.com/diebels727/readline"
         "os"
         "math/rand"
         "time"
        )
type Faker struct {
  FirstNames []string
  firstNamesFile string
  LastNames []string
  lastNamesFile string
  UserNames []string
  userNamesFile string
  JoinCharacters []string
}

func New() (*Faker) {
  rand.Seed( time.Now().UTC().UnixNano())

  f := Faker{firstNamesFile: "data/name/first_names",lastNamesFile: "data/name/last_names",userNamesFile:"data/username/usernames"}
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

  usernames_file,err := os.Open(f.userNamesFile)
  if err != nil {
    panic("Error opening usernames file")
  }
  reader = readline.Readline{File: usernames_file}
  defer usernames_file.Close()
  reader.Map(f.userNamesReader)

  return &f
}

func (f *Faker) userNamesReader(str string) {
  f.UserNames = append(f.UserNames,str)
  return
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
  return f.UserNames[rand.Intn(len(f.UserNames))]
}

