#!/bin/bash

PRIVILEGED=false

curl -v -X PUT $MARATHON_API_URL/v2/apps/shurenyun-$TASKENV-$SERVICE -H Content-Type:application/json -d \
'{
      "id": "shurenyun-'$TASKENV'-'$SERVICE'",
      "cpus": '$CPUS',
      "mem": '$MEM',
      "instances": '$INSTANCES',
      "constraints": [["hostname", "LIKE", "'$DEPLOYIP'"]],
      "container": {
                     "type": "DOCKER",
                     "docker": {
                                     "image": "'$SERVICE_IMAGE'",
                                     "network": "BRIDGE",
                                     "forcePullImage": '$FORCEPULLIMAGE',
                                     "privileged": '$PRIVILEGED',
                                     "portMappings": [
                                             { "containerPort": 80, "hostPort": 0, "protocol": "tcp"}
                                     ]
                                     "volumes": [
                                            {
                                              "containerPath": "/go/bin/primarykey",
                                              "hostPath": "/data/drone",
                                              "mode": "RW"
                                            }
                                    ]
                                }
                   },
      "healthChecks": [{
               "path": "/api/v3/health/harbor",
               "protocol": "HTTP",
               "gracePeriodSeconds": 300,
               "intervalSeconds": 60,
               "portIndex": 0,
               "timeoutSeconds": 20,
               "maxConsecutiveFailures": 3
           }],
      "env": {
                    "BAMBOO_PUBLIC": "true",
                    "BAMBOO_PROXY":"true",
                    "BAMBOO_HTTP_PROTOCOL":"true"
             },
      "uris": [
               "'$CONFIGSERVER'/config/demo/config/registry/docker.tar.gz"
       ]
}'
