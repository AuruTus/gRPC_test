#!/bin/bash

python -m grpc_tools.protoc -I./pb --python_out=./py_client/src --grpc_python_out=./py_client/src ./pb/helloworld.proto