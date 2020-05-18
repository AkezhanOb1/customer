#!/bin/bash
protoc customer.proto  --go_out=plugins=grpc:.