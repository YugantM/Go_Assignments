package main

import "github.com/fogleman/gg"

func main() {
    dc := gg.NewContext(16, 16)
    dc.DrawCircle(4, 4, 2)
    dc.SetRGB(250, 5,155)
    dc.Fill()
    dc.SavePNG("out.png")
}
