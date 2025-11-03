#https://www.codewars.com/kata/55bf01e5a717a0d57e0000ec/train/python
def persistence(n):
    persistence = 0
    while True:
        digits = [char for char in str(n)];
        if len(digits) == 1:
            return persistence
        else:
            product = 1
            for dig in digits:
                product = product * int(dig)
            n = product
            persistence += 1
