# Google Spreadsheets Golang API
![technology Go](https://img.shields.io/badge/technology-go-blue.svg)
[![License](https://img.shields.io/github/license/hprose/hprose-golang.svg)](http://opensource.org/licenses/MIT)

Google sheets API made with golang

## Features:

* Open a spreadsheet by **title**, **key** or **url**.
* Read, write, and format cell ranges.
* Sharing and access control.
* Batching updates.


## EndPoints

### Opportunities

- **GET** [/:sheet_name/:start_column/:start_row/:end_column/:end_row"]()

#### request

```json
{
  "sheet_name": "Sheet 1",
  "start_column": "A",
  "start_row": "1",
  "end_column": "B",
  "end_row": "2"
}
```

#### response

```json
{
  "cells": [
    {
      "cell_position": "A1",
      "information": "Cell A1 Information"
    },
    {
      "cell_position": "A2",
      "information": "Cell A2 Information"
    },
    {
      "cell_position": "B1",
      "information": "Cell B1 Information"
    },
    {
      "cell_position": "B2",
      "information": "Cell B2 Information"
    }
  ]
}
```


## Account Service
To make requests to a Google spreadsheet, you need credential authorization, the simplest way to do this is through an account service



## Pending features

## Contributing
Read our Contributing Guide to learn about reporting issues, contributing code, and more ways to contribute.


## Security
If you happen to find a security vulnerability, please let us know at 
XXX

and allow us to respond before disclosing the issue publicly.


# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details