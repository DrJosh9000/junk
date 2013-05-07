#!/usr/bin/env python

# Best Game Roster
# by Josh Deprez
# May 7, 2013
#
# N teams are playing a 3-way game.
# What is a roster of N games that evenly spreads the number of
# pairings that appear?
#
# This is a brute-force solution.
# It was done to get a job done, not to demonstrate good coding skills.
#

import itertools
import sys

games = 6
teams = range(games)

combos = list(itertools.combinations(teams,3))
pairs = list(itertools.combinations(teams,2))
meantc = 3


teamstemp = {}
for i in teams:
	teamstemp[i] = 0

pairstemp = {}
for pair in pairs:
	pairstemp[pair] = 0

besttc = sys.maxint
bestpc = sys.maxint
bestroster = []
for roster in itertools.combinations(combos,games):
	teamcount = dict(teamstemp)
	paircount = dict(pairstemp)
	for game in roster:
		for i in game:
			teamcount[i] += 1
		for pair in pairs:
			if pair[0] in game and pair[1] in game:
				paircount[pair] += 1
	tc = sum((meantc-x)*(meantc-x) for x in teamcount.values())
	pc = sum((1-x)*(1-x) for x in paircount.values())
	if tc < besttc or (tc == besttc and pc < bestpc):
		besttc = tc
		bestpc = pc
		bestroster = roster
	
print bestroster
print besttc
print bestpc
