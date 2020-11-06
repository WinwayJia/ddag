#!/bin/bash

go build -buildmode=plugin
mv ./decoder.so ../../lib 
