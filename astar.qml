import QtQuick 2.0
import QtQuick.Controls 1.0
import "../astar" as S

Rectangle {
    id: screen
    width: 400
    height: 400

    SystemPalette { id: activePalette }

    Text {
        id: statusText
        objectName: "statusText"
        text: "Click the grid cells to make a start, end, and walls."
        anchors { top: parent.top; margins: 5; horizontalCenter: parent.horizontalCenter }
    }

    Item {
        id: m
        width: parent.width
        anchors { top: statusText.bottom; bottom: toolBar.top }

        Flickable {
            anchors.fill: parent
            contentWidth: g.width
            contentHeight: g.height
            flickableDirection: Flickable.HorizontalAndVerticalFlick
            Grid {
                id: g
                objectName: "grid"
                anchors.centerIn: parent

                spacing: 1
            }
        }
    }

    Rectangle {
        id: toolBar
        width: parent.width
        height: buildBtn.height + 10
        color: "white"
        border.color: "black"
        border.width: 1
        anchors.bottom: screen.bottom

        Button {
            id: buildBtn
            objectName: "newBtn"
            anchors { left: parent.left; verticalCenter: parent.verticalCenter; margins: 5 }
            text: "New"
            width: 40
            onClicked: grid.buildGrid()
        }

        Button {
            id: runBtn
            objectName: "runBtn"
            anchors { left: buildBtn.right; verticalCenter: parent.verticalCenter; margins: 5 }
            text: "Run"
            width: 40
            enabled: false
            onClicked: grid.runAStar()
        }

        Text {
            id: rowText
            anchors { left: runBtn.right; verticalCenter: parent.verticalCenter; margins: 5 }
            text: "Rows:"
        }
        Text {
            id: colText
            anchors { left: rowRect.right; verticalCenter: parent.verticalCenter; margins: 5 }
            text: "Columns:"
        }

        SpinBox {
            id: colRect
            objectName: "cols"
            width: 45
            anchors { left: colText.right; verticalCenter: parent.verticalCenter; margins: 5 }
            maximumValue: 20
            minimumValue: 1
            value: 10
        }

        SpinBox {
            id: rowRect
            objectName: "rows"
            width: 45
            anchors { left: rowText.right; verticalCenter: parent.verticalCenter; margins: 5 }
            maximumValue: 20
            minimumValue: 1
            value: 10
        }
    }
}
