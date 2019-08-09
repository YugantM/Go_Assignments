package main

import "github.com/fogleman/gg"

func main() {
    dc := gg.NewContext(400, 400)

    dc.DrawCircle(40, 40, 20)
    dc.SetRGB(250,205,155)
    
    dc.SavePNG("out.png")
}
