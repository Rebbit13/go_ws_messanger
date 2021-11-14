function addMessage(author, time, text) {
    var messageList = document.getElementById("messages-list")

    var message = document.createElement("li")
    message.classList.add("message")

    var authorTag = document.createElement("p")
    authorTag.classList.add("message-author")
    var authorText = document.createTextNode(author)
    authorTag.appendChild(authorText)

    var createdAtTag = document.createElement('p')
    createdAtTag.classList.add("message-created-at")
    var createdAtText = document.createTextNode(time)
    createdAtTag.appendChild(createdAtText)

    var textTag = document.createElement('p')
    textTag.classList.add("message-text")
    var  textText = document.createTextNode(text)
    textTag.appendChild(textText)

    message.appendChild(authorTag)
    message.appendChild(createdAtTag)
    message.appendChild(textTag)

    messageList.appendChild(message)
}

function addRoomItem(text) {
    var roomList = document.getElementById("room-list")

    var room = document.createElement("form")
    room.classList.add("room-list-item")

    var button = document.createElement("button")
    var buttonText = document.createTextNode(text)
    button.appendChild(buttonText)

    room.appendChild(button)

    roomList.appendChild(room)
}
