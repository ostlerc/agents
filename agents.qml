import QtQuick 2.2
import QtQuick.Controls 1.0
import QtQuick.Layouts 1.1
import QtQuick.Dialogs 1.0

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
                            text: "Click the grid cells to make a Nest, food, and walls."
                        }
                        Text {
                            text: "count:"
                            visible: sb1.visible
                        }
                        SpinBox {
                            id: sb1
                            objectName: "countSpinner"
                            maximumValue: 999
                            minimumValue: 1
                            value: 1
                            onEditingFinished: {
                                grid.setCount(value)
                            }
                            visible: false
                        }
                        Text {
                            visible: sb2.visible
                            text: "lifetime:"
                        }
                        SpinBox {
                            id: sb2
                            objectName: "lifeSpinner"
                            maximumValue: 999
                            minimumValue: 1
                            value: 25
                            onEditingFinished: {
                                grid.setLife(value)
                            }
                            visible: false
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
                Rectangle {
                    id: toolBar
                    Layout.fillWidth: true
                    Layout.preferredHeight: newBtn.height + 10
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
                            property int type: 0
                            objectName: "fileDialog"
                            title: "Choose a filename"
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
                                fileDialog.close()
                                fileDialog.type = 1 //load
                                fileDialog.open()
                            }
                        }
                        Button {
                            objectName: "runBtn"
                            text: "Run"
                            visible: false
                            //onClicked: grid.runAStar()
                        }
                        Text {
                            id: rowText
                            text: "Rows:"
                        }
                        SpinBox {
                            id: colRect
                            objectName: "cols"
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
                            maximumValue: 99
                            minimumValue: 1
                            value: 25
                        }
                        Rectangle {
                            Layout.fillWidth: true
                        }
                        Text {
                            id: foodcnttxt
                            text: "Food Count:"
                        }
                        SpinBox {
                            id: foodcntbox
                            objectName: "defaultFoodCountCombo"
                            maximumValue: 999
                            minimumValue: 1
                            value: 75
                        }
                        Text {
                            id: foodexptxt
                            text: "Food Life:"
                        }
                        SpinBox {
                            id: foodexpbox
                            objectName: "defaultFoodLifetimeCombo"
                            maximumValue: 999
                            minimumValue: 1
                            value: 125
                        }
                    }
                }
            }
        }
        Tab {
            title: "stats"
        }
    }
}
