#!/bin/bash
for i in {1..3}
do
    for i in {1..10011}
    do
        echo "o" >> filechange.txt
    done
    git add .
    git commit -m"ok"
    git push origin main

done
