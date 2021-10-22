#!/bin/bash
curl -X PUT -H 'Content-Type: application/json' -d @put.demo.json  http://localhost:3000/pokemon
#http://ec2-3-21-40-16.us-east-2.compute.amazonaws.com:3000/pokemon