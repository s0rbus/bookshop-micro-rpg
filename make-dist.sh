#!/bin/bash

mkdir -p expansions
find expansion-src -name '*.so' -exec cp {} expansions/ \;
