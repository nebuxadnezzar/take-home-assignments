#!/bin/bash +x

# run it from several terminals simultaneously
for i in {1..100}
do
    curl -L "localhost:5555/api/memes?query=becky" -H "X-Token:secret2"
done
