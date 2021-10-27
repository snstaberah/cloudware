#!/bin/bash
set -ex

echo "START_FLAG=" $START_FLAG

./cloudware


exec "$@"
