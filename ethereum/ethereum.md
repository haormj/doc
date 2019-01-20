## ethereum

### 安装Geth

```shell
apt install software-properties-common
add-apt-repository -y ppa:ethereum/ethereum
apt update
apt install ethereum
```

### 搭建以太坊私链

配置文件

```json
{
    "config": {
        "chainId": 98,
        "homesteadBlock": 0,
        "eip155Block": 0,
        "eip158Block": 0
    },
    "difficulty": "200000000",
    "gasLimit": "2100000",
    "alloc": {
        "7df9a8751174b3bc565e6424a0050ebc1b2d1d82": {
            "balance": "300000"
        },
        "f41c74c9ae680c1aa78f42e5647a62f353b7bdde": {
            "balance": "400000"
        }
    }
}
```

运行

```
geth --datadir /path/to/data init config.json
```

进入命令行交互界面

```shell
# 进入交互界面
geth --datadir /path/to/data --networkid 98 console
# 创建账户 0x7d54e68845378fa63b76f63d8f4c0d3b2c521c2a
personal.newAccount()
# 获取账户余额
eth.getBalance("0x7d54e68845378fa63b76f63d8f4c0d3b2c521c2a")
# 退出终端
```

启动挖矿

```shell
# command
geth --datadir /path/to/data --networkid 98 --mine --minerthreads 1 --etherbase 0x7d54e68845378fa63b76f63d8f4c0d3b2c521c2a
```

挖矿日志

```
haormj@pc:~/Project/ethereum$ geth --datadir ./ --networkid 98 --mine --minerthreads 1 --etherbase 0x7d54e68845378fa63b76f63d8f4c0d3b2c521c2a 
INFO [06-21|23:05:59] Maximum peer count                       ETH=25 LES=0 total=25
INFO [06-21|23:05:59] Starting peer-to-peer node               instance=Geth/v1.8.11-stable-dea1ce05/linux-amd64/go1.10.1
INFO [06-21|23:05:59] Allocated cache and file handles         database=/home/haormj/Project/ethereum/geth/chaindata cache=768 handles=512
INFO [06-21|23:05:59] Initialised chain configuration          config="{ChainID: 98 Homestead: 0 DAO: <nil> DAOSupport: false EIP150: <nil> EIP155: 0 EIP158: 0 Byzantium: <nil> Constantinople: <nil> Engine: unknown}"
INFO [06-21|23:05:59] Disk storage enabled for ethash caches   dir=/home/haormj/Project/ethereum/geth/ethash count=3
INFO [06-21|23:05:59] Disk storage enabled for ethash DAGs     dir=/home/haormj/.ethash                      count=2
INFO [06-21|23:05:59] Initialising Ethereum protocol           versions="[63 62]" network=98
INFO [06-21|23:05:59] Loaded most recent local header          number=0 hash=6a8dfe…017773 td=200000000
INFO [06-21|23:05:59] Loaded most recent local full block      number=0 hash=6a8dfe…017773 td=200000000
INFO [06-21|23:05:59] Loaded most recent local fast block      number=0 hash=6a8dfe…017773 td=200000000
INFO [06-21|23:05:59] Loaded local transaction journal         transactions=0 dropped=0
INFO [06-21|23:05:59] Regenerated local transaction journal    transactions=0 accounts=0
INFO [06-21|23:05:59] Starting P2P networking 
INFO [06-21|23:06:01] UDP listener up                          self=enode://c0425d2a7f2c3f984d3932d7eab32d48263f715118e30325ab849bc823403e97f1aacd08d567d771db0baa36fea6fd76dc2b385be116795ba181bc22d282f4d0@[::]:30303
INFO [06-21|23:06:01] RLPx listener up                         self=enode://c0425d2a7f2c3f984d3932d7eab32d48263f715118e30325ab849bc823403e97f1aacd08d567d771db0baa36fea6fd76dc2b385be116795ba181bc22d282f4d0@[::]:30303
INFO [06-21|23:06:01] IPC endpoint opened                      url=/home/haormj/Project/ethereum/geth.ipc
INFO [06-21|23:06:01] Transaction pool price threshold updated price=18000000000
INFO [06-21|23:06:01] Starting mining operation 
INFO [06-21|23:06:01] Commit new mining work                   number=1 txs=0 uncles=0 elapsed=273.356µs
INFO [06-21|23:06:03] Generating DAG in progress               epoch=0 percentage=0 elapsed=1.390s
INFO [06-21|23:06:04] Generating DAG in progress               epoch=0 percentage=1 elapsed=2.771s
INFO [06-21|23:06:06] Generating DAG in progress               epoch=0 percentage=2 elapsed=4.147s

INFO [06-21|23:11:34] Generated ethash verification cache      epoch=1 elapsed=3m1.620s
INFO [06-21|23:23:45] Successfully sealed new block            number=1 hash=b94c1f…af697d
INFO [06-21|23:23:45] 🔨 mined potential block                  number=1 hash=b94c1f…af697d
INFO [06-21|23:23:45] Commit new mining work                   number=2 txs=0 uncles=0 elapsed=157.427µs

```

