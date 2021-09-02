#!/bin/bash

templateFile="./test.tmpl"
go run ../ $templateFile

expectedOut="./expected-out"
actualOut="./values.yaml"

if cmp -s "$expectedOut" "$actualOut"; then
    code=0;
else
    printf 'Expected output and actual output are not the same';
    code=1;
fi

rm $actualOut;

exit $code;
