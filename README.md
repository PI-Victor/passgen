Passgen
---

*Opinionated* random password generator package for your go application.  
Creates a random password that:  
* is not based on a time event
* at least 8 characters
* does not contain a complete word
* contains at least one character from each of these types:  
 - ` ~ ! @ # $ % ^ & * ( ) _ - + = { } [ ] \ | : ; " ' < > , . ? /
 - 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
 - abcdefghijklmnopqrstuvwxyz
 - ABCDEFGHIJKLMNOPQRSTUVWXYZ


```go
import "github.com/PI-Victor/passgen"

randomPass, err := passGen(10)
if err != nil {
  return panic(err)
}
```
