# json-config
Super simple json config reader, with post processing callback for env variable expansion

## Go get
go get -u github.com/paul-at-nangalan/json-config

## Import
import "github.com/paul-at-nangalan/json-config"

## Usage

### Setup
Call the reader.Setup() function with the config directory path. The path can be relative.

### Read
Call the reader.Read() function to read config files. Files must be under the config directory.

