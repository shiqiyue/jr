# json reader

read json file implemented by golang program.  


## usage
1、 install
```shell
go get -u github.com/shiqiyue/jr
```
2、 run command in shell
```shell
example1: get version attribute from a.json 
jr -f=a.json -p=version
example2: get lodash version attribute from package.json
jr -p=dependencies.lodash
```
