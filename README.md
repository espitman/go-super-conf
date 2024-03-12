# superConf

golang log library

## Installation

```bash
go get github.com/espitman/go-super-conf
```

## Contributing

This package use gookit/config/v2 lib

## Usage

This lib first read environment variables that passed to application<br/>
So, read json files in config folder in root of projects

```
1. config/default.json
2. config/{environment}.json [APP_ENV=...]
3. config/local.json
```

‍
**_Get config_**<br/>

```
PORT := superConf.Get("app.port")
```

**_Note:_** environment variables must be uppercase and user \_ only for nested levels

```
APP_PORT => "app": {"port": xxx}
```
