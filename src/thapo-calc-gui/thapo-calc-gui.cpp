#include <QGuiApplication>
#include <QQmlApplicationEngine>

int main(int argc, char *argv[]) {

    QCoreApplication::setAttribute(Qt::AA_EnableHighDpiScaling);

    QGuiApplication app(argc, argv);

    QQmlApplicationEngine engine;
    engine.load(QUrl(QStringLiteral("qrc:/thapo-calc-gui.qml")));
    if (engine.rootObjects().isEmpty())
        return -1;

    return app.exec();
}