智能合约

```
pragma solidity ^0.4.11;
contract helloworld {
    function renderHelloWorld() returns (string) {
        return 'helloworld';
    }
}
```

安装Solidity编译器

```
# 安装Solidity编译器
apt install -y solc
```

编译结果

```
haormj@pc:~/Project/ethereum$ solc --bin helloworld.sol 
helloworld.sol:3:5: Warning: No visibility specified. Defaulting to "public". 
    function renderHelloWorld() returns (string) {
    ^ (Relevant source part starts here and spans across multiple lines).
helloworld.sol:3:5: Warning: Function state mutability can be restricted to pure
    function renderHelloWorld() returns (string) {
    ^ (Relevant source part starts here and spans across multiple lines).

======= helloworld.sol:helloworld =======
Binary: 
608060405234801561001057600080fd5b5061013f806100206000396000f300608060405260043610610041576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063942ae0a714610046575b600080fd5b34801561005257600080fd5b5061005b6100d6565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561009b578082015181840152602081019050610080565b50505050905090810190601f1680156100c85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60606040805190810160405280600a81526020017f68656c6c6f776f726c64000000000000000000000000000000000000000000008152509050905600a165627a7a72305820647dc46dbd6d6620a55a3905ff500ae8ed9525faa2da1ad3b2cf112186b744670029
```

继续执行

```
haormj@pc:~/Project/ethereum$ solc --abi helloworld.sol 
helloworld.sol:3:5: Warning: No visibility specified. Defaulting to "public". 
    function renderHelloWorld() returns (string) {
    ^ (Relevant source part starts here and spans across multiple lines).
helloworld.sol:3:5: Warning: Function state mutability can be restricted to pure
    function renderHelloWorld() returns (string) {
    ^ (Relevant source part starts here and spans across multiple lines).

======= helloworld.sol:helloworld =======
Contract JSON ABI 
[{"constant":false,"inputs":[],"name":"renderHelloWorld","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"nonpayable","type":"function"}]
```

部署

```
code="0x608060405234801561001057600080fd5b5061013f806100206000396000f300608060405260043610610041576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063942ae0a714610046575b600080fd5b34801561005257600080fd5b5061005b6100d6565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561009b578082015181840152602081019050610080565b50505050905090810190601f1680156100c85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60606040805190810160405280600a81526020017f68656c6c6f776f726c64000000000000000000000000000000000000000000008152509050905600a165627a7a72305820647dc46dbd6d6620a55a3905ff500ae8ed9525faa2da1ad3b2cf112186b744670029"
abi=[{"constant":false,"inputs":[],"name":"renderHelloWorld","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"nonpayable","type":"function"}]
personal.unlockAccount("0x7d54e68845378fa63b76f63d8f4c0d3b2c521c2a")
myHelloWorld=eth.contract(abi)
contract=myHelloWorld.new({from:"0x7d54e68845378fa63b76f63d8f4c0d3b2c521c2a",data:code,gas:1000000})
```

