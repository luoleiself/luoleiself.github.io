import sqlite3

db = sqlite3.connect('./database.db')

cursor = db.cursor()
print('cursor', cursor)

r1 = cursor.execute('SELECT * FROM users')
print('r1.arraysize', r1.arraysize, r1.lastrowid, r1.description, r1.rowcount)
result = r1.fetchall()
print('result', result)
cursor.close()

# users = (('zhangsan', '123'), ('lisi', '123'), ('wangwu', '123'), ('zhaoliu', '123'))
# cursor = db.cursor()
# r2 = cursor.executemany(
#     'insert into users (username, password) values (?, ?)', users)
# print('r2.arraysize', r2.arraysize, r2.lastrowid, r2.description, r2.rowcount)
# cursor.close()

cursor = db.cursor()
r3 = cursor.execute('select * from users')
print('r3.arraysize', r3.arraysize, r3.lastrowid, r3.description, r3.rowcount)
cursor.close()
