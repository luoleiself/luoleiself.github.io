import sqlite3

conn = sqlite3.connect('database.db')
# 获取游标
cursor = conn.cursor()

# 查询
r1 = cursor.execute('SELECT * FROM users')
print(f'r1.arraysize {r1.arraysize} r1.lastrowid {r1.lastrowid} r1.rowcount {r1.rowcount}')
result = r1.fetchall()
print('result', result)
print('-' * 10)

# 清空表
r4 = cursor.execute('DELETE FROM users')
print(f'r4.arraysize {r4.arraysize} r4.lastrowid {r4.lastrowid} r4.rowcount {r4.rowcount}')
print('-' * 10)

# 批量插入
users = (('zhangsan', '123'), ('lisi', '123'), ('wangwu', '123'), ('zhaoliu', '123'))
r2 = cursor.executemany(
    'insert into users(username, password) values (?, ?)', users)
conn.commit()   # insert 隐式创建事务, commit 提交事务
print(f'r2.arraysize {r2.arraysize} r2.lastrowid {r2.lastrowid} r2.rowcount {r2.rowcount}')
print('-' * 10)

# 查询
r3 = cursor.execute('select * from users')
print(f'r3.arraysize {r3.arraysize} r3.lastrowid {r3.lastrowid} r3.rowcount {r3.rowcount}')
print('r3.fetchall()', r3.fetchall())
print('-' * 10)

# fetchmany
r5 = cursor.execute('select * from users')
result = r5.fetchmany(2)
print(f'r5.arraysize {r5.arraysize} r5.lastrowid {r5.lastrowid} r5.rowcount {r5.rowcount}')
print(f'r5.fetchmany(2) {result}')
conn.close()