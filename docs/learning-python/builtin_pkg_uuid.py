"""
Python uuid 模块完整指南 (Python 3.14)

本文件演示 Python 标准库 uuid 包的使用方法和最佳实践。
uuid 模块实现了 RFC 4122 / RFC 9562 中定义的 UUID 标准版本。

作者：基于 Python 3.14 完整使用指南
版本：Python 3.14+
"""

import uuid


# ============================================================================
# §1 UUID 版本对比
# ============================================================================

"""
UUID 版本选择矩阵:

| 方法 | 版本 | 基于 | 确定性？ | 可按时间排序？ | 泄露信息？ |
|------|------|------|----------|----------------|------------|
| uuid1() | 1 | 时间戳 + MAC 地址 | 否 | 部分 | 是 (MAC+ 时间) |
| uuid3(ns, name) | 3 | namespace+name 的 MD5 哈希 | 是 | 否 | 否 |
| uuid4() | 4 | 随机 (122 位随机) | 否 | 否 | 否 |
| uuid5(ns, name) | 5 | namespace+name 的 SHA-1 哈希 | 是 | 否 | 否 |
| uuid6() | 6 | 重排序的时间戳 (3.14+) | 否 | 是 | 是 (时间) |
| uuid7() | 7 | Unix 纪元毫秒 + 随机 (3.14+) | 否 | 是 | 是 (时间) |
| uuid8(a,b,c) | 8 | 完全自定义 (3.14+) | 视情况 | 视情况 | 视情况 |

说明:
- 确定性：相同的输入是否总是产生相同的输出
- 时间可排序：生成的 UUID 按字节/整数比较时是否保持时间顺序
- 泄露信息：是否暴露主机标识符 (如 MAC 地址) 或精确时间
"""


# ============================================================================
# §2 各版本详细说明
# ============================================================================

def demo_uuid_v1():
    """UUID Version 1:基于时间和硬件地址
    
    将当前时间 (自 1582 年起的 100 纳秒计数) 与机器的 MAC 地址组合。
    
    优点：实践中唯一，包含时间信息
    缺点：会暴露主机和网络地址，隐私敏感场景应避免
    适用场景：审计日志、需要追踪生成时间的场合
    """
    print("=" * 60)
    print("UUID v1:时间 + MAC 地址")
    print("=" * 60)
    u = uuid.uuid1()
    print(f"UUID: {u}")
    print(f"版本：{u.version}")  # 1
    print(f"节点 (MAC): {hex(u.node)}")
    print(f"时间戳：{u.time}")
    print()


def demo_uuid_v3():
    """UUID Version 3:基于 MD5 哈希
    
    使用 MD5 算法对 (namespace, name) 进行哈希。
    
    优点：确定性输出，相同输入永远产生相同结果
    缺点：MD5 密码学上已不安全；不可时间排序
    适用场景：稳定的名称映射 ID(不再推荐使用，建议用 v5)
    
    警告：MD5 已不推荐用于安全相关场景
    """
    print("=" * 60)
    print("UUID v3:MD5 哈希 (不推荐)")
    print("=" * 60)
    # 注意：这些输出是确定性的！相同输入总是产生相同结果
    u = uuid.uuid3(uuid.NAMESPACE_DNS, "python.org")
    print(f"uuid3(NAMESPACE_DNS, 'python.org'): {u}")
    # 再次生成相同输入应得到相同结果
    u2 = uuid.uuid3(uuid.NAMESPACE_DNS, "python.org")
    assert u == u2, "v3 必须是确定性的"
    print(f"验证确定性：{u == u2}")
    print(f"版本：{u.version}")  # 3
    print(f"哈希类型：MD5")
    print()


def demo_uuid_v4():
    """UUID Version 4:纯随机 UUID
    
    生成 122 位随机数的 UUID，是最常用的默认选择。
    
    优点：无隐私泄露风险;生成速度快
    缺点：不可排序;存在极小概率碰撞
    适用场景：会话 ID、令牌、密钥、通用唯一标识符
    
    统计：122 位随机空间，碰撞概率可忽略不计
    """
    print("=" * 60)
    print("UUID v4:纯随机 (默认选择)")
    print("=" * 60)
    for i in range(3):
        u = uuid.uuid4()
        print(f"{i+1}. {u}")
    print(f"版本：{u.version}")  # 4
    print()


