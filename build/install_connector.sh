#!/bin/bash

connector_version="10.7.5"

wget https://github.com/confluentinc/kafka-connect-jdbc/releases/download/v${connector_version}/confluentinc-kafka-connect-jdbc-${connector_version}.tar.gz

tar -xvf confluentinc-kafka-connect-jdbc-${connector_version}.tar.gz

mv confluentinc-kafka-connect-jdbc-${connector_version} /usr/share/confluent-hub-components
