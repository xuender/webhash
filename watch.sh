#!/bin/bash
# 监听网页变化
user=`whoami`
command="/home/$user/bin/webhash"
report=`$command -r`

if [ "$report" != "" ]; then
  XDG_RUNTIME_DIR=/run/user/$(id -u) notify-send "webhash提示" "$report" -u critical -i applications-internet
  $command update
fi
