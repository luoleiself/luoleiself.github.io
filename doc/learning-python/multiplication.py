print('-----9 * 9-----')
print('-----while-----')
i = 1
while i <= 9:
    j = 1
    while j <= i:
        print(f'{j} * {i} = {j * i}', end='\t')
        j += 1
    print()
    i += 1

print('---for-range---')
for i in range(1, 10):
    for j in range(1, i + 1):
        print(f'{j} * {i} = {j * i}', end='\t')
    print()
