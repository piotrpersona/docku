package main

import (
	"os"
	"log"
	"fmt"
	"sync"
	"math/rand"
	"time"
	"io/ioutil"
	"encoding/json"
)

type Image struct {
	Registry string `json:"registry"`
	Tag string `json:"tag"`
}

type ImagesSet map[string]Image

type ImagesMetadata struct {
	Registry string `json:"registry"`
	Images ImagesSet `json:"images"`
}

func readImages(imagesPath string) *ImagesMetadata {
	jsonFile, err := os.Open(imagesPath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	data, _ := ioutil.ReadAll(jsonFile)
	var imagesMetadata ImagesMetadata
	json.Unmarshal(data, &imagesMetadata)
	return &imagesMetadata
}

type networkFunction func(string) string

func measureTime(nf networkFunction) networkFunction {
	return func(argument string) string {
		start := time.Now()
		result := nf(argument)
		fmt.Printf("It took: %v\n", time.Since(start))
		return result
	}
}

func pullImage(image string) string {
	fmt.Printf("Pulling image %s\n", image)
	sleep := rand.Intn(800)
	time.Sleep(time.Millisecond * time.Duration(sleep))
	fmt.Printf("Pull done: %s\n", image)
	return image
}

func tagImage(sourceImage, destinationImage string) string {
	fmt.Printf("Tagged %s with %s\n", sourceImage, destinationImage)
	return destinationImage
}

func pushImage(image string) string {
	fmt.Printf("Pushing image %s\n", image)
	sleep := rand.Intn(800)
	time.Sleep(time.Millisecond * time.Duration(sleep))
	fmt.Printf("Push done: %s\n", image)
	return image
}

func uploadImage(sourceImageName, destinationImageName string, wg *sync.WaitGroup) {
	defer wg.Done()
	pullImageMeasured := measureTime(pullImage)
	pushImageMeasured := measureTime(pushImage)
	sourceImage := pullImageMeasured(sourceImageName)
	destinationImage := tagImage(sourceImage, destinationImageName)
	pushImageMeasured(destinationImage)
}

func imageURL(registry, name, tag string) string {
	return registry + "/" + name + ":" + tag
}

func uploadImages(imagesMetadata *ImagesMetadata) {
	var wg sync.WaitGroup
	wg.Add(len(imagesMetadata.Images))
	registry := imagesMetadata.Registry
	for imageName, imageMeta := range imagesMetadata.Images {
		sourceImage := imageURL(imageMeta.Registry, imageName, imageMeta.Tag)
		destinationImage := imageURL(registry, imageName, imageMeta.Tag)
		go uploadImage(sourceImage, destinationImage, &wg)
	}
	wg.Wait()
}

func postUploadActions(start time.Time) {
	fmt.Println("Done!")
	fmt.Printf("In time: %v\n", time.Since(start))
}

func main() {
	imagesMetadata := readImages("images.json")
	start := time.Now()
	uploadImages(imagesMetadata)
	postUploadActions(start)
}
