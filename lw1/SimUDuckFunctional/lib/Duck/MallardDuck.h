#ifndef MALLARDDUCK_H
#define MALLARDDUCK_H

#include "Duck.h"
#include "FunctionalBehaviors.h"

class MallardDuck : public Duck
{
public:
    MallardDuck()
        : Duck(Behaviors::makeFlyWithWings(),
               Behaviors::quackBehavior,
               Behaviors::danceWaltz)
    {
    }

    void Display() const override
    {
        std::cout << "I'm a real Mallard duck" << std::endl;
    }
};

#endif
