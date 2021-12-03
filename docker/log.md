# log



docker启动容器的标准输出默认是在/var/lib/docker/containers中，这与容器的镜像日志会在同一个磁盘下。

可以配置docker的参数，将日志的路径挂载到一个独立的磁盘上("data-root": "/mnt/docker-data")。

同时在管理端产品化时，提供logs能力的时候必须带上--tail选项。