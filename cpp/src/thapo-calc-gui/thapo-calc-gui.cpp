#include <QGuiApplication>
#include <QQmlApplicationEngine>
#include "Controller.h"

int main(int argc, char *argv[]) {

    QCoreApplication::setAttribute(Qt::AA_EnableHighDpiScaling);

    QGuiApplication app(argc, argv);

    qmlRegisterType<Controller>("thapo.calc.Controller", 1, 0, "Controller");
    QQmlApplicationEngine engine;
    engine.load(QUrl(QStringLiteral("qrc:/thapo-calc-gui.qml")));
    if (engine.rootObjects().isEmpty())
        return -1;

    return app.exec();
}
