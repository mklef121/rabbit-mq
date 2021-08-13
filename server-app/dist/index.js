"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const amqplib_1 = __importDefault(require("amqplib"));
let connection;
function rabbitConnection() {
    return __awaiter(this, void 0, void 0, function* () {
        if (connection)
            return connection;
        connection = yield amqplib_1.default.connect("amqp://miracool-dev:password@localhost/dev-vhost");
        return connection;
    });
}
let channel;
/**
 *
 * @param isNew {boolean} indicates if a new connection should be returned or reuse an existing one
 * @returns
 */
function rabbitChannel(isNew = false) {
    return __awaiter(this, void 0, void 0, function* () {
        if (channel && !isNew)
            return channel;
        channel = yield (yield rabbitConnection()).createChannel();
        return channel;
    });
}
function setupQueue() {
    return __awaiter(this, void 0, void 0, function* () {
        const queueRep = "taxi.1";
        const exchangeRep = 'taxi-direct';
        const raCh = yield rabbitChannel();
        const queue = yield raCh.assertQueue(queueRep, { durable: true });
        let exchange = yield raCh.assertExchange(exchangeRep, 'direct', { durable: true });
        //Bind queue to the particular exchange topic
        yield raCh.bindQueue(queue.queue, exchangeRep, "taxi.1");
        return exchange;
    });
}
setupQueue();
// const channel = rabbitMqInitialize();
const app = express_1.default();
const port = 5000;
app.get('/', (_, res) => __awaiter(void 0, void 0, void 0, function* () {
    res.status(200).send();
}));
console.log("His");
app.listen(port, () => console.log(`Running on port ${port}`));
//# sourceMappingURL=index.js.map