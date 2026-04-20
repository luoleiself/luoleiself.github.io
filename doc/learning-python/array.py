import array

'''
array 是可变序列, 元素类型统一

只支持以下类型代码:
| Type code | C Type                | Python Type       | Minimum size in bytes   | Notes |
| 'b'       | signed char           | int               | 1                       |       |
| 'B'       | unsigned char         | int               | 1                       |       |
| 'u'       | wchar_t               | Unicode character | 2                       | (1)   |
| 'w'       | Py_UCS4               | Unicode character | 4                       | (2)   |
| 'h'       | signed short          | int               | 2                       |       |
| 'H'       | unsigned short        | int               | 2                       |       |
| 'i'       | signed int            | int               | 2                       |       |
| 'I'       | unsigned int          | int               | 2                       |       |
| 'l'       | signed long           | int               | 4                       |       |
| 'L'       | unsigned long         | int               | 4                       |       |
| 'q'       | signed long long      | int               | 8                       |       |
| 'Q'       | unsigned long long    | int               | 8                       |       |
| 'f'       | float                 | float             | 4                       |       |
| 'd'       | double                | float             | 8                       |       |
'''

print('类型为: <class \'array.array\'>')

print('i 和 l, I 和 L 的取值范围相同')

# 有符号的字符数组, 取值范围 -128 到 127
signed_arr_b = array.array('b', [-128, 97, 98, 127])
print('有符号字符数组', signed_arr_b)
print(f'切片操作: {signed_arr_b[::-1]}')

# 无符号的字符数组, 取值范围 0 到 255
unsigned_arr_b = array.array('B', b'abc')
print('无符号字符数组', unsigned_arr_b)
print(f'切片操作: {unsigned_arr_b[::-1]}')

# 宽字符数组
wchar_arr = array.array('u', '中国')
print('宽字符数组', wchar_arr)
print(f'切片操作: {wchar_arr[::-1]}')

# 有符号的短整型数组, 取值范围 -32768 到 32767
signed_arr_h = array.array('h', [-32768, -32765, 32765, 32767])
print('有符号短整型数组', signed_arr_h)
print(f'切片操作: {signed_arr_h[::-1]}')

# 无符号的短整型数组, 取值范围 0 到 65535
unsigned_arr_h = array.array('H', [0, 65535])
print('无符号短整型数组', unsigned_arr_h)
print(f'切片操作: {unsigned_arr_h[::-1]}')

print('----------')
# 有符号的整型数组, 取值范围 -2147483648 到 2147483647
signed_arr_i = array.array(
    'i', [-2147483648, -2147483645, 2147483645, 2147483647])
print('有符号整型数组', signed_arr_i)
print(f'切片操作: {signed_arr_i[::-1]}')

# 无符号的整型数组, 取值范围 0 到 4294967295
unsigned_arr_i = array.array('I', [0, 4294967295])
print('无符号整型数组', unsigned_arr_i)
print(f'切片操作: {unsigned_arr_i[::-1]}')

# 有符号的长整型数组, 取值范围 -2147483648 到 2147483647
signed_arr_l = array.array('l', [-2147483648, 2147483647])
print('有符号长整型数组', signed_arr_l)
print(f'切片操作: {signed_arr_l[::-1]}')

# 无符号的长整型数组, 取值范围 0 到 4294967295
unsigned_arr_l = array.array('L', [0, 4294967295])
print('无符号长整型数组', unsigned_arr_l)
print(f'切片操作: {unsigned_arr_l[::-1]}')
print('----------')

# 有符号的长整型数组, 取值范围 -9223372036854775808 到 9223372036854775807
signed_arr_q = array.array(
    'q', [-9223372036854775808, -9223372036854775805, 9223372036854775805, 9223372036854775807])
print('有符号长整型数组', signed_arr_q)
print(f'切片操作: {signed_arr_q[::-1]}')

# 无符号的长整型数组, 取值范围 0 到 18446744073709551615
unsigned_arr_q = array.array('Q', [0, 18446744073709551615])
print('无符号长整型数组', unsigned_arr_q)
print(f'切片操作: {unsigned_arr_q[::-1]}')

# 浮点型数组
float_arr = array.array('f', [1.0, 2.0, 3.0])
print('浮点型数组', float_arr)
print(f'切片操作: {float_arr[::-1]}')

# 双精度型数组
double_arr = array.array('d', [1.0, 2.0, 3.0])
print('双精度型数组', double_arr)
print(f'切片操作: {double_arr[::-1]}')
