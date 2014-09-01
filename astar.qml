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
        width: parent.width
        anchors { top: parent.top; bottom: toolBar.top }

        Flickable {
            anchors.fill: parent
            contentWidth: g.width
            contentHeight: g.height
            flickableDirection: Flickable.HorizontalAndVerticalFlick
            Grid {
                id: g
                objectName: "grid"
                anchors.centerIn: parent

                columns: grid.columnCount
                rows: grid.rowCount
                spacing: 1
            }
        }
    }

    Rectangle {
        id: toolBar
        width: parent.width; height: runBtn.height
        color: "white"
        border.color: "black"
        border.width: 1
        anchors.bottom: screen.bottom

        Button {
            id: buildBtn
            anchors { left: parent.left; verticalCenter: parent.verticalCenter; margins: 5 }
            text: "Create Grid"
            onClicked: grid.buildGrid()
        }

        Button {
            id: runBtn
            anchors { left: buildBtn.right; verticalCenter: parent.verticalCenter; margins: 5 }
            text: "Run AStar"
            onClicked: grid.runAStar()
        }

        Text {
            id: colText
            anchors { left: runBtn.right; top: parent.top; margins: 5 }
            text: "Columns:"
        }
        Text {
            id: rowText
            anchors { left: colRect.right; top: parent.top; margins: 5 }
            text: "Rows:"
        }

        Rectangle {
            id: colRect
            border.color: "black"
            border.width: 1
            anchors { left: colText.right; top: parent.top; margins: 5 }
            color: "red"
            width: 30
            height: c.height
            radius: 5
            TextInput {
                id: r
                objectName: "rows"
                focus: true
                text: "10"
                validator: IntValidator{bottom:1; top: 20}
                anchors { verticalCenter: parent.verticalCenter; horizontalCenter: parent.horizontalCenter }
                onAccepted: grid.rowsClicked()
            }
            MouseArea {
                anchors.fill: parent
                onClicked: {
                    r.forceActiveFocus()
                }
            }
        }

        Rectangle {
            id: rowRect
            border.color: "black"
            border.width: 1
            anchors { left: rowText.right; top: parent.top; leftMargin: 5; topMargin: 5 }
            color: "lightsteelblue"
            width: 30
            height: c.height
            radius: 5
            TextInput {
                id: c
                objectName: "cols"
                validator: IntValidator{bottom:1; top: 20}
                focus: true
                text: "10"
                anchors { verticalCenter: parent.verticalCenter; horizontalCenter: parent.horizontalCenter }
                onAccepted: grid.colsClicked()
            }
            MouseArea {
                anchors.fill: parent
                onClicked: {
                    c.forceActiveFocus()
                }
            }
        }
    }
}
