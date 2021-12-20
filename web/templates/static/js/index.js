function addMessage(author, time, text) {
    var messageList = document.getElementById("messages-list");

    var message = document.createElement("li");
    message.classList.add("message");

    var authorTag = document.createElement("p");
    authorTag.classList.add("message-author");
    var authorText = document.createTextNode(author);
    authorTag.appendChild(authorText);

    var createdAtTag = document.createElement('p');
    createdAtTag.classList.add("message-created-at");
    var createdAtText = document.createTextNode(time);
    createdAtTag.appendChild(createdAtText);

    var textTag = document.createElement('p');
    textTag.classList.add("message-text");
    var  textText = document.createTextNode(text);
    textTag.appendChild(textText);

    message.appendChild(authorTag);
    message.appendChild(createdAtTag);
    message.appendChild(textTag);

    messageList.appendChild(message);
    messageList.parentElement.scrollTop = messageList.parentElement.scrollHeight
}

function flushMessages() {
    var messageList = document.getElementById("messages-list");
    while( messageList.firstChild ){
        messageList.removeChild( messageList.firstChild );
    }
}

function sendMessage(event) {
    event.preventDefault();
    var roomId = parseInt(localStorage.getItem("active_chat"))
    if (roomId === null || sub === null) {
        alert("Choose the chat or create new")
        return
    }
    var messageInput = document.getElementById("message-input");
    sub.publish(messageInput.value)
    messageInput.value = ""
}

function sendMessageIfEnterPressed(event) {
    if (event.keyCode === 13) {
        sendMessage(event);
    }
}


function refreshSendButton(roomId) {
    localStorage.setItem("active_chat", roomId)
}

function refreshAndDisabledButtons(roomId) {
    var oldRoom = localStorage.getItem("active_chat");
    if (oldRoom !== "0") {
        var oldButton = document.getElementById("connect_to_room" + oldRoom)
        oldButton.disabled = false
    }
    var newRoomButton = document.getElementById("connect_to_room" + roomId)
    newRoomButton.disabled = true
}

function createRoomConnectFunction(roomId) {
    return function (event) {
        if (sub !== null) {
            sub.unsubscribe();
        }
        flushMessages();
        event.preventDefault();
        refreshAndDisabledButtons(roomId);
        sub = centrifuge.subscribe(roomId.toString(), receiveMessage);
        centrifuge.connect();
        refreshSendButton(roomId);
    }
}


function addRoomItem(text, id, dissabled = false) {
    var roomList = document.getElementById("room-list");

    var room = document.createElement("form");
    room.classList.add("room-list-item");

    var button = document.createElement("button");
    var buttonText = document.createTextNode(text);
    button.appendChild(buttonText);
    button.id = "connect_to_room" + id
    button.addEventListener("click", createRoomConnectFunction(id))
    button.disabled = dissabled

    room.appendChild(button);
    roomList.appendChild(room);
}

function flushRooms() {
    var http = new XMLHttpRequest();
    var url = "/room";

    var roomList = document.getElementById("room-list");
    while( roomList.firstChild ){
        roomList.removeChild( roomList.firstChild );
    }

    http.open("GET", url, true);
    http.setRequestHeader("Content-Type", "application/json");
    http.setRequestHeader("Authorization", "Bearer " + localStorage.getItem("access_token"));

    http.onreadystatechange = function () {
        if (http.readyState === XMLHttpRequest.DONE) {
            switch (http.status) {
                case 200:
                    var rooms = JSON.parse(http.responseText);
                    for (let room of rooms){
                        addRoomItem(room["Title"], room["ID"])
                    }
                    break;
                default:
                    alert(http.responseText);
            }
        }
    }

    http.send();
}

function setTokens(access_token, refresh_token) {
    localStorage.setItem("access_token", access_token);
    localStorage.setItem("refresh_token", refresh_token);
    console.log(document.cookie)
    document.cookie = "access_token=" + access_token
    document.cookie = "refresh_token=" + refresh_token
    console.log(document.cookie)
    var registrationWindow = document.getElementById("registration");
    registrationWindow.classList.add("hidden");
    flushRooms();
}

