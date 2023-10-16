ifneq ($(MODE), production)
	include .env
	export
endif

OUT_DIR = ./out

MAIN_FILE = ./cmd/main
OUTPUT_FILE = main

MAIN_SEEDER_FILE = ./cmd/seeder
OUTPUT_SEEDER_FILE = seeder

env:
	echo ${ENV}

build:
	go build -o ./out/${OUTPUT_FILE} ${MAIN_FILE}

run: build
	${OUT_DIR}/${OUTPUT_FILE}

build-seeder:
	go build -o ./out/${OUTPUT_SEEDER_FILE} ${MAIN_SEEDER_FILE}

seeder: build-seeder
	${OUT_DIR}/${OUTPUT_SEEDER_FILE}

clean:
	go clean
	rm -rf out
