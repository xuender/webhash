# webhash

网页摘要计算

## 安装

curl

```shell
sh -c "$(curl -fsSL https://github.com/xuender/webhash/raw/master/install.sh)"
```

wget

```shell
sh -c "$(wget -O- https://github.com/xuender/webhash/raw/master/install.sh)"
```

## 增加摘要网页

```shell
webhash add https://api.github.com/repos/golang/go/milestones/72
```

## 检查网页内容修改

```shell
webhash check
# 或者
webhash
```

## 更新网页摘要

```shell
webhash update
```

## linux 通知

使用 `notify-send` 命令

```shell
watch.sh
```

## 定时执行

1. [安装 webhash](#安装)
1. 下载 [watch.sh](https://raw.githubusercontent.com/xuender/webhash/master/watch.sh)
1. 复制到 `~/bin` 目录
1. 设置定时任务 `crontab -e`

  ```shell
  # 每小时执行一次
  0 * * * * /home/$(whoami)/bin/watch.sh
  ```
