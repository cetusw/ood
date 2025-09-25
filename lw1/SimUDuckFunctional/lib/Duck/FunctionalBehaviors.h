#ifndef SIMUDUCKFUNCTIONAL_FUNCTIONALBEHAVIORS_H
#define SIMUDUCKFUNCTIONAL_FUNCTIONALBEHAVIORS_H
#include <functional>
#include <iostream>

namespace Behaviors
{
    inline const std::function<void()> &quackBehavior()
    {
        static const std::function behavior = [] { std::cout << "Quack" << std::endl; };
        return behavior;
    }

    inline const std::function<void()> &squeakBehavior()
    {
        static const std::function behavior = [] { std::cout << "Squeak" << std::endl; };
        return behavior;
    }

    const std::function MuteQuackBehavior = []
    {
    };

    inline const std::function<void()> &danceWaltz()
    {
        static const std::function behavior = [] { std::cout << "I'm dancing the waltz now!" << std::endl; };
        return behavior;
    }

    inline const std::function<void()> &danceMinuet()
    {
        static const std::function behavior = [] { std::cout << "I'm dancing a minuet now!" << std::endl; };
        return behavior;
    }

    const std::function DanceNoWay = []
    {
    };

    const std::function FlyNoWay = []
    {
    };

    inline auto makeFlyWithWings()
    {
        return [flightsCount = 0]() mutable // TODO: что такое?
        {
            flightsCount++;
            std::cout << "I'm flying with wings!! Flight number: " << flightsCount << std::endl;
        };
    }
} // namespace Behaviors

#endif //SIMUDUCKFUNCTIONAL_FUNCTIONALBEHAVIORS_H
