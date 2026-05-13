import redis
from redis.cluster import ClusterNode
import redis.asyncio as redis_asyncio

"""
单服务器:
    r = Redis() # 实例化
    r = Redis.from_url()    # 通过 url 获取实例
    r = Redis.from_pool()   # 通过连接池获取实例

哨兵模式:
    sentinel = Sentinel(sentinels) # 实例化
    master = sentinel.master_for('master_name') # 获取主节点实例
    slave = sentinel.slave_for('master_name')   # 获取从节点实例

    master_addr = sentinel.discover_master('master_name')   # 获取主节点地址
    slaves = sentinel.discover_slaves('master_name')    # 获取从节点

集群模式:
    cluster = RedisCluster.from_url() # 通过 url 获取节点连接

    nodes = ClusterNode()   # 实例化节点
    cluster = RedisCluster(nodes)  # 实例化

异步方式:
    redis.asyncio 包, 使用方式同上.
"""
# 单服务器
# 实例化 Redis
r1 = redis.Redis(host='172.31.218.169', port=6379, password='mypassword', db=0, decode_responses=True)

# # 使用 url 获取实例
# r2 = redis.Redis.from_url('redis://default:mypassword@172.31.218.169:6379/0')

# # 使用连接池获取实例
# pool = redis.connection.ConnectionPool(host='172.31.218.169', port=6379, password='mypassword', db=0, max_connections=20,
#                                        decode_responses=True)
# r3 = redis.Redis.from_pool(pool)

# 使用连接
pong = r1.ping()
print(f'pong: {pong}')

r1.set('name', 'tom', ex=20)
print(f'r1.get("name"): {r1.get("name")}')
print(f'r1.ttl("name"): {r1.ttl("name")}')
print('-' * 30)

# 哨兵模式
# 哨兵服务器地址列表
sentinel_list = [
    ('sentinel1.example.com', 26379),
    ('sentinel2.example.com', 26379),
    ('sentinel3.example.com', 26379)
]
# 实例化 Sentinel, password='sentinel_password'  # 哨兵密码
sentinel = redis.Sentinel(sentinel_list, password='', socket_timeout=0.5)
# 获取主节点连接, password='redis_password'    # Redis 密码
master = sentinel.master_for('mymaster', password='', socket_timeout=0.5, decode_responses=True)
# 获取从节点连接, password='redis_password'    # Redis 密码
slave = sentinel.slave_for('mymaster', password='', socket_timeout=0.5, decode_responses=True)

# 获取主从节点 IP
master_addr = sentinel.discover_master('mymaster', password='')
print(f'master_addr: {master_addr}')
# 获取从节点
slaves = sentinel.discover_slaves('mymaster', password='')
print(f'slaves: {slaves}')

# 使用连接
master.set('sentinel:name', 'python', ex=20)
print(f'slave.get("sentinel:name"): {slave.get("sentinel:name")}')
print(f'slave.ttl("sentinel:name"): {slave.ttl("sentinel:name")}')
print('-' * 30)

# 集群模式
# 使用 url 获取实例
# cluster = redis.RedisCluster.from_url('redis://user:pass@hostname:6379/0')

# 集群服务器地址列表
cluster_list = [
    ClusterNode('cluster1.example.com', 6379),
    ClusterNode('cluster2.example.com', 6379),
    ClusterNode('cluster3.example.com', 6379),
]
# 实例化 RedisCluster
cluster = redis.RedisCluster(startup_nodes=cluster_list, decode_responses=True, max_connections_per_node=20)
# 根据 key 计算节点
# node =  cluster.get_node_from_key('cluster:name')

# 使用连接
cluster.set('cluster:name', 'python', ex=20)
print(f'cluster.get("cluster:name"): {cluster.get("cluster:name")}')
print(f'cluster.ttl("cluster:name"): {cluster.ttl("cluster:name")}')
print('-' * 30)

# 异步方式
# 单服务器
r4 = redis_asyncio.Redis(host='localhost', port=6379, password='', db=0, decode_responses=True)

# 哨兵模式
sentinel = redis_asyncio.Sentinel(sentinel_list, password='', socket_timeout=0.5)
master = sentinel.master_for('mymaster', password='', socket_timeout=0.5, decode_responses=True)
slave = sentinel.slave_for('mymaster', password='', socket_timeout=0.5, decode_responses=True)

# 集群模式
cluster = redis_asyncio.RedisCluster(startup_nodes=cluster_list, decode_responses=True, max_connections=20)


async def set():
    await master.set('sentinel:name', 'python', ex=20)

    await cluster.set('cluster:name', 'python', ex=20)


print('-' * 30)

# 订阅
p = r1.pubsub()
p.subscribe('channel')
for message in p.listen():
    print(message)

# 发布消息
p.publish('channel', 'hello')
# 取消订阅
p.unsubscribe('channel')

# 订阅模式
p.psubscribe('_gg_:')
# 取消订阅模式
p.punsubscribe('_gg_:')

# 通过频道订阅的数量
r1.pubsub_numsub('channel')
# 获取所有订阅的频道
r1.pubsub_channels()
# 通过模式订阅的数量
r1.pubsub_numpat()

print('-' * 30)

lua_script = """
local value = redis.call('GET', KEYS[1])
value = tonumber(value)
return value * ARGV[1]
"""

# 脚本
r1.eval('return {KEYS[1], KEYS[2], ARGV[1], ARGV[2]}', 2, 'key1', 'key2', 'first', 'second')

# 实例化可调用脚本
script_sha = r1.register_script('return {KEYS[1], KEYS[2], ARGV[1], ARGV[2]}')
result = script_sha(keys=['key1', 'key2'], args=['first', 'second'])
