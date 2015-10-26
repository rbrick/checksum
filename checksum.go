package main

import (
    "flag"
    "fmt"
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "os"
    "io/ioutil"
)

func main() {

    // What sums we need
    checks := map[string]*bool{
        "md5":    flag.Bool("md5", false, "flags"),
        "sha1":   flag.Bool("sha1", false, "flags"),
        "sha256": flag.Bool("sha256", false, "flags"),
        "sha512": flag.Bool("sha512", false, "flags")}

    funcs := map[string]func([]byte) interface{}{
        "md5": func(bytes []byte) interface{} {return md5.Sum(bytes)},
        "sha1": func(bytes []byte) interface{} {return sha1.Sum(bytes)},
        "sha256": func(bytes []byte) interface{} {return sha256.Sum256(bytes)},
        "sha512": func(bytes []byte) interface{} {return sha512.Sum512(bytes)},
    }

    // Parse the flag
    flag.Parse()

    // Parse the arguments
    for _, i := range flag.Args() {
        for k, v := range checks {
            if *v {
                if exists(i) {
                    data, _ := ioutil.ReadFile(i)
                    fmt.Printf("%s=%x [%s]\n", i, funcs[k](data), k)

                }
            }
        }
    }

}

/*
  Tests if a file exists
*/
func exists(s string) bool {
    _, err := os.Stat(s)
    return err == nil
}


