#!/bin/bash

cd expansions
for d in * ; do
    if [ -d "$d" ]; then
       echo "building $d"
       cd $d
       go build -buildmode=plugin -o "$d.so"
       cd ..
    fi
done
