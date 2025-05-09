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
2025/05/09 10:49:52 [INFO] info1
2025/05/09 10:49:52 [INFO] infof2
2025/05/09 10:49:52 [INFO] infoln 3
2025/05/09 10:49:52 [ERROR] error1
2025/05/09 10:49:52 [ERROR] errorf2
2025/05/09 10:49:52 [ERROR] errorln 3
2025/05/09 10:49:52 [WARN] warn1
2025/05/09 10:49:52 [WARN] warnf2
2025/05/09 10:49:52 [WARN] warnln 3
```
