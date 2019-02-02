mkdir .build
flutter build bundle
go run engineDownloader.go
go build main.go
.\main.exe