#ifndef SIMUDUCK_MOCKFLYBEHAVIOR_H
#define SIMUDUCK_MOCKFLYBEHAVIOR_H

#include "gmock/gmock.h"
#include "lib/Duck/Fly/IFlyBehavior.h"

class MockFlyBehavior : public IFlyBehavior
{
public:
    MOCK_METHOD(bool, Fly, (), (override));
};

#endif //SIMUDUCK_MOCKFLYBEHAVIOR_H
