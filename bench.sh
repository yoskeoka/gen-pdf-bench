#!/bin/bash

attempt=10

echo "node version" $(node -v)

( 
    cd puppeteer
    for i in `seq 1 ${attempt}`
    do
        echo "puppeteer bench #$i"
        time node index.js
    done
)

for i in `seq 1 ${attempt}`
do
    echo "wkhtmltopdf bench #$i"
    time ./testwk
done
