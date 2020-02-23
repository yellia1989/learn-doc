#!/bin/bash

while true; do sleep 1;date;netstat -n -p tcp | grep 8888; done
