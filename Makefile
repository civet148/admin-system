#SHELL=/usr/bin/env bash

CLEAN:=
BINS:=
DATE_TIME=`date +'%Y%m%d %H:%M:%S'`
COMMIT_ID=`git rev-parse --short HEAD`
MANAGER_DIR=${PWD}
CONSOLE_CODE=/tmp/admin-system-frontend

build:
	rm -f admin-system
	go mod tidy && go build -ldflags "-s -w -X 'main.BuildTime=${DATE_TIME}' -X 'main.GitCommit=${COMMIT_ID}'" -o admin-system cmd/main.go
.PHONY: build
BINS+=admin-system

nodejs:
	curl -sL https://deb.nodesource.com/setup_14.x | sudo -E bash - && sudo apt update && sudo apt install -y nodejs build-essential && sudo npm install -g yarn

console:
	rm -rf ${CONSOLE_CODE} && git clone -b master https://git.your-enterprise.com/admin-system-frontend.git ${CONSOLE_CODE}
	cd ${CONSOLE_CODE} && git log -2 && npm install && npm run build:prod

docker:
	docker build --build-arg GIT_USER=${GIT_USER} --build-arg GIT_PASSWORD=${GIT_PASSWORD} --tag admin-system -f Dockerfile .

# 检查环境变量
env-%:
	@ if [ "${${*}}" = "" ]; then \
	    echo "Environment variable $* not set"; \
	    exit 1; \
	fi


clean:
	rm -rf $(CLEAN) $(BINS)
.PHONY: clean
