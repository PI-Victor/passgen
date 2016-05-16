Passgen
---

*Opinionated* random password generator package for go applications.

Creates a random password that:  
* is not based on a time event
* has at least 8 characters
* does not contain a complete word
* contains at least one character from each of these types:  
 - ` ~ ! @ # $ % ^ & * ( ) _ - + = { } [ ] \ | : ; " ' < > , . ? /
 - 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
 - abcdefghijklmnopqrstuvwxyz
 - ABCDEFGHIJKLMNOPQRSTUVWXYZ


```go
import "github.com/PI-Victor/passgen"

// passGen accepts an int as a password length (min. 8 characters). It will
// return a generated password as a string or an error if your desired password
// length is lower than 8 characters.

randomPass, err := passGen(10, "")
if err != nil {
  return panic(err)
}
```

If you really need a password that you can remember, the package can map some
alphabet letters to special characters to enhance its security. For example you
can send it one or more words:

```go
import "github.com/PI-Victor/passgen"

randomPass, err := passGen(10, "This Is My PassWord")
if err != nil {
  panic(err)
}

fmt.Println(randomPass)

```
becomes:
```
7h!5/!5+My\P455W0rd
```


the mappings are:
```
a:4
o:0
l:1
i:!
s:5
t:7
e:3

```

* NOTE: if the length of the desired password is bigger than the length of the word, it
will be mixed in with random characters till it reaches the desired length.
* NOTE: When sending multiple words if the contain blank spaces they will be replaced by a random special character.


Released under [Apache Vers. 2.0](http://www.apache.org/licenses/)
