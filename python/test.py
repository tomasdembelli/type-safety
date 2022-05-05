def add_two_things(a, b):
    c = a + b
    print(c)
    print(type(c))

add_two_things(1, 2) # int + int: expect int 3
add_two_things(1, 2.3) # int + float: expect float 3.3
add_two_things("a", "b") # str + str: expect string 'ab'

add_two_things(1, "2") # expect TypeError: unsupported operand type(s) for +: 'int' and 'str'