执行结果

```
INFO [06-21|23:37:26] Submitted contract creation              fullhash=0x3a22360e3d91f22422ae62ddc8658e5715dc3220d1283ff3e19a8f0cb7110714 contract=0x7C9b3A4767ae0e494b075B64d7D02Ee2CC65934D
{
  abi: [{
      constant: false,
      inputs: [],
      name: "renderHelloWorld",
      outputs: [{...}],
      payable: false,
      stateMutability: "nonpayable",
      type: "function"
  }],
  address: undefined,
  transactionHash: "0x3a22360e3d91f22422ae62ddc8658e5715dc3220d1283ff3e19a8f0cb7110714"
}
```

查看交易

```shell
# 查看交易状态
txpool.status

{
  pending: 1,
  queued: 0
}
```

重启启动挖矿

```
haormj@pc:~/Project/ethereum$ geth --datadir ./ --networkid 98 --mine --minerthreads 1 --etherbase 0x7d54e68845378fa63b76f63d8f4c0d3b2c521c2a
INFO [06-21|23:42:49] Maximum peer count                       ETH=25 LES=0 total=25
INFO [06-21|23:42:49] Starting peer-to-peer node               instance=Geth/v1.8.11-stable-dea1ce05/linux-amd64/go1.10.1
INFO [06-21|23:42:49] Allocated cache and file handles         database=/home/haormj/Project/ethereum/geth/chaindata cache=768 handles=512
INFO [06-21|23:42:49] Initialised chain configuration          config="{ChainID: 98 Homestead: 0 DAO: <nil> DAOSupport: false EIP150: <nil> EIP155: 0 EIP158: 0 Byzantium: <nil> Constantinople: <nil> Engine: unknown}"
INFO [06-21|23:42:49] Disk storage enabled for ethash caches   dir=/home/haormj/Project/ethereum/geth/ethash count=3
INFO [06-21|23:42:49] Disk storage enabled for ethash DAGs     dir=/home/haormj/.ethash                      count=2
INFO [06-21|23:42:49] Initialising Ethereum protocol           versions="[63 62]" network=98
INFO [06-21|23:42:49] Loaded most recent local header          number=1 hash=b94c1f…af697d td=390332056
INFO [06-21|23:42:49] Loaded most recent local full block      number=1 hash=b94c1f…af697d td=390332056
INFO [06-21|23:42:49] Loaded most recent local fast block      number=1 hash=b94c1f…af697d td=390332056
INFO [06-21|23:42:49] Loaded local transaction journal         transactions=1 dropped=0
INFO [06-21|23:42:49] Regenerated local transaction journal    transactions=1 accounts=1
WARN [06-21|23:42:49] Blockchain not empty, fast sync disabled 
INFO [06-21|23:42:49] Starting P2P networking 
INFO [06-21|23:42:51] UDP listener up                          self=enode://c0425d2a7f2c3f984d3932d7eab32d48263f715118e30325ab849bc823403e97f1aacd08d567d771db0baa36fea6fd76dc2b385be116795ba181bc22d282f4d0@[::]:30303
INFO [06-21|23:42:51] RLPx listener up                         self=enode://c0425d2a7f2c3f984d3932d7eab32d48263f715118e30325ab849bc823403e97f1aacd08d567d771db0baa36fea6fd76dc2b385be116795ba181bc22d282f4d0@[::]:30303
INFO [06-21|23:42:51] IPC endpoint opened                      url=/home/haormj/Project/ethereum/geth.ipc
INFO [06-21|23:42:51] Transaction pool price threshold updated price=18000000000
INFO [06-21|23:42:51] Starting mining operation 
INFO [06-21|23:42:51] Commit new mining work                   number=2 txs=1 uncles=0 elapsed=689.747µs
INFO [06-22|00:19:27] Successfully sealed new block            number=2 hash=e4761e…e91152
INFO [06-22|00:19:27] 🔨 mined potential block                  number=2 hash=e4761e…e91152
INFO [06-22|00:19:27] Commit new mining work                   number=3 txs=0 uncles=0 elapsed=211.09µs
INFO [06-22|00:21:51] Successfully sealed new block            number=3 hash=9f9d4e…3b5679
INFO [06-22|00:21:51] 🔨 mined potential block                  number=3 hash=9f9d4e…3b5679
INFO [06-22|00:21:51] Commit new mining work                   number=4 txs=0 uncles=0 elapsed=136.495µs
INFO [06-22|00:46:57] Regenerated local transaction journal    transactions=0 accounts=0
INFO [06-22|00:57:01] Successfully sealed new block            number=4 hash=1f9921…367407
INFO [06-22|00:57:01] 🔨 mined potential block                  number=4 hash=1f9921…367407
INFO [06-22|00:57:01] Commit new mining work                   number=5 txs=0 uncles=0 elapsed=160.226µs
INFO [06-22|01:14:13] Successfully sealed new block            number=5 hash=231a67…61e0c7
INFO [06-22|01:14:13] 🔨 mined potential block                  number=5 hash=231a67…61e0c7
INFO [06-22|01:14:13] Commit new mining work                   number=6 txs=0 uncles=0 elapsed=133.145µs
INFO [06-22|01:46:55] Successfully sealed new block            number=6 hash=d09bc1…a2acfa
INFO [06-22|01:46:55] 🔨 mined potential block                  number=6 hash=d09bc1…a2acfa
INFO [06-22|01:46:55] Commit new mining work                   number=7 txs=0 uncles=0 elapsed=184.717µs
INFO [06-22|01:46:57] Regenerated local transaction journal    transactions=0 accounts=0
INFO [06-22|02:28:27] Successfully sealed new block            number=7 hash=e1bbff…6be396
INFO [06-22|02:28:27] 🔗 block reached canonical chain          number=2 hash=e4761e…e91152
INFO [06-22|02:28:27] 🔨 mined potential block                  number=7 hash=e1bbff…6be396
INFO [06-22|02:28:27] Commit new mining work                   number=8 txs=0 uncles=0 elapsed=150.294µs
INFO [06-22|02:44:47] Successfully sealed new block            number=8 hash=c43a6d…c339a9
INFO [06-22|02:44:47] 🔗 block reached canonical chain          number=3 hash=9f9d4e…3b5679
INFO [06-22|02:44:47] 🔨 mined potential block                  number=8 hash=c43a6d…c339a9
INFO [06-22|02:44:47] Commit new mining work                   number=9 txs=0 uncles=0 elapsed=142.477µs
INFO [06-22|02:46:57] Regenerated local transaction journal    transactions=0 accounts=0
INFO [06-22|03:26:30] Successfully sealed new block            number=9 hash=f7e19e…891dee
INFO [06-22|03:26:30] 🔗 block reached canonical chain          number=4 hash=1f9921…367407
INFO [06-22|03:26:30] 🔨 mined potential block                  number=9 hash=f7e19e…891dee
INFO [06-22|03:26:30] Commit new mining work                   number=10 txs=0 uncles=0 elapsed=205.885µs
INFO [06-22|03:46:57] Regenerated local transaction journal    transactions=0 accounts=0
INFO [06-22|04:05:26] Successfully sealed new block            number=10 hash=30340f…a95e01
INFO [06-22|04:05:26] 🔗 block reached canonical chain          number=5  hash=231a67…61e0c7
INFO [06-22|04:05:26] 🔨 mined potential block                  number=10 hash=30340f…a95e01
INFO [06-22|04:05:26] Commit new mining work                   number=11 txs=0 uncles=0 elapsed=198.75µs
INFO [06-22|04:09:45] Successfully sealed new block            number=11 hash=cce2f2…8cafa3
INFO [06-22|04:09:45] 🔗 block reached canonical chain          number=6  hash=d09bc1…a2acfa
INFO [06-22|04:09:45] 🔨 mined potential block                  number=11 hash=cce2f2…8cafa3
INFO [06-22|04:09:45] Commit new mining work                   number=12 txs=0 uncles=0 elapsed=175.472µs
INFO [06-22|04:46:57] Regenerated local transaction journal    transactions=0 accounts=0
INFO [06-22|05:34:19] Successfully sealed new block            number=12 hash=4e48f2…9b256f
INFO [06-22|05:34:19] 🔗 block reached canonical chain          number=7  hash=e1bbff…6be396
INFO [06-22|05:34:19] 🔨 mined potential block                  number=12 hash=4e48f2…9b256f
INFO [06-22|05:34:19] Commit new mining work                   number=13 txs=0 uncles=0 elapsed=126.899µs
INFO [06-22|05:46:57] Regenerated local transaction journal    transactions=0 accounts=0
INFO [06-22|06:08:08] Successfully sealed new block            number=13 hash=abc5ff…9c8683
INFO [06-22|06:08:08] 🔗 block reached canonical chain          number=8  hash=c43a6d…c339a9
INFO [06-22|06:08:08] 🔨 mined potential block                  number=13 hash=abc5ff…9c8683
INFO [06-22|06:08:08] Commit new mining work                   number=14 txs=0 uncles=0 elapsed=362.678µs
INFO [06-22|06:42:19] Successfully sealed new block            number=14 hash=aa44f1…d4cf3c
INFO [06-22|06:42:19] 🔗 block reached canonical chain          number=9  hash=f7e19e…891dee
INFO [06-22|06:42:19] 🔨 mined potential block                  number=14 hash=aa44f1…d4cf3c
INFO [06-22|06:42:19] Commit new mining work                   number=15 txs=0 uncles=0 elapsed=133.344µs
INFO [06-22|06:46:57] Regenerated local transaction journal    transactions=0 accounts=0
INFO [06-22|07:37:39] Successfully sealed new block            number=15 hash=b05939…a03f31
INFO [06-22|07:37:39] 🔗 block reached canonical chain          number=10 hash=30340f…a95e01
INFO [06-22|07:37:39] 🔨 mined potential block                  number=15 hash=b05939…a03f31
INFO [06-22|07:37:39] Commit new mining work                   number=16 txs=0 uncles=0 elapsed=208.091µs
INFO [06-22|07:44:21] Successfully sealed new block            number=16 hash=35f9b0…cb68e7
INFO [06-22|07:44:21] 🔗 block reached canonical chain          number=11 hash=cce2f2…8cafa3
INFO [06-22|07:44:21] 🔨 mined potential block                  number=16 hash=35f9b0…cb68e7
INFO [06-22|07:44:21] Commit new mining work                   number=17 txs=0 uncles=0 elapsed=129.959µs
INFO [06-22|07:45:51] Successfully sealed new block            number=17 hash=6fad93…c6f1ad
INFO [06-22|07:45:51] 🔗 block reached canonical chain          number=12 hash=4e48f2…9b256f
INFO [06-22|07:45:51] 🔨 mined potential block                  number=17 hash=6fad93…c6f1ad
INFO [06-22|07:45:51] Commit new mining work                   number=18 txs=0 uncles=0 elapsed=207.78µs
INFO [06-22|07:46:57] Regenerated local transaction journal    transactions=0 accounts=0
```

查看交易

```shell
> txpool.status
{
  pending: 0,
  queued: 0
}
```

