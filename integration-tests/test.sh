#!/bin/bash

set -eo pipefail 

dirname=`dirname $0`
templateFile="$dirname/test.tmpl"
go run $1 $templateFile

cwd=$(pwd)
expectedOut="$dirname/expected-out"
actualOut="$cwd/values.yaml"

if cmp -s "$expectedOut" "$actualOut"; then
    code=0;
else
    printf 'Expected output and actual output are not the same';
    printf 'Expected:'
    cat $expectedOut
    printf 'ActualOut'
    cat $actualOut
    code=1;
fi

rm $actualOut;

exit $code;
