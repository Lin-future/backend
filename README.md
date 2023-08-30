# GO-SVC-TPL

go-svc-tpl 是为了快速构建 golang web 项目开发的项目框架，以及相关的一系列工具链。



## 框架使用

1. 未完成，请 clone 下来，在其基础上现改。
2. run `go generate ./api` in to update docs

## 设计思路

见文档 [design](./docs/design.md) (未更新)

Design-Drive 的后端开发模式，先进行 API 和 model 设计，后进行后端开发。

## 开发计划

`*` 为优先级较低的项目





- 使用 apifox 作为 api 文档编辑，导出 OpenAPI 文档并生成代码
- 寻找 go 中使用 `@` 做代码注解的实现





- [x] 完成一个最简示例版本
- [x] 解析 ts 生成代码的工具 (应该是使用node)
- [ ] 分析文档生成 api 的工具，或许还有模板代码的生成工具 (go开发)
- [ ] 支持 mongo





### ~~API-swag~~

- [ ] 支持 3.0
- [ ] 支持通过 model 中定义的 form/head tag 生成相应接口的参数
- [ ] model 解析格式向 json 规则对齐
- [ ] 支持更多校验规则
- [ ] 通过分析代码生成接口文档，而不是注释



### utils

封装一些常用功能，方便开发

或者更新一些作者已经明确标识不会维护的repo

- [ ] **fork stacktrace 并改进 current/code 等功能**
- [ ] *cache (generic interface)

