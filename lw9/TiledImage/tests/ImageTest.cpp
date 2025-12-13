#include "image.h"
#include "tile.h"
#include <gtest/gtest.h>

class ImageTest : public ::testing::Test
{
};

TEST_F(ImageTest, ConstructorValid)
{
	constexpr Size sz{ 10, 20 };
	constexpr Color bg = ToColor(10, 20, 30);
	const Image img(sz, bg);

	EXPECT_EQ(img.GetPixel({ 0, 0 }), bg);
	EXPECT_EQ(img.GetPixel({ 9, 19 }), bg);
}

TEST_F(ImageTest, OutOfBoundsAccess)
{
	Image img({ 10, 10 }, ToColor(255, 255, 255));
	EXPECT_EQ(img.GetPixel({ -1, 5 }), 0);

	img.SetPixel({ 100, 100 }, ToColor(1, 1, 1));
	EXPECT_EQ(img.GetSize().width, 10);
}

TEST_F(ImageTest, CoWModifyTriggersCopy)
{
	const int baseline = Tile::GetInstanceCount();

	Image img({ 16, 16 }, 0);
	EXPECT_EQ(Tile::GetInstanceCount(), baseline + 1);

	img.SetPixel({ 0, 0 }, ToColor(128, 128, 128));
	EXPECT_EQ(Tile::GetInstanceCount(), baseline + 2);
	EXPECT_EQ(img.GetPixel({ 0, 0 }), ToColor(128, 128, 128));
}

TEST_F(ImageTest, CoWReadDoesNotCopy)
{
	const int baseline = Tile::GetInstanceCount();
	const Image img({ 16, 16 }, 0);

	const Color c = img.GetPixel({ 5, 5 });
	EXPECT_EQ(c, 0);
	EXPECT_EQ(Tile::GetInstanceCount(), baseline + 1);
}