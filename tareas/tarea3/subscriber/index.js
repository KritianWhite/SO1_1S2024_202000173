const { createClient } = require('redis');

(async () => {
    const client = createClient({
        socket: {
            host: '10.22.69.155',
            port: 6379,
        }
    });
    
    client.on('error', (error) => {
        console.error("Error on redis: ",error);
    });

    await client.connect();
    console.log("Redis connected");
    setInterval(async () => {
        const msg = JSON.stringify({msg: "Hello, everyone!"});
        console.log("Publishing message: ",msg);l
        try{
            const result = await client.publish("test: ",msg);
            console.log("Published message successfully: ",result);
        }catch(error){
            console.error("Error publishing message: ",error);
        }
    }, 3000);
})();