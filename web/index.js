self.postMessage('Cool Dream');
self.postMessage({name: 'huiYe', age: 26});
self.onmessage=(res)=>{
    console.log(res.data)
}
self.close();