def demo_uuid_v5():
    """UUID Version 5:基于 SHA-1 哈希
    
    使用 SHA-1 算法对 (namespace, name) 进行哈希。
    
    优点：确定性输出;比 MD5 更安全;跨平台一致
    缺点：不可时间排序;SHA-1 也在逐步被弃用 (但仍是行业标准)
    适用场景：从域名/URL/标识符派生的稳定 ID(推荐使用)
    
    推荐理由：优先使用 v5 而非 v3，因为 SHA-1 比 MD5 更强
    """
    print("=" * 60)
    print("UUID v5:SHA-1 哈希 (推荐)")
    print("=" * 60)
    # 注意：这些输出是确定性的！相同输入总是产生相同结果
    u = uuid.uuid5(uuid.NAMESPACE_DNS, "python.org")
    print(f"uuid5(NAMESPACE_DNS, 'python.org'): {u}")
    # 再次生成相同输入应得到相同结果
    u2 = uuid.uuid5(uuid.NAMESPACE_DNS, "python.org")
    assert u == u2, "v5 必须是确定性的"
    print(f"验证确定性：{u == u2}")
    print(f"版本：{u.version}")  # 5
    print(f"哈希类型：SHA-1")
    print()


def demo_uuid_v6():
    """UUID Version 6:重排序时间戳 (Python 3.14+,RFC 9562)
    
    类似 uuid1()，但时间戳字段被重新排列，使其按字节可正确排序。
    
    优点：时间有序且可精确排序;保留 v1 的大部分信息
    缺点：仍然泄露时间信息;需要 MAC 地址支持
    适用场景：数据库主键、需要按创建时间排序的场景
    
    改进点：相比 v1，解决了时间字段字节序导致的排序问题
    """
    print("=" * 60)
    print("UUID v6:重排序时间戳 (3.14+)")
    print("=" * 60)
    u = uuid.uuid6()
    print(f"UUID: {u}")
    print(f"版本：{u.version}")  # 6
    print(f"时间戳：{u.time} (可直接比较)")
    print(f"节点 (MAC): {hex(u.node)}")
    print()


def demo_uuid_v7():
    """UUID Version 7:Unix 毫秒时间戳 + 随机 (Python 3.14+,RFC 9562)
    
    使用 48 位 Unix 纪元毫秒时间戳，配合随机数和计数器实现单调递增。
    
    优点：时间有序;对数据库索引友好;可控制熵源
    缺点：泄露时间信息;需要系统时钟
    适用场景：数据库主键 (强烈推荐)、现代应用 ID 生成
    
    特性:
    - 同一毫秒内的多次调用通过 42 位计数器保证唯一性
    - 时间戳在高位，便于按字节排序
    - 32 位随机数提供足够的熵
    """
    print("=" * 60)
    print("UUID v7:Unix 毫秒 + 随机 (3.14+, 推荐为主键)")
    print("=" * 60)
    for i in range(3):
        u = uuid.uuid7()
        print(f"{i+1}. {u}")
        print(f"   版本：{u.version} | 时间：{u.time} ms")
    print()


def demo_uuid_v8():
    """UUID Version 8:完全自定义 (Python 3.14+,RFC 9562)
    
    允许用户自定义 UUID 的各个组成部分。
    
    参数:
    - a: 前 48 位 (octets 0-5)
    - b: 中间 12 位 (octets 6-7)
    - c: 最后 62 位 (octets 8-15)
    
    优点：完全灵活;可实现任意方案
    缺点：需要手动管理格式;需遵守 variant/version 位规则
    适用场景：特殊需求、遗留系统兼容、实验性方案
    """
    print("=" * 60)
    print("UUID v8:自定义方案 (3.14+)")
    print("=" * 60)
    # 示例：完全自定义的 UUID
    a = 0x123456789abc  # 48 bits
    b = 0xdef           # 12 bits
    c = 0x123456789abcdef012  # 62 bits
    u = uuid.uuid8(a, b, c)
    print(f"uuid8(0x123456789abc, 0xdef, 0x123456789abcdef012)")
    print(f"=> {u}")
    print(f"版本：{u.version}")  # 8
    print()


