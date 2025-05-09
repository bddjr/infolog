# infolog

```
go get github.com/bddjr/infolog
```

```go
infolog.INFO("info")
infolog.INFOf("infof %v", 123)
infolog.INFOln("infoln")

infolog.ERROR("error")
infolog.ERRORf("errorf %v", errors.New("123"))
infolog.ERRORln("errorln")

infolog.WARN("warn")
infolog.WARNf("warnf %v", 123)
infolog.WARNln("warnln")
```

```
2025/05/09 10:34:58 [INFO] info
2025/05/09 10:34:58 [INFO] infof 123
2025/05/09 10:34:58 [INFO] infoln
2025/05/09 10:34:58 [ERROR] error
2025/05/09 10:34:58 [ERROR] errorf 123        
2025/05/09 10:34:58 [ERROR] errorln
2025/05/09 10:34:58 [WARN] warn
2025/05/09 10:34:58 [WARN] warnf 123
2025/05/09 10:34:58 [WARN] warnln
```
