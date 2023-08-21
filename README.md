<div align="center">

# 🍥 Mahiro CD

小寻轻量级 CI/CD 服务器自动部署工具

Mahiro lightweight CI/CD server automation deployment tool

![stats](https://stats.deeptrain.net/repo/zmh-program/mahirocd)

</div>

### 介绍 Introduction
Mahiro CD 是一个轻量级子托管CI/CD工具，用于在多服务器中进行自动构建和部署。
类似于 Jenkins，但是更加轻量级，更加适合多服务器的自动部署。不暴露服务器的端口，更加安全。

Mahiro CD is a lightweight sub-hosted CI/CD tool for automatic building and deployment in multiple servers.
Similar to Jenkins, but more lightweight and more suitable for automatic deployment of multiple servers. Do not expose the port of the server, more secure.

### 架构 Architecture
![Architecture](/docs/struct.png)

### 特性 Features
- [x] ⚡ 轻量级 Lightweight
- [x] ✨ 分布式 Distributed
- [x] 🛠 安全 Secure
- [x] 🎨 易上手 Easy to use
- [x] 🎈 跨平台 Cross platform
- [x] 🔮 高性能 High performance
- [x] 🔧 高扩展性 High scalability
- [x] ⛏ 任务调度 Task scheduling
- [x] 📋 日志存储 Log storage

### 安装 Installation
> **Note**
> 请在环境中安装 `git` 和 `go`
> 
> Please install `git` and `go` in the environment

### 主节点 Master node
```shell
git clone https://github.com/zmh-program/mahirocd.git
./master.sh
```

### 从节点 Slave node
```shell
git clone https://github.com/zmh-program/mahirocd.git
./slave.sh
```

### 配置 Configuration
> 主节点 Master node
```yaml
# transport/config.yaml

port: 306  # server port
secret: 114514  # secret key
```

> 从节点 Slave node
```yaml
# config.yaml
endpoint: ws://localhost:306 # master node address
```

### 使用 Usage
配置类似于 **GitHub Actions**。在**.flow**文件夹下新建任意文件名，后缀为 *.yaml* 或 *.yml* 即新建一个任务，选择一个仓库，填写构建脚本，点击构建，即可完成自动构建和部署。

The configuration is similar to **GitHub Actions**. Create any file name in the **.flow** folder, with the suffix *.yaml* or *.yml* to create a task, select a repository, fill in the build script, click build, and you can complete the automatic build and deployment.

e.g.
```yaml
name: mahirocd  # task name
repo: "zmh-program/mahirocd"  # repository
path: "/www/wwwroot/mahirocd"  # working directory
steps:
  - name: "build frontend"  # step name
    run: | # script
      pnpm install
      pnpm build

  - name: "build backend"
    run: go build .

```

### 日志 Log
日志文件存储在 `logs` 文件夹下，以任务名的hash命名。

The log file is stored in the `logs` folder and named after the hash of the task name.

### 开源协议 License
MIT License

### 贡献 Contributing
![contributions](https://stats.deeptrain.net/contributor/zmh-program/mahirocd)