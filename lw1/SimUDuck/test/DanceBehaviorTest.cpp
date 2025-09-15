#include "gtest/gtest.h"
#include "gmock/gmock.h"

#include "lib/Duck/MallardDuck.h"
#include "lib/Duck/Dance/DanceNoWay.h"

#include "Mock/MockFlyBehavior.h"
#include "Mock/MockQuackBehavior.h"
#include "Mock/MockDanceBehavior.h"

using testing::Exactly;

class DuckDanceTest : public testing::Test
{
protected:
    std::unique_ptr<MockFlyBehavior> mockFly;
    std::unique_ptr<MockQuackBehavior> mockQuack;
    std::unique_ptr<MockDanceBehavior> mockDance;

    MockFlyBehavior* mockFlyPtr = nullptr;
    MockQuackBehavior* mockQuackPtr = nullptr;
    MockDanceBehavior* mockDancePtr = nullptr;

    void SetUp() override
    {
        mockFly = std::make_unique<MockFlyBehavior>();
        mockQuack = std::make_unique<MockQuackBehavior>();
        mockDance = std::make_unique<MockDanceBehavior>();

        mockFlyPtr = mockFly.get();
        mockQuackPtr = mockQuack.get();
        mockDancePtr = mockDance.get();
    }
};

TEST_F(DuckDanceTest, DanceCallsDanceBehavior)
{
    EXPECT_CALL(*mockDancePtr, Dance()).Times(Exactly(1));

    MallardDuck duck;
    duck.SetFlyBehavior(std::move(mockFly));
    duck.SetQuackBehavior(std::move(mockQuack));
    duck.SetDanceBehavior(std::move(mockDance));

    duck.Dance();
}

TEST_F(DuckDanceTest, SetDanceBehaviorChangesBehavior)
{
    MallardDuck duck;
    duck.SetFlyBehavior(std::move(mockFly));
    duck.SetQuackBehavior(std::move(mockQuack));
    duck.SetDanceBehavior(std::make_unique<DanceNoWay>());

    auto newMockDance = std::make_unique<MockDanceBehavior>();
    const MockDanceBehavior* newMockDancePtr = newMockDance.get();

    EXPECT_CALL(*newMockDancePtr, Dance()).Times(Exactly(1));

    duck.SetDanceBehavior(std::move(newMockDance));

    duck.Dance();
}

TEST_F(DuckDanceTest, DanceCalledMultipleTimes)
{
    EXPECT_CALL(*mockDancePtr, Dance()).Times(Exactly(3));

    MallardDuck duck;
    duck.SetFlyBehavior(std::move(mockFly));
    duck.SetQuackBehavior(std::move(mockQuack));
    duck.SetDanceBehavior(std::move(mockDance));

    duck.Dance();
    duck.Dance();
    duck.Dance();
}