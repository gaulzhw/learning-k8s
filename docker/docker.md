# docker



## image

<img src="img/docker-image.png" alt="docker-image" style="zoom: 67%;" />

- image index：json文件，可选，包含了对image中所有manifest的描述，相当于一个manifest列表，包含每个manifest的media type、文件大小、sha256码、支持的平台、平台特殊的配置
- image manifest：json文件，包含了对filesystem layers和image config的描述
- image config：json文件，包含了对image的描述
  - https://github.com/opencontainers/image-spec/blob/main/config.md
- filesystem layers：包含了文件系统的信息，即该image包含了哪些文件/目录，以及它们的属性和数据
  - https://github.com/opencontainers/image-spec/blob/main/layer.md



## 相关文件

### 镜像管理

- ubuntu：官方提供的最新ubuntu镜像，对应的完整名称为docker.io/library/ubuntu:latest
- ubuntu:16.04：官方提供的ubuntu 16.04镜像，对应的完整名称为docker.io/library/ubuntu:16.04
- ubuntu:@sha256:abcdef...：官方提供的digest码为sha256:abcdef...的ubuntu镜像，对应的完整名称为docker.io/library/ubuntu@sha256:abcdef...

docker pull的流程：

- docker发送image的名称+tag（或者digest）给registry服务器，服务器根据收到的image的名称+tag（或者digest），找到相应image的manifest，然后将manifest返回给docker
- docker得到manifest后，读取里面image配置文件的digest(sha256)，这个sha256码就是image的ID
- 根据ID在本地找有没有存在同样ID的image，有的话就不用继续下载了
- 如果没有，那么会给registry服务器发请求（里面包含配置文件的sha256和media type），拿到image的配置文件（Image Config）
- 根据配置文件中的diff_ids（每个diffid对应一个layer tar包的sha256，tar包相当于layer的原始格式），在本地找对应的layer是否存在
- 如果layer不存在，则根据manifest里面layer的sha256和media type去服务器拿相应的layer（相当去拿压缩格式的包）。
- 拿到后进行解压，并检查解压后tar包的sha256能否和配置文件（Image Config）中的diff_id对的上，对不上说明有问题，下载失败
- 根据docker所用的后台文件系统类型，解压tar包并放到指定的目录
- 等所有的layer都下载完成后，整个image下载完成，就可以使用了



- repositories.json

镜像存放路径：/var/lib/docker/image/[storage-driver]

repositories.json记录了和本地image相关的repository信息，主要是name、image id的对应关系

```shell
root@CHNDSI-VS-208:/var/lib/docker/image/overlay2# cat repositories.json | python3 -m json.tool
{
    "Repositories": {
        "k8s.gcr.io/etcd": {
            "k8s.gcr.io/etcd:3.5.0-0": "sha256:0048118155842e4c91f0498dd298b8e93dc3aecc7052d9882b76f48e311a76ba",
            "k8s.gcr.io/etcd@sha256:9ce33ba33d8e738a5b85ed50b5080ac746deceed4a7496c550927a7a19ca3b6d": "sha256:0048118155842e4c91f0498dd298b8e93dc3aecc7052d9882b76f48e311a76ba"
        }
    }
}
```

- 配置文件 image config

根据镜像查询出来的sha256，在imagedb里查询对应的配置

/var/lib/docker/image/[storage-driver]/imagedb/content/sha256/[image-sha256]

```shell
# 这里只关心rootfs
# 从diff_ids里可以看出etcd:3.5.0-0这个image包含了5个layer，
# 从上到下依次是从底层到顶层，417cb9...是最底层，ce8b3e...是最顶层
root@CHNDSI-VS-208:/var/lib/docker/image/overlay2/imagedb/content/sha256# cat 0048118155842e4c91f0498dd298b8e93dc3aecc7052d9882b76f48e311a76ba  | python3 -m json.tool
{
		...
    "rootfs": {
        "type": "layers",
        "diff_ids": [
            "sha256:417cb9b79adeec55f58b890dc9831e252e3523d8de5fd28b4ee2abb151b7dc8b",
            "sha256:33158bca9fb5a5ac2884b9f220006d7c000b5e7c5eac49890a651902a8d09574",
            "sha256:13de6ee856e951f4a03d2a3efd38aaf2d83c98d6b6ab117186e6384c9f074c5a",
            "sha256:eb364b1a02cae2de904b549073f5c3fcddd2c8949697f36b58f3bd5bb739fea1",
            "sha256:ce8b3ebd2ee7ca142b968754b3314a9d0c7e60dd97dbd8bde04481b2a9f40a6f"
        ]
    }
}
```

- layer的diff_id和digest的对应关系

layer的diff_id存在image的配置文件中，而layer的digest存在image的manifest中，他们的对应关系被存储在了image/[storage-driver]/distribution目录下

​	diffid-by-digest：存放digest到diffid的对应关系

​	v2metadata-by-diffid：存放diffid到digest的对应关系

