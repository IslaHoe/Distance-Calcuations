#!/bin/sh
rm customers.json
curl https://s3.amazonaws.com/intercom-take-home-test/customers.txt >> customers.txt
sed '1s/^/[/; $!s/$/,/;$s/$/]/' customers.txt | jq >> customers.json 
rm customers.txt



go run main.go