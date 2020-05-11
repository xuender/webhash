# Webhash

网页摘要计算, 使用 [HighwayHash](https://github.com/google/highwayhash) 算法

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

## 监视网页修改

```shell
webhash watch
```

## 定时执行

设置定时任务 `crontab -e`

  ```shell
  # 每小时执行一次
  0 * * * * /home/$(whoami)/bin/webhash watch
  ```
