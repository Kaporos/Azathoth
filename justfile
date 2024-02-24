run:
	clear && sh watch.sh
connect:
	#!/bin/sh
	while true
	do
		ssh 127.0.0.1 -p 4022
		sleep 0.5
		clear 
	done
