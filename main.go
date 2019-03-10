package main

import (
	"image"
	_ "image/png"
	"os"
	"path"
	"runtime"	
	"fmt"

	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/embedder"
	"github.com/go-gl/glfw/v3.2/glfw"

	"github.com/ynsgnr/aria2go"

)

var downloader aria2go.Aria2go

func main() {

	_, currentFilePath, _, _ := runtime.Caller(0)
	dir := path.Dir(currentFilePath)

	initialApplicationHeight := 600
	initialApplicationWidth := 800

	options := []flutter.Option{
		flutter.ProjectAssetsPath(dir + "/build/flutter_assets"),

		// This path should not be changed. icudtl.dat is handled by engineDownloader.go
		flutter.ApplicationICUDataPath(dir + "/icudtl.dat"),

		flutter.ApplicationWindowDimension(initialApplicationWidth, initialApplicationHeight),
		flutter.WindowIcon(iconProvider),
		flutter.OptionVMArguments([]string{
			// "--disable-dart-asserts", // release mode flag
			// "--disable-observatory",
			"--observatory-port=50300",
		}),
		flutter.OptionAddPluginReceiver(aria2Plugin, "aria2Plugin"),
	}

	downloader = aria2go.New()
	downloader.KeepRunning()

	if err := flutter.Run(options...); err != nil {
		fmt.Printf("Failed running the Flutter app: %v\n", err)
		os.Exit(1)
	}

	downloader.Finalize()

}

func iconProvider() ([]image.Image, error) {
	_, currentFilePath, _, _ := runtime.Caller(0)
	dir := path.Dir(currentFilePath)
	imgFile, err := os.Open(dir + "/assets/icon.png")
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}
	return []image.Image{img}, nil
}

func aria2Plugin(
	platMessage *embedder.PlatformMessage,
	flutterEngine *embedder.FlutterEngine,
	window *glfw.Window,
) bool {
	return true
}
