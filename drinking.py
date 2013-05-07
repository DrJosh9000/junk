#!/usr/bin/env python

import random

print("""           D R I N K I N G

A super minimalistic game by Josh Deprez.

""")

bac = 0.0
drinks = 0
menu = [
	'Atom Bomb',
	'shot of Angostura bitters, neat',
	'Sazerac',
	'Negroni',
	'Margerita',
	'Tommy\'s Margerita',
	'Blood and Sand',
	'Vieux Carre',
	'cognac',
	'house red',
	'house white',
	'bubbly',
	'Blue Moon',
	'VB',
	'Stella',
	'Manhattan',
	'Bowmore single malt, with a single ice cube',
	'Monkey Shoulder triple malt, on the rocks',
	'Screwdriver',
	'Jellyfish shot',
	'Jelly Donut shot',
	'Wet Pussy shot',
	'Sex on the Beach',
	'Cosmopolitan',
	'666 Vodka shot',
	'Juicy Lucy',
	'shot of spiced rum',
]

print("""You are seated at a bar. The bartender, dressed in a white shirt and black bow tie, eyes you thoughtfully, rubbing the polished mahogany of the bar in slow strokes with a white towel.

""")

while bac < 0.3:
	print("""Your blood alcohol level is currently %.2f. You have had %d drinks in total.
	
Possible actions: DRINK, DRINK and DRINK.""" % (bac, drinks))
	
	i = raw_input("What do you do? ")
	
	if i.upper() == "DRINK":
		print("""You order a %s. Within moments the bartender serves it to you, and in seemingly less time, you have consumed it.\n""" % (random.choice(menu)))
		bac += 0.02
		drinks += 1
	else:
		print("""I'm sorry Dave, but I can't let you do that!\n""")
	
print("""You attempt to get off the bar stool. Some part of your brain recognises a sensation of falling, which is somehow relieving instead of terrifying. You realise you are now lying on your back. Something feels wrong. The booze is trying to escape your mouth, so you choke it back. Another part of your brain tells you choking is probably the wrong thing to do. Then you black out.

You are dead!""")