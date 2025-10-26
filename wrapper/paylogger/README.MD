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
  -

Struct 如下：

```golang
Log paylogger.LoggerOption `yaml:"log"`
```
