package printer

import (
    //"bufio"
    //"fmt"
    "os"
)

func TextToFile(text string, filename string) {
    file, err := os.Create(filename + ".txt")
      if err != nil {
      }
      defer file.Close()

      file.WriteString(text)
}
