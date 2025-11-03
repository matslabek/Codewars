#https://www.codewars.com/kata/550554fd08b86f84fe000a58/train/python
def in_array(array1, array2):
    r = []
    for x in array1:
        for y in array2:
            if x in y and x not in r:
                r.append(x)
    r.sort() 
    return r
