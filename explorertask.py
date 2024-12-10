def change_direction(direction):
  print(f"Current direction: {direction}")
  print("Enter new direction")
  direction = input()
  return direction

def change_speed(speed):
  print(f"Current speed: {speed}")
  print("Enter new speed")
  speed = input()
  return speed

def_speed = 10
def_speed = change_speed(def_speed)
def_direction = "North"
def_direction = change_direction(def_direction)

print(f"New direction: {def_direction}")
print(f"New speed: {def_speed}")