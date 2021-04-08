// main.go

package main

func main() {
    a := App{}
    a.Initialize(
        "davidschindler",
        "",
        "davidschindler")

    a.Run(":8010")
}
