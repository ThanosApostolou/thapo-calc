import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQml 2.3
import QtQuick.Layouts 1.13
import thapo.calc.Controller 1.0

ApplicationWindow {
    id: main_window
    visible: true
    title: qsTr("Thapo Calc")
    property int margin: 30
    width: columnLayout.implicitWidth + 2 * margin
    height: columnLayout.implicitHeight + 2 * margin
    minimumWidth: columnLayout.Layout.minimumWidth + 2 * margin
    minimumHeight: columnLayout.Layout.minimumHeight + 2 * margin

    Controller {
        id: controller
    }

    ColumnLayout {
        id: columnLayout
        anchors.fill: parent
        anchors.margins: margin
        Layout.minimumWidth: 400
        Layout.minimumHeight: 500

        TextArea {
            id: history
            Layout.fillWidth: true
            Layout.fillHeight: true
            Layout.preferredWidth: 30
            Layout.preferredHeight: 20
            readOnly: true
            text: controller.history
        }

        RowLayout {
            id: rowLayout1

            TextField {
                id: input
                Layout.fillWidth: true
                Layout.fillHeight: true
                Layout.preferredWidth: 70
                Layout.preferredHeight: 10
                text: controller.input
                onTextChanged: controller.input = text
            }

            TextField {
                id: output
                Layout.fillWidth: true
                Layout.fillHeight: true
                Layout.preferredWidth: 30
                Layout.preferredHeight: 10
                readOnly: true
                text: controller.output
            }

        }

        GridLayout {
            id: grid
            rows: 4
            columns: 4

            Repeater {
                id: buttons_line1
                model: ["1", "2", "3", "^",
                        "4", "5", "6", "*",
                        "7", "8", "9", "/",
                        "0", "+", "-"]
                Button {
                    Layout.fillWidth: true
                    Layout.fillHeight: true
                    Layout.preferredHeight: 90/4
                    text: modelData
                    onClicked: controller.input += this.text
                }
            }
            Button {
                id: save
                Layout.fillWidth: true
                Layout.fillHeight: true
                Layout.preferredHeight: 90/4
                text: "="
                onClicked: history.text += input.text + "" + output.text + "\n"
            }
        }

    }

}
