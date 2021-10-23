def order_weight(strng):
    """Empty string case"""
    if not strng:
        return ""
    string_weights = strng.split()
    sum_of_digits_list = []
    digits_dict = {}
    for string_number in string_weights:
        """Calculating digit sum for every number"""
        digits_sum = sum([int(digit) for digit in string_number])
        sum_of_digits_list.append(digits_sum)
        """Dictionary handling duplicate sums, has structure of sum_of_digits: [no1, no2]"""
        if digits_sum not in digits_dict:
            digits_dict[digits_sum] = []
            digits_dict[digits_sum].append(string_number)
        else:
            no_of_items = len(digits_dict[digits_sum])
            """When a number has each sum, it's appended to the dictionary in sorted way"""
            for index in range(no_of_items):
                if string_number <= digits_dict[digits_sum][index]:
                    digits_dict[digits_sum].insert(index, string_number)
                    break
                elif index == no_of_items - 1:
                    digits_dict[digits_sum].append(string_number)
    sum_of_digits_list.sort()
    final_sorted_list = []
    checked_sums = []
    for digit_sum in sum_of_digits_list:
        if digit_sum not in checked_sums:
            for digits in digits_dict[digit_sum]:
                final_sorted_list.append(digits)
        checked_sums.append(digit_sum)
    return " ".join(final_sorted_list)
