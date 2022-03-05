#!/bin/bash

cd expansions
for d in * ; do
    if [ -d "$d" ]; then
       echo "building $d"
       cd $d
       go build -trimpath -buildmode=plugin
       #go build -trimpath -buildmode=plugin -o "../../bin/expansions/$d.so"
       cd ..
    fi
done
