package main

import (
    "bufio"
    "fmt"
    "strconv"
    "os"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter the index of the fib number to caluclate: ");
  var index, _ = reader.ReadString('\n');
  var index2, err = strconv.Atoi(index);
  if(err){
    fmt.Println(err);
    return
  }
  var a = 1;
  var b = 0;
  var c = 0;
  var i = 0;
   for i < index2 {
       c = a + b;
       b = a;
       a = c;
       i++;
   }
   fmt.Print("The %dth fib number is: ", index)
   fmt.Println("%d", c);
}
