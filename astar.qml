import QtQuick 2.0

Rectangle {
    id: screen
    width: 490; height: 720

    SystemPalette { id: activePalette }

    Item {
        width: parent.width
        anchors { top: parent.top; bottom: toolBar.top }
    }

    Rectangle {
        id: toolBar
        width: parent.width; height: 30
        color: activePalette.window
        anchors.bottom: screen.bottom

        Button {
            anchors { left: parent.left; verticalCenter: parent.verticalCenter }
            text: "New Game"
            onClicked: grid.clicked()
        }

        Text {
            objectName: "text"
            id: score
            anchors { right: parent.right; verticalCenter: parent.verticalCenter }
            text: "0"
        }
    }
}
