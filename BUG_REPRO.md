# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

同一巡检单带着相同幂等键重试，本该重放第一次结果，现在却返回 500；数据库日志同时记录了唯一约束冲突。本次只做原因定位，代码必须保持不变。请说明错误从数据库到接口的传递过程中哪里失去了可识别身份，并用日志或运行结果支撑结论。

## 含 Bug 版本

- 仓库：11DingKing/chargeguard-bug-11
- 仓库地址：https://github.com/11DingKing/chargeguard-bug-11.git
- parent SHA：a870af88f8607ec2fcbfd39d676dcb602e506882

## 复现步骤

```bash
git clone -- https://github.com/11DingKing/chargeguard-bug-11.git bug-repro
cd bug-repro
git checkout --detach a870af88f8607ec2fcbfd39d676dcb602e506882
go test ./internal/httpapi -run TestTaskBehavior -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/httpapi -run TestTaskBehavior -count=1
--- FAIL: TestTaskBehavior (0.00s)
    task_behavior_test.go:13: status=500 headers=map[Content-Type:[text/plain; charset=utf-8] X-Content-Type-Options:[nosniff]] body=storage error
FAIL
FAIL	chargeguard/internal/httpapi	0.065s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/httpapi -run TestTaskBehavior -count=1
--- FAIL: TestTaskBehavior (0.00s)
    task_behavior_test.go:13: status=500 headers=map[Content-Type:[text/plain; charset=utf-8] X-Content-Type-Options:[nosniff]] body=storage error
FAIL
FAIL	chargeguard/internal/httpapi	0.003s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

根因结论必须准确写明出问题的 Go 文件、具体符号和完整失效机制，并由实际复现、源码调查和验证证据支撑；调查结束时目标仓库代码、测试和配置零改动。
