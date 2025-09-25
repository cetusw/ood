#ifndef DECOYDUCK_H
#define DECOYDUCK_H

#include "Duck.h"
#include <iostream>
#include <memory>
#include "FunctionalBehaviors.h"

class DecoyDuck : public Duck
{
public:
    DecoyDuck()
        : Duck(std::function(Behaviors::FlyNoWay),
               std::function(Behaviors::MuteQuackBehavior),
               std::function(Behaviors::DanceNoWay))
    {
    }

    void Display() const override
    {
        std::cout << "I'm decoy duck" << std::endl;
    }
};

#endif
