# Go HTTP Post Request CLI

This Go program accepts command-line parameters and makes an HTTP POST request based on the provided JSON data.

## Usage

### Build the Program

```bash
cd cmd && go build
```
### Run the Program 

```ex :
./cmd -p '{ "command": "POST", "url": "https://vp-test.requestcatcher.com/test", "body": "{ \"Id\": 12345, \"Customer\": \"John Smith\", \"Quantity\": 1, \"Price\": 10.00 }", "header": { "X-API-KEY": "test_KEY" }, "url_params": { "user_id": "1", "company": "10" } }'
```

## Command-Line Options
    -p: Specifies raw parameters in JSON format.


## Dependencies
    This program uses the standard Go packages and does not require any additional dependencies.
