#!/bin/bash

cd expansion-src
for d in * ; do
    if [ -d "$d" ]; then
       echo "building $d"
       cd $d
       go build -trimpath -buildmode=plugin
       cd ..
    fi
done
