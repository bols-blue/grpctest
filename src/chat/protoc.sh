#!/bin/bash -e

protoc --go_out=plugins=grpc:. ./*.proto
