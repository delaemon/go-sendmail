#!/bin/bash

XC_ARCH=${XC_ARCH:-386 amd64}
XC_OS=${XC_OS:-linux darwin windows}

rm -rf bin/
gox \
    -os="${XC_OS}" \
    -arch="${XC_ARCH}" \
    -output "bin/{{.OS}}_{{.Arch}}/{{.Dir}}"

