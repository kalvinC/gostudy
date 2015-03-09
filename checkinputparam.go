package main
import(
  "fmt"
  "os"
  )

func main(){
  if len(os.Args)>=2{
     fmt.Println("yes")
     os.Exit(0)
  }
  fmt.Println("No")
}

