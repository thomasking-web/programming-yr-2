passwords = [
  [],
  [],
  []
]

passwords[0].append("Raspbian")
# Add username
passwords[1].append("pi")
# Add password
passwords[2].append("raspberry")

def add_account():
  passwords[0].append(input("Enter the name of the service: "))
  passwords[1].append(input("Enter the username: "))
  passwords[2].append(input("Enter the password: "))

def view_password():
  account = input("Enter the name of the service you want to view the password for: ")
  if account in passwords[0]:
    index = passwords[0].index(account)
    print(f"Username: {passwords[1][index]}")
    print(f"Password: {passwords[2][index]}")
  else:
    print("Account not found")

def main_menu():
  while True:
    print("Menu:")
    print("1. Add a new account")
    print("2. View password for an account")
    print("3. Exit")
    choice = input("Enter your choice: ")
    if choice == "1":
      add_account()
    elif choice == "2":
      view_password()
    elif choice == "3":
      # Kill the program
      exit()
    else:
      print("Invalid choice")

main_menu()