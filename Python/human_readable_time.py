#https://www.codewars.com/kata/52685f7382004e774f0001f7/train/python
def make_readable(seconds):
    hours = seconds // 3600
    minutes = seconds // 60 - (hours * 60)
    secs = seconds - (3600 * hours) - (60 * minutes)
    time_string = ""
    for item in [hours, minutes, secs]:
        if item < 10:
            time_string += "0"
        time_string += str(item) + ":"
    return(time_string[:-1])
