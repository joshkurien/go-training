package main

import (
  "fmt"
  "io/ioutil"
  "os"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func encryp(pass, fname string) {
  var flag bool
  var i int

  fil, err := ioutil.ReadFile(fname)
  check(err)

  tmp := []uint8(pass)

  
  out := tmp[:]  


  flag = true
  for i = 0; i < len(fil) ;  {  
    
    for j := range tmp {
            
      if flag==true {
        fil[i]+=tmp[j]
      } else {
        fil[i]-=tmp[j]
      }
      flag = ! flag
      i++
      if i == len(fil) {
        break
      }
    }
  }


  flag = true
  for i = range tmp {
    if flag==true {
      tmp[i]+=1
    } else {
      tmp[i]-=1
    }
    flag = ! flag
    out[i]=tmp[i]
  }

  
  out = append(out,fil...)

  err = ioutil.WriteFile("output.txt", out, 0644)
  check(err)
}


func decryp(pass, fname string) bool {
  var flag bool
  var i int

  fil, err := ioutil.ReadFile(fname)
  check(err)

  tmp := []uint8(fil[:len(pass)])

  flag = true
  for i = range tmp {
    if flag==true {
      tmp[i]-=1
    } else {
      tmp[i]+=1
    }
    flag = ! flag
  }
  
  if(string(tmp)!=pass) {
    fmt.Println("Sorry, incorrect Password")
    return false
  }

  out := fil[len(pass):]

   flag = true

  for i = 0; i < len(out) ;  {  
    
    for j := range tmp {
            
      if flag==true {
        out[i]-=tmp[j]
      } else {
        out[i]+=tmp[j]
      }
      flag = ! flag
      i++

      if i == len(out) {
        break
      }
    }
  }

  err = ioutil.WriteFile("deciphered.txt", out, 0644)
  check(err)

  return true
}





func main() {
    
  var pass string
  var choice int

  fmt.Println("Enter Choice:\n1.Encrypt\n2.Decrypt")
  fmt.Scanf("%d",&choice)
  
  fmt.Print("Enter Password: ")
  fmt.Scanf("%s",&pass)

  switch choice {
  case 1:
    encryp(pass,string(os.Args[1]))
  case 2:
    decryp(pass,string(os.Args[1]))

  default:
    fmt.Println("Sorry, incorrect Choice, please choose between 1 and 2!!")
  }
}