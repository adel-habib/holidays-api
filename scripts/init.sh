#!/bin/bash

HOST=$POSTGRES_CONTAINER_NAME
PORT=$POSTGRES_INTERNAL_PORT

echo "Checking DB connection; Hostname: ${HOST}, Port: ${PORT} ..."
i=0
until [ $i -ge 10 ]
do
  # check if a tcp connection with the database is possible 
  nc -z $HOST $PORT && break

  i=$(( i + 1 ))

  echo "$i: Waiting for DB 1 second ..."
  sleep 1
done

if [ $i -eq 10 ]
then
  echo "DB connection refused, terminating ..."
  exit 1
fi

echo "DB is up ..."

chmod +x /app
./app