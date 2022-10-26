# Google Spreadsheets Golang API
![technology Go](https://img.shields.io/badge/technology-go-blue.svg)
[![License](https://img.shields.io/github/license/hprose/hprose-golang.svg)](http://opensource.org/licenses/MIT)

Google sheets API made with golang

## Features:

* Open a spreadsheet by account service.
* Read cell by ranges.

## EndPoints

### Read

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
To make requests to a Google spreadsheet, you need credential authorization, the simplest way to do this is through an 
account service

You must first create a project in the
[Google Cloud console](https://console.cloud.google.com/projectcreate?previousPage=%2Fwelcome%3Fproject%3Dgolang-sheets-365201&organizationId=0).

<img width="432" alt="2si" src="https://user-images.githubusercontent.com/52173621/197926639-25a1ba92-b3c1-4c48-927d-e17a4fd1ad3a.png">

Then go to the credentials/create credentials/account service section and follow the instructions for creating the 
account service.

<img width="953" alt="4si" src="https://user-images.githubusercontent.com/52173621/197927400-31b5a9e1-b96a-4314-ad89-db599be7710e.png">

You should see the following in your credentials panel.

<img width="742" alt="5" src="https://user-images.githubusercontent.com/52173621/197927737-0f38814f-845b-473e-8ad6-b6aa9851b340.png">

Click on that account service and create a new key of type json.

<img width="640" alt="6si" src="https://user-images.githubusercontent.com/52173621/197928323-b2208f5c-b506-4e98-8e52-10dfb89667d9.png">
<img width="404" alt="7" src="https://user-images.githubusercontent.com/52173621/197928578-08a8e157-6754-4613-8098-2885dfcef8b8.png">

Finally look for the Google Sheets API and enable it for your project.

<img width="432" alt="8si" src="https://user-images.githubusercontent.com/52173621/197928881-da667151-b541-4541-9d80-1eef09944e27.png">
<img width="462" alt="9" src="https://user-images.githubusercontent.com/52173621/197928846-e04beb5b-8a33-44d9-b546-c07ea066ebcf.png">

## Setup credentials

For security reasons this project ignores credential files. To set up credentials, you need to create the environments 
folder in ***internal/platform/sheets*** and add the json file you created in the previous section named ***credentials.json***.

<img width="228" alt="image" src="https://user-images.githubusercontent.com/52173621/197929396-c0b811fa-989e-4cfa-86e9-425bf2308942.png">

You should also create the spreadsheetsID.json file with the id of the sheet.

<img width="617" alt="10" src="https://user-images.githubusercontent.com/52173621/197930054-522c8411-3c63-438a-80a8-1fceaa277648.png">
<img width="453" alt="12" src="https://user-images.githubusercontent.com/52173621/197929999-8293c620-3255-4245-b256-dbdb4a32d97c.png">

## Pending features

* Write in cell.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
Please make sure to update tests as appropriate.

# License

This project is licensed under the [MIT](https://choosealicense.com/licenses/mit/) License - see the [LICENSE](LICENSE) file for details