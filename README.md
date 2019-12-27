# Getting Started

## Install Go Dependencies

```
go get github.com/stretchr/testify/assert github.com/renstrom/fuzzysearch/fuzzy
```

## Install React Dependencies

Install webpack:

```
npm i react react-dom webpack babel-loader babel-preset-es2015 babel-preset-react babel-core whatwg-fetch -S
```

## Run a React build

```
./node_modules/.bin/webpack -d
```

## Run the app
```
go run main.go
```