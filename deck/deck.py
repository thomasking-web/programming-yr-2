from random import shuffle
from sort_deck import sort_deck

deck = sort_deck()

shuffled = deck.copy()
shuffle(shuffled)

player1 = shuffled[:26]
player2 = shuffled[26:]

print(player1)
print(player2)