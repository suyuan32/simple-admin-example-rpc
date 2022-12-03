# simple admin example rpc

> 只需运行如下命令即可生成这个rpc项目模板

> You just need to run the command below to generate this project

```shell
goctls rpc new example --ent=true --moduleName=github.com/suyuan32/simple-admin-example-rpc --goZeroVersion=v1.4.2 --toolVersion=v0.0.9 --port=8080
```

> 修改 ent/schema/examle.go \
> Modify ent/schema/examle.go

```go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-core/pkg/ent/schema/mixins"
)

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("age"),
	}
}

func (Student) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return nil
}

```

> 运行如下命令生成 CRUD 代码 \
> Run the command below to generate CRUD code

```shell
make gen-ent

make gen-rpc-ent-logic model=Student
```