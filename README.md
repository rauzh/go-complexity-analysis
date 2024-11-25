# go-complexity-analysis
go-complexity-analysis calculates the Cyclomatic complexities, the Halstead complexities and the Maintainability indices of golang functions.

!! This is a refactored and brought to perfection fork of [shoooooman/go-complexity-analysis-action](https://github.com/shoooooman/go-complexity-analysis). Put star on his repo too!

# Install
```sh
$ go install github.com/rauzh/go-complexity-analysis/cmd/complexity@latest
```

# Usage
```sh
$ go vet -vettool=$(which complexity) [flags] [directory/file]
```

## Flags
`--cycloover`: show functions with the Cyclomatic complexity > N (default: 10)
`--halsteadover`: show functions with the Cyclomatic complexity > N (default: 50)
`--maintunder`: show functions with the Maintainability index < N (default: 20)
`--verbose`: show all functions complexity info

## Output
```
<filename>:<line>:<column>: func <funcname> seems to be complex (cyclomatic complexity=<cyclomatic complexity>)
<filename>:<line>:<column>: func <funcname> seems to have low maintainability (maintainability index=<maintainability index>)
```

## Examples
```go
$ go vet -vettool=$(which complexity) --cycloover 10 ./...
$ go vet -vettool=$(which complexity) --maintunder 20 main.go
$ go vet -vettool=$(which complexity) --cycloover 5 --maintunder 30 ./src
```

# Metrics
## Cyclomatic Complexity
The Cyclomatic complexity indicates the complexity of a program.

This program calculates the complexities of each function by counting idependent paths with the following rules.
```
Initial value: 1
+1: if, for, case, ||, &&
```

## Halstead Metrics

Calculation of each Halstead metrics can be found [here](https://www.verifysoft.com/en_halstead_metrics.html).

### Rules
1. Comments are not considered in Halstead Metrics
2. Operands and Operators are divided as follows:

#### Operands
- [Identifiers](!https://golang.org/ref/spec#Identifiers)
- [Constants](!https://golang.org/ref/spec#Constants)
- [Variables](!https://golang.org/ref/spec#Variables)

#### Operators
- [Operators](!https://golang.org/ref/spec#Operators_and_punctuation)
    - Parenthesis, such as "()", is counted as one operator
- [Keywords](!https://golang.org/ref/spec#Keywords)

## Maintainability Index
The Maintainability index represents maintainability of a program.

The value is calculated with the Cyclomatic complexity and the Halstead volume by using the following formula.
```
Maintainability Index = 171 - 5.2 * ln(Halstead Volume) - 0.23 * (Cyclomatic Complexity) - 16.2 * ln(Lines of Code)
```

This program shows normalized values instead of the original ones [introduced by Microsoft](https://docs.microsoft.com/en-us/archive/blogs/codeanalysis/maintainability-index-range-and-meaning).
```
Normalized Maintainability Index = MAX(0,(171 - 5.2 * ln(Halstead Volume) - 0.23 * (Cyclomatic Complexity) - 16.2 * ln(Lines of Code))*100 / 171)
```

The thresholds are as follows:
```
0-9 = Red
10-19 = Yellow
20-100 = Green
```


