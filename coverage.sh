#!/bin/bash

# Generates a  single test-coverage profile for all packages.
OUTFILE=coverage.out
rm -f $OUTFILE
touch $OUTFILE

for package in $(go list ./... | grep -Ev 'vendor'); do
    IFS='/' read -ra parts <<< "$package"

    TMP_OUTFILE=${parts[${#parts[@]}-1]}.out
    touch $TMP_OUTFILE

    go test -covermode=count -coverprofile=$TMP_OUTFILE $package
    if [ ! -s $OUTFILE ]; then
        cat $TMP_OUTFILE >> $OUTFILE
    else
        tail -n +2 $TMP_OUTFILE >> $OUTFILE
    fi
    rm $TMP_OUTFILE
done
