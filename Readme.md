# Go Unique code generator library

## Install 

```
go get github.com/sifer169966/unique-code-generator

## Usage

## Init generator
```go

quantityConfig := generator.Generator{
		Quantity: config.Quantity,
}
rand.Seed(time.Now().UnixNano())
generator := generator.NewGenerator()
code := generator.GenerateCodes(quantityConfig.Quantity)

 ## Ex.resulte code
CRDTGQOAKB
```
