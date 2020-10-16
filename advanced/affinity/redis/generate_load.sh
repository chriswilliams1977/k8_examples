# [START container_helloapp_redis_generate_load]
#!/bin/bash
# Usage: generate_load.sh <IP> <QPS>_

IP=$1
QPS=$2

while true
  do for N in $(seq 1 $QPS)
    do curl -I -m 5 -s -w "%{http_code}\n" -o /dev/null http://${IP}/ >> output &
    done
  sleep 1
done
# [END container_helloapp_redis_generate_load]