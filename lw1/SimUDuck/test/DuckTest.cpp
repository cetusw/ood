#include "DanceBehaviorMock.h"
#include "gtest/gtest.h"
#include "../lib/Duck/Duck.h"
#include "gmock/gmock.h"

class FlyNoWayBehavior : public IFlyBehavior {
public:
    void Fly() override {}
};

class MuteQuackBehavior : public IQuackBehavior {
public:
    void Quack() override {}
};

class TestDuck : public Duck {
public:
    using Duck::Duck;
    void Display() const override {}
};

class DuckTest : public ::testing::Test
{
protected:
    void SetUp() override
    {
        auto mockDanceBehaviorPtr = std::make_unique<MockDanceBehavior>();
        rawMockDanceBehavior = mockDanceBehaviorPtr.get();

        duck = std::make_unique<TestDuck>(
            std::make_unique<FlyNoWayBehavior>(),
            std::make_unique<MuteQuackBehavior>(),
            std::move(mockDanceBehaviorPtr));
    }

    MockDanceBehavior* rawMockDanceBehavior{};
    std::unique_ptr<Duck> duck;
};



TEST_F(DuckTest, DanceMethodDelegatesCallToDanceBehavior)
{
    EXPECT_CALL(*rawMockDanceBehavior, Dance()).Times(1);
    duck->Dance();
}

