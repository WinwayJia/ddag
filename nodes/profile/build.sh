#!/bin/bash

go build -buildmode=plugin
mv ./profile.so ../../lib 
