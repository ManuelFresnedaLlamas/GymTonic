#!/bin/bash

chmod 777 /dst/gym

if [ "$GYM_GO_GET" = true ] && [ "$GYM_GO_SERVER" != true ]
then
	cd /src && tail -f /dev/null
fi

if [ "$GYM_GO_VENDOR" = true ]
then
	echo "STARTING VENDORING"
	cd /src && go mod vendor
fi

if [ "$GYM_GO_SERVER" = true ] && [ "$GYM_GO_DEBUG" = true ]
then
	echo "STARTING SEVER WITH DEBUG"
	cd /src && go build -gcflags="all=-N -l" -o /dst/server -mod vendor ./cmd && cd /dst && dlv --listen=:40000 --headless=true --api-version=2 --accept-multiclient exec ./server
fi

if [ "$GYM_GO_SERVER" = true ] && [ "$GYM_GO_DEBUG" != true ]
then
	echo "STARTING SEVER"
	cd /src && go build -o /dst/server -mod vendor ./cmd && cd /dst && ./server
fi