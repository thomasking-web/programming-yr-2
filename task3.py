def change_colour(current_col):
  print(f"Your current colour is {current_col}")
  print("What would you like to change it to?")
  new_col = input()
  current_col = new_col
  print(f"Your new colour is {current_col}")
  return current_col

c_col = "Red"
c_col = change_colour(c_col)