github 地址： https://github.com/donnie4w/go-logger

配置如下：

```yaml
log:
  level: info
  output: file
  path: ./log/server.log
  time-mode: "day"
  max-size-mb: 100
  max-backup: 20
  compress: true
```

- level 等级，有 trace、debug、info、warn、error、painc
- output 输出格式，file 是文件；不配置即控制台；
- path 输出路径
- time-mode:
  - day: 按天分隔
  - hour: 按小时
  - month: 按月（默认）
- max-size-mb 文件最大大小，超过该大小时进行切割
- max-backup 保留的最大备份文件数量
- compress 是否启用备份文件的压缩

配置时间和文件大小后，都会按时间或文件大小进行分割

Struct 如下：

```golang
Log paylogger.LoggerOption `yaml:"log"`
```
