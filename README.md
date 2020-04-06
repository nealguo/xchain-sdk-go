## 工程简介
本工程src/xchain为SDK底层运行机制的目录，src/test为SDK的接口测试目录

## 包结构简介
src/xchain用于存放SDK底层运行所需的所有文件

* src/xchain/sdk为接口和运行入口，提供SDK API和RPC请求
* src/xchain/crypto为SM2的支持，来自github项目，有微调
* src/xchain/config为配置项定义和配置加载提供支持
* src/xchain/proto为合约相关*.proto生成的*.pb.go
* src/xchain/model为所有的数据模型定义

src/test用于存放SDK测试所需的所有文件

* src/test/main.go为测试的运行入口，提供SDK加载和调用
* src/test/contract为样例合约，包括单合约和跨合约
* src/test/config为配置项定义和配置加载提供支持
* src/test/res为测试所需的资源，包括yaml和私钥证书等
* src/test/service为SDK接口的简单封装
* src/test/yaml为yaml的支持，来自github项目，有微调

## 使用步骤简介
1.注册SDK对应的私钥和证书，将私钥拷贝到src/test/res

2.修改配置文件src/test/res/application.yml
  * 配置项contract，需要修改
  * 配置项baas，需要修改

3.直接运行src/test/main.go
