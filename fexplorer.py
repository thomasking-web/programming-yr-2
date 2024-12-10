# Further explorer task

# Write a password checker using a global variable in a function for a peer to try and convert.
# Global variable for the password
password = "securepassword"

def check_password(input_password):
  global password
  if input_password == password:
    return True
  else:
    return False

# Example usage
user_input = input("Enter the password: ")
if check_password(user_input):
  print("Access granted")
else:
  print("Access denied")