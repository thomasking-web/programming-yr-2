suits = ["♥", "♦", "♣", "♠"]
deck = []

def sort_deck():
  for suit in suits:
    for i in range(1, 14):
      # 0 = Ace, 11 = Jack, 12 = Queen, 13 = King
      if i == 1:
        deck.append("A" + suit)
      elif i == 11:
        deck.append("J" + suit)
      elif i == 12:
        deck.append("Q" + suit) 
      elif i == 13:
        deck.append("K" + suit)
      else:
        deck.append(str(i) + suit)

  return deck