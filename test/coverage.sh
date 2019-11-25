#!/usr/bin/env bash
if [ -z "$1" ]; then
    echo "Ha olvidado escribir la ruta del proyecto."
    exit 
fi

PROJECT=$1
SRC_DIRECTORY="/src"
TEST_DIRECTORY=$PROJECT"/test"

rm $TEST_DIRECTORY"/"coverage.txt
rm $TEST_DIRECTORY"/"coverage.html
echo "mode: set" >> $TEST_DIRECTORY"/"coverage.txt

for dt in $(go list ./... | grep -v bui); do
    BASENAME=$(basename $dt)
    cd $TEST_DIRECTORY"/"$BASENAME
    BASENAME_PKG=$(basename $PROJECT)$SRC_DIRECTORY"/"$BASENAME
    
    gotest -coverprofile=profile.out -coverpkg=$BASENAME_PKG -cover
    if [ -f profile.out ]; then
        tail -n +2 profile.out >> $TEST_DIRECTORY"/"coverage.txt
        rm profile.out
    fi
done

go tool cover -html=$TEST_DIRECTORY"/"coverage.txt -o $TEST_DIRECTORY"/"coverage.html