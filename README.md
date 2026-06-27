Este proyecto busca mostrar usando IaaC la configuración necesaria para hacer despliegues canario automatizados para la fase continuous deployment en DevOps

# Commands used:

## Create go project

```sh
go mod init github.com/rdzPedraos/AutomatedCanaryAnalisis
go get github.com/aws/aws-lambda-go/lambda
```

# Attachments

## How compile binary for aws lambda

Check [guide](https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html)

```sh
# build binary file
GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap ./src/functions/calculate

# zip file
zip myFunction.zip bootstrap

# upload to aws
aws lambda create-function --function-name myFunction \
--runtime provided.al2023 --handler bootstrap \
--role arn:aws:iam::111122223333:role/lambda-ex \
--zip-file fileb://myFunction.zip
```

## Add commitID in binary

Update command used for compiled binary

```sh
GOOS=linux GOARCH=amd64 \
go build \
-tags lambda.norpc \
-ldflags="-X github.com/rdzPedraos/AutomatedCanaryAnalisis/src/libraries/logger.commitID=$(git rev-parse HEAD)" \
-o bootstrap \
./src/functions/calculate
```

# [With docker](https://docs.aws.amazon.com/lambda/latest/dg/go-image.html)

## Build

```sh
docker build --build-arg FUNCTION_NAME=calculate --platform linux/amd64 --provenance=false -t docker-image:test .
```

## Test

```sh
# Run process
docker run -d -p 9000:8080 --entrypoint /usr/local/bin/aws-lambda-rie docker-image:test ./main

# Test endpoint
curl "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{}'
```
