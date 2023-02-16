protoc-gen-go-gorm
===

Generates GORM model from protobuf message.

## Installation

```
go install github.com/zgs225/protoc-gen-go-gorm@latest
```

## Features

* Support all field tags for model declaration
* Support all types of association
* Support enumeration

## Example

``` protobuf
message User {
    option (yuez.gorm.model) = {
        table_name: "users";
    };
    uint64 id = 1;
    string name = 2;
}
```

will generate model like below:

``` go
type UserModel struct {
    Id   uint64
    Name string
}

func (m UserModel) TableName() string {
    return "users"
}
```
