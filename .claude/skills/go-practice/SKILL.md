---
name: go-practice
description: 本仓库（Go 算法每日一题）的工作规范——新增题解、格式化验证、本地测试、Git 推送注意事项。当用户要求"做今天的题"、"新增题解"、"整理提交"、"推送到 GitHub"时使用。
---

# Go 算法练习仓库工作规范

本仓库是 LeetCode 每日一题的 Go 题解仓库。遵循以下约定工作。

## 目录约定

- 题解文件以当天日期命名，如 `20261020.go`（YYYYMMDD），包名一律为 `main`
- `main.go` 是测试入口，调用当前题目的函数并用题目示例做验证
- 所有题解文件平铺在仓库根目录，不分子目录

## 新增题解流程

1. 新建以当天日期命名的文件，如 `20261020.go`
2. 按题目要求的函数签名实现（函数名与 LeetCode 一致，便于对照）
3. 在 `main.go` 中调用新函数，传入题目示例用例
4. 验证（见下节）通过后再算完成

## 格式与验证（每次改完代码都要做）

```powershell
gofmt -w .        # 格式化（注意：PowerShell 下需用 Get-ChildItem *.go | ForEach-Object { gofmt -w $_.Name }）
go vet ./...      # 静态检查
go build ./...    # 编译检查
go run .          # 运行 main.go 的示例用例，人工核对输出
```

格式规范：

- 缩进用 Tab（gofmt 默认），换行符用 LF
- 项目已配 `.editorconfig`，新文件会自动遵循
- 提交前确保 `gofmt -l .` 输出为空

代码风格：与仓库现有题解保持一致——朴素直白的解法优先，变量名简短（`res`、`nums`、`tmp` 风格可接受），不过度抽象。

## Git 推送注意事项

- 提交信息用中文，简明描述本次改动（如 `新增 XXX 题解`、`整理 README`）
- 远端：`https://github.com/6xiancan/go-practice.git`，分支 `main`
- 本机已配置全局代理（因为国内直连 GitHub 会被重置）：
  ```powershell
  git config --global http.proxy http://127.0.0.1:15380
  git config --global https.proxy http://127.0.0.1:15380
  ```
- 如果 push 报 `Failed to connect to 127.0.0.1`：说明代理软件没开，先开代理或用
  `git config --global --unset http.proxy` / `--unset https.proxy` 临时取消
- 如果 push 报 `Connection was reset`：说明没走代理，检查上面的 proxy 配置是否存在（`git config --global -l`）
- Claude Code 的 shell 里无法弹出 GitHub 登录窗口；若遇到凭证问题，让用户在输入框用 `! git push` 自行执行

## 更新 README

每做一道新题，同步更新 `README.md` 的"题目列表"表格：加一行日期、函数名、中文题名、一句话思路。
