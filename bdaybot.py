#!/usr/bin/env python

import twitter
import datetime
import os
import time

LOG_DIR = '/home/jdeprez/lol/'
os.environ['TZ'] = 'Australia/Melbourne'
time.tzset()

bday_db = {
	(11,10): ['@DrJosh9000'],
}

api = twitter.Api(access_token_key='Get your own token key!', 
				  access_token_secret='Get your own secret!', 
				  consumer_key='Get your own consumer key!', 
				  consumer_secret='Get your own secret!')

now = datetime.datetime.now()

pal = bday_db.get((now.day,now.month))
if pal:
	message = ''
	if isinstance(pal,list):
		message = 'Happy birthday to %s!' % (', '.join(pal[0:-1]) + ' and ' + pal[-1])
	else:
		message = 'Happy birthday, %s!' % pal
	api.PostUpdate(message)
	with open(LOG_DIR + 'bdaybot.log', 'a+') as f:
		f.write('%s Posted message: %s\n' % (datetime.datetime.now(), message))
else:
	with open(LOG_DIR + 'bdaybot.log', 'a+') as f:
		f.write('%s No message.\n' % datetime.datetime.now())	


#api.PostUpdate('Well, this could be the best or worst thing ever.')

