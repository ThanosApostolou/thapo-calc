#ifndef CONTROLLER_H
#define CONTROLLER_H

#include <QObject>
#include <QString>

class Controller : public QObject {
    Q_OBJECT
    Q_PROPERTY(QString input READ getInput WRITE setInput NOTIFY inputChanged)
    Q_PROPERTY(QString output READ getOutput WRITE setOutput NOTIFY outputChanged)

public:
    explicit Controller(QObject *parent = nullptr);

    QString getInput();
    void setInput(const QString &input);
    QString getOutput();
    void setOutput(const QString &output);


signals:
    void inputChanged();
    void outputChanged();

private:
    QString m_input="";
    QString m_output="= ";
};

#endif // CONTROLLER_H
