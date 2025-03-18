#!/bin/bash
RECIPE_TABLE="Recipe"

awslocal dynamodb \
    create-table \
    --region na-east-1 \
    --table-name "$RECIPE_TABLE" \
    --attribute-definitions \
        AttributeName=category,AttributeType=S \
        AttributeName=id#updatedAt,AttributeType=S \
    --key-schema \
        AttributeName=category,KeyType=HASH \
        AttributeName=id#updatedAt,KeyType=RANGE \
    --provisioned-throughput \
        ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --table-class STANDARD

awslocal dynamodb batch-write-item --request-item "file:///tmp/recipes.json"

