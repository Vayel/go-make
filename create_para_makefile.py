N_CHILDREN = 40
SLEEP = 2

deps = ['rule' + str(i) for i in range(N_CHILDREN)]
print('all:', ' '.join(deps))
print('\tsleep', SLEEP)
print('\ttouch all')

for dep in deps:
    print(dep + ':')
    print('\tsleep', SLEEP)
    print('\ttouch', dep)
