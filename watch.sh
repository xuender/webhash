#!/bin/bash
command="/home/$(whoami)/bin/webhash"
report=`$command -r`

if [ "$report" != "" ]; then
  XDG_RUNTIME_DIR=/run/user/$(id -u) notify-send "Webhash 提示:" "$report" -u critical -i applications-internet
  $command update
fi
