# simple admin example rpc v1.1.5 gen by goctls v1.6.6

> 只需运行如下命令即可生成这个rpc项目模板

> You just need to run the command below to generate this project

```shell
goctls rpc new example --ent=true --module_name=github.com/suyuan32/simple-admin-example-rpc --port=8080 --desc=true --i18n=true

# or
# goctls rpc new example -e -m github.com/suyuan32/simple-admin-example-rpc -p 8080 -d -i

cd example

go mod tidy
```

> 修改 `ent/schema/examle.go` , 参考 example 中的 schema \
> Modify `ent/schema/examle.go`, you can find it in example's `ent/schema` directory


> 运行如下命令生成 CRUD 代码 \
> Run the command below to generate CRUD code

```shell
# 如果 make gen-ent 报错， 执行 go mod tidy | If there ig error when run make gen-ent command, run go mod tidy command.
make gen-ent

make gen-rpc-ent-logic model=Student group=student
make gen-rpc-ent-logic model=Teacher group=teacher

make gen-rpc
```

> 运行服务 | Run codes

```shell
go run example.go -f etc/example.yaml
```