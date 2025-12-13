#include "drawer.h"
#include "image.h"
#include "tile.h"
#include <iostream>

int main()
{
	constexpr Color skyColor = ToColor(135, 206, 235);
	constexpr Color wallColor = ToColor(222, 184, 135);
	constexpr Color roofColor = ToColor(139, 69, 19);
	constexpr Color doorColor = ToColor(92, 64, 51);
	constexpr Color windowColor = ToColor(255, 255, 0);
	constexpr Color chimneyColor = ToColor(112, 128, 144);
	constexpr Color smokeColor = ToColor(211, 211, 211);

	std::cout << "Generating a house image..." << std::endl;

	Image img({ 100, 100 }, skyColor);

	for (int y = 50; y < 90; ++y)
	{
		DrawLine(img, { 20, y }, { 80, y }, wallColor);
	}

	DrawLine(img, { 20, 50 }, { 50, 30 }, roofColor);
	DrawLine(img, { 50, 30 }, { 80, 50 }, roofColor);
	DrawLine(img, { 20, 50 }, { 80, 50 }, roofColor);

	for (int y = 70; y < 90; ++y)
	{
		DrawLine(img, { 45, y }, { 55, y }, doorColor);
	}

	for (int y = 60; y < 70; ++y)
	{
		DrawLine(img, { 25, y }, { 35, y }, windowColor);
	}

	for (int y = 28; y < 42; ++y)
	{
		DrawLine(img, { 65, y }, { 70, y }, chimneyColor);
	}

	DrawCircle(img, { 67, 25 }, 3, smokeColor);
	DrawCircle(img, { 72, 20 }, 4, smokeColor);
	DrawCircle(img, { 78, 15 }, 6, smokeColor);

	const std::string filename = "../house.ppm";
	img.SavePPM(filename);
	std::cout << "Image saved to " << filename << std::endl;

	std::cout << "\n--- CoW Statistics ---" << std::endl;
	std::cout << "Total tiles created during drawing: " << Tile::GetInstanceCount() << std::endl;

	return 0;
}