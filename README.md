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
