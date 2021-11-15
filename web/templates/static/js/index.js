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
}

function addRoomItem(text) {
    var roomList = document.getElementById("room-list");

    var room = document.createElement("form");
    room.classList.add("room-list-item");

    var button = document.createElement("button");
    var buttonText = document.createTextNode(text);
    button.appendChild(buttonText);

    room.appendChild(button);
    roomList.appendChild(room);
}

function startMessenging() {
    var rooms = document.getElementById("room-list")
}

function setTokens(access_token, refresh_token) {
    document.cookie = "access_token=" + access_token;
    document.cookie = "refresh_token=" + refresh_token
    var registrationWindow = document.getElementsByClassName("registration")[0]
    registrationWindow.classList.remove("hidden")
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


function addEventListeners() {
    var signUpButton = document.getElementById("sign_up");
    var signInButton = document.getElementById("sign_in");
    signUpButton.addEventListener("click", signUp);
    signInButton.addEventListener("click", signIn)
}

addEventListeners();
