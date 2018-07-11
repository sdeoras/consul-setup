#!/bin/sh
HOSTNAME=`hostname -I | awk '{print $1}'`
consul agent \
	-config-dir=/etc/consul.d/client \
	-bind=${HOSTNAME} \
	-enable-script-checks=true \
	-http-port=8500 \
	-client=127.0.0.1