# ============================================================================
# §3 预定义命名空间详解
# ============================================================================

def demo_namespaces():
    """预定义命名空间 (RFC 4122 §7.8)
    
    命名空间是标准的 UUID，用作 v3/v5 哈希的锚点，确保全球一致性。
    
    四大作用:
    1. 标准化：所有实现都同意这些 UUID(RFC 4122 规定)
    2. 确定性：相同的 (namespace, name) 永远产生相同的结果
    3. 跨平台一致：你的代码与他人使用相同命名空间会得到相同 UUID
    4. 语义意义：每个命名空间有预期的使用场景
    
    可用的预定义命名空间:
    - NAMESPACE_DNS: DNS 域名
    - NAMESPACE_URL: 绝对 URL  
    - NAMESPACE_OID: ISO OID 标识符
    - NAMESPACE_X500: X.500 可分辨名称 (LDAP)
    """
    print("=" * 60)
    print("预定义命名空间 (RFC 4122 §7.8)")
    print("=" * 60)
    
    namespaces = [
        ("NAMESPACE_DNS", uuid.NAMESPACE_DNS, "DNS 域名"),
        ("NAMESPACE_URL", uuid.NAMESPACE_URL, "绝对 URL"),
        ("NAMESPACE_OID", uuid.NAMESPACE_OID, "ISO OID 标识符"),
        ("NAMESPACE_X500", uuid.NAMESPACE_X500, "X.500 可分辨名称 (LDAP)"),
    ]
    
    for name, ns_id, desc in namespaces:
        print(f"\n{name}:")
        print(f"  UUID: {ns_id}")
        print(f"  用途：{desc}")
        
        # 展示使用该命名空间的示例
        if name == "NAMESPACE_DNS":
            example = uuid.uuid5(ns_id, "example.com")
            print(f"  示例：uuid5({name}, 'example.com') => {example}")
        elif name == "NAMESPACE_URL":
            example = uuid.uuid5(ns_id, "https://www.python.org/")
            print(f"  示例：uuid5({name}, 'https://www.python.org/') => {example}")
    
    print("\n原理：hash(namespace.bytes + name UTF-8 字节)")
    print("       namespace bytes 作为'锚',将 name 绑定到特定上下文")
    print()


# ============================================================================
# §4 NIL / MAX 哨兵 UUID (Python 3.14+,RFC 9562)
# ============================================================================

