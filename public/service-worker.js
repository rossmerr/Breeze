console.log(self,"test");

var evtSource = new EventSource("/stream");

self.addEventListener('activate', function(event) {
    console.log('activate');     

 
});

self.addEventListener('fetch', function(event){
    console.log("fetch" , event.request);
});

evtSource.onmessage = function(e) {
    console.log("message ", e);
  }