# go-benchmark-kafka-api

The purpose of this repository is to do a performance benchmark for Kafka Producer in GoLang via REST API. 

## Results

With a NodeJS producer, this repository was able to produce 10,000 messages in 26.4 seconds. If you are able to benchmark it with another tool like CLI or JMeter then kindly share the results. 

PRs are always welcome for benchmarking and improving the code.

## Getting started

First of all you need Go and Kafka to be installed on your local machine. Before you start make sure you have Kafka running on your local machine. The next step would be to clone the repository and run the following commands within the `go-benchmark-kafka-api` repository.

```shell
go build main.go
./main
```

Now to call the API you can use any tool right from JMeter or write down your own script in any language. A simple `curl` command to invoke this API is as follows

```shell
curl --location --request POST 'http://localhost:1234/produce' \
--header 'Content-Type: application/json' \
--data-raw '{
	"topic": "test_topic",
	"message": "Man what up!"
}'
```