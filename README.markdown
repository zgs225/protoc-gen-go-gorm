protoc-gen-go-gorm
===

Generates GORM model from protobuf message.

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
