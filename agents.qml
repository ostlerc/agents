import QtQuick 2.0
import QtQuick.Controls 1.0
import QtQuick.Layouts 1.1

ApplicationWindow {
    width: 800
    height: 600
    color: "white"

    TabView {
        anchors.fill: parent
        Tab {
            anchors.fill: parent
            title: "grid"
            ColumnLayout {
                id: screen
                anchors.fill: parent

                Rectangle {
                    id: statusRect
                    height: statusRow.height + 10
                    Layout.fillWidth: true
                    border.color: "black"
                    border.width: 1
                    RowLayout {
                        id: statusRow
                        anchors { verticalCenter: parent.verticalCenter; margins: 5; horizontalCenter: parent.horizontalCenter }
                        Text {
                            id: statusText
                            objectName: "statusText"
                            text: "Click the grid cells to make a start, end, and walls."
                        }

                        Text {
                            text: "text 2"
                        }
                    }
                }

                Rectangle {
                    Layout.fillHeight: true
                    Layout.fillWidth: true
                    Flickable {
                        clip: true
                        anchors.centerIn: parent
                        width: { return Math.min(parent.width, g.width) }
                        height: { return Math.min(parent.height, g.height) }
                        contentWidth: g.width; contentHeight: g.height
                        flickableDirection: Flickable.HorizontalAndVerticalFlick
                        Grid {
                            id: g
                            objectName: "grid"

                            spacing: 1
                        }
                    }
                }

                Rectangle {
                    id: botBorder
                    height: 1
                    Layout.minimumHeight: 1
                    Layout.fillWidth: true
                    color: "black"
                }

                RowLayout {
                    id: toolBar
                    Layout.fillWidth: true
                    Layout.preferredHeight: buildBtn.height + 10

                    Button {
                        id: buildBtn
                        objectName: "newBtn"
                        anchors {  margins: 5 }
                        text: "New"
                        width: 40
                        onClicked: grid.buildGrid()
                    }

                    Button {
                        id: runBtn
                        objectName: "runBtn"
                        anchors {  margins: 5 }
                        text: "Run"
                        width: 40
                        enabled: false
                        //onClicked: grid.runAStar()
                    }

                    Text {
                        id: rowText
                        anchors {  margins: 5 }
                        text: "Rows:"
                    }

                    SpinBox {
                        id: colRect
                        objectName: "cols"
                        width: 45
                        anchors {  margins: 5 }
                        maximumValue: 50
                        minimumValue: 1
                        value: 25
                    }

                    Text {
                        id: colText
                        anchors {  margins: 5 }
                        text: "Columns:"
                    }

                    SpinBox {
                        id: rowRect
                        objectName: "rows"
                        width: 45
                        anchors {  margins: 5 }
                        maximumValue: 50
                        minimumValue: 1
                        value: 25
                    }
                    
                    Rectangle {
                        Layout.fillWidth: true
                    }

                    Text {
                        id: foodcnttxt
                        anchors {  margins: 5 }
                        text: "Food Count:"
                    }

                    SpinBox {
                        id: foodcntbox
                        width: 65
                        anchors {  margins: 5 }
                        maximumValue: 999
                        minimumValue: 1
                        value: 250
                    }

                    Text {
                        id: foodexptxt
                        anchors {  margins: 5 }
                        text: "Food Lifetime:"
                    }

                    SpinBox {
                        id: foodexpbox
                        width: 65
                        anchors {  margins: 5 }
                        maximumValue: 999
                        minimumValue: 1
                        value: 25
                    }
                }
            }
        }
        Tab {
            title: "stats"
        }
    }
}
