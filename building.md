
> Applications that are going to exchange messages over RabbitMQ need to establish a permanent connection to the message broker. When this connection is established, a channel needs to be created so that message-oriented interactions, such as publishing and consuming messages, can be performed.

### Exchanges

- **direct exchange:** which delivers messages to a single queue
- **topic exchange:** which delivers messages to multiple queues based on pattern-matching routing keys.

Creating connections is a costly operation, very much like it is with database connections. Typically, database connections are pooled, where each instance of the pool is used by a single execution thread. AMQP is different in the sense that a single connection can be used by many threads through many multiplexed channels.


A routing strategy determines which queue (or queues) the message will be routed to. The routing strategy bases its decision on a routing key (a free-form string) and potentially on message meta-information.

### The direct exchange
A direct exchange delivers messages to queues based on a message routing key. A message goes to the queue(s) whose bindings routine key matches the routing key of the message.

An example use case of direct exchange could be as follows:
1. The customer orders the taxi named taxi.1. An HTTP request is sent from the customer's mobile application to the Application Service.
2. The Application Service sends a message to RabbitMQ with a routing key, taxi.1. The message routing key matches the name of the queue, so the message ends up in the taxi.1 queue.

### Other types of routing exchanges
- Fanout: Messages are routed to all queues bound to the fanout exchange. 
- Topic: Wildcards must form a match between the routing key and the binding's specific routing pattern.
- Headers: Use the message header attributes for routing.