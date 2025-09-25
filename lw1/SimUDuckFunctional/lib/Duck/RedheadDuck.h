#ifndef REDHEADDUCK_H
#define REDHEADDUCK_H

#include "Duck.h"
#include "FunctionalBehaviors.h"

class RedheadDuck : public Duck
{
public:
    RedheadDuck()
        : Duck(Behaviors::makeFlyWithWings(),
               std::function(Behaviors::quackBehavior),
               std::function(Behaviors::danceMinuet))
    {
    }

    void Display() const override
    {
        std::cout << "I'm redhead duck" << std::endl;
    }
};

#endif