```shell
root@CHNDSI-VS-208:/var/lib/docker/image/overlay2/distribution# tree .
.
├── diffid-by-digest
│   └── sha256
└── v2metadata-by-diffid
    └── sha256

# 根据diffid查询digest
root@CHNDSI-VS-208:/var/lib/docker/image/overlay2/distribution/v2metadata-by-diffid/sha256# cat 417cb9b79adeec55f58b890dc9831e252e3523d8de5fd28b4ee2abb151b7dc8b | python3 -m json.tool
[
    {
        "Digest": "sha256:5dea5ec2316d4a067b946b15c3c4f140b4f2ad607e73e9bc41b673ee5ebb99a3",
        "SourceRepository": "k8s.gcr.io/etcd",
        "HMAC": ""
    }
]

# 根据digest查询diffid
root@CHNDSI-VS-208:/var/lib/docker/image/overlay2/distribution/diffid-by-digest/sha256# cat 5dea5ec2316d4a067b946b15c3c4f140b4f2ad607e73e9bc41b673ee5ebb99a3
sha256:417cb9b79adeec55f58b890dc9831e252e3523d8de5fd28b4ee2abb151b7dc8b
```

- layer的元信息

layer的属性信息都放在了image/[store-driver]/layerdb目录下，目录名称是layer的chainid

chained，需要用到所有祖先layer的信息，从而保证根据chainid得到的rootfs唯一

底层的chainid，没有父层，chainid就是它自己

```shell
# tar-split.json.gz，layer压缩包的split文件，通过这个文件可以还原layer的tar包，
# 在docker save导出image的时候会用到
root@CHNDSI-VS-208:/var/lib/docker/image/overlay2/layerdb/sha256/417cb9b79adeec55f58b890dc9831e252e3523d8de5fd28b4ee2abb151b7dc8b# ls
cache-id  diff  size  tar-split.json.gz
```

- layer的数据

layer的文件放在/var/lib/docker/[stoge-driver]目录下，根据上面计算出来的chainid目录获取cache-id，在该路径下查询，在对应cache-id目录下的diff文件夹，即可看到每一层的文件信息

l文件夹包含了短签名的软链，链接到diff目录下

docker inspect [container]可以查看merge后的容器路径GraphDriver.Data

```json
{
  "GraphDriver": {
    "Data": {
      "LowerDir": "/var/lib/docker/overlay2/519e43614d77e173ebd53d4b5ca3d770cdf4cdaeebc06951b2bf1c8051bf9623-init/diff:/var/lib/docker/overlay2/90319a180aa85b21f68715d05bf146e2eb5773c62047f56bb28ec74479489a49/diff:/var/lib/docker/overlay2/74b09b2563b7bffad9eb130240ddcce0e3fc7f9ee617aa2246d3b54ef3769a27/diff:/var/lib/docker/overlay2/8506ee95a91c488225acff25d3a46b9b4e0bf408a6c2d55e50ef3820eba938ed/diff:/var/lib/docker/overlay2/185f0f1a65a3085d9b35a895f54495a40fb26492c383dabecba268a1fd4d0e31/diff:/var/lib/docker/overlay2/03be68990e5d5f43b6395720a93840428f56cdd5c2c3fb54f2243f86952cc6d5/diff:/var/lib/docker/overlay2/7561da8529d7d28d1a8aced13975a4ae9724f8d5e97cd5378b03e6c0a6407e35/diff:/var/lib/docker/overlay2/f112beeeea674e0c1bc3fbdba8e7404a21bb0eef3a755b66ca94d0cb52a1577d/diff:/var/lib/docker/overlay2/c82cddb579c703c9509f1978986f19c145f5bc6049175654c3d9faf1ce8352c5/diff:/var/lib/docker/overlay2/857a568744f619eb587245b1c5f0f8945746f6b18ac736dff962c455572cd4ff/diff:/var/lib/docker/overlay2/74a7d017a0556c8f1b4de3d5fb82fbaaf1e4fc3e6b18b201faa527915d55d507/diff:/var/lib/docker/overlay2/2e8d49c207cefcfc105bd43c7b50b2ce42cd6a9e19475bb7291ed69748a96f11/diff:/var/lib/docker/overlay2/8751467a51b13d676bd1f72d6b117310b0f7c93732d5033164d525ce23d9033a/diff:/var/lib/docker/overlay2/69374c04c2e8436c65fc27dcadf6e00ec486eec855edda75098e0b45101f8b61/diff:/var/lib/docker/overlay2/01574bd97d55397bee2bfd3cfc291df07edc64e11deb7078cdc32e50c1f84ca5/diff:/var/lib/docker/overlay2/4d340379817f2ed8d2a47b304dfa9a6698f6d337559186f62d3e949e37e7e889/diff:/var/lib/docker/overlay2/27cc7030ff8eb06dcf108d1a9ff412e3242b66ee8b311d83f8623af3b99561c8/diff",
      "MergedDir": "/var/lib/docker/overlay2/519e43614d77e173ebd53d4b5ca3d770cdf4cdaeebc06951b2bf1c8051bf9623/merged",
      "UpperDir": "/var/lib/docker/overlay2/519e43614d77e173ebd53d4b5ca3d770cdf4cdaeebc06951b2bf1c8051bf9623/diff",
      "WorkDir": "/var/lib/docker/overlay2/519e43614d77e173ebd53d4b5ca3d770cdf4cdaeebc06951b2bf1c8051bf9623/work"
    },
    "Name": "overlay2"
  }
}
```



### create管理的文件



### start管理的文件



## References

https://segmentfault.com/a/1190000009309276
