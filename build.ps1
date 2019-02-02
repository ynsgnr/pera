mkdir .build
flutter build bundle
go run engineDownloader.go
CMD /C set CGO_LDFLAGS=-L%cd% #taken from go-flutter-desktop-embedder
$env:LIBRARY_PATH=(Get-Item -Path ".\").FullName #set the path for mingw to find flutter_engine.dll
go build main.go
.\main.exe