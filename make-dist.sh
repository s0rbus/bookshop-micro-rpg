#!/bin/bash

mkdir -p bin/expansions
find expansions -name '*.so' -exec cp {} bin/expansions/ \;
