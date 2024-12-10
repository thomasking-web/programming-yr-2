def increase_score(answer, current_score): 
  if answer == "Yes": 
    current_score += 1
  return current_score

c_score = 0
c_score = increase_score("Yes", c_score)
print(c_score) 