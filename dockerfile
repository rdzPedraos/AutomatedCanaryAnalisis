FROM golang:1.24.3 as build

ARG FUNCTION_NAME

WORKDIR /lambda

COPY go.mod go.sum ./
COPY src ./src

RUN cd src/functions/${FUNCTION_NAME} \
    && go build -tags lambda.norpc -o /lambda/bootstrap .

# Copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2023
COPY --from=build /lambda/bootstrap ./main
ENTRYPOINT [ "./main" ]
