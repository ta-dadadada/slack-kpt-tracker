#!/bin/bash

until mysqladmin ping -h db --silent; do
    echo 'waiting for mysqld to be connectable...'
    sleep 3
done

./api
