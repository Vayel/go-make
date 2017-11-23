N_CHILDREN = 14
SLEEP = 1

deps = ['rule' + str(i) for i in range(N_CHILDREN)]
print('all:', ' '.join(deps))
print('\tsleep', SLEEP)
print('\ttouch all')

for dep in deps:
    print(dep + ':')
    print('\tsleep', SLEEP)
    print('\ttouch', dep)
