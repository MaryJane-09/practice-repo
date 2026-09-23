# Mutability
print("Mutability")
lst = [5,6,7,]
updated = lst
updated.append(8)
print(updated) # output = [5, 6, 7, 8]
print()

# Value vs reference — struct/array
print("Value vs reference — struct/array")
tpl = (3, 6, 9)
tpl[0] = 0
print(tpl) # output = TypeError: 'tuple' object does not support item assignment
print()

# Pointers
print("Pointers")
x = [4]
print(id(x))
y = x
y.append(5)
print(id(y))
print(x) 