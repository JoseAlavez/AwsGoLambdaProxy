# Motivation
An aws go lambda to support aws gateway LAMBDA_PROXY async executions.

# Configuration
Required Variables:
* **region**: Required to create a new AWS session for invoking the proxied lambda.
* **functionName**: Lambda function name to invoke.
* **invocationType**: Lambda invocation type to use.
* **logType**: Lambda log type to use.

Optional Variables:
* **qualifier**: Version of lambda function to use.
* **clientContext**: Additional client data to append in request.

More info in: https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html

# Build Example on Linux
```
cd $GOPATH/src/AwsGoLambdaProxy
env GOOS=linux GOARCH=amd64 go build -o /tmp/AwsGoLambdaProxy
zip -j /tmp/AwsGoLambdaProxy.zip /tmp/AwsGoLambdaProxy
aws lambda update-function-code --function-name FUNCTION_NAME --zip-file fileb:///tmp/AwsGoLambdaProxy.zip
```