function signIn(event) {
    event.preventDefault();

    var usernameInput = document.getElementById("registration-form-username");
    var passwordInput = document.getElementById("registration-form-password");

    var http = new XMLHttpRequest();
    var url = "/sign_in";
    http.open("POST", url, true);
    http.setRequestHeader("Content-Type", "application/json");

    http.onreadystatechange = function () {
        if (http.readyState === XMLHttpRequest.DONE) {
            switch (http.status) {
                case 200:
                    var tokens = JSON.parse(http.responseText);
                    setTokens(tokens["access_token"], tokens["refresh_token"]);
                    break;
                case 401:
                    alert("Wrong username or password");
                    break;
                default:
                    alert(http.responseText);
            }
        }
    };

    var data = JSON.stringify({"username": usernameInput.value, "password": passwordInput.value});
    http.send(data);
}

function signUp(event) {
    event.preventDefault();

    var usernameInput = document.getElementById("registration-form-username");
    var passwordInput = document.getElementById("registration-form-password");

    var http = new XMLHttpRequest();
    var url = "/sign_up";
    http.open("POST", url, true);
    http.setRequestHeader("Content-Type", "application/json");

    http.onreadystatechange = function () {
        if (http.readyState === XMLHttpRequest.DONE) {
            switch (http.status) {
                case 200:
                    var tokens = JSON.parse(http.responseText);
                    setTokens(tokens["access_token"], tokens["refresh_token"]);
                    break;
                default:
                    alert(http.responseText);
            }
        }
    };

    var data = JSON.stringify({"username": usernameInput.value, "password": passwordInput.value});
    http.send(data);
}

function showNewRoomForm(event) {
    event.preventDefault();

    var newRoomForm = document.getElementById("new_room_form")
    newRoomForm.classList.remove("hidden")

}

function createNewRoom(event) {
    event.preventDefault();

    var nameInput = document.getElementById("new-room-form-name");

    var http = new XMLHttpRequest();
    var url = "/room";
    http.open("POST", url, true);
    http.setRequestHeader("Content-Type", "application/json");
    http.setRequestHeader("Authorization", "Bearer " + localStorage.getItem("access_token"));

    http.onreadystatechange = function () {
        if (http.readyState === XMLHttpRequest.DONE) {
            switch (http.status) {
                case 200:
                    flushRooms();
                    break;
                default:
                    alert(http.responseText);
            }
        }
    }

    var data = JSON.stringify({"title": nameInput.value,});
    http.send(data);

    var newRoomWindow = document.getElementById("new_room_form")
    newRoomWindow.classList.add("hidden")
}

function addBasicEventListeners() {
    var signUpButton = document.getElementById("sign_up");
    signUpButton.addEventListener("click", signUp);

    var signInButton = document.getElementById("sign_in");
    signInButton.addEventListener("click", signIn);

    var newRoomButton = document.getElementById("new_room");
    newRoomButton.addEventListener("click", showNewRoomForm);

    var createNewRoomButton = document.getElementById("create-new-room");
    createNewRoomButton.addEventListener("click", createNewRoom);

    var sendMessageButton = document.getElementById("send-message");
    var sendMessageInput = document.getElementById("message-input")
    sendMessageInput.addEventListener("keyup", sendMessageIfEnterPressed)
    sendMessageButton.addEventListener("click", sendMessage)
}

function refreshLocalStorage() {
    localStorage.setItem("active_chat", "0")
}

function receiveMessage(context) {
    var body = context.data;
    addMessage(
        body["User"]["username"],
        body["CreatedAt"],
        body["Text"]
    )
}

function start() {
    addBasicEventListeners();
    refreshLocalStorage();
}

const centrifuge = new Centrifuge('ws://localhost:8080/connection/websocket');
var sub = null
start();
