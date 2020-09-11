#! /bin/sh

ab -t 10 -n 100 -c 100 http://0.0.0.0:3000/status & \
# ab -p addPlayer.json -T application/json  -c 10 -n 2000 http://0.0.0.0:3000/add & \
ab -p killPlayer.json -T application/json  -c 10 -n 2000 http://0.0.0.0:3000/kill & \
ab -p movePlayer.json -T application/json  -c 10 -n 2000 http://0.0.0.0:3000/move 
