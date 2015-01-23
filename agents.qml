import QtQuick 2.2
import QtQuick.Controls 1.0
import QtQuick.Layouts 1.1
import QtQuick.Dialogs 1.2

ApplicationWindow {
    width: 850
    height: 600
    color: "white"
    ColumnLayout {
        id: screen
        anchors.fill: parent
        spacing: 0
        Rectangle {
            id: statusRect
            Layout.fillWidth: true
            Layout.preferredHeight: statusRow.height + 10
            border.color: "black"
            border.width: 1
            RowLayout {
                id: statusRow
                anchors { verticalCenter: parent.verticalCenter; margins: 5; horizontalCenter: parent.horizontalCenter }
                Text {
                    id: statusText
                    objectName: "statusText"
                    text: "Click the grid cells to edit a puzzle"
                }
            }
        }
        Rectangle {
            Layout.fillHeight: true
            Layout.fillWidth: true
            Flickable {
                clip: true
                boundsBehavior: Flickable.StopAtBounds
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
        Rectangle {
            id: toolBar
            Layout.fillWidth: true
            Layout.preferredHeight: 70
            color: "white"
            TabView {
                id: btmtab
                anchors.fill: parent
                Tab {
                    title: "editor"
                    Rectangle {
                        color: "white"
                        RowLayout {
                            id: bottomLayout
                            anchors.fill: parent
                            Button {
                                id: newBtn
                                text: "New"
                                onClicked: grid.buildGrid()
                            }
                            Button {
                                text: "Save"
                                onClicked: {
                                    fileDialog.close()
                                    fileDialog.type = 0 //save
                                    fileDialog.open()
                                }
                            }
                            FileDialog {
                                id: fileDialog
                                title: "Choose a filename"
                                objectName: "fileDialog"

                                property int type: 0

                                onAccepted: {
                                    console.log("You chose: " + fileDialog.fileUrls)
                                    if(!type) {
                                        grid.saveGrid(fileDialog.fileUrl.toString())
                                    } else {
                                        grid.loadGrid(fileDialog.fileUrl.toString())
                                    }
                                }
                            }
                            Button {
                                text: "Load"
                                onClicked: {
                                    console.log(fileDialog)
                                    fileDialog.close()
                                    fileDialog.type = 1 //load
                                    fileDialog.open()
                                }
                            }
                            Text {
                                id: rowText
                                text: "Rows:"
                            }
                            SpinBox {
                                id: colRect
                                objectName: "cols"
                                Layout.minimumWidth: 42
                                maximumValue: 99
                                minimumValue: 1
                                value: 25
                            }
                            Text {
                                id: colText
                                text: "Columns:"
                            }
                            SpinBox {
                                id: rowRect
                                objectName: "rows"
                                Layout.minimumWidth: 42
                                maximumValue: 99
                                minimumValue: 1
                                value: 25
                            }
                            Rectangle {
                                Layout.fillWidth: true
                            }
                        }
                    }
                }

                Tab {
                    title: "simulator"
                    Rectangle {
                        color: "white"
                        RowLayout {
                            anchors.fill: parent
                        }
                    }
                }
            }
        }
    }
}
