package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"PPPgo/lib"

	"github.com/disintegration/imaging"
)

func main() {

	cpy := 0
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		os.Exit(1)
	}
	imagePath := os.Args[1]
	if len(os.Args) == 2 {
		fmt.Print("Enter NOC >>> ")
		fmt.Scanf("%d", &cpy)
	} else {
		cpy, _ = strconv.Atoi(os.Args[2])
	}

	paper := a29.Paper{}
	paper.New(2480, 3508)

	//imaging.Save(im.img, "out_example.jpg")

	im, _ := imaging.Open(imagePath)
	paper.Add(im, 27.0, 0.0, cpy, "Unknown")
	paper.PrintInfo()
	imaging.Save(paper.Render(3.0, 0, 1.0, true), filepath.Join(dir, "Paper.jpg"))

	if len(os.Args) > 3 {
		cmd := exec.Command(filepath.Join(dir, "Printer29"), filepath.Join(dir, "Paper.jpg"), os.Args[3])
		cmd.Run()
	} else {
		cmd := exec.Command(filepath.Join(dir, "Printer29"), filepath.Join(dir, "Paper.jpg"))
		cmd.Run()
	}

}
