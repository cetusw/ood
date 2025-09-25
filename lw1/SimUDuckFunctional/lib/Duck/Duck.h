#ifndef DUCK_H
#define DUCK_H

#include <functional>

class Duck
{
public:
    Duck(std::function<void()> &&flyBehavior,
         std::function<void()> &&quackBehavior,
         std::function<void()> &&danceBehavior)
        : m_flyBehavior(std::move(flyBehavior))
          , m_quackBehavior(std::move(quackBehavior))
          , m_danceBehavior(std::move(danceBehavior))
    {
    }

    void Fly() const { m_flyBehavior(); }
    void Quack() const { m_quackBehavior(); }
    void Dance() const { m_danceBehavior(); }

    void SetFlyBehavior(std::function<void()> &&flyBehavior)
    {
        m_flyBehavior = std::move(flyBehavior);
    }

    virtual void Display() const = 0;

    virtual ~Duck() = default;

private:
    std::function<void()> m_flyBehavior;
    std::function<void()> m_quackBehavior;
    std::function<void()> m_danceBehavior;
};

#endif
