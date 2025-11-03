#https://www.codewars.com/kata/514a024011ea4fb54200004b/train/python
import re

def domain_name(url):
    splitted_url = re.split(r"[/.]", url)
    for p in splitted_url:
        if p and p not in ["https:", "http:", "www"]:
            return p
    return