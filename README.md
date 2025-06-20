# Metric

A Go library for handling measurements, units, and quantities with type safety and precision.

## Overview

The `metric` library provides a comprehensive system for working with measurements and units in Go. It implements the International System of Units (SI) and extends to financial calculations with a specialized money package.

Key features:
- Type-safe operations on quantities with unit checking
- Support for all SI base and derived units
- Mathematical operations (add, subtract, multiply, divide) with proper unit handling
- Comparison operations (equals, greater than, less than)
- Financial calculations with precise decimal arithmetic
- Currency support with ISO 4217 standard
- Tax calculation functionality

## Installation

```bash
go get github.com/IAmRadek/metric
```

## Usage

### Working with SI Units and Quantities

```go
package main

import (
    "fmt"

    "github.com/IAmRadek/metric"
)

func main() {
    // Create quantities with SI units
    mass := metric.NewQuantity(75.0, metric.Kilogram)
    height := metric.NewQuantity(1.8, metric.Meter)

    // Calculate area (height squared)
    heightSquared, _ := height.MultiplyBy(height)

    // Calculate BMI (mass / height²)
    bmi, _ := mass.DivideBy(heightSquared)

    fmt.Printf("Mass: %v\n", mass)
    fmt.Printf("Height: %v\n", height)
    fmt.Printf("BMI: %v\n", bmi)
}
```

### Working with Money and Currencies

```go
package main

import (
    "fmt"

    "github.com/IAmRadek/metric/money"
)

func main() {
    // Create money values
    price := money.NewMoney(1999, money.USD)  // $19.99
    tax := money.NewTax(8.25, money.VAT)      // 8.25% VAT

    // Calculate price with tax
    priceWithTax := price.AfterTax(tax)

    fmt.Printf("Price: %v\n", price)
    fmt.Printf("Tax: %v\n", tax)
    fmt.Printf("Price with tax: %v\n", priceWithTax)

    // Perform arithmetic operations
    item1 := money.NewMoney(1099, money.USD)  // $10.99
    item2 := money.NewMoney(2499, money.USD)  // $24.99

    total, _ := item1.Add(item2)
    fmt.Printf("Total: %v\n", total)
}
```

### Creating Custom Units

```go
package main

import (
    "fmt"

    "github.com/IAmRadek/metric"
)

func main() {
    // Create a custom unit
    mySystem := metric.NewSystemOfUnits("MySystem", "MyOrganization")

    myUnit := metric.NewDerivedUnit(
        "myUnit",
        "My custom unit for measuring something",
        "mu",
        mySystem,
        metric.NewDerivedUnitTerm(metric.Meter, 1),
        metric.NewDerivedUnitTerm(metric.Second, -1),
    )

    // Create a quantity with the custom unit
    myQuantity := metric.NewQuantity(42.0, myUnit)

    fmt.Printf("My quantity: %v\n", myQuantity)
}
```

## API Documentation

### Core Interfaces

- `Metric`: Describes a standard of measurement with name, definition, and symbol
- `Unit`: Extends `Metric` and belongs to a system of units
- `SystemOfUnits`: Represents a standardized collection of units
- `Quantity`: Represents a value with an associated unit of measurement
- `DerivedUnit`: Represents a unit composed of other units with exponents

### Money Package

- `Currency`: Represents a monetary unit with code and decimal precision
- `Money`: Represents a monetary value with a specific currency
- `Tax`: Represents a tax rate with a specific type
- `TaxType`: Represents a type of tax (e.g., VAT)

## Dependencies

- Go 1.22.6 or higher
- github.com/govalues/decimal v0.1.33 (for precise decimal arithmetic)
- github.com/matryer/is v1.4.1 (for testing)

## License

MIT License
