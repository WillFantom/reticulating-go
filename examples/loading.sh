#!/bin/bash
set -e

for ((k = 0; k <= 20 ; k++))
do
    quote=$(curl --silent https://sims.willfantom.com/api/messages/random | jq .message )
    echo -n "[ "
    for ((i = 0 ; i <= k; i++)); do echo -n "#"; done
    for ((j = i ; j <= 20 ; j++)); do echo -n " "; done
    v=$((k * 5))
    echo -n " ] "
    echo -n "$v % || $quote                 " $'\r'
    sleep 0.75
done
