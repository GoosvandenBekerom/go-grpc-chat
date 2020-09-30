let ws_url = "ws://localhost:9999/chat";
let messages = [];

function setupWebSocketConnection() {
  if (!("WebSocket" in window)) {
    alert("Sorry, your browser doesn't support WebSockets, so the app wont work.");
    return;
  }

  let socket = new WebSocket(ws_url)

  socket.onopen = function() {
    addAdminMessage("Succesfully opened WebSocket to " + ws_url)
    addAdminMessage("Happy chatting!")
  }

  socket.onmessage = function (e) { 
    var msg = e.data;
    messages.push(msg);
    console.log("Received WebSocket message: " + msg);
    renderMessages()
 };

  socket.onclose = function() {
    addAdminMessage("WebSocket closed")
  }
}

function renderMessages() {
  let chat = document.getElementById("chat")
  chat.innerHTML = '';
  
  for(let i = 0; i < messages.length; i++) {
    let msg = messages[i]
    let msgTime = document.createElement("span")
    msgTime.classList = ["time"]
    let hh = (""+msg.time.getHours()).padStart(2, "0")
    let mm = (""+msg.time.getMinutes()).padStart(2, "0")
    msgTime.innerHTML = hh + ":" + mm + " / ";

    let msgUser = document.createElement("span")
    msgUser.classList = ["user"]
    msgUser.innerHTML = msg.user

    let msgContent = document.createElement("span")
    msgContent.classList = ["content"]
    msgContent.innerHTML = msg.content
    
    let msgRoot = document.createElement("div")
    msgRoot.classList = ["message"]
    msgRoot.appendChild(msgTime)
    msgRoot.appendChild(msgUser)
    msgRoot.appendChild(msgContent)

    chat.appendChild(msgRoot)
  }
  chat.scrollTop = chat.scrollHeight;
}

function sendMessage() {
  let input = document.getElementById("new-message-content")
  messages.push({
    time: new Date(),
    user: "Goos",
    content: input.value
  })
  renderMessages()
  input.value = '';
}

function addAdminMessage(content) {
  messages.push({
    time: new Date(),
    user: "Admin",
    content: content
  });
  renderMessages()
}