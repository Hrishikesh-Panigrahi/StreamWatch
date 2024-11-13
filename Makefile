GOCMD=go
BUILD_DIR=./tmp
BINARY_NAME=main.exe

all: run

build: 
	$(GOCMD) build -v -o $(BUILD_DIR)/$(BINARY_NAME) .

run: build
	./$(BUILD_DIR)/$(BINARY_NAME)
	
clean:
	$(GOCMD) clean
	rm -rf $(BUILD_DIR)

help:
	@echo "make build - build the binary file"
	@echo "make run - run the application"
	@echo "make clean - remove the binary file"

.PHONY: run/live
run/live:
    go run github.com/cosmtrek/air@v1.43.0 \
        --build.cmd "make build" --build.bin "/tmp/bin/${binary_name}" --build.delay "100" \
        --build.exclude_dir "" \
        --build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
        --misc.clean_on_exit "true"

.PHONY: all run clean