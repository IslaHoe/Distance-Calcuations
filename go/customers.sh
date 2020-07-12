#!/bin/sh
rm customers.json
curl https://s3.amazonaws.com/intercom-take-home-test/customers.txt >> customers.json 

go run main.go
code output.txt