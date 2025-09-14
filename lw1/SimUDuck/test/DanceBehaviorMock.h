#ifndef SIMUDUCK_DANCEBEHAVIORMOCK_H
#define SIMUDUCK_DANCEBEHAVIORMOCK_H

#include "../lib/Duck/Dance/IDanceBehavior.h"
#include "gmock/gmock.h"
#include "gtest/gtest.h"

class MockDanceBehavior : public IDanceBehavior
{
public:
    MOCK_METHOD(void, Dance, (), (override));
};

#endif //SIMUDUCK_DANCEBEHAVIORMOCK_H
