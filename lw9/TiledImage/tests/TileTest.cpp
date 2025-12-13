#include <gtest/gtest.h>
#include "tile.h"

class TileTest : public ::testing::Test {
};

TEST_F(TileTest, ConstructorFillsWithColor) {
	const Color bg = ToColor(100, 100, 100);
	Tile tile(bg);

	for (int y = 0; y < Tile::SIZE; ++y) {
		for (int x = 0; x < Tile::SIZE; ++x) {
			EXPECT_EQ(tile.GetPixel({x, y}), bg);
		}
	}
}

TEST_F(TileTest, SetAndGetPixelValid) {
	Tile tile(0);
	Point p = {4, 4};
	Color color = ToColor(255, 0, 0);

	tile.SetPixel(p, color);
	EXPECT_EQ(tile.GetPixel(p), color);
	EXPECT_EQ(tile.GetPixel({3, 4}), 0);
}

TEST_F(TileTest, OutOfBoundsAccess) {
	Tile tile(ToColor(255, 255, 255));
	std::vector<Point> outPoints = {
		{-1, 0}, {0, -1},
		{Tile::SIZE, 0}, {0, Tile::SIZE}
	};

	for (const auto& p : outPoints) {
		tile.SetPixel(p, ToColor(0, 0, 0));
		// GetPixel returns 0 (black) out of bounds logic in Tile::GetPixel
		EXPECT_EQ(tile.GetPixel(p), 0);
	}
}

TEST_F(TileTest, CopyConstructorDeepCopy) {
	Color c1 = ToColor(1, 1, 1);
	Color c2 = ToColor(2, 2, 2);

	Tile t1(c1);
	t1.SetPixel({0, 0}, c2);

	Tile t2 = t1;
	EXPECT_EQ(t2.GetPixel({0, 0}), c2);

	t2.SetPixel({0, 0}, c1);
	EXPECT_EQ(t2.GetPixel({0, 0}), c1);
	EXPECT_EQ(t1.GetPixel({0, 0}), c2);
}

TEST_F(TileTest, InstanceCounting) {
	int initialCount = Tile::GetInstanceCount();
	{
		Tile t1(0);
		EXPECT_EQ(Tile::GetInstanceCount(), initialCount + 1);
	}
	EXPECT_EQ(Tile::GetInstanceCount(), initialCount);
}