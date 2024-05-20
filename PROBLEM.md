## Trading System Leaderboard

### Problem Statement

You are a software engineer at a high-frequency trading firm. We have a robust trading system in place, and we want to develop a real-time leaderboard to showcase the top 10 traders based on their daily trading volume for a specific symbol. This leaderboard will be displayed at the landing page of an upcoming customer acquisition campaign by the internal Business Development team.

### Your task:
Write an efficient and scalable program that dynamically updates the leader-board with the latest ranking every minute, ensuring accuracy and performance even under high trading activity.

### Assumptions

- The trading system provides an API or data stream to access real-time trade data.
- Trade data includes fields such as trader ID, symbol, timestamp, and volume (quantity).
- You don’t need to consider trade direction (buy/sell) for this assignment.
- The leaderboard displays only the top 10 traders for the specified symbol.
- The symbol can be easily configured or passed as an argument to your program.
- Focus on the core functionality of the leaderboard. You don’t need to consider data persistence at this stage.

### Rules of the Game

- Use any programming language of your choice.
- Prioritize verlocity over scalability. Typically, we will have many such requirements from BD team in a fast-paced environment.
- Consider potential edge cases and handle them appropriately.
- Feel free to use libraries or frameworks that align with the language you choose and enhance your solution.
- Add a readme with clear instructions on how to run the finished solution.
- Be prepared to discuss your thought process, design decisions, and trade- offs during the interview.

### Bonus Points

- Implement few basic tests to demonstrate the correctness of your code.
- Design your code to be easily extendable to accommodate additional fea- tures or symbols.
- If you can roll up a rudimentary UI that consumes your leaderboard APIs, that would be great.