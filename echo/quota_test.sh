#!/bin/bash
for i in `seq 1 200`
do
  curl -X POST "http://104.197.225.150/echo?key=AIzaSyDazsnB8uE5dB8RrvHG_plHTrLWdI3pFRA" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"message\": \"string\", \"latency\": 0}";echo $i
done
