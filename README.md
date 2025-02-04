# Go SQL Parser

A production-grade SQL parser written in Go, with special support for Apache Iceberg DDL statements.

## Project Structure

```
go-sql-parser/
├── cmd/
│   └── parser/          # Command-line interface
├── pkg/
│   ├── ast/            # Abstract Syntax Tree definitions
│   ├── lexer/          # Lexical analyzer
│   ├── parser/         # SQL parser implementation
│   ├── token/          # Token definitions
│   └── iceberg/        # Iceberg-specific implementations
├── internal/
│   └── util/           # Internal utilities
└── README.md
```

## Features

- Production-grade SQL parsing
- Special support for Apache Iceberg DDL statements
- Efficient lexical analysis
- Comprehensive error reporting
- Support for:
  - CREATE TABLE statements
  - ALTER TABLE statements
  - DROP statements
  - Iceberg-specific table properties
  - Complex data types
  - Table partitioning
  - Sort orders
  - Custom properties

## Supported SQL Syntax

### Basic DDL
```sql
CREATE TABLE namespace.table_name (
    id bigint,
    data string,
    created_at timestamp
);
```

### Iceberg-Specific DDL
```sql
CREATE TABLE namespace.table_name (
    id bigint,
    data string,
    created_at timestamp
)
USING iceberg
PARTITIONED BY (date(created_at))
SORTED BY (id)
LOCATION 's3://bucket/path'
TBLPROPERTIES (
    'format-version' = '2',
    'write.format.default' = 'parquet'
);
```

## Usage

```go
import "github.com/sql/go-sql-parser/pkg/parser"

func main() {
    sql := `CREATE TABLE namespace.table_name (id bigint, data string);`
    p := parser.New(sql)
    stmt, err := p.Parse()
    if err != nil {
        log.Fatal(err)
    }
    // Use the parsed statement
}
```
