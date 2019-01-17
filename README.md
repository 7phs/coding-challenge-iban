## Requirements

1. Go 1.12 beta2+

2. Libraries which used in a project will downloaded by ```go mod vendor``` command

## Build

### Manual

A developer should load all dependencies library before building a binary, run command:
```bash
GO111MODULE=on go mod vendor
``` 

Then building an executable:
```bash
go build -o ./iban-validator
``` 

### Using makefile

Run a command
```bash
make build
```

### Using docker

Run a command
```bash
make image
```

## Testing

### Manual

A developer should load all dependencies library before testing, run command:
```bash
GO111MODULE=on go mod vendor
``` 

Then building an executable:
```bash
go test ./...
``` 

### Using makefile

Run a command
```bash
make testing
```

## Console command

A command supported two command:

- **run** - Running a service with a REST API
- **validate** - Console command to validate an IBAN

### Service run

#### Quick

Run a service using a makefile command
```bash
make run
```

Or using built docker image

```bash
make image-run
```

#### Manual

Environment variable:

- **ADDR** - address which a service listening. Default: ':8080'

- **STAGE** - a name of running staging. Default: 'dev'. Supported: 'prod', 'dev', 'test'

- **LOG_LEVEL** - a level of logging messages. Default: 'info'. Supported: 'debug', 'info', 'warning', 'error'

- **DB_PATH** - a path of a yaml file with countries format description. Default: './data/countries-iban.yaml'

- **LIMIT_TEXT_LENGTH** - a maximum length of a text parameters. Default: 1024

Command to run:
```bash
STAGE=prod ./iban-validator run 
```

### Console command validate

Using follow console command to validate several IBAN codes:
```bash
./iban-validator validate {IBAN} {IBAN2} {IBAN3} 
```

Example:
```bash
./iban-validator validate "GB82 WEST 1234 5698 7654 31" "BR97 00360305 00001 0009795493 P 1"
```
And getting an output:
```bash
GB82 WEST 1234 5698 7654 31: invalid
BR97 00360305 00001 0009795493 P 1: valid
```

## REST API

### GET /validate/{IBAN}

**Parameters**:

- **{IBAN}** - International Bank Account Number

**Return**:

- Status: **200 OK** - IBAN is valid 
- Status: **412 Precondition Failed** - IBAN is invalid 

**Description**:

A method are implementing general rule of IBAN validation.

**Example**:
```http request
GET /validate/GB82%20WEST%201234%205698%207654%2031

200 OK
{"iban":"valid"}
```
 
## Implementing

An implementation of the service base [International Bank Account Number](https://en.wikipedia.org/wiki/International_Bank_Account_Number#cite_note-IBANRegistry-1)
description and describing an IBAN validation algorithm. A yaml file 'data/countries-iban.yaml' was building also using Wikipedia description of IBAN.

Validation is supporting the following rules:

- A general checking of IBAN string symbols. Allowed symbols are from A to Z (in any case) and from 0 to 9;
- Country format checking: length of a string and a specific country IBAN template;
- Cheksum: calculating a numeric representation of IBAN string and calculating mod by 97.

Each rule is implementing in a separated validator. It helps to support the single responsibility principle.
