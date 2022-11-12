import client, {Connection, Channel, ConsumeMessage} from 'amqplib';
import nodemailer from 'nodemailer';
import hbs from 'nodemailer-express-handlebars';
import 'dotenv/config'

interface myPayload {
    email: string;
    userName: string;
    verifyToken: string;
}

const { RABBITMQ_URL, RABBITMQ_QUEUE } = process.env;

const consumer = (channel: Channel) => async  (msg: ConsumeMessage | null): Promise<void> => {

  if (msg) {
    let obj: myPayload = JSON.parse(msg.content.toString());
    await sendEmail(obj.email, obj.userName, obj.verifyToken);    
    channel.ack(msg)
  }
}

const sendEmail = async (email: string, userName: string, verifyToken: string) => {

    try {
        
        const GMAIL_USER = process.env.GMAIL_USER || "";
        const GMAIL_PASS = process.env.GMAIL_PASS || ""; 

        const transporter = nodemailer.createTransport({
            service: 'gmail',
            auth: {
                user: GMAIL_USER,
                pass: GMAIL_PASS
            }
        });

        var options = {
            viewEngine : {
                extname: '.hbs',
                defaultLayout: "",
                layoutsDir: './views/',
                partialsDir: './views/', 
            },
            viewPath: './views/',
            extName: '.hbs'
            };
        
        transporter.use('compile', hbs(options));

        let mailOptions = {
            from: 'Grupo 4 | Steamlike Platform ',
            to:  email, 
            subject: `Verificar cuenta: ${userName}`,
            text: 'Grupo 4 | SA-USAC | Fase 2 - Microservicios :D',
            template: 'index',
            context: {
                verifyURL: `http://localhost/verified-account/${userName}/${verifyToken}` 
            }
        };

        transporter.sendMail(mailOptions, (err: any, data: any) => {
            if (err) {
                console.log(err.message);
            }
            console.log('Email sent!!!');
        });
    }
    catch (e) {
        console.log(e);
    }

}

const delay = (ms: number) => {
    return new Promise( resolve => setTimeout(resolve, ms) );
}

const main = async () => {

    const RABBITMQ_URL_CONN: string = RABBITMQ_URL ? RABBITMQ_URL : "amqp://guest:guest@localhost:5672";

    let RABBITMQ_QUEUE_NAME: string = RABBITMQ_QUEUE ? RABBITMQ_QUEUE : 'test';

    const connection: Connection = await client.connect(RABBITMQ_URL_CONN)

    // Create a channel
    const channel: Channel = await connection.createChannel()

    // Makes the queue available to the client
    await channel.assertQueue(RABBITMQ_QUEUE_NAME)
    
    await delay(6000);
    // Start the consumer
    while (true) {
        console.log("leyendo queue")
        await channel.consume(RABBITMQ_QUEUE_NAME, consumer(channel))
    }
}

main()

