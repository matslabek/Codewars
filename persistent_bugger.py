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
