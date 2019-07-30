package main

import "github.com/fogleman/gg"

func main() {
    dc := gg.NewContext(1000, 1000)
    dc.DrawCircle(500, 500, 200)
    dc.SetRGB(250, 5,155)
    dc.Fill()
    dc.SavePNG("out.png")
}
