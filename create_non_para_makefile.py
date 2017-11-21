N_RULES = 10
SLEEP = 2

PATH = '/tmp/go-makes/outputfiles/'

print('all: rule0')
print('\tsleep', SLEEP)
print('\ttouch ' + PATH + 'all')
print('')

for i in range(N_RULES):
    name = 'rule' + str(i)
    dep = 'rule' + str(i+1)
    print(name + ':', dep)
    print('\tsleep', SLEEP)
    print('\ttouch', PATH+name, end='\n\n')

name = 'rule' + str(N_RULES)
print(name + ':')
print('\tsleep', SLEEP)
print('\ttouch', PATH+name)
