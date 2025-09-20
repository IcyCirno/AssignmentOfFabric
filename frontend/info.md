{
	code:
	error:
}

Group("/api")
{
	user := c.Group("/user")
	{
		/register
			{
				name
				email
				password
			}
			{
				
			}
		/login
			{
				email
				password
			}
			{
				token
			}
		/profile
			{

			}
			{
				name
				email
				gocoin
			}
	}

	card "/card"
	{
		/mint
			{
				name
				avatar
				profile
			}
			{
				name
				profile
				avatar
				hashid
				rarity
				cost
				attack
				blood
			}
		/destroy
			{
				hashid
			}
			{

			}
		/query
			{

			}
			{
				data: {},{}
			}
		/sell
			{
				hashid
				cost
			}
			{

			}
		/cancel
			{
				hashid
			}
			{

			}
	}

	market "/market"
	{
		/query
			{
				
			}
			{
				data:{},{}
			}
		/buy
			{
				hashid
			}
			{

			}
	}
}