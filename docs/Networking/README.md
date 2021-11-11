<div align='center'><h1>Go Based Networking</h1></div>

<div align="center"><h4>前序</h4></div>

- 本目录下的内容以介绍Go语言网络通信技术为主，辅以介绍基本的计算机网络知识。同时，在重要章节将refer一些常见代码demo或进阶项目实例以加深对相关网络技术的阐释。
- 为了不引起争议，目录将以英文为主，翻译供辅助参考。
- 我们会持续改进，增加更多丰富的内容，对于已有的内容，如果您有疑问或者改进的建议，可以通过 issues 与我们联系，当然我们欢迎你通过 PR 参与到项目的共建中来。



## [一、网络基础知识](./Basic-Networking/content.md)

#### 1、基本术语

#### 2、Protocol Hierarchies 网络协议

- 协议的层次结构
- 协议层的设计问题
- 面向连接/无连接服务
- 服务和协议的关系



## 二、TCP/IP Reference Model

#### 1、Physical Layer 物理层

- 基本术语
- 信道公式
- 数组调制与多路复用
- 交换技术



#### 2、Data Link Layer 数据链路层

- 基本功能与实现
  - 帧的定界
  - 差错控制
  - 检纠错问题
- 基本协议
- 以太网
- 无线网络

#### 3、Network Layer 网络层

- 虚电路与数据报
- 路由算法
- Internet的网络层
  - IPV4
  - IPV6
  - 路由转发过程
  - 网络控制协议

#### 4、Transport Layer 传输层

- 传输层服务
  Socket 套接字

- 传层协议详解

  TCP/IP连接建立与断开过程（三次握手、四次挥手）

- TCP和UDP协议

- 拥塞控制

#### 5、Application Layer 应用层

- 网络应用与协议
- C/S、B/S模式
- HTTP协议



## 三、Go-HTTP

#### 1、HTTP服务

#### 2、Go标准库实现HTTP

#### 3、net/http库源码详解

#### 4、第三方库介绍（例：Req）



## 四、Go-RPC

#### 1、RPC简介及原理架构

#### 2、Go标准库实现RPC

#### 3、Go语言的RPC框架（例：GRPC）

