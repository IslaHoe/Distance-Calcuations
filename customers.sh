#!/bin/sh
rm customers.txt
curl https://s3.amazonaws.com/intercom-take-home-test/customers.txt >> customers.txt

go build
./codingChallange
open output.txt
