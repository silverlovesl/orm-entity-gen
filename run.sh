#!/bin/bash

# echo -e "Input table name"
# read TABLE_NAME
TABLE_NAME="$1"

OUTPUT=$GOPATH/src/bitbucket.org/beecomb-grid/grid-ai/gridai_api/entities

# Output file name
if [ -z "$1" ]
then
	TABLE_NAME=all
fi

go build ./cmd/orm-entity-gen
./orm-entity-gen \
-user=grow \
-password=grow2015 \
-database=grow_development \
-table_name=${TABLE_NAME} \
-orm_name=gorm \
-output=${OUTPUT}

goimports -w $OUTPUT