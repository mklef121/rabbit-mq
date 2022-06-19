### Taxi Application

**CC** is a new taxi agency with huge potential. Today, the company has just two taxi drivers and two developers, but they want to expand a lot in the upcoming year. **CC** has already built a website in Go and started out with a backend, also written in Go, that stores CC trips in a database. CC also has some scripts, written in Javascript, that generate route reports.

So far, CC's system runs as follows:

- The company's website and blog runs on Go.
- The Rich Internet Application that stores route data, such as the starting point and the endpoint of the trip, is written in Go.
- There's a back-office that sends route updates to drivers and is written in Go.
- Multiple ad hoc Javascript scripts are used to extract and message data to generate route reports.
Taxi applications are written in Javascript.

The main reason is because of a new
feature CC wants to offer to their customers – they want to build a taxi application that handles reservations on the go. CC also wants to be able to scale painlessly. The plan is to build an app where users can book a car via a smartphone, receive booking confirmation, and view the car approaching the starting point of the trip.

Since CC already has some services in different languages, and since CC wants to be able to scale easily, they decided to use a ready-made message-oriented middleware such as RabbitMQ for asynchronous communication between the app, the client, and the backend.
