# simple admin example rpc

> 只需运行如下命令即可生成这个rpc项目模板

> You just need to run the command below to generate this project

```shell
goctls rpc new example --ent=true --moduleName=github.com/suyuan32/simple-admin-example-rpc --go_zero_version=v1.4.3 --tool_version=v0.1.7 --port=8080
```

> 修改 ent/schema/examle.go \
> Modify ent/schema/examle.go


> 运行如下命令生成 CRUD 代码 \
> Run the command below to generate CRUD code

```shell
make gen-ent

make gen-rpc-ent-logic model=Student group=student
make gen-rpc-ent-logic model=Teacher group=teacher
```