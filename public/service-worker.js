console.log(self,"test");

var wsuri = "ws://" + self.location.host + "/socket";
sock = new WebSocket(wsuri);

self.addEventListener('activate', function(event) {
    console.log('activate');     

    sock.onopen = function() {
        console.log("connected to " + wsuri);
    }

    sock.onmessage = function(e) {
        console.log("message received: " + e.data);
    }
});

self.addEventListener('push', function(event) {
    console.log('push', event.data.text());    
    sock.send(event.data.text()); 
});   

self.addEventListener('fetch', function(event) {
    console.log('fetch', event.data.text());   
    event.respondWith(e => e
        // magic goes here
    );
});