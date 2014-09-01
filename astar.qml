import QtQuick 2.0
import QtQuick.Controls 1.0
import "../astar" as S

Rectangle {
    id: screen
    width: 400
    height: 300

    SystemPalette { id: activePalette }

    Item {
        id: m
        anchors.fill: parent

        Flickable {
            anchors.fill: parent
            contentWidth: g.width
            contentHeight: g.height
            flickableDirection: Flickable.HorizontalAndVerticalFlick
            Grid {
                id: g
                anchors.centerIn: parent

                columns: 10
                spacing: 1

                Repeater {
                    model: 100
                    delegate: S.Tile { type: index % 4; }
                }
            }
        }
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
