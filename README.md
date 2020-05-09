# webhash

网页摘要计算

## 增加摘要网页

```shell
webhash add https://pinyin.sogou.com/linux/changelog.php
```

## 检查网页内容变化

```shell
webhash check
```

## 更新网页变化

```shell
webhash update
```

## linux 通知

使用 `notify-send` 命令

```shell
watch.sh
```

## 定时执行

1. 下载 webhash
1. 下载 watch.sh
1. 复制到 `~/bin` 目录
1. 设置定时任务 `crontab -e`

  ```shell
  # 每小时执行一次
  0 * * * * /home/$(whoami)/bin/watch.sh
  ```
