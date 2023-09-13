# change log

## v4.1.49

- add support for framerate & datatype

## v4.1.44

```
NodeFrom      string `json:"node_from,omitempty"`
Epoch         string `json:"epoch,omitempty"`
```

## v4.1.38

- 唤醒寻优逻辑，资源下载调整为：/awake-optimize-xxxx/threshold_file.txt

## v4.1.36: under development

## v4.1.35

- 退出码由 exitCode 调整为 ret_status
- export unloaddelay = "nio" for gts_perceive

## v4.1.30

- rm version:v1 for gts_awaken

## v4.1.29

- soft link to v4.1.28

## v4.1.28

- 异常 project 、 group 时响应信息不可读处理，eg：发现服务时如果服务不存在报的错误信息不可读
- 错误信息冗余修正，eg：有重复的错误信息
- 加载卸载日志信息不正确，eg：project和group重复了
- 服务提取逻辑，不存在 project、group 时保持原有的服务发现
- 多 patch_id 支持

## v4.1.24

- fix group bug

## v4.1.23

- gts_perceive use target's project and group

```go
p.tools.Logger().Infow("patchLoadTaskIat:HSet successfully",
	"sid", p.mess.Params.Sid, "HSetKey", fmt.Sprintf("%v_%v_%v", group, group, engService), "patchId", patchId, "patchIdPath", patchIdPath)
span.WithTag("patchLoadTaskIat:HSet successfully",
	fmt.Sprintf("HSetKey:%v,patchId:%v,patchIdPath:%v", fmt.Sprintf("%v_%v_%v", group, group, engService), patchId, patchIdPath))
```

## v4.1.17

- fix odeon down crash

## v4.1.16

- support ist/ost flag
- fix odeon down crash

## v4.1.14.dev

- add patch lic

## v4.1.13

- support ignore appid for ivw process

## v4.1.12

- update base img
- adjust awaken

## v4.1.11

- code optimization
- support awaken plugin

## v4.1.10

- support go 1.17

## v4.1.7

- update git.iflytek.com/AIaaS/gts_sdk v1.2.15

## v4.1.5

- update git.iflytek.com/AIaaS/gts_sdk v1.2.10

## v4.1.4

- update gts_sdk 1.2.4

## v4.1.3

- git.iflytek.com/AIaaS/gts_sdk v1.2.0

## v4.1.2

- update git.iflytek.com/AIaaS/gts_sdk/v2 v2.1.2

## v4.1.0

- support gts_perceive tts

## v4.0.4

- add namexpFlag for namexp

## v4.0.1

- rewrite log for iat.so

## v4.0.0

- migration vendor to module
- update gts_sdk v2.0.1(1、fix hbase driver bugs,2、load balance to nodes)
- storage mini support upload fallback 
- local generalmessage migrate to git.iflytek.com/HY_Castamere/generalmessage
- local cb_message migrate to git.iflytek.com/AIaaS/cb_message

## v3.2.1

- gts_perceive add trace between gts and engine

## v3.2.0

- aiot add version for 触发训练

## v3.1.0

- storageMini support namexp

## v2.22.11

- support specify the buckets
- 支持指定oss type，移除冗余的oss实例

## v2.22.9

- rewrite sc type from int to float64

## v3.0.0

- support refresh
- gts_perceive support check status code
- enable refresh for storage storageex storageMini

## v2.22.7

- set wedge dump to false

## v2.22.6

- fix storageMini data limit

## v2.22.5

- fix gts perceive bugs

## v2.22.4

- 增加perceive load/unload

## v2.22.3

- enable url2 to refresh queue
- fix mq exception

## v2.22.2

- delete rmq queue in storagemini

## v2.22.1

- Fix crashes in the detection GBK phase

```
panic: runtime error: index out of range [10] with length 10

goroutine 27471 [running]:
root/utils/encode.isGBK(0xc0004740f0, 0xa, 0x10, 0x21)
	/root/gts/src/root/utils/encode/encode.go:34 +0x13f
root/utils/encode.IsGBK(...)
	/root/gts/src/root/utils/encode/encode.go:14
plugin/unnamed-83621a50feac08d1f1f68130450ec9e24673ba34.(*Print).FilterData(0x1c9fe50, 0xc00051c1c0, 0x165a840, 0xc000346a50, 0x0, 0x0)
	/root/gts/src/root/impl/proxy/filter/print/print.go:25 +0x232
root/impl/proxy/process.(*Task).runDataFilter(0xc000346a50, 0xc00051c1c0, 0xc0002180a0, 0x14eae3e)
	/root/gts/src/root/impl/proxy/process/process.go:109 +0x6d9
root/impl/proxy/process.(*Task).Run(0xc000346a50, 0xc00051c1c0, 0x5)
...
```

## v2.22.0

- support perceive task

## v2.18.0

- add nunum support

## v2.17.0

- add hbase row key hash

## v2.16.5

- add patch.so

## v2.16.4

- add some logs about rmq

## v2.16.3

- adjust proc

## v2.16.2

- add some logs for a task_list

## v2.16.1

- vfe filter use tags[3] for ent instead of default ent

## v2.16.0

- iat filter add support proc option# gdk
