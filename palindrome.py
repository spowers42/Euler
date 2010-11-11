#!/usr/bin/python
import math

def is_palindrome(sequence):
	sequence = str(sequence)
	half = math.trunc(len(sequence)/2)
	if half>0:
		for x in range(0, half):
			if sequence[x]==sequence[(-1-x)]:
				pass
			else:
				return False
	return True