def demo_sentinel_uuids():
    """Nil UUID 和 Max UUID(RFC 9562 §5.9/§5.10,3.14+)
    
    Nil UUID(全零):
    - 值：00000000-0000-0000-0000-000000000000
    - 含义：表示"无值"或"未初始化"
    - 特性：最小可能 UUID，可哈希、可排序
    - 用法：占位符、数据库 NULL 代理、标志未分配 ID
    
    Max UUID(全 F):
    - 值：ffffffff-ffff-ffff-ffff-ffffffffffff
    - 含义：表示"最大可能值"
    - 特性：最大可能 UUID，可哈希、可排序
    - 用法：范围查询上界、排序终止标记、区间查询哨兵
    """
    print("=" * 60)
    print("哨兵 UUID(3.14+)")
    print("=" * 60)
    
    print("\n--- Nil UUID ---")
    nil = uuid.NIL
    print(f"UUID: {nil}")
    print(f"int 值：{nil.int}")  # 0
    print(f"用途：未初始化的占位符，替代 NULL")
    
    print("\n--- Max UUID ---")
    max_ = uuid.MAX
    print(f"UUID: {max_}")
    print(f"int 值：{max_.int}")  # 2^128 - 1
    print(f"用途：范围查询上界，排序终止标记")
    
    # Nil/MAX 实际用法示例
    print("\n--- 用法示例 ---")
    
    # 1. Nil 占位符检查
    user_id = uuid.NIL
    if user_id == uuid.NIL:
        print("✓ 用户尚未分配 ID")
    
    # 2. MAX 范围查询上界
    latest_uuid = uuid.uuid7()
    print(f"✓ 最新 UUID: {latest_uuid}")
    print(f"  查询条件：id <= {latest_uuid} AND id > start_time")
    print(f"  (MAX 可用于 WHERE created_at < MAX)")
    
    # 3. 排序示例 (NIL 最小，MAX 最大)
    sample_ids = [uuid.uuid4(), uuid.NIL, uuid.MAX, uuid.uuid4()]
    sorted_ids = sorted(sample_ids)
    print(f"\n--- 排序验证 ---")
    print(f"原始列表：{sample_ids[:2]} ... {sample_ids[2:]}")
    print(f"排序后第一个：{sorted_ids[0]} (应该是 NIL)")
    print(f"排序后最后一个：{sorted_ids[-1]} (应该是 MAX)")
    assert sorted_ids[0] == uuid.NIL, "NIL 应该最小"
    assert sorted_ids[-1] == uuid.MAX, "MAX 应该最大"
    print("✓ 排序正确")
    print()


# ============================================================================
# §5 UUID 属性访问
# ============================================================================

def demo_uuid_properties():
    """UUID 对象的所有常用属性
    
    UUID 对象是不可变的、可哈希的，可以用作字典键或集合成员。
    
    可用属性:
    - hex: 32 位十六进制字符串 (无分隔符)
    - bytes: 16 字节原始数据 (大端序)
    - bytes_le: 16 字节原始数据 (小端序，仅时间字段)
    - fields: 六个字段的元组
    - int: 128 位整数值
    - time: 时间字段 (仅 uuid1/uuid6/uuid7)
    - clock_seq: 时钟序列 (仅 uuid1/uuid6)
    - node: 节点地址 (仅 uuid1/uuid6)
    - variant: 变体标识符
    - version: 版本号 (1-8)
    - urn: URN 格式字符串
    """
    print("=" * 60)
    print("UUID 属性访问")
    print("=" * 60)
    
    u = uuid.uuid7()
    
    print(f"\n示例 UUID: {u}")
    print("-" * 60)
    
    properties = [
        ("version", u.version, "版本号 (1-8)"),
        ("variant", u.variant, "变体标识 ('rfc4122')"),
        ("hex", u.hex, "32 位十六进制字符串 (无分隔符)"),
        ("urn", u.urn, "URN 格式：'urn:uuid:...'",),
        ("int", f"{u.int} (十进制)", "128 位整数值"),
        ("bytes", f"{u.bytes.hex()} (hex)", "16 字节原始数据"),
        ("bytes_le", f"{u.bytes_le.hex()} (hex)", "小端序 16 字节"),
        ("fields", str(u.fields), "六个字段元组"),
    ]
    
    if u.version in (1, 6, 7):
        extra = [
            ("time", f"{u.time}", "时间戳 (单位依赖版本)"),
        ]
        properties.extend(extra)
        
    if u.version in (1, 6):
        extra2 = [
            ("clock_seq", u.clock_seq, "时钟序列"),
            ("node", f"{hex(u.node)}", "节点地址 (MAC)"),
        ]
        properties.extend(extra2)
    
    for prop_name, prop_value, description in properties:
        print(f"  {prop_name:12s} = {str(prop_value):40s}  # {description}")
    
    # 展示 fields 结构
    print("\n--- fields 详细结构 ---")
    tl, tm, thv, chv, cl, node = u.fields
    print(f"  time_low:              {tl:#x}")
    print(f"  time_mid:              {tm:#x}")
    print(f"  time_hi_version:       {thv:#x}")
    print(f"  clock_seq_hi_variant:  {chv:#x}")
    print(f"  clock_seq_low:         {cl:#x}")
    print(f"  node:                  {node:#x}")
    print()


