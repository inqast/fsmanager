# Family Subscription Manager telegram bot

### Setup
- create config.yml from example
- run `docker-compose up --build`

### Example
@fsmanager

### Usage

Domains: /me /sub /help

	/me commands:
		/me subs
			shows list of user subscriptions and its statuses
			no args

	/sub commands:
		/sub add
			adds new subscription, sets unpaid
			args: 
				name - name of service(string with no spaces)
				cost - monthly paid price(positive int)
				cap -  capacity (positive int)
				payday - positive int not grater then 31
			usage pattern: /sub add name=my_awesome_sub cost=500 cap=5 payday=15

		/sub status
			shows status of the subscription, executed only by members
			args:
				id - id of subscription(implicit)
			usage pattern: /sub status 5

		/sub edit
			edits the subscription, executed only by owner
			args:
				id - id of subscription(implicit)
				name - name of service(string with no spaces)(optional)
				cost - monthly paid price(positive int)(optional)
				cap -  capacity (positive int)(optional)
				payday - positive int not grater then 31(optional)
			usage pattern: /sub edit 5 name=new_awesome_name

		/sub drop
			drops the subscription, executed only by owner
			args:
				id - id of subscription(implicit)
			usage pattern: /sub drop 5

		/sub join
			enter to subscription membership
			args:
				id - id of subscription(implicit)
			usage pattern: /sub join 5

		/sub leave
			leave subscription membership, executed only by members
			args:
				id - id of subscription(implicit)
			usage pattern: /sub leave 5

		/sub pay
			marks that member is paid his share, executed only by members
			args:
				id - id of subscription(implicit)
			usage pattern: /sub pay 5

		/sub reset
			resets payment statuses of members, executed only by owner
			args:
				id - id of subscription(implicit)
			usage pattern: /sub reset 5

	/help commands:
		just call this domain to get manual