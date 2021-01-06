# fastcrud
Create you api fast and easily

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub license](https://img.shields.io/github/license/FerdinaKusumah/fastcrud.svg)](https://github.com/FerdinaKusumah/fastcrud/blob/master/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/FerdinaKusumah/fastcrud.svg)](https://GitHub.com/FerdinaKusumah/fastcrud/issues/)
[![GitHub issues-closed](https://img.shields.io/github/issues-closed/FerdinaKusumah/fastcrud.svg)](https://GitHub.com/FerdinaKusumah/fastcrud/issues?q=is%3Aissue+is%3Aclosed)
[![GitHub pull-requests](https://img.shields.io/github/issues-pr/Naereen/StrapDown.js.svg)](https://GitHub.com/Naereen/StrapDown.js/pull/)

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Resource</summary>
  <ol>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>

<!-- GETTING STARTED -->
## Getting Started
`fastcrud` fastcrud is an API with dynamic routes, perform CRUD fast and easy. 
You can enter any json data very quickly. 
save all data in memory very quickly, also available sample data that can be created quickly and easily.
The data will store in memory as default and will expired in `24` hours.

### Prerequisites
Need go `1.12+` or later.

## Installation

### from GitHub and make executable

```bash
▶ git clone https://github.com/FerdinaKusumah/fastcrud
▶ cd fastcrud
▶ go build .
▶ (sudo) mv fastcrud /usr/local/bin
```

<!-- USAGE EXAMPLES -->
## Usage

| **Flag**          	| **Description**                                                 	                                |
|-------------------	|-----------------------------------------------------------------------------------------------	|
| -h, --host         	| Running host server, default is `localhost`                 	                                    |
| -p  --port         	| Running port server, default is `8080`                 	                                        |
| -t  --ttl         	| How long data will keep in memory, unit in `minutes`, default is `24 hour` or `1440 minute`       |
| -h, --help        	| Display its helps                                               	                                |


### Examples

### Examples with command line
```shell script
  fastcrud --host "localhost" --port 8080 --ttl 60
``` 

### Examples with config file
example configuration with json file you can save with any name with `json` ext. 
```json
{
    "app": {
        "host": "localhost",
        "port": "9090",
        "ttl": 60, # unit in minute, change to -1 or 0 to disable expiration
    }
}
```  

## API pattern

#### Load example resource
[GET] `http://localhost:8080/api/seed-example` is to seed all example data.
Check this [link](https://jsonplaceholder.typicode.com/)

Availabel key name is:
* users -> `http://localhost:8080/api/users`
* todos -> `http://localhost:8080/api/todos`
* photos -> `http://localhost:8080/api/photos`
* albums -> `http://localhost:8080/api/albums`
* comments -> `http://localhost:8080/api/comments`
* posts -> `http://localhost:8080/api/posts`

#### Standart Crud
* [GET] `http://localhost:8080/api/{keyName}` is to retrieve all resource based on a key name.
* [POST] `http://localhost:8080/api/{keyName}` is to insert new single resource data json.
* [POST] `http://localhost:8080/api/{keyName}/bulk` is to insert new list of data json.
* [GET] `http://localhost:8080/api/{keyName}/{id}` is to retrieve specified resource based on a key name and id.
* [PUT] `http://localhost:8080/api/{keyName}/{id}` is to update specified resource based on a key name and id.
* [DELETE] `http://localhost:8080/api/{keyName}/{id}` is to delete specified resource based on a key name and id. 

## TODOs
- [ ] Supporting many key value database
- [ ] Dynamic file config and configuration

## Help & Bugs

[![contributions welcome](https://img.shields.io/badge/contributions-welcome-blue.svg)](https://github.com/FerdinaKusumah/fastcrud/issues)

If you are still confused or found a bug, please [open the issue](https://github.com/FerdinaKusumah/fastcrud/issues). All bug reports are appreciated, some features have not been tested yet due to lack of free time.

## License

[![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

**fastcrud** released under MIT. See `LICENSE` for more details.
