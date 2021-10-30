# 操作系统

## 前序
- 结合安全性、性能和复杂性三个因素的考虑，用 Go 语言编写 Kernel 并不占有优势。
- 此目录中的内容将以介绍操作系统基本概念为主，针对特定章节的内容会有相关 Go 语言实践
- 为了不引起争议，目录将以英文为主，翻译供辅助参考
- 我们会持续改进，增加更多丰富的内容，对于已有的内容，如果您有疑问或者改进的建议，可以通过 issues 与我们联系，当然我们欢迎你通过 PR 参与到项目的共建中来。

## Virtualization [虚拟化]

1. [Process and process API [进程与进程相关的API]](./Process-and-process-Api/Content.md)


2. Direct Execution [CPU 直接执行]
   - 直接执行
   - 受限直接执行

3. Scheduling [调度]
   - 单 CPU 调度
   - 多 CPU 调度
   - 彩票调度

4. Address-related [地址相关]
   - 地址空间
   - 地址转换

5. Space-related [空间相关]
   - 内存及内存管理分段技术
   - 自由空间管理

6. Complete VM System [虚拟系统]

## Concurrency 并发性

1.  Concurrency and Threads [并发和线程]

2.  Locks [锁]

3. Condition Variables and Semaphores [条件变量和信号量]

## Persistence 持久性

1. File System and devices [文件系统和硬件设备]
   - 输入/输出设备
   - 硬盘驱动
   - 普通文件系统及实现，FFS，LFS，NFS，AFS

2. Flash-based SSDs [闪存 SSD]

3. Distributed Systems [分布式系统]

## 引用
- [The benefits and costs of writing a POSIX kernel in a high-level language](https://www.usenix.org/system/files/osdi18-cutler.pdf)