package main

import (
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "math"
    "os"
)

func main() {
    width, height := 300, 300

    // Create a new RGBA image with the given dimensions
    img := image.NewRGBA(image.Rect(0, 0, width, height))

    // Define some colors
    bgColor := color.RGBA{50, 50, 150, 255}
    lineColor := color.RGBA{255, 215, 0, 255}

    // Fill the background with bgColor
    draw.Draw(img, img.Bounds(), &image.Uniform{bgColor}, image.Point{}, draw.Src)

    // Draw some pattern or design
    for i := 0; i < width; i++ {
        y := int(math.Sin(float64(i)/float64(width)*2*math.Pi)*float64(height)/2 + float64(height)/2)
        img.Set(i, y, lineColor)
    }

    // Draw some circles
    for r := 0; r < width/2; r += 20 {
        for theta := 0.0; theta < 2*math.Pi; theta += 0.01 {
            x := width/2 + int(float64(r)*math.Cos(theta))
            y := height/2 + int(float64(r)*math.Sin(theta))
            if x >= 0 && x < width && y >= 0 && y < height {
                img.Set(x, y, lineColor)
            }
        }
    }

    // Save to file
    file, err := os.Create("amazing_logo.png")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    png.Encode(file, img)
}