# ============================================================================
# §6 最佳实践与决策指南
# ============================================================================

def recommend_usage():
    """UUID 选用建议
    
    通用决策流程:
    1. 是否需要稳定性/确定性？→ 是 → uuid5(namespace, name)
    2. 是否需要时间排序？→ 是 → uuid7()
    3. 只需要通用唯一 ID？→ uuid4()
    4. 需要审计/追踪时间？→ uuid1() (接受隐私代价) 或 uuid6()/uuid7()
    """
    print("=" * 60)
    print("UUID 选用指南")
    print("=" * 60)
    
    scenarios = [
        (
            "场景：通用/随机 ID\n"
            "推荐：uuid4()\n"
            "理由：无隐私风险;生成快;碰撞概率极低\n"
            "示例：会话 ID、令牌、临时标识\n",
            "uuid4()"
        ),
        (
            "场景：数据库主键 (插入顺序重要)\n"
            "推荐：uuid7()(3.14+)\n"
            "理由：时间有序;对 B+ 树索引友好;可并行生成\n"
            "示例：MySQL/PostgreSQL 主键\n",
            "uuid7()"
        ),
        (
            "场景：从名称/URL 派生稳定 ID\n"
            "推荐：uuid5(namespace, name)\n"
            "理由：确定性输出;跨平台一致;语义清晰\n"
            "示例：url.id、domain_id\n",
            "uuid5(NAMESPACE_DNS, 'example.com')"
        ),
        (
            "场景：需要审计/时间追踪\n"
            "推荐：uuid6() 或 uuid7()\n"
            "理由：两者都时间有序;uuid6 含 MAC(uuid1 兼容性);uuid7 更现代\n"
            "注意：避免 uuid1() 除非必须\n",
            "uuid7()"
        ),
        (
            "场景：遗留系统兼容\n"
            "推荐：uuid6()\n"
            "理由：保留 v1 的大部分结构;解决排序问题\n",
            "uuid6()"
        ),
    ]
    
    for scenario_text, code_snippet in scenarios:
        print(scenario_text)
        print(f"代码：{code_snippet}")
        print("-" * 60)
    
    print("\n重要提醒:")
    print("-" * 60)
    print("1. 没有 uuid2(): DCE Security 版本在 RFC 中有定义，")
    print("   但 uuid 模块故意未实现——调用会抛 AttributeError")
    print()
    print("2. 避免 uuid1(): 除非确实需要时间 + 节点信息并接受其泄露")
    print("   它会将 MAC 地址和时间戳暴露给外界")
    print()
    print("3. 优先 uuid5 而非 uuid3: SHA-1 比 MD5 更安全，虽然")
    print("   两者都非加密级别，但行业趋势如此")
    print()
    print("4. 3.14+ 强烈推荐 uuid7(): 现代首选，兼顾时间有序性和安全性")
    print()


# ============================================================================
# §7 综合示例
# ============================================================================

def main():
    """运行所有 UUID 演示"""
    print("\n" + "#" * 70)
    print("# Python uuid 模块完整演示 (Python 3.14)")
    print("#" * 70 + "\n")
    
    print("正在加载 uuid 模块...")
    print(f"uuid 模块路径：{uuid.__file__}")
    print()
    
    # 演示所有版本
    demo_uuid_v1()
    demo_uuid_v3()
    demo_uuid_v4()
    demo_uuid_v5()
    demo_uuid_v6()
    demo_uuid_v7()
    demo_uuid_v8()
    
    # 命名空间
    demo_namespaces()
    
    # 哨兵 UUID
    demo_sentinel_uuids()
    
    # 属性
    demo_uuid_properties()
    
    # 最佳实践
    recommend_usage()
    
    print("\n" + "=" * 70)
    print("所有演示完成!")
    print("=" * 70)


# ============================================================================
# 模块入口
# ============================================================================

if __name__ == "__main__":
    main()
