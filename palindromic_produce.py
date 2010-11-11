#!/usr/bin/python
''' Project Euler problem 4
find the largest palindromic produce to two 3 digit numbers
'''

import palindrome as pal

palindromes = []

a = range(100, 1000, 1)
for n in a:
	for m in a:
		p=n*m
		if pal.is_palindrome(p):
			palindromes.append(p)
palindromes.sort()
print palindromes.pop()
