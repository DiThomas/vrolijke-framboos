#!/bin/bash
# build project for raspberry
env GOOS=linux GOARCH=arm GOARM=6 go build

#copy to
mv framboos ./bin/framboos
scp ./bin/framboos dithomas@192.168.3.230:/home/dithomas/framboos