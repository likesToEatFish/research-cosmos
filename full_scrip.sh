#!/bin/bash
for i in {1..1}
do
    for i in {1..10011}
    do
        echo "o" >> filechange.txt
    done
    ./git.sh

done
