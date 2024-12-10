nums = open('numbers.txt', 'r')
result = 0
for line in nums:
  result += int(line)
print(result)