all: deps

deps:
	go get -d -u github.com/hybridgroup/gobot/... && go install github.com/hybridgroup/gobot/platforms/raspi

run-pub:
	go run ./main.go

build:
	go build

restart_pi:
	ssh pi@$(RASPI_HOST) "sudo reboot"

deploy:
	GOARM=6 GOARCH=arm GOOS=linux go build
	scp ./delorean-core pi@$(RASPI_HOST):/home/pi/
	ssh pi@$(RASPI_HOST) "./delorean-core"

