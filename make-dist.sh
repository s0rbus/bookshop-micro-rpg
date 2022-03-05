#!/bin/bash

mkdir -p dist/expansions
find expansions -name '*.so' -exec cp {} dist/expansions/ \;
cp $(basename `pwd`) dist/
