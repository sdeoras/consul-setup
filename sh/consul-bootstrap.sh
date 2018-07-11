#!/bin/sh
HOSTNAME=`hostname -I | awk '{print $1}'`
consul agent \
	-config-dir=/etc/consul.d/bootstrap \
	-bind=${HOSTNAME} \
	-enable-script-checks=true
