#!/bin/bash

code_climate_binary_url=''

if [ "${TRAVIS_OS_NAME}" = 'osx' ]; then
    # Change to binary url build for macOS
    code_climate_binary_url='https://codeclimate.com/downloads/test-reporter/test-reporter-0.10.1-darwin-amd64'
else
    # Change to binary url build for Linux
    code_climate_binary_url='https://codeclimate.com/downloads/test-reporter/test-reporter-latest-linux-amd64'
fi

curl -L "${code_climate_binary_url}" > ./cc-test-reporter
