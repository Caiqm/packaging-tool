# Go 执行文件打包工具

### 使用步骤

当然，该程序能够灵活地为多种环境配置和架构打包 Go 应用程序。为了方便快捷部署与执行，你不仅可利用此工具生成特定平台与架构的二进制文件，还可将其打包成独立的可执行程序，并将其路径添加到操作系统的**环境变量PATH**中。这样一来，用户只需通过命令行界面，在任何目录下都能直接调用此打包工具。

**执行命令示例：**

```shell
./main -path "/go/test" -alias "test1" -arch "arm64" -platform "linux"
```

### 参数说明

- -path go执行文件目录，即main函数文件目录，默认值是当前目录
- -alias 生成go二进制文件名称别名，默认值是 main
- -arch 相应系统信息，arm64、amd64等等，默认值是 amd64
- -platform 目标操作系统平台，默认值是 linux

### 参数详解

- -path Go源码目录： 指定包含主函数的 Go 源码文件所在的目录，默认情况下是当前运行 main 可执行文件的目录。

```shell
./main -path /go/test
```

- -alias 生成的二进制文件名称： 设置编译后生成的 Go 二进制文件的别名，默认为 main。这意味着如果你不指定这个参数，输出的二进制文件将被命名为 main。

```shell
./main -alias test1
```

- -arch 目标系统架构： 指定要为目标平台编译的系统架构，例如 arm64（适用于 ARM64 架构）或 amd64（适用于 x86_64 架构）。默认设置是 amd64。

```shell
./main -arch arm64
```

- -platform 目标操作系统平台： 指定要编译的目标操作系统平台，例如 linux、darwin（Mac OS）、windows 等。这将与 GOOS 环境变量对应，确保编译出适用于特定操作系统的可执行文件。请根据实际需求提供正确的平台名称。

```shell
./main -platform linux
```

### 示例说明

在实际应用中，上述 main 可执行文件可能是一个自定义打包脚本或工具，它会根据提供的 -arch 和 -platform 参数进行交叉编译和打包。务必确保正确设置了相应的 GOOS 和 GOARCH 环境变量，以便 Go 编译器识别目标平台。

```shell
./main -path /go/test -alias test1 -arch arm64 -platform linux
```

### 注意事项

在实际应用中，上述 main 可执行文件可能是一个自定义打包脚本或工具，它会根据提供的 -arch 和 -platform 参数进行交叉编译和打包。务必确保正确设置了相应的 GOOS 和 GOARCH 环境变量，以便 Go 编译器识别目标平台。

完成打包后，可以将生成的二进制文件添加至系统的环境变量 PATH 中，方便用户在任何位置直接调用。


# Go 进程管理与重启工具

## 功能概述

本Go程序提供了一种跨平台（Windows和Linux）的进程管理功能，允许用户通过指定服务名称查找并操作对应的进程。主要功能包括：

1. **获取进程ID**：通过进程名称获取对应的PID。
2. **终止进程**：如果找到相应的进程ID，则发送SIGKILL信号终止进程。
3. **重新运行主程序**：切换到指定的Golang主文件路径，并以`nohup`命令重新启动该程序。

**执行命令示例：**

```shell
./main -p "main" -m "./test1" -pid-only 1
```

### 参数说明

- -p：该参数用于指定服务器线程的名称。如果不提供 **-p** 参数，但指定了 **-m** 参数，则程序会直接执行 -m 参数所指定的可执行文件。如果同时提供了 **-p** 和 **-m** 参数，则程序首先会尝试查找并终止与指定线程名称匹配的进程，然后执行 -m 参数所指示的可执行文件。

- -m：这是一个路径参数，它指定了要运行或重新启动的Golang主程序的完整文件路径（包括文件名）。当需要替换或重启一个已存在的服务时，配合 **-p** 参数使用此选项，确保正确指向新的可执行文件。

- -pid-only：该标志不直接展示服务器线程名称对应的端口，而是用来表示是否仅需显示与 **-p** 参数关联的进程ID。若设置了 -pid-only 标志（值为**1**），在提供了 **-p** 参数的情况下，程序将只显示对应进程的**PID**，而不进行任何进一步的操作（如终止和重启进程）。