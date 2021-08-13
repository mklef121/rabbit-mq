import express from 'express'
import amqplib, { Channel, Connection } from 'amqplib'

let connection: Connection;
async function rabbitConnection(){
    
    if(connection) return connection;
     connection = await amqplib.connect("amqp://miracool-dev:password@localhost/dev-vhost");

     return connection;
}

let channel: Channel
/**
 * 
 * @param isNew {boolean} indicates if a new connection should be returned or reuse an existing one
 * @returns 
 */
async function rabbitChannel(isNew = false){
    if(channel && !isNew) return channel;
    
    channel = await (await rabbitConnection()).createChannel();

    return channel;
}

async function setupQueue(){
    const queueRep = "taxi.1"
    const exchangeRep = 'taxi-direct'
    const raCh = await rabbitChannel();
    const queue = await raCh.assertQueue(queueRep,{durable: true});

    let exchange = await raCh.assertExchange(exchangeRep, 'direct',{durable: true});
    //Bind queue to the particular exchange topic
    await raCh.bindQueue(queue.queue, exchangeRep, "taxi.1");

    return exchange
}

setupQueue()

// const channel = rabbitMqInitialize();

const app = express()
const port = 5000

app.get('/', async (_, res) => {
  
  res.status(200).send()
})

console.log("His");


app.listen(port, () => console.log(`Running on port ${port}`))