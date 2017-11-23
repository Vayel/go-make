N_RULES = 10
SLEEP = 1

print('all: rule0')
print('\tsleep', SLEEP)
print('\ttouch all')
print('')

for i in range(N_RULES):
    name = 'rule' + str(i)
    dep = 'rule' + str(i+1)
    print(name + ':', dep)
    print('\tsleep', SLEEP)
    print('\ttouch', name, end='\n\n')

name = 'rule' + str(N_RULES)
print(name + ':')
print('\tsleep', SLEEP)
print('\ttouch', name)