```
backend
├─ .git
│  ├─ COMMIT_EDITMSG
│  ├─ config
│  ├─ description
│  ├─ FETCH_HEAD
│  ├─ HEAD
│  ├─ hooks
│  │  ├─ applypatch-msg.sample
│  │  ├─ commit-msg.sample
│  │  ├─ fsmonitor-watchman.sample
│  │  ├─ post-update.sample
│  │  ├─ pre-applypatch.sample
│  │  ├─ pre-commit.sample
│  │  ├─ pre-merge-commit.sample
│  │  ├─ pre-push.sample
│  │  ├─ pre-rebase.sample
│  │  ├─ pre-receive.sample
│  │  ├─ prepare-commit-msg.sample
│  │  ├─ push-to-checkout.sample
│  │  └─ update.sample
│  ├─ index
│  ├─ info
│  │  └─ exclude
│  ├─ logs
│  │  ├─ HEAD
│  │  └─ refs
│  │     ├─ heads
│  │     │  └─ main
│  │     └─ remotes
│  │        └─ origin
│  │           ├─ HEAD
│  │           └─ main
│  ├─ objects
│  │  ├─ 00
│  │  │  └─ 6319fab8cf9a282595606c542da250b0b20ee1
│  │  ├─ 01
│  │  │  ├─ 2a4dcda64e527a647b6d1541ece0c62e726ea2
│  │  │  └─ 75063ca25d9ad59dc8113ce6727db71d554d34
│  │  ├─ 07
│  │  │  ├─ 10d5ff187051d3c94f561e7d9938c4aa21b53e
│  │  │  └─ a0cc0fc69fdc13eda603a897b9116884ad9a61
│  │  ├─ 0a
│  │  │  └─ 0a0733abeede97558c13a5956964cdbf6e1e87
│  │  ├─ 0b
│  │  │  └─ a1b2ead2d3de3e2878caef6c83c558cf76e346
│  │  ├─ 0e
│  │  │  └─ d3d1798380e14a7bce3d737884f0eb33816afe
│  │  ├─ 10
│  │  │  └─ 3cc4dbc26e25fd27408617826584b79f9e278e
│  │  ├─ 14
│  │  │  ├─ 2f75e1c199927a90382b1f0d30634c04813c4a
│  │  │  └─ ba0cc6346ecd18a573993d08c04c891bd549a8
│  │  ├─ 15
│  │  │  └─ b8adf0a98eec267dbf3d827d7b198163c029b9
│  │  ├─ 18
│  │  │  └─ 9d75b1f9681c6db88ee5005b40c3b54fea21ca
│  │  ├─ 1b
│  │  │  └─ b1b7ca1a56c621a3b861a439ed94dd31a1e960
│  │  ├─ 1f
│  │  │  └─ 469bfbc611ca3182b79af1947656c6377058f2
│  │  ├─ 20
│  │  │  └─ defcf4d81434e76d046b58a45d4881185059de
│  │  ├─ 28
│  │  │  └─ 6d25917c9076bafa3bfc76815a12d6547ed135
│  │  ├─ 29
│  │  │  ├─ 420745339ade6960c24ea70f60c66b49c59620
│  │  │  └─ aa6594115743b934d9c257b715b8dee6fff3b7
│  │  ├─ 33
│  │  │  └─ 86b8a0d3bfc927794374c02378b7bd0f3cd0c5
│  │  ├─ 38
│  │  │  └─ f7cee1b6c59d41f84f9c6a2fb269b6c766d3ad
│  │  ├─ 39
│  │  │  ├─ b8bdf26334ff10ffdcd1d7334486705f93438e
│  │  │  └─ d632fbda60046b8233c0b2618dcb9db8cf0468
│  │  ├─ 3b
│  │  │  └─ 79ce5f78fd4f0b06ea06468394c4911802d48b
│  │  ├─ 3d
│  │  │  └─ faf1f15ac6f7ec2a029b356dedbad241e2bf1f
│  │  ├─ 3e
│  │  │  ├─ 74d6e97b9648d80369ce678a61bce8b316e5f4
│  │  │  └─ a64c89ca829a7867c70fe9baac059c5016ff86
│  │  ├─ 3f
│  │  │  └─ 396bcfd2b60f4590cb66f7c37bea2a3e2b2363
│  │  ├─ 41
│  │  │  └─ d3572e976d5463e6e1106de82553f2339abf02
│  │  ├─ 42
│  │  │  └─ acbe7c99e2fda7fee07be20ddb940a5ec795dc
│  │  ├─ 50
│  │  │  ├─ 323cc70c7f21584a24c36bb7318c1c550ed9bc
│  │  │  └─ ff8f9a44c43a352f6cbc26edbccff0b2006051
│  │  ├─ 53
│  │  │  └─ bd23cb8bcfe1f3687895d690a9858f825181a8
│  │  ├─ 55
│  │  │  └─ e9cffd6b08e1f2e22c87898c3a396654c7e195
│  │  ├─ 56
│  │  │  └─ e64cc19bf223b88933354ce793169e184f81ed
│  │  ├─ 58
│  │  │  └─ 2a77e41225a5b3043fba98d3cce0f9c01d2ad4
│  │  ├─ 5c
│  │  │  └─ 79327493adb151dada54670d92b95a2135a8b4
│  │  ├─ 61
│  │  │  └─ 513bd456fa6672511313628c75b1e39ac4c1a4
│  │  ├─ 66
│  │  │  └─ d513e8095630be2d8e0601d7eef1436a9ae99a
│  │  ├─ 6b
│  │  │  └─ e73fe9b46c8cdc62ad3e727be7dfdd0e14fc5b
│  │  ├─ 6c
│  │  │  └─ 6199e4827e8aa4984c0b626f9ea95bc6c64a0d
│  │  ├─ 70
│  │  │  └─ bfb48d0647ae659321092c694dc781c2eb2e52
│  │  ├─ 75
│  │  │  └─ 56e4c1221880b4c5966c12c76fde8613c4691d
│  │  ├─ 76
│  │  │  └─ 81ebd55ccc1f466c72d770e9ad276697d110bb
│  │  ├─ 79
│  │  │  └─ 41782b870066d871585040ab9a6843ac42ebdb
│  │  ├─ 7d
│  │  │  └─ c331ea568c9a25826aef5f1416fc8ac5f90bf5
│  │  ├─ 7e
│  │  │  └─ 1b53f3fb618991787e03407d9f47767095ad2a
│  │  ├─ 83
│  │  │  └─ f1cad9da365846ed3bb5b16cc8009c587027ce
│  │  ├─ 84
│  │  │  └─ f2f627e474893f4fb982e4688b802a4bc8324d
│  │  ├─ 86
│  │  │  └─ 067c38dd4c35053de1470a36b436f2117defa8
│  │  ├─ 8d
│  │  │  └─ ada3edaf50dbc082c9a125058f25def75e625a
│  │  ├─ 91
│  │  │  └─ 5d0bdcc9912c9503457486ecf5e076a4c3c99b
│  │  ├─ 93
│  │  │  └─ 2dd11c3b7730a20834c0cfb13c90a4fe252543
│  │  ├─ 95
│  │  │  └─ 5ed4ddad6979761061e03d54cfe6cce08ab7ff
│  │  ├─ 99
│  │  │  └─ 8c5eccd2fa0f2e24b8dc173e2c5c9b54b533ab
│  │  ├─ 9a
│  │  │  └─ 4c660072cb0d6cbbf1096f476b835640145ed0
│  │  ├─ 9c
│  │  │  └─ 01e6dc0ce14511234df8a9081af49013fc56cc
│  │  ├─ 9d
│  │  │  └─ b71614c4f79da3a86e4cb3ac0457c18f08b816
│  │  ├─ a4
│  │  │  └─ 2a85f804d81a3df20c9ab92fae09f4e09195f4
│  │  ├─ a6
│  │  │  ├─ 0f41078d17e307f34539d7b20a9b9cea051e47
│  │  │  └─ fe8ebcec220bc8efd300444d6a06d8594bb7cf
│  │  ├─ aa
│  │  │  └─ 59f9433f510b008d1f4cdd589cdecbc11ad10d
│  │  ├─ ab
│  │  │  └─ 15d552d629cc84918c516565ca93cdbf90facb
│  │  ├─ ad
│  │  │  └─ b78ea3dcbeb363a21bdf574b7861d3a2730fb1
│  │  ├─ ae
│  │  │  └─ ee0d910638e28b56062186f793d1bc71caad49
│  │  ├─ b1
│  │  │  └─ bd0de1b016312da0499130b3f69650c3ea27b7
│  │  ├─ b3
│  │  │  └─ 5cf6071d6f992c72c5e2681cea4baeef3682d9
│  │  ├─ b5
│  │  │  ├─ 792c006f4f3bea83c7fa5bf089e3176ac00b78
│  │  │  └─ f49b390d508201f0a55e6c2fdba174886e1f29
│  │  ├─ bd
│  │  │  └─ 0037542f9ec241f9aad55e67bc34b5b4979373
│  │  ├─ c3
│  │  │  ├─ 5e25661b8ce9f03b0eb81012c3b170dbd98481
│  │  │  └─ d76e5c374dc4cb1dc20c39a8143e0f0a6138d4
│  │  ├─ c4
│  │  │  └─ 054ded6ae13e566f42f2aea26c145ed3abc3e0
│  │  ├─ cb
│  │  │  ├─ 59bdc1dbe20143807bca1a9a065662e3f6243c
│  │  │  └─ ca8bb2da087fff7c945da8d5afaee17da50546
│  │  ├─ cc
│  │  │  ├─ 06c1f722bc8e02bc079c6c030beea20e183918
│  │  │  ├─ a4c4ea4be5a19dbcf19121fbabf7589d34d82b
│  │  │  └─ a99ee25f97fd9079b8b51bda6440c8be69ffda
│  │  ├─ d5
│  │  │  └─ 64d0bc3dd917926892c55e3706cc116d5b165e
│  │  ├─ d6
│  │  │  ├─ 32583e3a0584bf639acb32ff6fa52c096cb77c
│  │  │  └─ 44be469e0b8bd80d1ab1e90181cec693d86cef
│  │  ├─ da
│  │  │  └─ 6c7b812c4109a0062db924256f7ca396d53ff0
│  │  ├─ dc
│  │  │  ├─ 830b76af4204005a34e9fb59c44f8366699c8a
│  │  │  └─ 8b36de1a42fe92a0a4f1434c87e78d91ef6958
│  │  ├─ e1
│  │  │  └─ e3d567f337feb4b270dc0135da93ee81bb6512
│  │  ├─ e3
│  │  │  └─ 447117e430ac32335ea543158485c9d0da8efa
│  │  ├─ e4
│  │  │  └─ f2b97ad722c7991b1b09438a52b8cb2118aed0
│  │  ├─ e6
│  │  │  ├─ 5e639971c3fd2e41c8971bdbad3e12907959b0
│  │  │  └─ 9de29bb2d1d6434b8b29ae775ad8c2e48c5391
│  │  ├─ ea
│  │  │  ├─ 76d8938b7047b04b5b6d843b1d95e21a04a6fb
│  │  │  └─ ef0aaf44b79fb07ec5ad07f3910a39c8e37990
│  │  ├─ f0
│  │  │  └─ e8a56d238a199158d164ae6ec560a7a6575ffb
│  │  ├─ f1
│  │  │  ├─ 57a876207f7fb04fed98a4bb4e0a15c777eab1
│  │  │  └─ fc26815960cda87f65daef7a97e1520ec8f04c
│  │  ├─ f2
│  │  │  └─ 51489b316fac5c20a9126e16964dbfeb6afad8
│  │  ├─ f3
│  │  │  └─ 113e02b7fa60ed086708a9b1b05778e411f385
│  │  ├─ f6
│  │  │  └─ de5bdf4c5c0d7a1529025d119fee037af4372a
│  │  ├─ f8
│  │  │  └─ 8dfaf38c534f02225e13e3c32d16f98ef69f33
│  │  ├─ f9
│  │  │  └─ edc404cafe2884fed8e7f20de34be7c63e10d5
│  │  ├─ info
│  │  └─ pack
│  │     ├─ pack-a3e4badc67d264fdd23257998e97110d409579cf.idx
│  │     └─ pack-a3e4badc67d264fdd23257998e97110d409579cf.pack
│  ├─ packed-refs
│  └─ refs
│     ├─ heads
│     │  └─ main
│     ├─ remotes
│     │  └─ origin
│     │     ├─ HEAD
│     │     └─ main
│     └─ tags
├─ .gitignore
├─ api
│  ├─ dto
│  │  ├─ errcode.go
│  │  ├─ foo.go
│  │  ├─ general.go
│  │  ├─ link.go
│  │  ├─ server.go
│  │  └─ user.go
│  ├─ init.go
│  └─ route
│     ├─ foo.go
│     ├─ link.go
│     ├─ route.go
│     ├─ server.go
│     └─ user.go
├─ cmd
│  ├─ init.go
│  └─ root.go
├─ config_sample.yaml
├─ dockerfile
├─ docs
│  └─ design.md
├─ go.mod
├─ go.sum
├─ internal
│  ├─ controller
│  │  ├─ foo.go
│  │  ├─ link.go
│  │  ├─ mid.go
│  │  ├─ server.go
│  │  └─ user.go
│  ├─ dao
│  │  ├─ complex_crud.go
│  │  ├─ dao.go
│  │  └─ model
│  │     ├─ foo.go
│  │     ├─ link.go
│  │     └─ user.go
│  └─ service
│     └─ .gitkeep
├─ main.go
├─ README.md
└─ utils
   ├─ logger
   │  └─ logger.go
   └─ stacktrace
      ├─ cause.go
      ├─ cause_test.go
      ├─ cleanpath
      │  ├─ gopath.go
      │  └─ gopath_test.go
      ├─ codes_for_test.go
      ├─ doc.go
      ├─ format.go
      ├─ format_test.go
      ├─ functions_for_test.go
      ├─ LICENSE
      ├─ README.md
      ├─ stacktrace.go
      └─ stacktrace_test.go

```