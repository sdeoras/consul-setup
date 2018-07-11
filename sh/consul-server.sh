#!/bin/sh
HOSTNAME=`hostname -I | awk '{print $1}'`
consul agent \
	-config-dir=/etc/consul.d/server \
	-bind=${HOSTNAME} \
	-enable-script-checks=true
