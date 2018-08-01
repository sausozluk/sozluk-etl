EXECUTABLE := etl
TARGET := main.go

make:
	GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -o $(EXECUTABLE)-linux $(TARGET)
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o $(EXECUTABLE)-windows.exe $(TARGET)
	GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -o $(EXECUTABLE)-darwin $(TARGET)
