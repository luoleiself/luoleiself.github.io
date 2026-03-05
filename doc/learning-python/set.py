print('intersection_update()')
s1 = {1, 2}
s2 = {2, 3}
print(f"s1: {s1}, s2: {s2}")
s1.intersection_update(s2)  # 更新原集合, 只保留交集
print(f"s1: {s1}")
print(f"s2: {s2}")
print('-----------')

print('difference_update()')
s3 = {1, 2}
s4 = {2, 3}
print(f"s3: {s3}, s4: {s4}")
s3.difference_update(s4)    # 更新原集合, 只保留差集
print(f"s3: {s3}")
print(f"s4: {s4}")
print('-----------')

print('symmetric_difference_update()')
s5 = {1, 2}
s6 = {2, 3}
print(f"s5: {s5}, s6: {s6}")
s5.symmetric_difference_update(s6)  # 更新原集合, 只保留异或集
print(f"s5: {s5}")
print(f"s6: {s6}")
print('-----------')
