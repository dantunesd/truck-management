# Golang Backend Challenge

First of all, we’d like to thank you for choosing Loadsmart. As a next step in the recruitment process, we have an exercise for you.

This is an opportunity to demonstrate your front end skills. Just be clear and do it to the best of your ability - the faster you complete it, of course, the quicker we can move forward.

## Context
We are a company that helps carriers automate their workflow while moving loads. This is super important, because most of the carriers are small and have a lot of bookkeeping tasks to do. Such as updating locations of their trucks on the road, calculating how many hours a driver drove a truck, anticipating any problems with cargos they are hauling, etc.

In the United States, the Department of Transportation (DOT) requires that all trucks on the road to operate with an Electronic Logging Device (ELD) turned on. Most of the manufacturers of this device provide active APIs to update each truck's location. All of the carrier's trucks run with an ELD device.

Since most of these carriers are small, they run on very tight operations, with small margins and not a lot of employees to do all of these operational tasks. The problem we want to solve is at the core of this challenge: create an API to allow external parties (ELD providers)  to update our truck location for us.

## How we will evaluate this challenge
The idea of this challenge is to understand and gauge where you are professionally as a Golang developer.

We will evaluate the following characteristics in this challenge:

### Requirements interpretation
- Correctness
- Unit Testing
- Idiomatic Writing
- Completeness of all tasks
- Design Patterns in writing common code (database calls, etc)
- Capacity for abstraction and high level reasoning
- Organization
- Reproducibility (we need to be able to run this test ourselves, after you submit it to us)
- Regarding the API: documentation, reliability and scalability

Other than the directions of each proposal, you are free to choose formats and schemas for each endpoint as you prefer. Bear in mind that this system needs to store state between calls, otherwise it’s not possible to complete all tasks.
We need to be able to run your solution when we evaluate your project, so provide up-to-date instructions for installing / compiling and running the API binary.
## Truck API
We need an API that can solve some issues for us. First and foremost we need to create an API that allows carriers - that are using our SaaS to manage their trucks. There are several operations the carriers can do to their trucks, but basically they are:

- Create a new truck
- Update details of a truck
- Delete a truck

Keep in mind that this is a Software as a Service product, so we are in fact, managing trucks for several different carriers inside our databases.

Trucks can be uniquely identified by their license plate, which is composed of 7 alphanumeric characters (ABC1234, A1B1234, etc) and have the following interesting properties (which may or not be relevant for your challenge)

- Truck Type (Dryvan, Reefer)
- Size (in feet - from 20 to 53)
- Color (blue, red, green, yellow)
- Make
- Model
- Year of manufacture
- ELD ID

## Location API
Not only our customers are interested in managing their trucks and fleet, they are also interested in receiving the locations from the ELD manufacturers to update their current position. Our system later on will use these updates to trigger all sorts of different side effects, so this integration with ELD is pretty important

This API needs to allow external parties to our company to post to our system the current location of a truck. The ELDs can provide the following information (that may or may not be relevant for your challenge)

- ELD ID
- Engine State (On/Off)
- Current Speed
- Current Bearing
- Latitude
- Longitude
- Engine hours
- Odometer

Not only do we need a way for external parties to post the locations for our customers' trucks, we need an easy way to fetch the latest information about each truck. We need a way to get the last known position for any truck we have.

## Bonus Round (optional)
Our SaaS service provides reports to our customers. One of the most important reports we have is the trip report. The trip report allows us to output to the user summaries of each trip, considering the state of the engine that came from the ELD.

Ideally this endpoint should receive the truck id we are looking for and return to us the summary of the last trip. If the trip is ongoing, the report should state so. If there are no known trips, the report should output a different value.

This report should contain:

- Origin (trip starting point)
- Destination (trip finishing point or last known point for ongoing trips)
- State (ongoing, finished)
- Odometer delta (how many miles this truck moved , calculated by the odometer from start to finish)
- Engine hours delta (how many engine hours were used, calculated by the ELD)
- Average Speed

## When finished

- Send an email to the _recruiter_ notifying you are done
- Please label the email subject line (Your Full Name - Loadsmart Go Back End Test)
- Please, do not post your solution to any kind of public repository

