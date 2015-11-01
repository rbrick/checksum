package main

import (
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "flag"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {

    // What sums we need
    checks := map[string]*bool{
        "md5":    flag.Bool("md5", false, "Get the md5 sum"),
        "sha1":   flag.Bool("sha1", false, "Get the sha1 sum"),
        "sha256": flag.Bool("sha256", false, "Get the sha256 sum"),
        "sha512": flag.Bool("sha512", false, "Get the sha512 sum")}

    funcs := map[string]func([]byte) interface{}{
        "md5":    func(bytes []byte) interface{} { return md5.Sum(bytes) },
        "sha1":   func(bytes []byte) interface{} { return sha1.Sum(bytes) },
        "sha256": func(bytes []byte) interface{} { return sha256.Sum256(bytes) },
        "sha512": func(bytes []byte) interface{} { return sha512.Sum512(bytes) },
    }

    // Parse the flag
    flag.Parse()

    if flag.NFlag() < 1 || flag.NArg() < 1 {
        fmt.Println("checksum [-hashingAlgorithm...] [files...]")
        flag.PrintDefaults()
        return
    }

    // Parse the arguments
    for _, i := range flag.Args() {
        for k, v := range checks {
            if *v {
                if exists(i) {
                    data, _ := ioutil.ReadFile(i)

                    result := make(chan interface{})
                    go func() {
                       result <- funcs[k](data)
                    } ()

                    fmt.Printf("%s=%x [%s]\n", i, <-result, k)
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
