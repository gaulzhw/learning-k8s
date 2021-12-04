# docker



## 配置路径

docker的默认使用路径

- default pid file: /var/run/docker.pid
- default data root: /var/lib/docker
- default exec root: /var/run/docker

https://github.com/moby/moby/blob/master/cmd/dockerd/config_common_unix.go

可以通过命令`docker info | grep "Docker Root Dir"`查看运行时的根路径

问题：通常情况下与宿主机根目录在同一磁盘上，而容器的镜像、日志会持续增长，为了不影响系统服务使用的必要磁盘，有必要单独磁盘挂载

几种方式可以分离挂载：

- 软链接

  ```shell
  mv /var/lib/docker /data/docker
  ln -s /data/docker /var/lib/docker
  ```

- root-data，在/etc/docker/daemon.json中添加`"data-root": "/data/docker"`配置



## 日志

https://sematext.com/guides/docker-logs/

docker默认采用的是json-file的日志引擎，标准输出在/var/lib/docker/containers/[container-id]/[container-id]-json.log

日志可以采取docker的日志引擎做rotate

在/etc/docker/daemon.json中

```json
{
  "log-driver": "json-file",
  "log-opts": {
    "mode": "non-blocking",
    "max-size": "100m",
    "max-file": "3"
  }
}
```



## 管理系统

管理端查询日志，可以加上--tail，指定查询日志的条数
